<template>
  <div class="h-1 w-full overflow-hidden rounded-full bg-muted">
    <div
      class="h-full rounded-full transition-all duration-300"
      :class="barColorClass"
      :style="{ width: `${clampedPercent}%` }"
    ></div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  percent: number
  warnAt?: number
  dangerAt?: number
}

const props = withDefaults(defineProps<Props>(), {
  warnAt: 70,
  dangerAt: 90,
})

const clampedPercent = computed(() => Math.min(100, Math.max(0, props.percent)))

const barColorClass = computed(() => {
  if (clampedPercent.value >= props.dangerAt) return 'bg-danger'
  if (clampedPercent.value >= props.warnAt) return 'bg-warning'
  return 'bg-accent'
})
</script>
