import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import Landing from './views/Landing.vue'
import Dashboard from './views/Dashboard.vue'

const routes = [
  { path: '/', component: Landing },
  { path: '/dashboard', component: Dashboard }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

createApp(App).use(router).mount('#app')

