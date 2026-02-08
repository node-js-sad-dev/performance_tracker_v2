<script setup lang="ts">

</script>

<template>
  <button class="btn btn-primary">+ Log New Lap</button>
</template>

<style>
.btn {
  padding: 0.6rem 1.2rem;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 600;
  font-size: 0.9rem;
  transition: opacity 0.2s;
}

.btn-primary {
  background: var(--accent);
  color: var(--bg-dark);
}

.btn-primary:hover {
  background: var(--accent-hover);
}

.btn-secondary {
  background: var(--bg-hover);
  color: var(--text-main);
}
</style>