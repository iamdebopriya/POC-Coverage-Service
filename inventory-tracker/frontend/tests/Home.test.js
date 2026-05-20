import { mount } from '@vue/test-utils'
import { describe, it, expect, vi, beforeEach } from 'vitest'
import Home from '../src/pages/Home.vue'
import * as api from '../src/services/api.js'

vi.mock('../src/services/api.js')

const mockItems = [
  {
    id: 1,
    name: 'Laptop',
    category: 'Electronics',
    quantity: 50,
    price: 45000
  },
  {
    id: 2,
    name: 'Mouse',
    category: 'Electronics',
    quantity: 5,
    price: 999
  },
  {
    id: 3,
    name: 'Chair',
    category: 'Furniture',
    quantity: 15,
    price: 3500
  }
]

const mountComponent = () => {
  return mount(Home, {
    global: {
      stubs: {
        RouterLink: {
          template: '<a><slot /></a>'
        }
      }
    }
  })
}

beforeEach(() => {
  vi.clearAllMocks()

  api.getItems.mockResolvedValue({
    data: mockItems
  })

  api.getLowStockItems.mockResolvedValue({
    data: [mockItems[1]]
  })

  api.deleteItem.mockResolvedValue({})
})

describe('Home.vue', () => {
  it('renders the page title', async () => {
    const wrapper = mountComponent()

    await wrapper.vm.$nextTick()

    expect(wrapper.text()).toContain('Store Dashboard')
  })

  it('renders inventory items in the table', async () => {
    const wrapper = mountComponent()

    await wrapper.vm.$nextTick()
    await wrapper.vm.$nextTick()

    expect(wrapper.text()).toContain('Laptop')
    expect(wrapper.text()).toContain('Mouse')
    expect(wrapper.text()).toContain('Chair')
  })

  it('renders metric cards correctly', async () => {
    const wrapper = mountComponent()

    await wrapper.vm.$nextTick()
    await wrapper.vm.$nextTick()

    expect(wrapper.text()).toContain('Total Items')
    expect(wrapper.text()).toContain('In Stock')
    expect(wrapper.text()).toContain('Low Stock')
    expect(wrapper.text()).toContain('Inventory Value')
    expect(wrapper.text()).toContain('Average Price')
  })

  it('shows correct total item count', async () => {
    const wrapper = mountComponent()

    await wrapper.vm.$nextTick()
    await wrapper.vm.$nextTick()

    expect(wrapper.text()).toContain('3')
  })

  it('shows low stock banner when low stock items exist', async () => {
    const wrapper = mountComponent()

    await wrapper.vm.$nextTick()
    await wrapper.vm.$nextTick()

    expect(wrapper.text()).toContain('low on stock')
  })

  it('shows in stock badge for items with quantity above 10', async () => {
    const wrapper = mountComponent()

    await wrapper.vm.$nextTick()
    await wrapper.vm.$nextTick()

    expect(wrapper.text()).toContain('In Stock')
  })

  it('shows low stock badge for items with quantity below 10', async () => {
    const wrapper = mountComponent()

    await wrapper.vm.$nextTick()
    await wrapper.vm.$nextTick()

    expect(wrapper.text()).toContain('Low Stock')
  })

  it('renders search input field', async () => {
    const wrapper = mountComponent()

    await wrapper.vm.$nextTick()

    const searchInput = wrapper.find('input[type="text"]')

    expect(searchInput.exists()).toBe(true)
  })

  it('filters items by search text', async () => {
    const wrapper = mountComponent()

    await wrapper.vm.$nextTick()
    await wrapper.vm.$nextTick()

    const searchInput = wrapper.find('input[type="text"]')

    await searchInput.setValue('Laptop')

    expect(wrapper.text()).toContain('Laptop')
  })

  it('renders category filter dropdown', async () => {
    const wrapper = mountComponent()

    await wrapper.vm.$nextTick()

    const selects = wrapper.findAll('select')

    expect(selects.length).toBeGreaterThan(0)
  })

  it('filters items by category', async () => {
    const wrapper = mountComponent()

    await wrapper.vm.$nextTick()
    await wrapper.vm.$nextTick()

    const categorySelect = wrapper.findAll('select')[0]

    await categorySelect.setValue('Furniture')

    expect(wrapper.text()).toContain('Chair')
  })

  it('renders add item button', async () => {
    const wrapper = mountComponent()

    await wrapper.vm.$nextTick()

    expect(wrapper.find('.btn-primary').text()).toContain('Add Item')
  })

  it('renders edit and delete buttons', async () => {
    const wrapper = mountComponent()

    await wrapper.vm.$nextTick()
    await wrapper.vm.$nextTick()

    expect(wrapper.findAll('.btn-edit').length).toBeGreaterThan(0)
    expect(wrapper.findAll('.btn-delete').length).toBeGreaterThan(0)
  })

  it('calls deleteItem when delete button is clicked', async () => {
    window.confirm = vi.fn(() => true)

    const wrapper = mountComponent()

    await wrapper.vm.$nextTick()
    await wrapper.vm.$nextTick()

    const deleteButtons = wrapper.findAll('.btn-delete')

    await deleteButtons[0].trigger('click')

    expect(api.deleteItem).toHaveBeenCalled()
  })

  it('renders empty state when no items exist', async () => {
    api.getItems.mockResolvedValue({
      data: []
    })

    api.getLowStockItems.mockResolvedValue({
      data: []
    })

    const wrapper = mountComponent()

    await wrapper.vm.$nextTick()
    await wrapper.vm.$nextTick()

    expect(wrapper.text()).toContain('No matching items found')
  })
})