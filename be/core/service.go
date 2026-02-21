package core

import (
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

func ApplyFilters(
	queryBuilder *strings.Builder,
	filters map[string][]string,
	filterRules map[string]FilterRule,
	args *[]interface{},
	argCounter *int,
) {
	var whereClauses []string

	for key, values := range filters {
		rule, isAllowed := filterRules[key]

		if !isAllowed || len(values) == 0 {
			continue
		}

		var orConditions []string
		for _, val := range values {
			if val == "" {
				continue
			}

			var dbValue string
			if rule.IsFuzzy {
				dbValue = "%" + val + "%"
			} else {
				dbValue = val
			}

			orConditions = append(orConditions, fmt.Sprintf("%s %s $%d", rule.DBColumn, rule.Operator, *argCounter))

			*args = append(*args, dbValue)

			*argCounter++
		}

		if len(orConditions) > 0 {
			whereClauses = append(whereClauses, "("+strings.Join(orConditions, " OR ")+")")
		}
	}

	if len(whereClauses) > 0 {
		queryBuilder.WriteString(" WHERE " + strings.Join(whereClauses, " AND "))
	}
}

func GetEntityList(payload GetEntityListPayload) (pgx.Rows, error) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString(`SELECT `)
	queryBuilder.WriteString(strings.Join(payload.SelectFields, ", "))
	queryBuilder.WriteString(` FROM ` + payload.TableName)

	var args []interface{}
	argCounter := 1

	ApplyFilters(&queryBuilder, payload.Filters, payload.FilterRules, &args, &argCounter)

	var sortOrder string
	if strings.ToUpper(payload.Sort.SortOrder) == "ASC" {
		sortOrder = "ASC"
	} else {
		sortOrder = "DESC"
	}

	queryBuilder.WriteString(fmt.Sprintf(" ORDER BY %s %s", payload.Sort.SortBy, sortOrder))

	queryBuilder.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d", argCounter, argCounter+1))
	args = append(args, payload.Pagination.Limit, (payload.Pagination.Page-1)*payload.Pagination.Limit)

	return payload.Pool.Query(payload.Context, queryBuilder.String(), args...)
}

func GetEntityCount(payload GetEntityCountPayload) (int64, error) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString(`SELECT COUNT(*) as count FROM ` + payload.TableName)

	var args []interface{}
	argCounter := 1

	ApplyFilters(&queryBuilder, payload.Filters, payload.FilterRules, &args, &argCounter)

	row := payload.Pool.QueryRow(payload.Context, queryBuilder.String(), args...)

	var count int64

	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func GetUpdateQueryArgs(updates map[string]IOptionalField) (string, []interface{}, int) {
	var setClauses []string
	var args []interface{}
	argID := 1

	for col, field := range updates {
		if !field.GetIsSet() {
			continue
		}

		setClauses = append(setClauses, fmt.Sprintf(`"%s" = $%d`, col, argID))

		if field.GetIsNull() {
			args = append(args, nil)
		} else {
			args = append(args, field.GetValue())
		}

		argID++
	}

	return strings.Join(setClauses, ", "), args, argID
}

func UpdateEntity(payload UpdateEntityByIdPayload) error {
	clauses, args, argID := GetUpdateQueryArgs(payload.Updates)

	if len(clauses) == 0 {
		return nil
	}

	query := fmt.Sprintf(
		`UPDATE "%s" SET %s WHERE id = $%d`,
		payload.Table,
		clauses,
		argID,
	)

	args = append(args, payload.ID)

	_, err := payload.Executor.Exec(payload.Context, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update entity in table %s: %w", payload.Table, err)
	}

	return nil
}
