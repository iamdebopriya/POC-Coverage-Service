import { defineConfig } from 'vitest/config'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  test: {
  environment: 'jsdom',
  globals: true,
  coverage: {
    provider: 'v8',
    reporter: ['text', 'json-summary', 'html'],
  },
},
  server: {
    port: 5173,
    proxy: {
      '/api': {
        target: 'http://localhost:8082',
        changeOrigin: true,
      },
    },
  },
})