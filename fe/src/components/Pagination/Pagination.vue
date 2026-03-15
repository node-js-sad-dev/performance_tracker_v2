<script setup lang="ts">
const props = defineProps<{
  page: number;
  totalPages: number;
}>();

const emit = defineEmits<{
  (event: 'page-change', page: number): void;
}>();

function buildPages(totalPages: number, currentPage: number) {
  if (totalPages <= 5) {
    return Array.from({ length: totalPages }, (_, index) => index + 1);
  }

  const pages = new Set<number>([1, totalPages, currentPage, currentPage - 1, currentPage + 1]);

  return Array.from(pages)
    .filter((value) => value >= 1 && value <= totalPages)
    .sort((left, right) => left - right);
}

function changePage(page: number) {
  if (page < 1 || page > props.totalPages || page === props.page) {
    return;
  }

  emit('page-change', page);
}
</script>

<template>
  <div class="pagination" aria-label="Pagination">
    <button class="btn-page" :disabled="props.page <= 1" @click="changePage(props.page - 1)">
      Prev
    </button>

    <button
      v-for="pageNumber in buildPages(props.totalPages, props.page)"
      :key="pageNumber"
      class="btn-page"
      :class="{ active: pageNumber === props.page }"
      @click="changePage(pageNumber)">
      {{ pageNumber }}
    </button>

    <button
      class="btn-page"
      :disabled="props.page >= props.totalPages"
      @click="changePage(props.page + 1)">
      Next
    </button>
  </div>
</template>

<style scoped>
.pagination {
  display: flex;
  gap: 0.5rem;
  justify-content: flex-end;
  flex-wrap: wrap;
}

.btn-page {
  background: var(--bg-panel);
  border: 1px solid var(--border);
  color: var(--text-main);
  min-width: 42px;
  height: 38px;
  padding: 0 0.85rem;
  border-radius: 10px;
  cursor: pointer;
}

.btn-page.active {
  background: var(--accent);
  border-color: var(--accent);
  color: #101010;
  font-weight: 700;
}

.btn-page:disabled {
  opacity: 0.45;
  cursor: not-allowed;
}
</style>
