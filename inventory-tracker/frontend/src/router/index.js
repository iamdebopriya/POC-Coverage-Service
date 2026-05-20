import { createRouter, createWebHistory } from 'vue-router'
import Landing from '../pages/Landing.vue'
import Login from '../pages/Login.vue'
import Register from '../pages/Register.vue'
import Home from '../pages/Home.vue'
import AddItem from '../pages/AddItem.vue'
import EditItem from '../pages/EditItem.vue'
import Dashboard from '../pages/Dashboard.vue'

const routes = [
  {
    path: '/',
    name: 'Landing',
    component: Landing
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/inventory',
    name: 'Home',
    component: Home,
    meta: { requiresAuth: true }
  },
  {
    path: '/add',
    name: 'AddItem',
    component: AddItem,
    meta: { requiresAuth: true }
  },
  {
    path: '/edit/:id',
    name: 'EditItem',
    component: EditItem,
    meta: { requiresAuth: true }
  },
  // {
  //   path: '/dashboard',
  //   name: 'Dashboard',
  //   component: Dashboard,
  //   meta: { requiresAuth: true }
  // }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const user = localStorage.getItem('user')
  if (to.meta.requiresAuth && !user) {
    next('/login')
  } else {
    next()
  }
})

export default router