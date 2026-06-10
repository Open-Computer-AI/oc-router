<template>
  <!-- Row 1: Core Stats -->
  <div class="grid grid-cols-2 gap-4 lg:grid-cols-4">
    <!-- Balance -->
    <div v-if="!isSimple" class="card p-4">
      <div class="flex items-center gap-3">
        <div class="rounded-lg bg-success-muted p-2">
          <svg class="h-5 w-5 text-success" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.25 18.75a60.07 60.07 0 0115.797 2.101c.727.198 1.453-.342 1.453-1.096V18.75M3.75 4.5v.75A.75.75 0 013 6h-.75m0 0v-.375c0-.621.504-1.125 1.125-1.125H20.25M2.25 6v9m18-10.5v.75c0 .414.336.75.75.75h.75m-1.5-1.5h.375c.621 0 1.125.504 1.125 1.125v9.75c0 .621-.504 1.125-1.125 1.125h-.375m1.5-1.5H21a.75.75 0 00-.75.75v.75m0 0H3.75m0 0h-.375a1.125 1.125 0 01-1.125-1.125V15m1.5 1.5v-.75A.75.75 0 003 15h-.75M15 10.5a3 3 0 11-6 0 3 3 0 016 0zm3 0h.008v.008H18V10.5zm-12 0h.008v.008H6V10.5z" />
          </svg>
        </div>
        <div>
          <p class="text-xs font-medium text-text-muted">{{ t('dashboard.balance') }}</p>
          <p class="text-xl font-bold text-success">${{ formatBalance(balance) }}</p>
          <p class="text-xs text-text-muted">{{ t('common.available') }}</p>
        </div>
      </div>
    </div>

    <!-- API Keys -->
    <div class="card p-4">
      <div class="flex items-center gap-3">
        <div class="rounded-lg bg-info-muted p-2">
          <Icon name="key" size="md" class="text-info" :stroke-width="2" />
        </div>
        <div>
          <p class="text-xs font-medium text-text-muted">{{ t('dashboard.apiKeys') }}</p>
          <p class="text-xl font-bold text-text">{{ stats?.total_api_keys || 0 }}</p>
          <p class="text-xs text-success">{{ stats?.active_api_keys || 0 }} {{ t('common.active') }}</p>
        </div>
      </div>
    </div>

    <!-- Today Requests -->
    <div class="card p-4">
      <div class="flex items-center gap-3">
        <div class="rounded-lg bg-success-muted p-2">
          <Icon name="chart" size="md" class="text-success" :stroke-width="2" />
        </div>
        <div>
          <p class="text-xs font-medium text-text-muted">{{ t('dashboard.todayRequests') }}</p>
          <p class="text-xl font-bold text-text">{{ stats?.today_requests || 0 }}</p>
          <p class="text-xs text-text-muted">{{ t('common.total') }}: {{ formatNumber(stats?.total_requests || 0) }}</p>
        </div>
      </div>
    </div>

    <!-- Today Cost -->
    <div class="card p-4">
      <div class="flex items-center gap-3">
        <div class="rounded-lg bg-[rgba(126,87,194,0.08)] p-2 dark:bg-[rgba(149,117,205,0.12)]">
          <Icon name="dollar" size="md" class="text-[#7e57c2] dark:text-[#b39ddb]" :stroke-width="2" />
        </div>
        <div>
          <p class="text-xs font-medium text-text-muted">{{ t('dashboard.todayCost') }}</p>
          <p class="text-xl font-bold text-text">
            <span class="text-[#7e57c2] dark:text-[#b39ddb]" :title="t('dashboard.actual')">${{ formatCost(stats?.today_actual_cost || 0) }}</span>
            <span class="text-sm font-normal text-text-dim" :title="t('dashboard.standard')"> / ${{ formatCost(stats?.today_cost || 0) }}</span>
          </p>
          <p class="text-xs">
            <span class="text-text-muted">{{ t('common.total') }}: </span>
            <span class="text-[#7e57c2] dark:text-[#b39ddb]" :title="t('dashboard.actual')">${{ formatCost(stats?.total_actual_cost || 0) }}</span>
            <span class="text-text-dim" :title="t('dashboard.standard')"> / ${{ formatCost(stats?.total_cost || 0) }}</span>
          </p>
        </div>
      </div>
    </div>
  </div>

  <!-- Row 2: Token Stats -->
  <div class="grid grid-cols-2 gap-4 lg:grid-cols-4">
    <!-- Today Tokens -->
    <div class="card p-4">
      <div class="flex items-center gap-3">
        <div class="rounded-lg bg-warning-muted p-2">
          <Icon name="cube" size="md" class="text-warning" :stroke-width="2" />
        </div>
        <div>
          <p class="text-xs font-medium text-text-muted">{{ t('dashboard.todayTokens') }}</p>
          <p class="text-xl font-bold text-text">{{ formatTokens(stats?.today_tokens || 0) }}</p>
          <p class="text-xs text-text-muted">{{ t('dashboard.input') }}: {{ formatTokens(stats?.today_input_tokens || 0) }} / {{ t('dashboard.output') }}: {{ formatTokens(stats?.today_output_tokens || 0) }}</p>
        </div>
      </div>
    </div>

    <!-- Total Tokens -->
    <div class="card p-4">
      <div class="flex items-center gap-3">
        <div class="rounded-lg bg-[rgba(126,87,194,0.08)] p-2 dark:bg-[rgba(149,117,205,0.12)]">
          <Icon name="database" size="md" class="text-[#7e57c2] dark:text-[#b39ddb]" :stroke-width="2" />
        </div>
        <div>
          <p class="text-xs font-medium text-text-muted">{{ t('dashboard.totalTokens') }}</p>
          <p class="text-xl font-bold text-text">{{ formatTokens(stats?.total_tokens || 0) }}</p>
          <p class="text-xs text-text-muted">{{ t('dashboard.input') }}: {{ formatTokens(stats?.total_input_tokens || 0) }} / {{ t('dashboard.output') }}: {{ formatTokens(stats?.total_output_tokens || 0) }}</p>
        </div>
      </div>
    </div>

    <!-- Performance (RPM/TPM) -->
    <div class="card p-4">
      <div class="flex items-center gap-3">
        <div class="rounded-lg bg-[rgba(126,87,194,0.08)] p-2 dark:bg-[rgba(149,117,205,0.12)]">
          <Icon name="bolt" size="md" class="text-[#7e57c2] dark:text-[#b39ddb]" :stroke-width="2" />
        </div>
        <div class="flex-1">
          <p class="text-xs font-medium text-text-muted">{{ t('dashboard.performance') }}</p>
          <div class="flex items-baseline gap-2">
            <p class="text-xl font-bold text-text">{{ formatTokens(stats?.rpm || 0) }}</p>
            <span class="text-xs text-text-muted">RPM</span>
          </div>
          <div class="flex items-baseline gap-2">
            <p class="text-sm font-semibold text-[#7e57c2] dark:text-[#b39ddb]">{{ formatTokens(stats?.tpm || 0) }}</p>
            <span class="text-xs text-text-muted">TPM</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Avg Response Time -->
    <div class="card p-4">
      <div class="flex items-center gap-3">
        <div class="rounded-lg bg-danger-muted p-2">
          <Icon name="clock" size="md" class="text-danger" :stroke-width="2" />
        </div>
        <div>
          <p class="text-xs font-medium text-text-muted">{{ t('dashboard.avgResponse') }}</p>
          <p class="text-xl font-bold text-text">{{ formatDuration(stats?.average_duration_ms || 0) }}</p>
          <p class="text-xs text-text-muted">{{ t('dashboard.averageTime') }}</p>
        </div>
      </div>
    </div>
  </div>

  <!-- Row 3: Per-platform breakdown -->
  <div v-if="!isSimple && platformCards.length > 0" class="card p-4">
    <div class="mb-3 flex items-center justify-between">
      <h3 class="text-sm font-semibold text-text">{{ t('dashboard.platformBreakdown') }}</h3>
      <span class="text-xs text-text-muted">
        {{ t('dashboard.platformCount', { count: sortedPlatforms.length }) }}
      </span>
    </div>
    <div class="grid grid-cols-1 gap-3 sm:grid-cols-2 lg:grid-cols-4">
      <div
        v-for="item in platformCards"
        :key="item.platform"
        :class="[
          'rounded-lg border p-3',
          item.isOther
            ? 'border-dashed border-border bg-muted'
            : 'border-border'
        ]"
      >
        <div class="flex items-center justify-between">
          <span class="text-sm font-semibold text-text">
            {{ item.isOther ? t('dashboard.platformOther') : platformLabel(item.platform) }}
          </span>
          <span class="font-mono text-sm text-[#7e57c2] dark:text-[#b39ddb]" :title="t('dashboard.actual')">
            ${{ formatCost(item.total_actual_cost) }}
          </span>
        </div>
        <div class="mt-2 space-y-1 text-xs">
          <div class="flex items-center justify-between">
            <span class="text-text-muted">{{ t('dashboard.todayCost') }}</span>
            <span class="font-mono text-text">${{ formatCost(item.today_actual_cost) }}</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-text-muted">{{ t('dashboard.requests') }}</span>
            <span class="font-mono text-text-muted">
              {{ item.total_requests > 0 ? formatNumber(item.total_requests) : '-' }}
            </span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-text-muted">{{ t('dashboard.tokens') }}</span>
            <span class="font-mono text-text-muted">
              {{ item.total_tokens > 0 ? formatTokens(item.total_tokens) : '-' }}
            </span>
          </div>
        </div>

        <!-- Quota section: only shown when quota config exists, not __other__, and at least one window has a limit -->
        <div v-if="hasAnyLimit(item.quota) && !item.isOther" class="mt-3 space-y-1.5 border-t border-border pt-2">
          <p class="text-[10px] uppercase tracking-wide text-text-dim">
            {{ t('dashboard.platformQuota.title') }}
          </p>
          <template v-for="w in (['daily', 'weekly', 'monthly'] as const)" :key="w">
            <div v-if="quotaVal(item.quota, `${w}_limit_usd`) != null" class="space-y-0.5">
              <!-- limit=0: fully disabled -->
              <template v-if="(quotaVal(item.quota, `${w}_limit_usd`) as number) === 0">
                <div class="flex items-center justify-between text-xs">
                  <span class="text-text-muted">{{ t(`dashboard.platformQuota.${w}`) }}</span>
                  <span class="font-mono text-danger">{{ t('dashboard.platformQuota.disabled') }}</span>
                </div>
                <div class="h-1.5 w-full overflow-hidden rounded-full bg-muted">
                  <div class="h-full w-full rounded-full bg-danger" />
                </div>
              </template>
              <!-- limit>0: normal usage progress bar -->
              <template v-else>
                <div class="flex items-center justify-between text-xs">
                  <span class="text-text-muted">{{ t(`dashboard.platformQuota.${w}`) }}</span>
                  <span class="font-mono text-text">
                    ${{ formatUsd((quotaVal(item.quota, `${w}_usage_usd`) as number) ?? 0) }} / ${{ formatUsd(quotaVal(item.quota, `${w}_limit_usd`) as number) }}
                  </span>
                </div>
                <div class="h-1.5 w-full overflow-hidden rounded-full bg-muted">
                  <div
                    class="h-full rounded-full transition-all"
                    :class="quotaBarClass(calcPercent((quotaVal(item.quota, `${w}_usage_usd`) as number) ?? 0, quotaVal(item.quota, `${w}_limit_usd`) as number))"
                    :style="{ width: calcPercent((quotaVal(item.quota, `${w}_usage_usd`) as number) ?? 0, quotaVal(item.quota, `${w}_limit_usd`) as number) + '%' }"
                  />
                </div>
                <p v-if="quotaVal(item.quota, `${w}_window_resets_at`)" class="text-[10px] text-text-dim">
                  {{ t('dashboard.platformQuota.resetsAt', { time: formatResetTime(quotaVal(item.quota, `${w}_window_resets_at`) as string) }) }}
                </p>
              </template>
            </div>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import Icon from '@/components/icons/Icon.vue'
import type { UserDashboardStats as UserStatsType } from '@/api/usage'
import type { PlatformQuotaItem } from '@/types'

interface FusedPlatformCard {
  platform: string
  total_actual_cost: number
  today_actual_cost: number
  total_requests: number
  total_tokens: number
  isOther?: boolean
  quota?: PlatformQuotaItem
}

const props = defineProps<{
  stats: UserStatsType
  balance: number
  isSimple: boolean
  platformQuotas?: PlatformQuotaItem[] | null
}>()
const { t } = useI18n()

const PLATFORM_LABELS: Record<string, string> = {
  anthropic: 'Claude',
  openai: 'OpenAI',
  gemini: 'Gemini',
  antigravity: 'Antigravity'
}

const platformLabel = (p: string) => PLATFORM_LABELS[p] ?? p

const sortedPlatforms = computed(() => {
  const list = props.stats?.by_platform ?? []
  return [...list].sort((a, b) => b.total_actual_cost - a.total_actual_cost)
})

const OTHER_THRESHOLD = 0.0001
const platformCards = computed<FusedPlatformCard[]>(() => {
  const byPlat = new Map<string, (typeof sortedPlatforms.value)[number]>()
  for (const item of props.stats?.by_platform ?? []) byPlat.set(item.platform, item)

  const byQuota = new Map<string, PlatformQuotaItem>()
  for (const q of props.platformQuotas ?? []) byQuota.set(q.platform, q)

  const platforms = new Set<string>([...byPlat.keys(), ...byQuota.keys()])

  const PLATFORM_ORDER = ['anthropic', 'openai', 'gemini', 'antigravity']
  const cards: FusedPlatformCard[] = []

  for (const p of platforms) {
    const stat = byPlat.get(p)
    cards.push({
      platform: p,
      total_actual_cost: stat?.total_actual_cost ?? 0,
      today_actual_cost: stat?.today_actual_cost ?? 0,
      total_requests: stat?.total_requests ?? 0,
      total_tokens: stat?.total_tokens ?? 0,
      quota: byQuota.get(p),
    })
  }

  cards.sort((a, b) => {
    const ai = PLATFORM_ORDER.indexOf(a.platform)
    const bi = PLATFORM_ORDER.indexOf(b.platform)
    if (ai === -1 && bi === -1) return a.platform.localeCompare(b.platform)
    if (ai === -1) return 1
    if (bi === -1) return -1
    return ai - bi
  })

  const total = props.stats?.total_actual_cost ?? 0
  const today = props.stats?.today_actual_cost ?? 0
  const sumTotal = cards.reduce((s, c) => s + c.total_actual_cost, 0)
  const sumToday = cards.reduce((s, c) => s + c.today_actual_cost, 0)
  const diffTotal = Math.max(0, total - sumTotal)
  const diffToday = Math.max(0, today - sumToday)

  if (diffTotal > OTHER_THRESHOLD || diffToday > OTHER_THRESHOLD) {
    cards.push({
      platform: '__other__',
      total_actual_cost: diffTotal,
      today_actual_cost: diffToday,
      total_requests: 0,
      total_tokens: 0,
      isOther: true,
    })
  }

  return cards
})

// Quota helpers

type QuotaWindow = 'daily' | 'weekly' | 'monthly'
type QuotaField = `${QuotaWindow}_limit_usd` | `${QuotaWindow}_usage_usd` | `${QuotaWindow}_window_resets_at`

function quotaVal(q: PlatformQuotaItem | undefined, key: QuotaField): PlatformQuotaItem[QuotaField] {
  return q?.[key]
}

function hasAnyLimit(q: PlatformQuotaItem | undefined): boolean {
  if (!q) return false
  return q.daily_limit_usd != null || q.weekly_limit_usd != null || q.monthly_limit_usd != null
}

function calcPercent(usage: number, limit: number): number {
  if (!limit || limit <= 0) return 0
  return Math.min(100, Math.max(0, Math.round((usage / limit) * 100)))
}

function quotaBarClass(p: number): string {
  if (p >= 95) return 'bg-danger'
  if (p >= 75) return 'bg-warning'
  return 'bg-success'
}

const usdFormatter = new Intl.NumberFormat('en-US', {
  minimumFractionDigits: 2,
  maximumFractionDigits: 2,
})
function formatUsd(n: number): string {
  if (!Number.isFinite(n)) return '0.00'
  return usdFormatter.format(n)
}

function formatResetTime(iso: string | null | undefined): string {
  if (!iso) return ''
  const d = new Date(iso)
  if (Number.isNaN(d.getTime())) return iso
  return d.toLocaleString(undefined, {
    month: 'numeric',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    hour12: false,
  })
}

const formatBalance = (b: number) =>
  new Intl.NumberFormat('en-US', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2
  }).format(b)

const formatNumber = (n: number) => n.toLocaleString()
const formatCost = (c: number) => c.toFixed(4)
const formatTokens = (t: number) => {
  if (t >= 1_000_000) return `${(t / 1_000_000).toFixed(1)}M`
  if (t >= 1000) return `${(t / 1000).toFixed(1)}K`
  return t.toString()
}
const formatDuration = (ms: number) => ms >= 1000 ? `${(ms / 1000).toFixed(2)}s` : `${ms.toFixed(0)}ms`
</script>
