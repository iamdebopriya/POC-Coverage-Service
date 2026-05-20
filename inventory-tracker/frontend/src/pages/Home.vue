<template>
  <div class="home">

    <!-- Header -->
    <div class="page-header">

  <div class="header-content">

    <div class="welcome-badge">
      Store Dashboard
    </div>

    <h1 class="page-title">
      Welcome back,
      <span class="store-name">
        {{ user?.store_name }}
      </span>
    </h1>

    <p class="page-subtitle">
      Manage your inventory, monitor stock levels,
      and organize your store items efficiently.
    </p>

  </div>

      <router-link to="/add" class="btn-primary">
        + Add Item
      </router-link>
    </div>

    <!-- Metrics -->
    <div class="metrics-grid">

      <div class="metric-card">
        <div class="metric-icon neutral">
          <i class="fas fa-boxes-stacked"></i>
        </div>

        <div class="metric-info">
          <div class="metric-label">Total Items</div>
          <div class="metric-value">{{ items.length }}</div>
        </div>
      </div>

      <div class="metric-card">
        <div class="metric-icon neutral">
          <i class="fas fa-circle-check"></i>
        </div>

        <div class="metric-info">
          <div class="metric-label">In Stock</div>
          <div class="metric-value">{{ inStockCount }}</div>
        </div>
      </div>

      <div class="metric-card">
        <div class="metric-icon neutral">
          <i class="fas fa-triangle-exclamation"></i>
        </div>

        <div class="metric-info">
          <div class="metric-label">Low Stock</div>
          <div class="metric-value">{{ lowStockItems.length }}</div>
        </div>
      </div>

      <div class="metric-card">
        <div class="metric-icon neutral">
          <i class="fas fa-indian-rupee-sign"></i>
        </div>

        <div class="metric-info">
          <div class="metric-label">Inventory Value</div>
          <div class="metric-value">₹{{ totalValue }}</div>
        </div>
      </div>

    </div>

   
    <div class="extra-metrics">

      <div class="mini-metric">
        <span class="mini-label">Average Price</span>
        <strong>₹{{ averagePrice }}</strong>
      </div>

      <div class="mini-metric">
        <span class="mini-label">Highest Stock</span>
        <strong>{{ highestStock }}</strong>
      </div>

      <div class="mini-metric">
        <span class="mini-label">Categories</span>
        <strong>{{ categories.length }}</strong>
      </div>

      <div class="mini-metric">
        <span class="mini-label">Filtered Results</span>
        <strong>{{ filteredItems.length }}</strong>
      </div>

    </div>

    <div
      class="low-stock-banner"
      v-if="lowStockItems.length > 0"
    >
      {{ lowStockItems.length }} item(s) are low on stock!
    </div>

    <div class="toolbar">

      <div class="search-box">
        <i class="fas fa-search"></i>

        <input
          v-model="searchText"
          type="text"
          placeholder="Search items..."
        />
      </div>

      <select v-model="selectedCategory" class="filter-select">
        <option value="">All Categories</option>

        <option
          v-for="category in categories"
          :key="category"
          :value="category"
        >
          {{ category }}
        </option>
      </select>

      <select v-model="stockFilter" class="filter-select">
        <option value="">All Stock</option>
        <option value="instock">In Stock</option>
        <option value="low">Low Stock</option>
      </select>

    </div>


    <div class="table-wrapper">

      <table class="inventory-table">

        <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Category</th>
            <th>Quantity</th>
            <th>Price</th>
            <th>Status</th>
            <th>Actions</th>
          </tr>
        </thead>

        <tbody>

          <tr
            v-for="item in filteredItems"
            :key="item.id"
          >
            <td>{{ item.id }}</td>
            <td>{{ item.name }}</td>
            <td>{{ item.category }}</td>
            <td>{{ item.quantity }}</td>
            <td>₹{{ item.price }}</td>

            <td>
              <span
                class="badge"
                :class="item.quantity < 10 ? 'badge-danger' : 'badge-success'"
              >
                {{ item.quantity < 10 ? 'Low Stock' : 'In Stock' }}
              </span>
            </td>

            <td class="actions">

              <router-link
                :to="`/edit/${item.id}`"
                class="btn-edit"
              >
                Edit
              </router-link>

              <button
                @click="handleDelete(item.id)"
                class="btn-delete"
              >
                Delete
              </button>

            </td>
          </tr>

          <tr v-if="filteredItems.length === 0">
            <td colspan="7" class="empty-state">
              No matching items found.
            </td>
          </tr>

        </tbody>

      </table>

    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import {
  getItems,
  deleteItem,
  getLowStockItems
} from '../services/api.js'

const items = ref([])
const lowStockItems = ref([])
const user = ref(JSON.parse(localStorage.getItem('user') || 'null'))

const searchText = ref('')
const selectedCategory = ref('')
const stockFilter = ref('')

const fetchItems = async () => {
  const res = await getItems()
  items.value = res.data
}

const fetchLowStock = async () => {
  const res = await getLowStockItems()
  lowStockItems.value = res.data
}

const handleDelete = async (id) => {
  if (confirm('Are you sure you want to delete this item?')) {
    await deleteItem(id)

    await fetchItems()
    await fetchLowStock()
  }
}

const inStockCount = computed(() => {
  return items.value.filter(i => i.quantity >= 10).length
})

const totalValue = computed(() => {
  return items.value
    .reduce((sum, i) => sum + (i.price * i.quantity), 0)
    .toLocaleString()
})

const averagePrice = computed(() => {
  if (items.value.length === 0) return 0

  const total = items.value.reduce((sum, i) => sum + i.price, 0)

  return (total / items.value.length).toFixed(0)
})

const highestStock = computed(() => {
  if (items.value.length === 0) return 0

  return Math.max(...items.value.map(i => i.quantity))
})

const categories = computed(() => {
  return [...new Set(items.value.map(i => i.category))]
})

const filteredItems = computed(() => {
  return items.value.filter(item => {

    const matchesSearch =
      item.name.toLowerCase().includes(searchText.value.toLowerCase())

    const matchesCategory =
      !selectedCategory.value ||
      item.category === selectedCategory.value

    const matchesStock =
      !stockFilter.value ||
      (stockFilter.value === 'low' && item.quantity < 10) ||
      (stockFilter.value === 'instock' && item.quantity >= 10)

    return (
      matchesSearch &&
      matchesCategory &&
      matchesStock
    )
  })
})

onMounted(() => {
  fetchItems()
  fetchLowStock()
})
</script>

<style scoped>
.home {
  padding: 0.5rem;
}



.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 1.5rem;
  margin-bottom: 2rem;
  padding: 1.5rem;
  border-radius: 24px;
  background:
    linear-gradient(
      135deg,
      #f8fafc,
      #eef2ff
    );
  border: 1px solid #e5e7eb;
}

.header-content {
  display: flex;
  flex-direction: column;
}

.welcome-badge {
  width: fit-content;
  padding: 6px 14px;
  border-radius: 999px;
  background: #ede9fe;
  color: #6d28d9;
  font-size: 0.75rem;
  font-weight: 700;
  margin-bottom: 0.9rem;
  letter-spacing: 0.04em;
}

.page-title {
  font-size: 2rem;
  font-weight: 800;
  color: #111827;
  margin-bottom: 0.5rem;
  line-height: 1.2;
}

.store-name {
  color: #07095f;
  font-size: 2.1rem;
  font-weight: 1500;
}

.page-subtitle {
  color: #6b7280;
  font-size: 0.95rem;
  max-width: 650px;
  line-height: 1.6;
}

.page-header h2 {
  font-size: 2rem;
  color: #111827;
  font-weight: 700;
}

.btn-primary {
  background: #374151;
  color: white;
  padding: 0.75rem 1.3rem;
  border-radius: 12px;
  text-decoration: none;
  font-size: 0.92rem;
  font-weight: 600;
  transition: 0.2s ease;
}

.btn-primary:hover {
  background: #1f2937;
}


.metrics-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 1rem;
  margin-bottom: 1rem;
}

.metric-card {
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 18px;
  padding: 1.3rem;
  display: flex;
  align-items: center;
  gap: 1rem;
  transition: 0.2s ease;
}

.metric-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(0,0,0,0.06);
}

.metric-icon {
  width: 52px;
  height: 52px;
  border-radius: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f3f4f6;
  color: #374151;
  font-size: 1.1rem;
}

.metric-label {
  font-size: 0.82rem;
  color: #6b7280;
  margin-bottom: 4px;
}

.metric-value {
  font-size: 1.5rem;
  font-weight: 700;
  color: #111827;
}



.extra-metrics {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.mini-metric {
  background: #fafafa;
  border: 1px solid #e5e7eb;
  border-radius: 14px;
  padding: 1rem;
}

.mini-label {
  display: block;
  color: #6b7280;
  font-size: 0.82rem;
  margin-bottom: 6px;
}

.mini-metric strong {
  font-size: 1.2rem;
  color: #111827;
}



.low-stock-banner {
  background: #edc5c5;
  border: 1px solid #d1d5db;
  color: #eb0404;
  padding: 0.9rem 1rem;
  border-radius: 12px;
  margin-bottom: 1.5rem;
  font-size: 0.92rem;
}


.toolbar {
  display: flex;
  gap: 1rem;
  margin-bottom: 1.5rem;
  flex-wrap: wrap;
}

.search-box {
  flex: 1;
  min-width: 240px;
  position: relative;
}

.search-box i {
  position: absolute;
  left: 14px;
  top: 50%;
  transform: translateY(-50%);
  color: #9ca3af;
}

.search-box input {
  width: 100%;
  padding: 0.82rem 1rem 0.82rem 2.5rem;
  border-radius: 12px;
  border: 1px solid #d1d5db;
  font-size: 0.92rem;
  outline: none;
  transition: 0.2s;
}

.search-box input:focus {
  border-color: #6b7280;
}

.filter-select {
  padding: 0.82rem 1rem;
  border-radius: 12px;
  border: 1px solid #d1d5db;
  background: white;
  font-size: 0.92rem;
  color: #374151;
  min-width: 180px;
}


.table-wrapper {
  background: white;
  border-radius: 18px;
  overflow: hidden;
  border: 1px solid #e5e7eb;
}

.inventory-table {
  width: 100%;
  border-collapse: collapse;
}

.inventory-table th {
  background: #f9fafb;
  padding: 1rem;
  text-align: left;
  font-size: 0.86rem;
  color: #6b7280;
  font-weight: 600;
  border-bottom: 1px solid #e5e7eb;
}

.inventory-table td {
  padding: 1rem;
  border-bottom: 1px solid #f3f4f6;
  color: #374151;
  font-size: 0.92rem;
}



.badge {
  padding: 5px 12px;
  border-radius: 999px;
  font-size: 0.78rem;
  font-weight: 600;
}

.badge-success {
  background: #ecfdf5;
  color: #047857;
}

.badge-danger {
  background: #fef2f2;
  color: #b91c1c;
}


.actions {
  display: flex;
  gap: 0.6rem;
}

.btn-edit {
  background: #374151;
  color: white;
  padding: 0.45rem 0.9rem;
  border-radius: 8px;
  text-decoration: none;
  font-size: 0.8rem;
}

.btn-delete {
  background: #ef4444;
  color: white;
  border: none;
  padding: 0.45rem 0.9rem;
  border-radius: 8px;
  cursor: pointer;
  font-size: 0.8rem;
}

.empty-state {
  text-align: center;
  color: #9ca3af;
  padding: 2rem;
}



@media (max-width: 900px) {

  .page-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 1rem;
  }

  .toolbar {
    flex-direction: column;
  }

  .filter-select {
    width: 100%;
  }
}
</style>