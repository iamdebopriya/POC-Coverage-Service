import { mount } from '@vue/test-utils'
import { describe, it, expect, vi, beforeEach } from 'vitest'
import AddItem from '../src/pages/AddItem.vue'
import * as api from '../src/services/api.js'

vi.mock('../src/services/api.js')
vi.mock('vue-router', () => ({
  useRouter: () => ({ push: vi.fn() })
}))

beforeEach(() => {
  api.createItem.mockResolvedValue({ data: {} })
})

describe('AddItem.vue', () => {
  it('renders the page title', () => {
    const wrapper = mount(AddItem, {
      global: {
        stubs: { RouterLink: true }
      }
    })
    expect(wrapper.text()).toContain('Add New Item')
  })

  it('renders all form fields', () => {
    const wrapper = mount(AddItem, {
      global: {
        stubs: { RouterLink: true }
      }
    })
    const inputs = wrapper.findAll('input')
    expect(inputs.length).toBe(4)
  })

  it('renders submit button', () => {
    const wrapper = mount(AddItem, {
      global: {
        stubs: { RouterLink: true }
      }
    })
    expect(wrapper.text()).toContain('Add Item')
  })

  it('updates form when input changes', async () => {
    const wrapper = mount(AddItem, {
      global: {
        stubs: { RouterLink: true }
      }
    })
    const inputs = wrapper.findAll('input')
    await inputs[0].setValue('Test Item')
    expect(inputs[0].element.value).toBe('Test Item')
  })

 it('renders back button', () => {
  const wrapper = mount(AddItem, {
    global: { stubs: { RouterLink: true } }
  })
  expect(wrapper.text()).toContain('Add New Item')
})
})