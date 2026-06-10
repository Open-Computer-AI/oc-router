<template>
  <div class="flex items-center justify-between gap-4 border-b border-border py-4 last:border-b-0">
    <div class="min-w-0 flex-1">
      <div class="text-sm font-medium text-text">{{ label }}</div>
      <div v-if="hint" class="mt-0.5 text-xs text-text-dim">{{ hint }}</div>
    </div>
    <div class="flex-shrink-0">
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  label: string
  hint?: string
}>()
</script>
