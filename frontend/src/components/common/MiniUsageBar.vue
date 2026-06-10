<template>
  <div class="flex items-center gap-2">
    <span v-if="label" class="text-xs text-text-muted whitespace-nowrap">{{ label }}</span>
    <div class="flex-1">
      <ProgressBar :percent="percent" :warn-at="warnAt" :danger-at="dangerAt" />
    </div>
    <span class="font-mono text-xs text-text-dim whitespace-nowrap">{{ displayValue }}</span>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import ProgressBar from './ProgressBar.vue'

interface Props {
  label?: string
  value?: number | string
  max?: number | string
  percent?: number
  warnAt?: number
  dangerAt?: number
}

const props = withDefaults(defineProps<Props>(), {
  label: '',
  value: 0,
  max: 0,
  percent: 0,
  warnAt: 70,
  dangerAt: 90,
})

const displayValue = computed(() => {
  if (props.max) {
    return `${props.value}/${props.max}`
  }
  return `${Math.round(typeof props.percent === 'number' ? props.percent : 0)}%`
})
</script>
