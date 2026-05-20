<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'

import RunPanel from '../components/RunPanel.vue'
import StatCard from '../components/StatCard.vue'
import CoverageTable from '../components/CoverageTable.vue'
import FilterBar from '../components/FilterBar.vue'

import {
  fetchRegisteredServices,
  fetchCoverages,
  fetchCoverageServices,
  downloadCoverageUrl
} from '../api'

import type { Coverage, RegisteredService } from '../types'

const router = useRouter()

const registeredServices = ref<RegisteredService[]>([])
const coverages = ref<Coverage[]>([])
const filterServices = ref<string[]>(([]))
const selectedService = ref('')
const from = ref('')
const to = ref('')
const loading = ref(false)
const error = ref('')

const stats = computed(() => {
  if (!coverages.value.length) return null

  const rows = coverages.value

  return {
    total: rows.reduce((s, r) => s + r.total_tests, 0),
    passed: rows.reduce((s, r) => s + r.passed_tests, 0),
    failed: rows.reduce((s, r) => s + r.failed_tests, 0),
    avgBE:
      rows.reduce((s, r) => s + r.backend_coverage, 0) / rows.length,
    avgFE:
      rows.reduce((s, r) => s + r.frontend_coverage, 0) / rows.length
  }
})

const latest = computed(() => coverages.value[0] ?? null)

async function loadCoverages() {
  loading.value = true
  error.value = ''

  try {
    const [covs, svcs] = await Promise.all([
      fetchCoverages({
        service: selectedService.value || undefined,
        from: from.value || undefined,
        to: to.value || undefined
      }),
      fetchCoverageServices()
    ])

    coverages.value = covs
    filterServices.value = svcs
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'Failed to load'
  } finally {
    loading.value = false
  }
}

async function init() {
  const [svcs] = await Promise.all([
    fetchRegisteredServices(),
    loadCoverages()
  ])

  registeredServices.value = svcs
}

function download() {
  window.open(
    downloadCoverageUrl({
      service: selectedService.value || undefined,
      from: from.value || undefined,
      to: to.value || undefined
    }),
    '_blank'
  )
}

function goBack() {
  router.push('/')
}

onMounted(init)
</script>

<template>
  <div class="dashboard-page">

    <!-- TOP -->
    <header class="topbar">

      <div class="top-left">

        <button class="back-btn" @click="goBack">
          ← Back
        </button>

        <div class="heading-wrap">
          <span class="page-label">
            TEST COVERAGE TRACKER
          </span>

          <h1 class="page-title">
            Coverage Dashboard
          </h1>
        </div>

      </div>

      <div class="live-wrap">
        <span class="live-dot"></span>
        <span class="live-text">Live</span>
      </div>

    </header>

    <!-- STATS -->
    <section v-if="stats" class="stats-wrap">

      <div class="section-label">
        Overview
      </div>

      <div class="stats-grid">

        <StatCard
          label="Avg BE Coverage"
          :value="stats.avgBE.toFixed(1) + '%'"
          color="#737373"
        />

        <StatCard
          label="Avg FE Coverage"
          :value="stats.avgFE.toFixed(1) + '%'"
          color="#737373"
        />

        <StatCard
          label="Total Tests"
          :value="stats.total"
          color="#737373"
        />

        <StatCard
          label="Passed"
          :value="stats.passed"
          sub="all runs"
          color="#737373"
        />

        <StatCard
          label="Failed"
          :value="stats.failed"
          color="#737373"
        />
        <StatCard
          label="Pass Rate"
          :value="((stats.passed / stats.total) * 100).toFixed(1) + '%'"
          color="#737373"
        />
        <StatCard
          label="Runs"
          :value="coverages.length"
          color="#737373"
        />
        
        <StatCard
          label="Services"
          :value="registeredServices.length"
          color="#737373"
        />
       
      </div>

    </section>

    <!-- RUN -->
    <section class="section-card">

      <div class="section-label">
        Run Tests
      </div>

      <RunPanel
        :services="registeredServices"
        @done="loadCoverages"
      />

    </section>

    <!-- LATEST -->
    <section
      v-if="latest"
      class="latest-banner"
    >

      <div class="section-label">
        Latest Run
      </div>

      <div class="latest-row">

        <span>
          <strong>{{ latest.service_name }}</strong>
        </span>

        <span>
          BE {{ latest.backend_coverage.toFixed(1) }}%
        </span>

        <span>
          FE {{ latest.frontend_coverage.toFixed(1) }}%
        </span>

        <span>
          {{ latest.passed_tests }}/{{ latest.total_tests }} Passed
        </span>

        <span>
          {{ new Date(latest.timestamp).toLocaleString() }}
        </span>

      </div>

    </section>

    <!-- FILTER -->
    <section class="section-card">

      <div class="section-label">
        Filters
      </div>

      <FilterBar
        v-model:selectedService="selectedService"
        v-model:from="from"
        v-model:to="to"
        :services="filterServices"
        @filter="loadCoverages"
        @download="download"
      />

    </section>

    <!-- TABLE -->
    <section class="table-section">

      <div class="section-label">
        Coverage History
      </div>

      <div v-if="loading" class="msg">
        Loading...
      </div>

      <div v-else-if="error" class="msg error">
        {{ error }}
      </div>

      <CoverageTable
        v-else
        :rows="coverages"
      />

    </section>

  </div>
</template>

<style scoped>
.dashboard-page {
  min-height: 100vh;
  background: #f5f5f4;
  color: #111111;
  padding: 1.5rem 2rem;
  font-family:
    Inter,
    system-ui,
    sans-serif;
}

/* TOP */

.topbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 2rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid #dddddd;
}

.top-left {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.back-btn {
  border: 1px solid #d4d4d4;
  background: white;
  color: #111;
  padding: 0.55rem 1rem;
  border-radius: 10px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: 0.2s;
}

.back-btn:hover {
  background: #f0f0f0;
}

.heading-wrap {
  display: flex;
  flex-direction: column;
  gap: 0.2rem;
}

.page-label {
  font-size: 0.72rem;
  letter-spacing: 0.15em;
  color: #737373;
  font-weight: 700;
}

.page-title {
  margin: 0;
  font-size: 1.4rem;
  font-weight: 700;
}

.live-wrap {
  display: flex;
  align-items: center;
  gap: 0.45rem;
}

.live-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #737373;
}

.live-text {
  font-size: 0.8rem;
  color: #525252;
  font-weight: 600;
}

/* LABELS */

.section-label {
  font-size: 0.75rem;
  font-weight: 700;
  letter-spacing: 0.12em;
  color: #737373;
  margin-bottom: 0.8rem;
  text-transform: uppercase;
}

/* STATS */

.stats-wrap {
  margin-bottom: 1.5rem;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(170px, 1fr));
  gap: 0.8rem;
}

/* CARDS */

.section-card {
  background: #fafafa;
  border: 1px solid #e4e4e4;
  border-radius: 16px;
  padding: 1rem;
  margin-bottom: 1rem;
}

/* LATEST */

.latest-banner {
  background: #eeeeee;
  border: 1px solid #d6d6d6;
  border-radius: 14px;
  padding: 0.9rem 1rem;
  margin-bottom: 1rem;
}

.latest-row {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
  font-size: 0.9rem;
  color: #444;
}

.latest-row strong {
  color: #111;
}

/* TABLE */

.table-section {
  background: #fafafa;
  border: 1px solid #e4e4e4;
  border-radius: 16px;
  padding: 1rem;
}

.msg {
  text-align: center;
  padding: 2rem;
  color: #666;
}

.msg.error {
  color: #dc2626;
}

/* RESPONSIVE */

@media (max-width: 768px) {

  .dashboard-page {
    padding: 1rem;
  }

  .topbar {
    flex-direction: column;
    align-items: flex-start;
    gap: 1rem;
  }

  .latest-row {
    flex-direction: column;
    gap: 0.5rem;
  }
}
</style>