import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { mount, flushPromises } from '@vue/test-utils'
import HelloWorld from '../components/HelloWorld.vue'

describe('HelloWorld.vue', () => {
  beforeEach(() => {
    vi.stubGlobal('fetch', vi.fn())
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('renders title and subtitle', () => {
    const wrapper = mount(HelloWorld)
    expect(wrapper.find('h1').text()).toBe('Hello World App')
    expect(wrapper.find('.subtitle').text()).toBe('Vue 3 + Go backend')
  })

  it('shows "Say Hello" button initially', () => {
    const wrapper = mount(HelloWorld)
    const btn = wrapper.find('button')
    expect(btn.text()).toBe('Say Hello')
    expect(btn.attributes('disabled')).toBeUndefined()
  })

  it('displays message on successful fetch', async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      ok: true,
      json: async () => ({ message: 'Hello, World!', status: 'ok' }),
    })
    vi.stubGlobal('fetch', mockFetch)

    const wrapper = mount(HelloWorld)
    await wrapper.find('button').trigger('click')
    await flushPromises()

    expect(wrapper.find('[data-testid="message"]').text()).toBe('Hello, World!')
    expect(wrapper.find('[data-testid="error"]').exists()).toBe(false)
  })

  it('displays error when fetch fails', async () => {
    const mockFetch = vi.fn().mockRejectedValue(new Error('Network error'))
    vi.stubGlobal('fetch', mockFetch)

    const wrapper = mount(HelloWorld)
    await wrapper.find('button').trigger('click')
    await flushPromises()

    expect(wrapper.find('[data-testid="error"]').text()).toContain('Network error')
    expect(wrapper.find('[data-testid="message"]').exists()).toBe(false)
  })

  it('displays error on non-ok HTTP response', async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      ok: false,
      status: 500,
    })
    vi.stubGlobal('fetch', mockFetch)

    const wrapper = mount(HelloWorld)
    await wrapper.find('button').trigger('click')
    await flushPromises()

    expect(wrapper.find('[data-testid="error"]').text()).toContain('HTTP error: 500')
  })

  it('disables button while loading', async () => {
    let resolve!: (v: unknown) => void
    const mockFetch = vi.fn().mockReturnValue(new Promise(r => { resolve = r }))
    vi.stubGlobal('fetch', mockFetch)

    const wrapper = mount(HelloWorld)
    await wrapper.find('button').trigger('click')
    await vi.waitFor(() => {
      expect(wrapper.find('button').attributes('disabled')).toBeDefined()
    })

    resolve({ ok: true, json: async () => ({ message: 'Hello, World!', status: 'ok' }) })
    await flushPromises()
    expect(wrapper.find('button').attributes('disabled')).toBeUndefined()
  })

  it('calls /api/hello endpoint', async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      ok: true,
      json: async () => ({ message: 'Hello, World!', status: 'ok' }),
    })
    vi.stubGlobal('fetch', mockFetch)

    const wrapper = mount(HelloWorld)
    await wrapper.find('button').trigger('click')
    await flushPromises()

    expect(mockFetch).toHaveBeenCalledWith('/api/hello')
  })
})
