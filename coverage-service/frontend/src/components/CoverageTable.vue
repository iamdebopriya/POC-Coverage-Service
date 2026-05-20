<script setup lang="ts">
import type { Coverage } from '../types'
defineProps<{ rows: Coverage[] }>()
function passRate(r: Coverage) { return r.total_tests ? ((r.passed_tests/r.total_tests)*100).toFixed(1)+'%' : '—' }
function fmt(ts: string) { return new Date(ts).toLocaleString() }
function pill(v: number) { return v >= 85 ? 'green' : v >= 50 ? 'orange' : 'red' }
</script>
<template>
  <div class="wrap">
    <table>
      <thead><tr>
        <th>Service</th><th>BE Cov.</th><th>FE Cov.</th>
        <th>Total</th><th>Passed</th><th>Failed</th><th>Flaky</th>
        <th>Pass Rate</th><th>Avg Time (s)</th><th>Timestamp</th>
      </tr></thead>
      <tbody>
        <tr v-if="!rows.length"><td colspan="10" class="empty">No results yet — run tests above.</td></tr>
        <tr v-for="r in rows" :key="r.id">
          <td class="svc">{{ r.service_name }}</td>
          <td><span class="pill" :class="pill(r.backend_coverage)">{{ r.backend_coverage.toFixed(1) }}%</span></td>
          <td><span class="pill" :class="pill(r.frontend_coverage)">{{ r.frontend_coverage.toFixed(1) }}%</span></td>
          <td>{{ r.total_tests }}</td>
          <td class="ok">{{ r.passed_tests }}</td>
          <td class="bad">{{ r.failed_tests }}</td>
          <td class="warn">{{ r.flaky_tests }}</td>
          <td>{{ passRate(r) }}</td>
          <td>{{ r.avg_execution_time.toFixed(2) }}</td>
          <td class="ts">{{ fmt(r.timestamp) }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<style scoped>
.wrap { overflow-x:auto; border-radius:12px; box-shadow:0 2px 12px rgba(0,0,0,0.07); background:#fff; }
table { width:100%; border-collapse:collapse; font-size:0.875rem; }
th { background:#f9fafb; padding:0.7rem 1rem; text-align:left; font-size:0.72rem; text-transform:uppercase;
  letter-spacing:0.05em; color:#6b7280; border-bottom:1px solid #e5e7eb; white-space:nowrap; }
td { padding:0.7rem 1rem; border-bottom:1px solid #f3f4f6; color:#374151; }
tr:last-child td { border-bottom:none; } tr:hover td { background:#f9fafb; }
.svc { font-weight:600; color:#111827; }
.ok  { color:#059669; font-weight:600; }
.bad { color:#dc2626; font-weight:600; }
.warn{ color:#d97706; font-weight:600; }
.ts  { color:#9ca3af; white-space:nowrap; font-size:0.78rem; }
.empty { text-align:center; padding:2rem; color:#9ca3af; }
.pill { display:inline-block; padding:0.18rem 0.5rem; border-radius:999px; font-weight:600; font-size:0.78rem; }
.pill.green  { background:#d1fae5; color:#065f46; }
.pill.orange { background:#fef3c7; color:#92400e; }
.pill.red    { background:#fee2e2; color:#991b1b; }
</style>
