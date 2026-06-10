<template>
  <span
    class="inline-flex items-center gap-1.5 rounded-full border border-border px-2.5 py-0.5 text-xs font-medium text-text-muted"
  >
    <span
      class="inline-block h-2 w-2 rounded-full"
      :style="{ backgroundColor: dotColor }"
    ></span>
    <slot>{{ label }}</slot>
  </span>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  platform?: string
  label?: string
  color?: string
}

const props = withDefaults(defineProps<Props>(), {
  platform: '',
  label: '',
  color: '',
})

const platformColors: Record<string, string> = {
  openai: '#10a37f',
  anthropic: '#d97757',
  google: '#4285f4',
  azure: '#0078d4',
  aws: '#ff9900',
  mistral: '#ff7000',
  cohere: '#39594d',
  deepseek: '#4d6bfe',
}

const dotColor = computed(() => {
  if (props.color) return props.color
  return platformColors[props.platform?.toLowerCase()] || '#9b9b9b'
})
</script>
