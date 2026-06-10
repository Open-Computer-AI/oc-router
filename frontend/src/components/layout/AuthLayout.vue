<template>
  <div class="flex min-h-screen items-center justify-center bg-bg p-4">
    <!-- Content Container -->
    <div class="relative z-10 w-full max-w-[420px]">
      <!-- Logo/Brand -->
      <div class="mb-8 text-center">
        <template v-if="settingsLoaded">
          <!-- OC Router Wordmark -->
          <div class="mb-4 inline-flex items-center justify-center">
            <img :src="siteLogo || '/oc-router-logo.png'" alt="OC Router" class="h-10 w-auto object-contain dark:invert dark:brightness-200" />
          </div>
          <!-- Mono subtitle -->
          <p class="font-mono text-[11px] uppercase tracking-widest text-text-dim">
            {{ siteSubtitle }}
          </p>
        </template>
      </div>

      <!-- Card Container -->
      <div class="rounded-2xl border border-border bg-bg-elevated p-8 shadow-modal">
        <slot />
      </div>

      <!-- Footer Links -->
      <div class="mt-6 text-center text-sm">
        <slot name="footer" />
      </div>

      <!-- Copyright -->
      <div class="mt-8 text-center text-xs text-text-dim">
        &copy; {{ currentYear }} {{ siteName }}. All rights reserved.
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useAppStore } from '@/stores'
import { sanitizeUrl } from '@/utils/url'

const appStore = useAppStore()

const siteName = computed(() => appStore.siteName || 'OC Router')
const siteLogo = computed(() => sanitizeUrl(appStore.siteLogo || '', { allowRelative: true, allowDataUrl: true }))
const siteSubtitle = computed(() => appStore.cachedPublicSettings?.site_subtitle || 'An Open Computer product')
const settingsLoaded = computed(() => appStore.publicSettingsLoaded)

const currentYear = computed(() => new Date().getFullYear())

onMounted(() => {
  appStore.fetchPublicSettings()
})
</script>
