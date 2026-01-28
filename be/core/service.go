package core

import (
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

func GetEntityList(payload GetEntityListPayload) (pgx.Rows, error) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString(`SELECT `)
	queryBuilder.WriteString(strings.Join(payload.SelectFields, ", "))
	queryBuilder.WriteString(` FROM ` + payload.TableName)

	var args []interface{}
	var whereClauses []string
	argCounter := 1

	for key, values := range payload.Filters {
		rule, isAllowed := payload.FilterRules[key]

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

			orConditions = append(orConditions, fmt.Sprintf("%s %s $%d", rule.DBColumn, rule.Operator, argCounter))
			args = append(args, dbValue)
			argCounter++
		}

		if len(orConditions) > 0 {
			whereClauses = append(whereClauses, "("+strings.Join(orConditions, " OR ")+")")
		}
	}

	if len(whereClauses) > 0 {
		queryBuilder.WriteString(" WHERE " + strings.Join(whereClauses, " AND "))
	}

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
