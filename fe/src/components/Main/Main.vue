<script setup lang="ts">
import TopBar from '../TopBar/TopBar.vue';
import Table from '../Table/Table.vue';
import Pagination from '../Pagination/Pagination.vue';
import type { Props, State } from './interfaces.ts';

import { reactive } from 'vue';

const props = defineProps<Props>();

const state = reactive<State>({
  laps: [],
});
</script>

<template>
  <main class="main-content" :entity="props.entity">
    <TopBar :entity="props.entity"></TopBar>
    <section class="content-view" id="view-list">
      <div class="data-table-container">
        <Table :laps="state.laps"> </Table>
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
