<script setup lang="ts">
import './assets/main.css';

import { computed } from 'vue';
import { useRoute } from 'vue-router';

import Main from './components/Main/Main.vue';
import Sidebar from './components/Sidebar/Sidebar.vue';
import { Entity } from './core/enums.ts';

const route = useRoute();

const activeEntity = computed(
  () => (route.meta.entity as Entity | undefined) ?? Entity.LAPS
);
</script>

<template>
  <div class="app-container">
    <Sidebar :entity="activeEntity" />
    <Main />
  </div>
</template>
