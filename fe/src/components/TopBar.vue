<script setup lang="ts">
import NewEntryButton from "./NewEntryButton.vue";
</script>

<template>
  <header class="top-bar">
    <h1>Laps Telemetry</h1>
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