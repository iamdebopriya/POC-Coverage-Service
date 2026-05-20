import { mount } from '@vue/test-utils'
import { describe, it, expect, vi, beforeEach } from 'vitest'
import Login from '../src/pages/Login.vue'
import * as api from '../src/services/api.js'

vi.mock('../src/services/api.js')
vi.mock('vue-router', () => ({
  useRouter: () => ({ push: vi.fn() })
}))

beforeEach(() => {
  api.login.mockResolvedValue({
    data: { id: 1, store_name: 'TestStore' }
  })
})

describe('Login.vue', () => {
  it('renders the page title', () => {
    const wrapper = mount(Login, {
      global: { stubs: { RouterLink: true } }
    })
    expect(wrapper.text()).toContain('Inventory Tracker')
  })

  it('renders store name input', () => {
    const wrapper = mount(Login, {
      global: { stubs: { RouterLink: true } }
    })
    const inputs = wrapper.findAll('input')
    expect(inputs.length).toBe(2)
  })

  it('renders sign in button', () => {
    const wrapper = mount(Login, {
      global: { stubs: { RouterLink: true } }
    })
    expect(wrapper.text()).toContain('Sign In')
  })

  it('updates form when input changes', async () => {
    const wrapper = mount(Login, {
      global: { stubs: { RouterLink: true } }
    })
    const inputs = wrapper.findAll('input')
    await inputs[0].setValue('TestStore')
    expect(inputs[0].element.value).toBe('TestStore')
  })

  it('renders register link', () => {
    const wrapper = mount(Login, {
      global: { stubs: { RouterLink: true } }
    })
    expect(wrapper.text()).toContain("Don't have an account?")
  })
})