<template>
  <div class="register-page">
    <div class="register-card">
      <div class="register-header">
        <h1>Inventory Tracker</h1>
        <p>Create your store account</p>
      </div>

      <form @submit.prevent="handleRegister">
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

        <div class="form-group">
          <label>Confirm Password</label>
          <input
            v-model="form.confirm_password"
            type="password"
            placeholder="Confirm password"
            required
          />
        </div>

        <div class="error-msg" v-if="error">{{ error }}</div>
        <div class="success-msg" v-if="success">{{ success }}</div>

        <button type="submit" class="btn-register">Create Account</button>
      </form>

      <div class="register-footer">
        Already have an account?
        <router-link to="/login">Sign in here</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { register } from '../services/api.js'

const router = useRouter()
const error = ref('')
const success = ref('')

const form = ref({
  store_name: '',
  password: '',
  confirm_password: ''
})

const handleRegister = async () => {
  try {
    error.value = ''
    success.value = ''

    if (form.value.password !== form.value.confirm_password) {
      error.value = 'Passwords do not match'
      return
    }

    await register({
      store_name: form.value.store_name,
      password: form.value.password
    })

    success.value = 'Account created successfully! Redirecting to login...'
    setTimeout(() => {
      router.push('/login')
    }, 1500)

  } catch (err) {
    error.value = 'Store name already exists'
  }
}
</script>

<style scoped>
.register-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f5f6fa;
}

.register-card {
  background: white;
  border-radius: 12px;
  padding: 2.5rem;
  width: 100%;
  max-width: 420px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.08);
}

.register-header {
  text-align: center;
  margin-bottom: 2rem;
}

.register-header h1 {
  font-size: 1.6rem;
  font-weight: 700;
  color: #2c3e50;
  margin-bottom: 6px;
}

.register-header p {
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

.success-msg {
  background: #e8f8f0;
  color: #27ae60;
  padding: 8px 12px;
  border-radius: 6px;
  font-size: 0.85rem;
  margin-bottom: 1rem;
}

.btn-register {
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

.btn-register:hover {
  background: #1a252f;
}

.register-footer {
  text-align: center;
  margin-top: 1.5rem;
  font-size: 0.85rem;
  color: #888;
}

.register-footer a {
  color: #2c3e50;
  font-weight: 500;
  text-decoration: none;
}
</style>