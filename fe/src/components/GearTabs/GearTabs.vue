<script setup lang="ts">
import { RouterLink } from 'vue-router';

import { GEAR_TABS, buildResourcePath } from '../../core/consts.ts';
import type { GearResourceKey } from '../../core/enums.ts';

const props = defineProps<{
  resourceKey: GearResourceKey;
}>();
</script>

<template>
  <nav class="gear-tabs" aria-label="Gear sections">
    <RouterLink
      v-for="tab in GEAR_TABS"
      :key="tab.key"
      :to="buildResourcePath(tab.key, 'list')"
      class="gear-tab"
      :class="{ active: props.resourceKey === tab.key }">
      {{ tab.label }}
    </RouterLink>
  </nav>
</template>

<style scoped>
.gear-tabs {
  display: flex;
  gap: 0.75rem;
  flex-wrap: wrap;
  margin-bottom: 1rem;
}

.gear-tab {
  padding: 0.7rem 1rem;
  border: 1px solid var(--border);
  border-radius: 10px;
  color: var(--text-muted);
  background: rgba(255, 255, 255, 0.02);
}

.gear-tab.active {
  color: var(--accent);
  border-color: rgba(0, 229, 255, 0.25);
  background: rgba(0, 229, 255, 0.08);
}
</style>
