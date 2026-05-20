<script setup lang="ts">
import { ref } from 'vue'

interface HelloResponse {
  message: string
  status: string
}

const message = ref<string>('')
const error = ref<string>('')
const loading = ref<boolean>(false)
const fetched = ref<boolean>(false)

async function fetchHello() {
  loading.value = true
  error.value = ''
  message.value = ''

  try {
    const res = await fetch('/api/hello')
    if (!res.ok) throw new Error(`HTTP error: ${res.status}`)
    const data: HelloResponse = await res.json()
    message.value = data.message
    fetched.value = true
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'Unknown error'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="card">
    <h1 class="title">Hello World App</h1>
    <p class="subtitle">Vue 3 + Go backend</p>

    <button class="btn" :disabled="loading" @click="fetchHello">
      {{ loading ? 'Loading…' : 'Say Hello' }}
    </button>

    <p v-if="message" class="message success" data-testid="message">
      {{ message }}
    </p>
    <p v-if="error" class="message error" data-testid="error">
      Error: {{ error }}
    </p>
  </div>
</template>

<style scoped>
.card {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1.2rem;
  padding: 3rem 2.5rem;
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 4px 32px rgba(0, 0, 0, 0.08);
  min-width: 320px;
}

.title {
  font-size: 2rem;
  font-weight: 700;
  margin: 0;
  color: #1a1a2e;
}

.subtitle {
  margin: 0;
  color: #6b7280;
  font-size: 0.95rem;
}

.btn {
  padding: 0.65rem 2rem;
  border: none;
  border-radius: 8px;
  background: #4f46e5;
  color: #fff;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s, transform 0.1s;
}

.btn:hover:not(:disabled) {
  background: #4338ca;
  transform: translateY(-1px);
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.message {
  margin: 0;
  font-size: 1.1rem;
  font-weight: 500;
  padding: 0.6rem 1.2rem;
  border-radius: 8px;
}

.success {
  background: #ecfdf5;
  color: #065f46;
}

.error {
  background: #fef2f2;
  color: #991b1b;
}
</style>
