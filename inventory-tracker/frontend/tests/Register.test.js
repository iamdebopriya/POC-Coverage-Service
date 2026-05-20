import { mount } from '@vue/test-utils'
import { describe, it, expect, vi, beforeEach } from 'vitest'
import Register from '../src/pages/Register.vue'
import * as api from '../src/services/api.js'

vi.mock('../src/services/api.js')
vi.mock('vue-router', () => ({
  useRouter: () => ({ push: vi.fn() })
}))

beforeEach(() => {
  api.register.mockResolvedValue({
    data: { id: 1, store_name: 'TestStore' }
  })
})

describe('Register.vue', () => {
  it('renders the page title', () => {
    const wrapper = mount(Register, {
      global: { stubs: { RouterLink: true } }
    })
    expect(wrapper.text()).toContain('Inventory Tracker')
  })

  it('renders all form fields', () => {
    const wrapper = mount(Register, {
      global: { stubs: { RouterLink: true } }
    })
    const inputs = wrapper.findAll('input')
    expect(inputs.length).toBe(3)
  })

  it('renders create account button', () => {
    const wrapper = mount(Register, {
      global: { stubs: { RouterLink: true } }
    })
    expect(wrapper.text()).toContain('Create Account')
  })

  it('updates form when input changes', async () => {
    const wrapper = mount(Register, {
      global: { stubs: { RouterLink: true } }
    })
    const inputs = wrapper.findAll('input')
    await inputs[0].setValue('TestStore')
    expect(inputs[0].element.value).toBe('TestStore')
  })

  it('shows error when passwords do not match', async () => {
    const wrapper = mount(Register, {
      global: { stubs: { RouterLink: true } }
    })
    const inputs = wrapper.findAll('input')
    await inputs[1].setValue('password123')
    await inputs[2].setValue('differentpassword')
    await wrapper.find('form').trigger('submit')
    await wrapper.vm.$nextTick()
    expect(wrapper.text()).toContain('Passwords do not match')
  })
})