<template>
  <div id="app">

    <nav class="navbar">


      <div class="nav-left">

        <div class="logo-box">
          <i class="fas fa-boxes-stacked"></i>
        </div>

        <div class="nav-brand">
          <h1>Inventory Tracker</h1>
          <span>Store Management</span>
        </div>

      </div>



      <div class="nav-links">

        <router-link to="/">
          Home
        </router-link>

        <router-link
          to="/inventory"
          v-if="user"
        >
          Inventory
        </router-link>

        <!-- <router-link
          to="/dashboard"
          v-if="user"
        >
          Dashboard
        </router-link> -->



        <div
          class="store-badge"
          v-if="user"
        >
          <i class="fas fa-store"></i>

          <span>
            {{ user.store_name }}
          </span>
        </div>



        <button
          @click="handleLogout"
          class="btn-logout"
          v-if="user"
        >
          Logout
        </button>


        <router-link
          to="/login"
          class="btn-nav-login"
          v-if="!user"
        >
          Login
        </router-link>


        <router-link
          to="/register"
          class="btn-nav-register"
          v-if="!user"
        >
          Sign Up
        </router-link>

      </div>

    </nav>


    <main class="main-content">
      <router-view />
    </main>

  </div>
</template>

<script setup>

import {
  ref,
  watch
} from 'vue'

import {
  useRouter,
  useRoute
} from 'vue-router'

const router = useRouter()
const route = useRoute()

const user = ref(
  JSON.parse(
    localStorage.getItem('user') || 'null'
  )
)

const handleLogout = () => {

  localStorage.removeItem('user')

  user.value = null

  router.push('/')
}

watch(route, () => {

  user.value = JSON.parse(
    localStorage.getItem('user') || 'null'
  )
})

</script>

<style>

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family:
    'Segoe UI',
    sans-serif;

  background: #f5f7fb;
  color: #333;
}

#app {
  width: 100%;
}


.navbar {

  position: sticky;
  top: 0;
  z-index: 100;

  display: flex;
  justify-content: space-between;
  align-items: center;

  width: 100%;

  padding: 1rem 2.5rem;

  background:
    linear-gradient(
      135deg,
      #1e3a5f,
      #2c3e50
    );

  border-bottom:
    1px solid rgba(255,255,255,0.08);

  box-shadow:
    0 8px 24px rgba(0,0,0,0.08);
}


.nav-left {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.logo-box {

  width: 48px;
  height: 48px;

  border-radius: 14px;

  background:
    linear-gradient(
      135deg,
      #60a5fa,
      #2563eb
    );

  display: flex;
  align-items: center;
  justify-content: center;

  color: white;
  font-size: 1.1rem;

  box-shadow:
    0 6px 16px rgba(37,99,235,0.35);
}


.nav-brand {
  display: flex;
  flex-direction: column;
}

.nav-brand h1 {

  color: white;

  font-size: 1.35rem;
  font-weight: 700;

  letter-spacing: 0.02em;
}

.nav-brand span {

  color: #cbd5e1;

  font-size: 0.76rem;

  margin-top: 2px;
}



.nav-links {

  display: flex;
  align-items: center;
  gap: 1rem;
}


.nav-links a {

  position: relative;

  color: #dbe4ee;

  text-decoration: none;

  font-size: 0.92rem;
  font-weight: 500;

  transition: 0.25s ease;
}

.nav-links a:hover,
.nav-links a.router-link-active {

  color: white;
}


.nav-links a::after {

  content: "";

  position: absolute;

  left: 0;
  bottom: -6px;

  width: 0%;
  height: 2px;

  background: #60a5fa;

  transition: 0.25s ease;
}

.nav-links a:hover::after,
.nav-links a.router-link-active::after {

  width: 100%;
}


.store-badge {

  display: flex;
  align-items: center;
  gap: 0.5rem;

  padding:
    0.45rem 0.9rem;

  border-radius: 999px;

  background:
    rgba(255,255,255,0.12);

  color: #e0e7ff;

  font-size: 0.84rem;
  font-weight: 600;

  backdrop-filter: blur(10px);

  border:
    1px solid rgba(255,255,255,0.08);
}

.store-badge i {
  color: #93c5fd;
}


/* Logout */

.btn-logout {

  background:
    linear-gradient(
      135deg,
      #ef4444,
      #dc2626
    );

  color: white;

  border: none;

  padding:
    0.55rem 1rem;

  border-radius: 10px;

  font-size: 0.84rem;
  font-weight: 600;

  cursor: pointer;

  transition: 0.25s ease;
}

.btn-logout:hover {

  transform: translateY(-2px);

  box-shadow:
    0 8px 18px rgba(239,68,68,0.25);
}



.btn-nav-login {

  padding:
    0.55rem 1rem !important;

  border-radius: 10px;

  border:
    1px solid rgba(255,255,255,0.18);

  color: white !important;

  background:
    rgba(255,255,255,0.05);

  transition: 0.25s ease;
}

.btn-nav-login:hover {

  background:
    rgba(255,255,255,0.12);
}


.btn-nav-register {

  padding:
    0.55rem 1rem !important;

  border-radius: 10px;

  background:
    linear-gradient(
      135deg,
      #3b82f6,
      #2563eb
    );

  color: white !important;

  font-weight: 600;

  box-shadow:
    0 6px 18px rgba(37,99,235,0.25);

  transition: 0.25s ease;
}

.btn-nav-register:hover {

  transform: translateY(-2px);

  box-shadow:
    0 10px 20px rgba(37,99,235,0.35);
}


.main-content {

  width: 100%;

  padding:
    2rem 2.5rem;
}


@media (max-width: 900px) {

  .navbar {

    flex-direction: column;
    align-items: flex-start;

    gap: 1rem;

    padding: 1rem 1.2rem;
  }

  .nav-links {

    flex-wrap: wrap;
    gap: 0.8rem;
  }

  .main-content {

    padding: 1.2rem;
  }
}

</style>