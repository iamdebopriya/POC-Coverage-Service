import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  headers: {
    'Content-Type': 'application/json'
  }
})

// Add X-User-ID header to every request
api.interceptors.request.use((config) => {
  const user = JSON.parse(localStorage.getItem('user') || '{}')
  if (user.id) {
    config.headers['X-User-ID'] = user.id
  }
  return config
})

// Item APIs
export const getItems = () => api.get('/items')
export const getItem = (id) => api.get(`/items/${id}`)
export const createItem = (item) => api.post('/items', item)
export const updateItem = (id, item) => api.put(`/items/${id}`, item)
export const deleteItem = (id) => api.delete(`/items/${id}`)
export const getLowStockItems = () => api.get('/items/low-stock')

// Auth APIs
export const register = (data) => api.post('/auth/register', data)
export const login = (data) => api.post('/auth/login', data)

// Coverage APIs
export const saveCoverage = (data) => api.post('/coverage', data)
export const getCoverage = (from, to) => {
  if (from && to) {
    return api.get(`/coverage?from=${from}&to=${to}`)
  }
  return api.get('/coverage')
}
export const downloadCoverage = (from, to) => {
  if (from && to) {
    return api.get(`/coverage/download?from=${from}&to=${to}`, { responseType: 'blob' })
  }
  return api.get('/coverage/download', { responseType: 'blob' })
}

export default api