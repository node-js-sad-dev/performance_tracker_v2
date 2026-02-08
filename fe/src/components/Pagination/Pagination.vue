<script setup lang="ts">

</script>

<template>
  <div class="pagination">
    <button class="btn-page" disabled>&lt;</button>
    <button class="btn-page active">1</button>
    <button class="btn-page">2</button>
    <button class="btn-page">3</button>
    <button class="btn-page">&gt;</button>
  </div>
</template>

<style>
.pagination {
  display: flex;
  justify-content: flex-end;
  gap: 5px;
}

.btn-page {
  background: var(--bg-panel);
  border: 1px solid var(--border);
  color: var(--text-main);
  width: 32px;
  height: 32px;
  border-radius: 4px;
  cursor: pointer;
}

.btn-page.active {
  background: var(--accent);
  color: var(--bg-dark);
  border-color: var(--accent);
}

.btn-page:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
</style>