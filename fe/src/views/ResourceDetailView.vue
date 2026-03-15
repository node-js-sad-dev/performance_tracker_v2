<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import GearTabs from '../components/GearTabs/GearTabs.vue';
import NewEntryButton from '../components/NewEntryButton/NewEntryButton.vue';
import TopBar from '../components/TopBar/TopBar.vue';
import {
  RESOURCE_CONFIG,
  buildResourcePath,
  getDetailEntries,
  getImageFromRecord,
  getResourceHeading,
} from '../core/consts.ts';
import { deleteResource, getResourceById, getSelectOptions } from '../core/fetcher.ts';
import { Entity, type GearResourceKey, type ResourceKey } from '../core/enums.ts';
import type { LookupCollection, PlainRecord } from '../core/interfaces.ts';

const route = useRoute();
const router = useRouter();

const loading = ref(false);
const deleting = ref(false);
const error = ref<string | null>(null);
const record = ref<PlainRecord | null>(null);
const lookups = ref<LookupCollection>({});

const resourceKey = computed(() => route.meta.resourceKey as ResourceKey);
const config = computed(() => RESOURCE_CONFIG[resourceKey.value]);
const id = computed(() => Number(route.params.id ?? 0));
const isGearSection = computed(() => config.value.navEntity === Entity.GEAR);
const gearResourceKey = computed(() => resourceKey.value as GearResourceKey);
const existingImage = computed(() => (record.value ? getImageFromRecord(record.value) : null));
const title = computed(() =>
  record.value ? getResourceHeading(resourceKey.value, record.value) : config.value.labelSingular
);
const detailEntries = computed(() =>
  record.value ? getDetailEntries(resourceKey.value, record.value, lookups.value) : []
);

async function loadLookups() {
  if (resourceKey.value !== 'lap') {
    lookups.value = {};
    return;
  }

  const lookupKeys: ResourceKey[] = ['car', 'track', 'game', 'wheel', 'pedals', 'cockpit'];
  const entries = await Promise.all(
    lookupKeys.map(async (lookupKey) => [lookupKey, await getSelectOptions(lookupKey)] as const)
  );

  lookups.value = Object.fromEntries(entries) as LookupCollection;
}

async function loadView() {
  loading.value = true;
  error.value = null;

  try {
    if (!Number.isInteger(id.value) || id.value <= 0) {
      throw new Error('Invalid record identifier.');
    }

    await loadLookups();
    const entity = await getResourceById(resourceKey.value, id.value);
    record.value = entity as unknown as PlainRecord;
  } catch (loadError) {
    error.value = loadError instanceof Error ? loadError.message : 'Failed to load record.';
    record.value = null;
  } finally {
    loading.value = false;
  }
}

watch([resourceKey, id], () => {
  void loadView();
}, { immediate: true });

async function handleDelete() {
  const shouldDelete = window.confirm(
    `Delete ${config.value.labelSingular.toLowerCase()} #${id.value}? This action cannot be undone.`
  );

  if (!shouldDelete) {
    return;
  }

  deleting.value = true;
  error.value = null;

  try {
    await deleteResource(resourceKey.value, id.value);
    await router.push(buildResourcePath(resourceKey.value, 'list'));
  } catch (deleteError) {
    error.value = deleteError instanceof Error ? deleteError.message : 'Failed to delete record.';
  } finally {
    deleting.value = false;
  }
}
</script>

<template>
  <section>
    <TopBar
      eyebrow="Detail view"
      :title="title"
      :subtitle="'Read-only view of the selected backend record with quick navigation for follow-up actions.'">
      <template #actions>
        <NewEntryButton :to="buildResourcePath(resourceKey, 'list')" label="Back to list" variant="secondary" />
        <NewEntryButton :to="buildResourcePath(resourceKey, 'edit', id)" label="Edit" variant="secondary" />
        <NewEntryButton :label="deleting ? 'Deleting...' : 'Delete'" variant="danger" :disabled="deleting" @click="handleDelete" />
      </template>
    </TopBar>

    <GearTabs v-if="isGearSection" :resource-key="gearResourceKey" />

    <div v-if="error" class="status-banner error">
      {{ error }}
    </div>

    <div v-if="loading" class="detail-card">
      <div class="loading-state">Loading details...</div>
    </div>

    <template v-else-if="record">
      <section class="hero-card" :style="!existingImage ? { gridTemplateColumns: '1fr' } : undefined">
        <div class="hero-copy">
          <p class="hero-kicker">{{ config.labelSingular }} details</p>
          <h2 class="hero-title">{{ title }}</h2>
          <p class="subtitle">Record ID #{{ record.id }}</p>
        </div>
        <img
          v-if="existingImage"
          :src="existingImage"
          :alt="`${config.labelSingular} preview`"
          class="hero-image" />
      </section>

      <section class="detail-card">
        <div class="detail-grid">
          <article v-for="entry in detailEntries" :key="entry.label" class="detail-item">
            <p class="detail-label">{{ entry.label }}</p>
            <p class="detail-value" :class="`tone-${entry.tone}`">{{ entry.value }}</p>
          </article>
        </div>
      </section>
    </template>
  </section>
</template>

