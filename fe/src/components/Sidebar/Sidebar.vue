<script setup lang="ts">
import { CURRENT_VIEW } from '../../core/consts.ts';

const currentView = defineModel<keyof typeof CURRENT_VIEW>({ required: true });

const menuItems = [
  { id: 'laps', label: 'Laps', icon: '⏱', currentView: CURRENT_VIEW.LAPS },
  { id: 'cars', label: 'Cars', icon: '🏎', currentView: CURRENT_VIEW.CARS },
  {
    id: 'tracks',
    label: 'Tracks',
    icon: '🏁',
    currentView: CURRENT_VIEW.TRACKS,
  },
  { id: 'games', label: 'Games', icon: '🎮', currentView: CURRENT_VIEW.GAMES },
  { id: 'gear', label: 'Gear', icon: '⚙️', currentView: CURRENT_VIEW.GEAR },
];
</script>

<template>
  <aside class="sidebar">
    <div class="brand">
      <h2>Sim<span class="accent">Track</span></h2>
    </div>

    <nav class="nav-menu">
      <a
        v-for="item in menuItems"
        :key="item.id"
        :href="`/${item.id}`"
        class="nav-item"
        :class="{ active: currentView === item.id }"
        @click.prevent="currentView = item.currentView">
        <span class="icon">{{ item.icon }}</span> {{ item.label }}
      </a>
    </nav>

    <div class="user-profile">
      <div class="avatar">D</div>
      <span>Later this block will display current user</span>
    </div>
  </aside>
</template>

<style>
.sidebar {
  background-color: var(--bg-panel);
  border-right: 1px solid var(--border);
  display: flex;
  flex-direction: column;
  padding: 1.5rem;
}

.brand h2 {
  font-size: 1.5rem;
  margin-bottom: 3rem;
  letter-spacing: -0.5px;
}

.brand .accent {
  color: var(--accent);
}

.nav-menu {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.nav-item {
  display: flex;
  align-items: center;
  padding: 0.75rem 1rem;
  text-decoration: none;
  color: var(--text-muted);
  border-radius: 8px;
  transition: all 0.2s ease;
  font-weight: 600;
}

.nav-item:hover {
  background-color: var(--bg-hover);
  color: var(--text-main);
}

.nav-item.active {
  background-color: rgba(0, 229, 255, 0.1);
  color: var(--accent);
  border-left: 3px solid var(--accent);
}

.nav-item .icon {
  margin-right: 12px;
  width: 20px;
  text-align: center;
}

.user-profile {
  border-top: 1px solid var(--border);
  padding-top: 1rem;
  display: flex;
  align-items: center;
  gap: 10px;
}

.avatar {
  width: 32px;
  height: 32px;
  background: var(--accent);
  color: var(--bg-dark);
  border-radius: 50%;
  display: grid;
  place-items: center;
  font-weight: bold;
}
</style>
