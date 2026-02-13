<script setup lang="ts">
import NewEntryButton from '../NewEntryButton/NewEntryButton.vue';
import { computed } from 'vue';
import type { Props } from './interfaces.ts';
import { EntityConfig } from '../../core/consts.ts';

const props = defineProps<Props>();

const title = computed(
  () => EntityConfig[props.entity].headerTitle ?? 'Telemetry'
);
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
