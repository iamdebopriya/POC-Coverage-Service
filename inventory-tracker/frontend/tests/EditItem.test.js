import { mount } from '@vue/test-utils'
import { describe, it, expect, vi, beforeEach } from 'vitest'
import EditItem from '../src/pages/EditItem.vue'
import * as api from '../src/services/api.js'

vi.mock('../src/services/api.js')
vi.mock('vue-router', () => ({
  useRouter: () => ({ push: vi.fn() }),
  useRoute: () => ({ params: { id: '1' } })
}))

beforeEach(() => {
  api.getItem.mockResolvedValue({
    data: {
      id: 1,
      name: 'Laptop',
      category: 'Electronics',
      quantity: 50,
      price: 45000
    }
  })
  api.updateItem.mockResolvedValue({ data: {} })
})

describe('EditItem.vue', () => {
  it('renders the page title', () => {
    const wrapper = mount(EditItem, {
      global: {
        stubs: { RouterLink: true }
      }
    })
    expect(wrapper.text()).toContain('Edit Item')
  })

  it('renders all form fields', () => {
    const wrapper = mount(EditItem, {
      global: {
        stubs: { RouterLink: true }
      }
    })
    const inputs = wrapper.findAll('input')
    expect(inputs.length).toBe(4)
  })

  it('renders update button', () => {
    const wrapper = mount(EditItem, {
      global: {
        stubs: { RouterLink: true }
      }
    })
    expect(wrapper.text()).toContain('Update Item')
  })

 it('renders back button', () => {
  const wrapper = mount(EditItem, {
    global: { stubs: { RouterLink: true } }
  })
  expect(wrapper.text()).toContain('Edit Item')
})

  it('loads item data on mount', async () => {
    const wrapper = mount(EditItem, {
      global: {
        stubs: { RouterLink: true }
      }
    })
    await wrapper.vm.$nextTick()
    await wrapper.vm.$nextTick()
    const inputs = wrapper.findAll('input')
    expect(inputs[0].element.value).toBe('Laptop')
  })
})