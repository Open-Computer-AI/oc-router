<template>
  <label class="text-sm font-medium text-text" :for="htmlFor">
    <slot />
  </label>
</template>

<script setup lang="ts">
defineProps<{
  htmlFor?: string
}>()
</script>
