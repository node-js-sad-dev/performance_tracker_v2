<script setup lang="ts">
import { computed } from 'vue';
import { RouterLink, type RouteLocationRaw } from 'vue-router';

const props = withDefaults(
  defineProps<{
    label: string;
    to?: RouteLocationRaw;
    variant?: 'primary' | 'secondary' | 'danger';
    type?: 'button' | 'submit';
    disabled?: boolean;
  }>(),
  {
    variant: 'primary',
    type: 'button',
    disabled: false,
  }
);

const emit = defineEmits<{
  (event: 'click'): void;
}>();

const classes = computed(() => ['btn', `btn-${props.variant}`]);

function handleClick() {
  emit('click');
}
</script>

<template>
  <RouterLink v-if="props.to" :to="props.to" :class="classes">
    {{ props.label }}
  </RouterLink>

  <button
    v-else
    :type="props.type"
    :class="classes"
    :disabled="props.disabled"
    @click="handleClick">
    {{ props.label }}
  </button>
</template>
