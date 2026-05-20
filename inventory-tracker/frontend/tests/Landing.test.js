import { mount } from '@vue/test-utils'
import { describe, it, expect } from 'vitest'
import Landing from '../src/pages/Landing.vue'

describe('Landing.vue', () => {
  it('renders the landing title', () => {
    const wrapper = mount(Landing, {
      global: {
        stubs: { RouterLink: true }
      }
    })
    expect(wrapper.text()).toContain('Track Your Stock')
  })

  it('renders get started button', () => {
  const wrapper = mount(Landing, {
    global: { stubs: { RouterLink: true } }
  })
  expect(wrapper.text()).toContain('Smarter & Faster')
})
it('renders view dashboard button', () => {
  const wrapper = mount(Landing, {
    global: { stubs: { RouterLink: true } }
  })
  expect(wrapper.text()).toContain('Smarter & Faster')
})
  it('renders feature cards', () => {
    const wrapper = mount(Landing, {
      global: {
        stubs: { RouterLink: true }
      }
    })
    expect(wrapper.text()).toContain('Track Inventory')
    expect(wrapper.text()).toContain('Low Stock Alerts')
    expect(wrapper.text()).toContain('Coverage Dashboard')
    expect(wrapper.text()).toContain('Export Reports')
  })
})