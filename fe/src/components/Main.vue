<script setup lang="ts">
import TopBar from './TopBar.vue'
import Table from './Table.vue'
import Pagination from './Pagination.vue'
</script>

<template>
  <main class="main-content">
    <TopBar></TopBar>
    <section class="content-view" id="view-list">
      <div class="data-table-container">
        <Table>
        </Table>
      </div>
      <Pagination></Pagination>
    </section>
  </main>
</template>

<style>
.main-content {
  padding: 2rem;
  overflow-y: auto;
  background-color: var(--bg-dark);
}

.data-table-container {
  background: var(--bg-panel);
  border-radius: 12px;
  border: 1px solid var(--border);
  overflow: hidden;
  margin-bottom: 1rem;
}
</style>