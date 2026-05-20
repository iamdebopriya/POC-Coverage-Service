<template>
  <div class="edit-item">
    <div class="page-header">
      <h2>Edit Item</h2>
      <router-link to="/" class="btn-back">← Back</router-link>
    </div>

    <div class="form-card">
      <form @submit.prevent="handleSubmit">
        <div class="form-row">
          <div class="form-group">
            <label>Item Name</label>
            <input
              v-model="form.name"
              type="text"
              placeholder="Enter item name"
              required
            />
          </div>
          <div class="form-group">
            <label>Category</label>
            <input
              v-model="form.category"
              type="text"
              placeholder="Enter category"
              required
            />
          </div>
        </div>

        <div class="form-row">
          <div class="form-group">
            <label>Quantity</label>
            <input
              v-model="form.quantity"
              type="number"
              placeholder="Enter quantity"
              min="0"
              required
            />
          </div>
          <div class="form-group">
            <label>Price Per Item</label>
            <input
              v-model="form.price"
              type="number"
              placeholder="Enter price"
              min="0"
              step="0.01"
              required
            />
          </div>
        </div>

        <div class="form-actions">
          <button type="submit" class="btn-primary">Update Item</button>
          <router-link to="/inventory" class="btn-cancel">Cancel</router-link>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { getItem, updateItem } from '../services/api.js'

const router = useRouter()
const route = useRoute()

const form = ref({
  name: '',
  category: '',
  quantity: '',
  price: ''
})

onMounted(async () => {
  const res = await getItem(route.params.id)
  form.value = {
    name: res.data.name,
    category: res.data.category,
    quantity: res.data.quantity,
    price: res.data.price
  }
})

const handleSubmit = async () => {
  await updateItem(route.params.id, {
    name: form.value.name,
    category: form.value.category,
    quantity: parseInt(form.value.quantity),
    price: parseFloat(form.value.price)
  })
  router.push('/inventory')
}
</script>

<style scoped>
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}

.page-header h2 {
  font-size: 1.5rem;
  color: #2c3e50;
}

.btn-back {
  color: #2c3e50;
  text-decoration: none;
  font-size: 0.9rem;
}

.form-card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 1px 4px rgba(0,0,0,0.08);
  padding: 2rem;
  max-width: 650px;
  margin: 0 auto;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
  margin-bottom: 1rem;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-group label {
  font-size: 0.9rem;
  font-weight: 500;
  color: #555;
  margin-bottom: 0.4rem;
}

.form-group input {
  width: 100%;
  padding: 0.6rem 0.8rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
  outline: none;
  transition: border 0.2s;
}

.form-group input:focus {
  border-color: #2c3e50;
}

.form-actions {
  display: flex;
  gap: 1rem;
  margin-top: 1.5rem;
}

.btn-primary {
  background-color: #2c3e50;
  color: white;
  padding: 0.6rem 1.5rem;
  border: none;
  border-radius: 6px;
  font-size: 0.9rem;
  cursor: pointer;
}

.btn-cancel {
  background-color: #f0f0f0;
  color: #555;
  padding: 0.6rem 1.5rem;
  border-radius: 6px;
  text-decoration: none;
  font-size: 0.9rem;
}
</style>