<script setup lang="ts">
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <div class="app">
    <HelloWorld />
  </div>
</template>

<style>
* {
  box-sizing: border-box;
}

body {
  margin: 0;
  font-family: 'Segoe UI', system-ui, -apple-system, sans-serif;
  background: #f3f4f6;
}

.app {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
