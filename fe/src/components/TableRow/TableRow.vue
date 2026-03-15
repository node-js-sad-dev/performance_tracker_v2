<script setup lang="ts">
import type { Tone } from '../../core/interfaces.ts';
import type { RowProps } from '../Table/interfaces.ts';

const props = defineProps<RowProps>();

const emit = defineEmits<{
  (event: 'view'): void;
  (event: 'edit'): void;
  (event: 'delete'): void;
}>();

function resolveTone(columnIndex: number): Tone {
  const column = props.columns[columnIndex];
  return column?.getTone?.(props.item) ?? 'default';
}

function resolveRawValue(columnIndex: number): string {
  const column = props.columns[columnIndex];
  const value = column?.getValue(props.item);
  if (value === null || value === undefined || value === '') {
    return '--';
  }

  return String(value);
}

function formatTimeValue(value: string): string {
  const match = value.match(/^(\d+):(\d{2})\.(\d{3})$/);
  if (!match) {
    return value;
  }

  const minutes = Number(match[1]);
  if (minutes < 60) {
    return value;
  }

  const hours = Math.floor(minutes / 60);
  const remainingMinutes = minutes % 60;

  return `${hours}:${String(remainingMinutes).padStart(2, '0')}:${match[2]}.${match[3]}`;
}

function resolveDisplayValue(columnIndex: number): string {
  const value = resolveRawValue(columnIndex);
  const column = props.columns[columnIndex];

  if (column?.kind === 'time') {
    return formatTimeValue(value);
  }

  return value;
}
</script>

<template>
  <tr class="data-row">
    <td
      v-for="(column, index) in props.columns"
      :key="column.key"
      :class="[column.align, { 'time-cell': column.kind === 'time' }]">
      <span
        v-if="column.kind === 'badge'"
        class="badge"
        :class="`badge-${resolveTone(index)}`">
        {{ resolveDisplayValue(index) }}
      </span>
      <span
        v-else-if="column.kind === 'time'"
        class="time-pill mono"
        :title="resolveRawValue(index)">
        {{ resolveDisplayValue(index) }}
      </span>
      <span v-else :class="{ mono: column.kind === 'datetime' }">
        {{ resolveDisplayValue(index) }}
      </span>
    </td>
    <td class="actions-cell">
      <div class="actions-group">
        <button class="action-link" type="button" @click="emit('view')">View</button>
        <button class="action-link accent" type="button" @click="emit('edit')">Edit</button>
        <button class="action-link danger" type="button" @click="emit('delete')">Delete</button>
      </div>
    </td>
  </tr>
</template>

<style scoped>
.data-row td {
  padding: 1rem 1.15rem;
  vertical-align: middle;
  background: rgba(255, 255, 255, 0.025);
  border-top: 1px solid rgba(255, 255, 255, 0.05);
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
  transition:
    background-color 0.2s ease,
    border-color 0.2s ease,
    transform 0.2s ease;
}

.data-row td:first-child {
  border-left: 1px solid rgba(255, 255, 255, 0.05);
  border-top-left-radius: 14px;
  border-bottom-left-radius: 14px;
}

.data-row td:last-child {
  border-right: 1px solid rgba(255, 255, 255, 0.05);
  border-top-right-radius: 14px;
  border-bottom-right-radius: 14px;
}

.data-row:hover td {
  background: rgba(255, 255, 255, 0.045);
  border-color: rgba(0, 229, 255, 0.16);
}

.center {
  text-align: center;
}

.center .badge {
  margin-inline: auto;
}

.right {
  text-align: right;
}

.actions-cell {
  white-space: nowrap;
}

.actions-group {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 0.9rem;
}

.action-link {
  background: transparent;
  border: none;
  color: var(--text-muted);
  cursor: pointer;
  padding: 0;
}

.action-link:hover {
  color: var(--text-main);
}

.action-link.accent:hover {
  color: var(--accent);
}

.action-link.danger:hover {
  color: var(--danger);
}

.badge {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-width: 5.75rem;
  padding: 0.45rem 0.8rem;
  border-radius: 999px;
  border: 1px solid transparent;
  font-size: 0.78rem;
  font-weight: 700;
  letter-spacing: 0.02em;
}

.badge-success {
  background: rgba(12, 198, 107, 0.12);
  border-color: rgba(12, 198, 107, 0.2);
  color: var(--success);
}

.badge-danger {
  background: rgba(255, 95, 95, 0.12);
  border-color: rgba(255, 95, 95, 0.2);
  color: var(--danger);
}

.badge-accent {
  background: rgba(0, 229, 255, 0.1);
  border-color: rgba(0, 229, 255, 0.22);
  color: var(--accent);
}

.badge-muted,
.badge-default {
  background: rgba(255, 255, 255, 0.04);
  border-color: rgba(255, 255, 255, 0.08);
  color: var(--text-muted);
}

.time-cell {
  text-align: right;
}

.time-pill {
  display: inline-flex;
  align-items: center;
  justify-content: flex-end;
  min-width: 10.5ch;
  padding: 0.45rem 0.75rem;
  border-radius: 999px;
  background: rgba(0, 229, 255, 0.08);
  border: 1px solid rgba(0, 229, 255, 0.18);
  color: #b5f9ff;
  letter-spacing: 0.03em;
  font-feature-settings: 'tnum' 1;
}

.mono {
  font-family: var(--font-mono);
}
</style>


