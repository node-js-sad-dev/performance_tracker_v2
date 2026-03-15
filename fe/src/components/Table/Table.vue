<script setup lang="ts">
import TableRow from '../TableRow/TableRow.vue';
import type { Props } from './interfaces.ts';

const props = defineProps<Props>();

const emit = defineEmits<{
  (event: 'sort', column: string): void;
  (event: 'view', item: Record<string, unknown>): void;
  (event: 'edit', item: Record<string, unknown>): void;
  (event: 'delete', item: Record<string, unknown>): void;
}>();

function getSortIndicator(columnSortBy?: string) {
  if (!columnSortBy || columnSortBy !== props.sortBy) {
    return '';
  }

  return props.sortOrder === 'asc' ? '^' : 'v';
}
</script>

<template>
  <div class="table-shell">
    <table class="data-table">
      <thead>
        <tr>
          <th v-for="column in props.columns" :key="column.key" :class="column.align">
            <button
              v-if="column.sortBy"
              class="sort-button"
              type="button"
              @click="emit('sort', column.sortBy)">
              <span>{{ column.label }}</span>
              <span class="sort-indicator">{{ getSortIndicator(column.sortBy) }}</span>
            </button>
            <span v-else>{{ column.label }}</span>
          </th>
          <th class="right-align">Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-if="props.loading">
          <td :colspan="props.columns.length + 1" class="table-state loading-state">Loading data...</td>
        </tr>
        <tr v-else-if="props.rows.length === 0">
          <td :colspan="props.columns.length + 1" class="table-state empty-state">{{ props.emptyLabel }}</td>
        </tr>
        <TableRow
          v-for="item in props.rows"
          v-else
          :key="String(item.id ?? JSON.stringify(item))"
          :item="item"
          :columns="props.columns"
          @view="emit('view', item)"
          @edit="emit('edit', item)"
          @delete="emit('delete', item)" />
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.table-shell {
  background: var(--bg-panel);
  border: 1px solid var(--border);
  border-radius: 16px;
  overflow-x: auto;
  margin-bottom: 1rem;
  padding: 0 0.75rem 0.75rem;
}

.data-table {
  width: 100%;
  border-collapse: separate;
  border-spacing: 0 0.75rem;
}

th {
  padding: 0.95rem 1.15rem 0.6rem;
  text-align: left;
  vertical-align: middle;
  font-size: 0.82rem;
  letter-spacing: 0.06em;
  text-transform: uppercase;
  color: var(--text-muted);
  white-space: nowrap;
}

thead tr th {
  border-bottom: 1px solid var(--border);
}

.right-align,
th.right-align {
  text-align: right;
}

th.center {
  text-align: center;
}

th.right {
  text-align: right;
}

.sort-button {
  display: inline-flex;
  align-items: center;
  gap: 0.45rem;
  background: transparent;
  border: none;
  color: inherit;
  font: inherit;
  cursor: pointer;
  text-transform: inherit;
  letter-spacing: inherit;
}

.sort-indicator {
  color: var(--accent);
  min-width: 0.7rem;
}

.table-state {
  padding: 1.35rem 1.25rem;
  text-align: center;
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-radius: 14px;
}

@media (max-width: 900px) {
  .table-shell {
    padding-bottom: 0.5rem;
  }

  .data-table {
    min-width: 780px;
  }
}
</style>
