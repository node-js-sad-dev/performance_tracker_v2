<script setup lang="ts">
import NewEntryButton from '../NewEntryButton/NewEntryButton.vue';
import { CURRENT_VIEW } from '../../core/consts.ts';
import { computed } from 'vue';
import type { Props } from './interfaces.ts';

const props = defineProps<Props>();

const TITLES: Record<keyof typeof CURRENT_VIEW, string> = {
  [CURRENT_VIEW.LAPS]: 'Laps Telemetry',
  [CURRENT_VIEW.CARS]: 'Cars Telemetry',
  [CURRENT_VIEW.TRACKS]: 'Tracks Telemetry',
  [CURRENT_VIEW.GAMES]: 'Games Telemetry',
  [CURRENT_VIEW.GEAR]: 'Gear Telemetry',
};

const title = computed(() => TITLES[props.currentView] ?? 'Telemetry');
</script>

<template>
  <header class="top-bar">
    <h1>{{ title }}</h1>
    <div class="actions">
      <NewEntryButton></NewEntryButton>
    </div>
  </header>
</template>

<style>
.top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}
</style>
