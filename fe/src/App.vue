<script setup lang="ts">
import './assets/main.css'

import Sidebar from './components/Sidebar.vue'
import Main from './components/Main.vue'
</script>

<template>
  <div class="app-container">
    <Sidebar></Sidebar>
    <Main></Main>
  </div>
</template>

<style>
.app-container {
  display: grid;
  grid-template-columns: 250px 1fr; /* Sidebar | Content */
  height: 100vh;
}
</style>