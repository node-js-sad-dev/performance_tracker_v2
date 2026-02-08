<script setup lang="ts">
import TableRow from "../TableRow/TableRow.vue";
</script>

<template>
  <table class="data-table">
    <thead>
    <tr>
      <th>Date</th>
      <th>Game</th>
      <th>Car</th>
      <th>Track</th>
      <th>Time</th>
      <th>Clear?</th>
      <th class="text-right">Actions</th>
    </tr>
    </thead>
    <tbody>
    <tr>
      <td class="mono">2023-10-27</td>
      <td>Assetto Corsa</td>
      <td>Porsche 911 GT3</td>
      <td>Nürburgring</td>
      <td class="mono time-highlight">
        1:54.302
      </td>
      <td>
        <span class="badge success">Yes</span>
      </td>
      <td class="actions-cell">
        <button
            class="icon-btn edit"
            title="Edit"
        >
          ✎
        </button>
        <button
            class="icon-btn delete"
            title="Delete"
        >
          🗑
        </button>
        <button
            class="icon-btn detail"
            title="Details"
        >
          👁
        </button>
      </td>
    </tr>

    <TableRow></TableRow>

    </tbody>
  </table>
</template>

<style>
.data-table {
  width: 100%;
  border-collapse: collapse;
  text-align: left;
}

.data-table th,
.data-table td {
  padding: 1rem 1.5rem;
  border-bottom: 1px solid var(--border);
}

.data-table th {
  background-color: rgba(255, 255, 255, 0.03);
  color: var(--text-muted);
  font-size: 0.85rem;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.data-table tbody tr:hover {
  background-color: var(--bg-hover);
}

.data-table tbody tr:last-child td {
  border-bottom: none;
}

/* Table Specific Typography */
.mono {
  font-family: var(--font-mono);
  font-size: 0.9rem;
}

.time-highlight {
  color: var(--accent);
  font-weight: bold;
}

.text-right {
  text-align: right;
}

/* Status Badges */
.badge {
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 0.75rem;
  font-weight: bold;
  text-transform: uppercase;
}

.badge.success {
  background: rgba(0, 200, 81, 0.2);
  color: var(--success);
}

.badge.danger {
  background: rgba(255, 68, 68, 0.2);
  color: var(--danger);
}

/* Actions */
.actions-cell {
  text-align: right;
  white-space: nowrap;
}

.icon-btn {
  background: none;
  border: none;
  color: var(--text-muted);
  cursor: pointer;
  font-size: 1.1rem;
  margin-left: 8px;
  padding: 4px;
  transition: color 0.2s;
}

.icon-btn:hover {
  color: var(--text-main);
}

.icon-btn.edit:hover {
  color: var(--accent);
}

.icon-btn.delete:hover {
  color: var(--danger);
}
</style>