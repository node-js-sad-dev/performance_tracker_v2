<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import GearTabs from '../components/GearTabs/GearTabs.vue';
import NewEntryButton from '../components/NewEntryButton/NewEntryButton.vue';
import Pagination from '../components/Pagination/Pagination.vue';
import Table from '../components/Table/Table.vue';
import TopBar from '../components/TopBar/TopBar.vue';
import { RESOURCE_CONFIG, buildResourcePath } from '../core/consts.ts';
import { deleteResource, getResourceList } from '../core/fetcher.ts';
import { Entity, type GearResourceKey, type ResourceKey, type SortOrder } from '../core/enums.ts';
import type { PlainRecord } from '../core/interfaces.ts';

const PAGE_SIZE = 10;

const route = useRoute();
const router = useRouter();

const rows = ref<PlainRecord[]>([]);
const totalCount = ref(0);
const loading = ref(false);
const error = ref<string | null>(null);
const searchDraft = ref('');

function readSingleQueryValue(value: unknown): string | undefined {
  if (typeof value === 'string') {
    return value;
  }

  if (Array.isArray(value) && typeof value[0] === 'string') {
    return value[0];
  }

  return undefined;
}

function readPositiveInt(value: unknown, fallback: number): number {
  const parsed = Number(readSingleQueryValue(value));
  return Number.isInteger(parsed) && parsed > 0 ? parsed : fallback;
}

function readSortOrder(value: unknown, fallback: SortOrder): SortOrder {
  const parsed = readSingleQueryValue(value);
  return parsed === 'asc' || parsed === 'desc' ? parsed : fallback;
}

const resourceKey = computed(() => route.meta.resourceKey as ResourceKey);
const config = computed(() => RESOURCE_CONFIG[resourceKey.value]);
const page = computed(() => readPositiveInt(route.query.page, 1));
const sortBy = computed(
  () => readSingleQueryValue(route.query.sortBy) ?? config.value.defaultSortBy
);
const sortOrder = computed(() =>
  readSortOrder(route.query.sortOrder, config.value.defaultSortOrder)
);
const activeSearch = computed(() => readSingleQueryValue(route.query.search) ?? '');
const isGearSection = computed(() => config.value.navEntity === Entity.GEAR);
const gearResourceKey = computed(() => resourceKey.value as GearResourceKey);
const totalPages = computed(() => Math.max(1, Math.ceil(totalCount.value / PAGE_SIZE)));
const subtitle = computed(() => {
  const label = totalCount.value === 1
    ? config.value.labelSingular.toLowerCase()
    : config.value.labelPlural.toLowerCase();

  return `${totalCount.value} ${label} available in the current dataset.`;
});

watch(
  activeSearch,
  (value) => {
    searchDraft.value = value;
  },
  { immediate: true }
);

function buildQueryPatch(patch: Record<string, string | number | undefined>) {
  const nextQuery: Record<string, string> = {};

  Object.entries(route.query).forEach(([key, value]) => {
    const normalized = readSingleQueryValue(value);
    if (normalized) {
      nextQuery[key] = normalized;
    }
  });

  Object.entries(patch).forEach(([key, value]) => {
    if (value === undefined || value === '') {
      delete nextQuery[key];
      return;
    }

    nextQuery[key] = String(value);
  });

  void router.push({ path: route.path, query: nextQuery });
}

async function loadList() {
  loading.value = true;
  error.value = null;

  try {
    const result = await getResourceList(resourceKey.value, {
      page: page.value,
      limit: PAGE_SIZE,
      sortBy: sortBy.value,
      sortOrder: sortOrder.value,
      search: activeSearch.value,
    });

    rows.value = result.items as unknown as PlainRecord[];
    totalCount.value = result.totalCount;
  } catch (loadError) {
    error.value = loadError instanceof Error ? loadError.message : 'Failed to load data.';
    rows.value = [];
    totalCount.value = 0;
  } finally {
    loading.value = false;
  }
}

watch([resourceKey, page, sortBy, sortOrder, activeSearch], () => {
  void loadList();
}, { immediate: true });

function getItemId(item: PlainRecord): number {
  return Number(item.id ?? 0);
}

function handleSort(columnSortBy: string) {
  if (sortBy.value === columnSortBy) {
    buildQueryPatch({
      sortOrder: sortOrder.value === 'asc' ? 'desc' : 'asc',
      page: 1,
    });
    return;
  }

  buildQueryPatch({
    sortBy: columnSortBy,
    sortOrder: 'asc',
    page: 1,
  });
}

function handleSearch() {
  buildQueryPatch({
    search: searchDraft.value.trim() || undefined,
    page: 1,
  });
}

function handlePageChange(nextPage: number) {
  buildQueryPatch({ page: nextPage });
}

function handleView(item: PlainRecord) {
  void router.push(buildResourcePath(resourceKey.value, 'detail', getItemId(item)));
}

function handleEdit(item: PlainRecord) {
  void router.push(buildResourcePath(resourceKey.value, 'edit', getItemId(item)));
}

async function handleDelete(item: PlainRecord) {
  const label = item.name ? String(item.name) : `${config.value.labelSingular} #${getItemId(item)}`;
  const shouldDelete = window.confirm(`Delete ${label}? This action cannot be undone.`);
  if (!shouldDelete) {
    return;
  }

  try {
    await deleteResource(resourceKey.value, getItemId(item));

    if (rows.value.length === 1 && page.value > 1) {
      buildQueryPatch({ page: page.value - 1 });
      return;
    }

    await loadList();
  } catch (deleteError) {
    error.value = deleteError instanceof Error ? deleteError.message : 'Failed to delete record.';
  }
}
</script>

<template>
  <section>
    <TopBar
      :eyebrow="isGearSection ? 'Gear section' : 'Telemetry data'"
      :title="config.headerTitle"
      :subtitle="'Browse, sort, inspect, and maintain records from the finished backend API.'">
      <template #actions>
        <NewEntryButton
          :to="buildResourcePath(resourceKey, 'create')"
          :label="config.createActionLabel" />
      </template>
    </TopBar>

    <GearTabs v-if="isGearSection" :resource-key="gearResourceKey" />

    <div v-if="error" class="status-banner error">
      {{ error }}
    </div>

    <div class="view-card">
      <div class="section-toolbar">
        <p class="page-meta">{{ subtitle }}</p>

        <form v-if="config.searchKey" class="search-form" @submit.prevent="handleSearch">
          <input
            v-model="searchDraft"
            class="search-input"
            :placeholder="config.searchPlaceholder"
            type="search" />
          <NewEntryButton label="Search" variant="secondary" type="submit" />
        </form>
      </div>

      <Table
        :columns="config.tableColumns"
        :rows="rows"
        :sort-by="sortBy"
        :sort-order="sortOrder"
        :loading="loading"
        :empty-label="config.emptyStateLabel"
        @sort="handleSort"
        @view="handleView"
        @edit="handleEdit"
        @delete="handleDelete" />

      <Pagination
        :page="page"
        :total-pages="totalPages"
        @page-change="handlePageChange" />
    </div>
  </section>
</template>

