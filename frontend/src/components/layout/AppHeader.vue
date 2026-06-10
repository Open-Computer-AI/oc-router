<template>
  <!-- Mobile-only top bar for menu toggle and key controls -->
  <header class="flex items-center justify-between px-2 py-2 lg:hidden">
    <button
      @click="toggleMobileSidebar"
      class="btn btn-ghost btn-icon"
      aria-label="Toggle Menu"
    >
      <Icon name="menu" size="md" />
    </button>

    <div class="flex items-center gap-2">
      <AnnouncementBell v-if="user" />
      <LocaleSwitcher />
    </div>
  </header>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAppStore, useAuthStore } from '@/stores'
import LocaleSwitcher from '@/components/common/LocaleSwitcher.vue'
import AnnouncementBell from '@/components/common/AnnouncementBell.vue'
import Icon from '@/components/icons/Icon.vue'

const appStore = useAppStore()
const authStore = useAuthStore()

const user = computed(() => authStore.user)

function toggleMobileSidebar() {
  appStore.toggleMobileSidebar()
}
</script>
