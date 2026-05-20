<script setup lang="ts">
import { ref, nextTick, onUnmounted } from 'vue'
import { runServiceTests } from '../api'
import type { RegisteredService } from '../types'

defineProps<{ services: RegisteredService[] }>()
const emit = defineEmits<{ (e: 'done'): void }>()

const selected = ref('')
const lines = ref<string[]>([])
const running = ref(false)
const status = ref<'idle' | 'running' | 'success' | 'error'>('idle')
const terminalEl = ref<HTMLElement | null>(null)

let es: EventSource | null = null

onUnmounted(() => es?.close())

async function scrollBottom() {
  await nextTick()
  if (terminalEl.value) terminalEl.value.scrollTop = terminalEl.value.scrollHeight
}

function run() {
  if (!selected.value || running.value) return

  lines.value = []
  running.value = true
  status.value = 'running'

  es = runServiceTests(
    selected.value,
    (line) => {
      lines.value.push(line)
      scrollBottom()
    },
    () => {
      running.value = false
      status.value = 'success'
      emit('done')
    },
    (msg) => {
      lines.value.push('ERROR: ' + msg)
      running.value = false
      status.value = 'error'
      scrollBottom()
    },
  )
}

function stop() {
  es?.close()
  running.value = false
  status.value = 'idle'
  lines.value.push('--- stopped by user ---')
}

function lineClass(line: string) {
  if (/PASS|✓|✔|passed/i.test(line) && !/FAIL/i.test(line)) return 'pass'
  if (/FAIL|✗|✘|failed|ERROR/i.test(line)) return 'fail'
  if (/coverage:|▶|===|---|\d+\.\d+%|✓ Results saved/.test(line)) return 'info'
  return ''
}
</script>

<template>
  <div class="panel">
    <!-- Header row -->
    <div class="panel-header">
      <span class="panel-title">Run Tests</span>
      <div class="controls">
        <select
          v-model="selected"
          :disabled="running"
          class="svc-select"
          data-testid="service-select"
        >
          <option value="" disabled>Select a service…</option>
          <option v-for="s in services" :key="s.name" :value="s.name">
            {{ s.display_name }}
          </option>
        </select>

        <button
          v-if="!running"
          class="btn btn-run"
          :disabled="!selected"
          data-testid="btn-run"
          @click="run"
        >
          ▶ Run Tests
        </button>
        <button v-else class="btn btn-stop" data-testid="btn-stop" @click="stop">
          ■ Stop
        </button>

        <span v-if="status === 'running'" class="badge badge-running">Running…</span>
        <span v-else-if="status === 'success'" class="badge badge-ok">✓ Done</span>
        <span v-else-if="status === 'error'" class="badge badge-err">✗ Failed</span>
      </div>
    </div>

    <!-- Terminal output -->
    <div v-if="lines.length || running" class="terminal" ref="terminalEl" data-testid="terminal">
      <div v-if="!lines.length && running" class="term-waiting">Connecting…</div>
      <div
        v-for="(line, i) in lines"
        :key="i"
        class="term-line"
        :class="lineClass(line)"
      >{{ line || '\u00a0' }}</div>
      <div v-if="running" class="term-cursor">█</div>
    </div>

    <!-- Empty state -->
    <div v-else class="empty-state">
      Select a service above and click <strong>Run Tests</strong> to see live output here.
    </div>
  </div>
</template>

<style scoped>
.panel {
  background: #fafafa;
  border: 1px solid #e5e5e5;
  border-radius: 18px;
  overflow: hidden;
}

/* HEADER */

.panel-header {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem 1.2rem;
  border-bottom: 1px solid #ececec;
  flex-wrap: wrap;
}

.panel-title {
  font-weight: 700;
  font-size: 0.95rem;
  color: #27272a;
  white-space: nowrap;
  letter-spacing: -0.02em;
}

/* CONTROLS */

.controls {
  display: flex;
  align-items: center;
  gap: 0.7rem;
  flex: 1;
  flex-wrap: wrap;
}

/* SELECT */

.svc-select {
  flex: 1;
  min-width: 220px;
  max-width: 340px;

  height: 42px;

  padding: 0 0.9rem;

  border: 1px solid #d4d4d8;
  border-radius: 12px;

  font-size: 0.88rem;

  color: #18181b;
  background: #ffffff;

  outline: none;

  transition:
    border-color 0.18s ease,
    background 0.18s ease;
}

.svc-select:focus {
  border-color: #a1a1aa;
  background: #fcfcfc;
}

.svc-select:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

/* BUTTONS */

.btn {
  height: 42px;

  padding: 0 1.15rem;

  border-radius: 12px;

  font-size: 0.84rem;
  font-weight: 600;

  cursor: pointer;

  transition: 0.18s ease;

  white-space: nowrap;
}

/* RUN */

.btn-run {
  background: #111111;
  color: white;
  border: 1px solid #111111;
}

.btn-run:hover:not(:disabled) {
  opacity: 0.92;
}

.btn-run:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* STOP */

.btn-stop {
  background: #f3f4f6;
  color: #3f3f46;
  border: 1px solid #d4d4d8;
}

.btn-stop:hover {
  background: #ebebeb;
}

/* BADGES */

.badge {
  font-size: 0.72rem;
  font-weight: 700;

  padding: 0.3rem 0.7rem;

  border-radius: 999px;

  border: 1px solid #dddddd;

  background: #f4f4f5;

  color: #52525b;
}

.badge-running {
  background: #f4f4f5;
  color: #52525b;
}

.badge-ok {
  background: #f4f4f5;
  color: #52525b;
}

.badge-err {
  background: #f4f4f5;
  color: #52525b;
}

/* TERMINAL */

.terminal {
  background: #0f1117;

  padding: 0.95rem 1rem;

  font-family:
    'Fira Mono',
    'Consolas',
    'Courier New',
    monospace;

  font-size: 0.8rem;

  line-height: 1.58;

  max-height: 380px;

  overflow-y: auto;

  scroll-behavior: smooth;
}

.term-waiting {
  color: #6b7280;
  font-style: italic;
}

.term-line {
  color: #d1d5db;

  white-space: pre-wrap;

  word-break: break-word;
}

/* KEEP TERMINAL COLORS */

.term-line.pass {
  color: #34d399;
}

.term-line.fail {
  color: #f87171;
}

.term-line.info {
  color: #a78bfa;
}

.term-cursor {
  display: inline-block;

  color: #d4d4d8;

  animation: blink 1s step-end infinite;
}

@keyframes blink {
  0%,100% {
    opacity: 1;
  }

  50% {
    opacity: 0;
  }
}

/* EMPTY */

.empty-state {
  padding: 1.5rem 1.25rem;

  color: #8a8a8a;

  font-size: 0.875rem;
}

/* SCROLLBAR */

.terminal::-webkit-scrollbar {
  width: 8px;
}

.terminal::-webkit-scrollbar-track {
  background: #18181b;
}

.terminal::-webkit-scrollbar-thumb {
  background: #3f3f46;
  border-radius: 999px;
}

/* MOBILE */

@media (max-width: 768px) {

  .panel-header {
    align-items: flex-start;
  }

  .controls {
    width: 100%;
  }

  .svc-select {
    width: 100%;
    max-width: 100%;
  }

  .btn {
    flex: 1;
  }
}
</style>