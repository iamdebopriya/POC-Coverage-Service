<template>
  <div class="login-page">
    <div class="login-card">
      <div class="login-header">
        <h1>Inventory Tracker</h1>
        <p>Sign in to your store</p>
      </div>

      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label>Store Name</label>
          <input
            v-model="form.store_name"
            type="text"
            placeholder="Enter store name"
            required
          />
        </div>

        <div class="form-group">
          <label>Password</label>
          <input
            v-model="form.password"
            type="password"
            placeholder="Enter password"
            required
          />
        </div>

        <div class="error-msg" v-if="error">{{ error }}</div>

        <button type="submit" class="btn-login">Sign In</button>
      </form>

      <div class="login-footer">
        Don't have an account?
        <router-link to="/register">Register here</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { login } from '../services/api.js'

const router = useRouter()
const error = ref('')

const form = ref({
  store_name: '',
  password: ''
})

const handleLogin = async () => {
  try {
    error.value = ''
    const res = await login(form.value)
    localStorage.setItem('user', JSON.stringify(res.data))
    router.push('/inventory')
  } catch (err) {
    error.value = 'Invalid store name or password'
  }
}
</script>

<style scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f5f6fa;
}

.login-card {
  background: white;
  border-radius: 12px;
  padding: 2.5rem;
  width: 100%;
  max-width: 420px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.08);
}

.login-header {
  text-align: center;
  margin-bottom: 2rem;
}

.login-header h1 {
  font-size: 1.6rem;
  font-weight: 700;
  color: #2c3e50;
  margin-bottom: 6px;
}

.login-header p {
  font-size: 0.9rem;
  color: #888;
}

.form-group {
  margin-bottom: 1.25rem;
}

.form-group label {
  display: block;
  font-size: 0.85rem;
  font-weight: 500;
  color: #555;
  margin-bottom: 6px;
}

.form-group input {
  width: 100%;
  padding: 0.7rem 1rem;
  border: 1px solid #ddd;
  border-radius: 8px;
  font-size: 0.9rem;
  outline: none;
  transition: border 0.2s;
}

.form-group input:focus {
  border-color: #2c3e50;
}

.error-msg {
  background: #fdf0f0;
  color: #e74c3c;
  padding: 8px 12px;
  border-radius: 6px;
  font-size: 0.85rem;
  margin-bottom: 1rem;
}

.btn-login {
  width: 100%;
  background: #2c3e50;
  color: white;
  padding: 0.75rem;
  border: none;
  border-radius: 8px;
  font-size: 0.95rem;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s;
}

.btn-login:hover {
  background: #1a252f;
}

.login-footer {
  text-align: center;
  margin-top: 1.5rem;
  font-size: 0.85rem;
  color: #888;
}

.login-footer a {
  color: #2c3e50;
  font-weight: 500;
  text-decoration: none;
}
</style>