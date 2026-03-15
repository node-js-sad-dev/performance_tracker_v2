<script setup lang="ts">
import type { Props } from './interfaces.ts';

const props = defineProps<Props>();
</script>

<template>
  <header class="top-bar">
    <div class="title-stack">
      <p v-if="props.eyebrow" class="eyebrow">{{ props.eyebrow }}</p>
      <h1>{{ props.title }}</h1>
      <p v-if="props.subtitle" class="subtitle">{{ props.subtitle }}</p>
    </div>

    <div v-if="$slots.actions" class="actions">
      <slot name="actions" />
    </div>
  </header>
</template>
