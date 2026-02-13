<script setup lang="ts">
import './assets/main.css';

import Sidebar from './components/Sidebar/Sidebar.vue';
import Main from './components/Main/Main.vue';

import { reactive } from 'vue';
import { Entity } from './core/enums.ts';

const state = reactive<{
  entity: Entity;
}>({
  entity: Entity.LAPS,
});
</script>

<template>
  <div class="app-container">
    <Sidebar v-model="state.entity"></Sidebar>
    <Main :entity="state.entity"></Main>
  </div>
</template>

<style>
.app-container {
  display: grid;
  grid-template-columns: 250px 1fr; /* Sidebar | Content */
  height: 100vh;
}
</style>
