<template>
  <AuthLayout>
    <div class="space-y-6">
      <!-- Title -->
      <div class="text-center">
        <h2 class="font-display text-2xl font-semibold text-text">
          {{ t('auth.welcomeBack') }}
        </h2>
        <p class="mt-2 text-sm text-text-muted">
          Sign in with your GitHub account
        </p>
      </div>

      <!-- GitHub Sign-In Button -->
      <div class="space-y-4">
        <EmailOAuthButtons
          :disabled="!publicSettingsLoaded"
          :github-enabled="true"
          :google-enabled="false"
          :show-divider="false"
        />
      </div>

      <!-- Error message from OAuth callback -->
      <div v-if="oauthError" class="rounded-lg border border-red-200 bg-red-50 p-4 text-sm text-red-700 dark:border-red-800 dark:bg-red-900/20 dark:text-red-400">
        {{ oauthError }}
      </div>
    </div>
  </AuthLayout>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { AuthLayout } from '@/components/layout'
import EmailOAuthButtons from '@/components/auth/EmailOAuthButtons.vue'
import { getPublicSettings } from '@/api/auth'

const { t } = useI18n()
const route = useRoute()

const publicSettingsLoaded = ref<boolean>(false)
const oauthError = ref<string>('')

onMounted(async () => {
  // Check for OAuth error in query params (redirected back from failed OAuth)
  const errorMsg = route.query.error_message as string
  if (errorMsg) {
    oauthError.value = errorMsg
  }

  try {
    await getPublicSettings()
  } catch (error) {
    console.error('Failed to load public settings:', error)
  } finally {
    publicSettingsLoaded.value = true
  }
})
</script>
