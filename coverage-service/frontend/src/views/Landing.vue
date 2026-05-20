<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

function goToDashboard() {
  router.push('/dashboard')
}
</script>

<template>
  <div class="landing">
    <!-- Geometric background accents -->
    <div class="geo geo-circle"></div>
    <div class="geo geo-ring"></div>
    <div class="geo geo-dot-grid"></div>

    <div class="content">
      <div class="eyebrow">
        <span class="eyebrow-line"></span>
        <span class="eyebrow-text">Test Coverage Tracker</span>
      </div>

      <h1>
        <span class="line line-1">Coverage</span>
        <span class="line line-2">Dashboard</span>
      </h1>

      <p class="subtitle">
        Run tests, track BE &amp; FE coverage, and monitor<br />
        pass/fail results across your services.
      </p>

      <div class="actions">
        <button @click="goToDashboard" class="btn-primary">
          <span>Go to Dashboard</span>
          <svg class="btn-arrow" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M5 12H19M19 12L12 5M19 12L12 19" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
        </button>
        <span class="btn-hint">Select a service to get started</span>
      </div>

      <div class="stats-strip">
        <div class="stat">
          <span class="stat-value">70% + </span>
          <span class="stat-label">Avg BE Coverage</span>
        </div>
        <div class="stat-divider"></div>
        <div class="stat">
          <span class="stat-value stat-value--accent">70% +</span>
          <span class="stat-label">Avg FE Coverage</span>
        </div>
      </div>
    </div>

    <!-- Feature highlights panel -->
    <div class="panel-right">
      <div class="panel-card card-1">
        <div class="feature-icon">
          <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M5 12H19M19 12L12 5M19 12L12 19" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
        </div>
        <div class="card-label">Run Tests</div>
        <div class="feature-desc">Select any service and trigger a test run directly from the dashboard. Watch live output stream in real time.</div>
      </div>

      <div class="panel-card card-2">
        <div class="feature-icon icon-dark">
          <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M9 19v-6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2zm0 0V9a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v10m-6 0a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2m0 0V5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
        </div>
        <div class="card-label">BE &amp; FE Coverage</div>
        <div class="feature-desc">Track backend and frontend coverage separately per service, with pass rate and average run time.</div>
      </div>

      <div class="panel-card card-3">
        <div class="feature-icon">
          <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M4 16v1a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-1m-4-4-4 4m0 0-4-4m4 4V4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
        </div>
        <div class="card-label">Export &amp; Filter</div>
        <div class="feature-desc">Filter runs by service or date range. Export the full history as a CSV with one click.</div>
      </div>
    </div>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=DM+Serif+Display:ital@0;1&family=DM+Sans:wght@300;400;500;600&display=swap');

*, *::before, *::after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

:root {
  --navy: #0B1F3A;
  --navy-mid: #162d50;
  --navy-light: #1e3d6e;
  --white: #FFFFFF;
  --off-white: #F4F6F9;
  --accent: #2563EB;
  --accent-soft: #DBEAFE;
  --muted: #94A3B8;
  --border: rgba(11, 31, 58, 0.10);
}

.landing {
  font-family: 'DM Sans', sans-serif;
  width: 100%;
  min-height: 100vh;
  background-color: var(--white);
  display: grid;
  grid-template-columns: 1fr 420px;
  align-items: center;
  padding: 0 80px;
  gap: 60px;
  position: relative;
  overflow: hidden;
}

/* ─── Geometric decorations ─── */
.geo {
  position: absolute;
  pointer-events: none;
}

.geo-circle {
  width: 520px;
  height: 520px;
  border-radius: 50%;
  background: var(--navy);
  top: -180px;
  right: -100px;
  opacity: 0.04;
}

.geo-ring {
  width: 300px;
  height: 300px;
  border-radius: 50%;
  border: 40px solid var(--navy);
  bottom: -80px;
  left: -80px;
  opacity: 0.05;
}

.geo-dot-grid {
  width: 200px;
  height: 200px;
  background-image: radial-gradient(circle, var(--navy) 1.5px, transparent 1.5px);
  background-size: 18px 18px;
  top: 60px;
  right: 500px;
  opacity: 0.07;
}

/* ─── Left content ─── */
.content {
  display: flex;
  flex-direction: column;
  gap: 0;
  animation: fadeUp 0.7s ease both;
}

.eyebrow {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 28px;
}

.eyebrow-line {
  display: block;
  width: 32px;
  height: 2px;
  background: var(--accent);
}

.eyebrow-text {
  font-family: 'DM Sans', sans-serif;
  font-size: 0.75rem;
  font-weight: 600;
  letter-spacing: 0.15em;
  text-transform: uppercase;
  color: var(--accent);
}

h1 {
  font-family: 'DM Serif Display', serif;
  color: var(--navy);
  display: flex;
  flex-direction: column;
  line-height: 1;
  margin-bottom: 28px;
}

.line {
  display: block;
}

.line-1 {
  font-size: clamp(4rem, 8vw, 7rem);
  animation: slideIn 0.6s ease both 0.1s;
}

.line-2 {
  font-size: clamp(4rem, 8vw, 7rem);
  color: var(--accent);
  animation: slideIn 0.6s ease both 0.2s;
}

.subtitle {
  font-size: 1.1rem;
  color: var(--muted);
  line-height: 1.7;
  font-weight: 400;
  margin-bottom: 44px;
  animation: fadeUp 0.6s ease both 0.3s;
}

.actions {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-bottom: 60px;
  animation: fadeUp 0.6s ease both 0.4s;
}

.btn-primary {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  padding: 14px 28px;
  background: var(--navy);
  color: var(--white);
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-family: 'DM Sans', sans-serif;
  font-size: 0.95rem;
  font-weight: 600;
  letter-spacing: 0.01em;
  transition: background 0.25s ease, transform 0.2s ease, box-shadow 0.25s ease;
}

.btn-primary:hover {
  background: var(--navy-light);
  transform: translateY(-2px);
  box-shadow: 0 12px 30px rgba(11, 31, 58, 0.18);
}

.btn-primary:active {
  transform: translateY(0);
  box-shadow: none;
}

.btn-arrow {
  width: 18px;
  height: 18px;
  transition: transform 0.2s ease;
}

.btn-primary:hover .btn-arrow {
  transform: translateX(3px);
}

.btn-hint {
  font-size: 0.8rem;
  color: var(--muted);
  font-weight: 400;
}

/* ─── Stats strip ─── */
.stats-strip {
  display: flex;
  align-items: center;
  gap: 28px;
  padding-top: 28px;
  border-top: 1px solid var(--border);
  animation: fadeUp 0.6s ease both 0.5s;
}

.stat {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.stat-value {
  font-size: 1.75rem;
  font-weight: 700;
  color: var(--navy);
  line-height: 1;
  font-family: 'DM Sans', sans-serif;
}

.stat-value--accent {
  color: var(--accent);
}

.stat-label {
  font-size: 0.82rem;
  font-weight: 500;
  color: var(--muted);
}

.stat-divider {
  width: 1px;
  height: 44px;
  background: var(--border);
}

/* ─── Right panel ─── */
.panel-right {
  display: flex;
  flex-direction: column;
  gap: 16px;
  animation: fadeUp 0.7s ease both 0.2s;
}

.panel-card {
  background: var(--navy);
  border-radius: 14px;
  padding: 24px 26px;
  color: var(--white);
}

.card-1 { background: var(--navy); }
.card-2 { background: var(--off-white); color: var(--navy); }
.card-3 { background: var(--navy-mid); }

.feature-icon {
  width: 36px;
  height: 36px;
  background: rgba(255,255,255,0.12);
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 14px;
}

.feature-icon svg {
  width: 18px;
  height: 18px;
  stroke: currentColor;
}

/* Dark cards → white icon */
.card-1 .feature-icon svg,
.card-3 .feature-icon svg {
  color: white;
}

/* Light card → navy icon */
.card-2 .feature-icon svg {
  color: var(--navy);
}

.icon-dark {
  background: rgba(11, 31, 58, 0.1);
}

.card-label {
  font-size: 1rem;
  font-weight: 600;
  margin-bottom: 6px;
}

.feature-desc {
  font-size: 0.82rem;
  line-height: 1.6;
  opacity: 0.7;
  margin-top: 6px;
}

.card-2 .card-label { color: var(--navy); }
.card-2 .feature-desc { color: var(--navy); }

/* ─── Animations ─── */
@keyframes fadeUp {
  from { opacity: 0; transform: translateY(20px); }
  to   { opacity: 1; transform: translateY(0); }
}

@keyframes slideIn {
  from { opacity: 0; transform: translateX(-24px); }
  to   { opacity: 1; transform: translateX(0); }
}

/* ─── Responsive ─── */
@media (max-width: 900px) {
  .landing {
    grid-template-columns: 1fr;
    padding: 60px 32px;
    gap: 48px;
  }

  .panel-right {
    flex-direction: row;
    flex-wrap: wrap;
  }

  .panel-card {
    flex: 1 1 140px;
  }
}
</style>