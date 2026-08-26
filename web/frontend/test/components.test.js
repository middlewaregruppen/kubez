import { config, flushPromises, shallowMount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import axios from 'axios'
import ApiEndpoint from '@/components/ApiEndpoint.vue'
import Instructions from '@/components/Instructions.vue'

vi.mock('axios', () => ({
  default: { get: vi.fn() }
}))

config.global.renderStubDefaultSlot = true
config.global.stubs = {
  VSelect: { template: '<div><slot /></div>' },
  VTextField: { template: '<div><slot /></div>' }
}

describe('ApiEndpoint', () => {
  const props = {
    namespace: 'tools',
    name: 'echo',
    mindelay: '10',
    maxdelay: '25',
    requestrate: '2500',
    responsetype: 'echo'
  }

  it('presents its endpoint summary and computed delay', () => {
    const wrapper = shallowMount(ApiEndpoint, { props })

    expect(wrapper.text()).toContain('echo')
    expect(wrapper.text()).toContain('namespace: tools')
    expect(wrapper.vm.prettyBytesRate(wrapper.vm.aep.requestrate)).toBe('2.5 kB/s')
    expect(wrapper.vm.delayRange).toBe('10-25 ms')
  })

  it('starts expanded when requested', () => {
    const wrapper = shallowMount(ApiEndpoint, { props: { ...props, expanded: true } })
    expect(wrapper.vm.collapsed).toBe(false)
  })

  it('emits the edited endpoint and collapses on save', async () => {
    const wrapper = shallowMount(ApiEndpoint, { props })
    await wrapper.setData({ collapsed: false })
    wrapper.vm.aep.mindelay = '15'

    wrapper.vm.saveAndCollapse()

    expect(wrapper.emitted('update')[0][0]).toMatchObject({ name: 'echo', mindelay: '15' })
    expect(wrapper.vm.collapsed).toBe(true)
  })
})

describe('Instructions', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('loads and renders bundled markdown documentation', async () => {
    axios.get.mockResolvedValue({ data: '# Kubez help\n\nUse `curl`.' })
    const wrapper = shallowMount(Instructions, { props: { href: '/help.md' } })
    await flushPromises()

    expect(axios.get).toHaveBeenCalledWith('/help.md')
    expect(wrapper.vm.loading).toBe(false)
    expect(wrapper.find('article').html()).toContain('<h1>Kubez help</h1>')
    expect(wrapper.find('article').html()).toContain('<code>curl</code>')
  })

  it('shows a useful request error instead of documentation', async () => {
    axios.get.mockRejectedValue({ response: { data: 'Not found' } })
    const wrapper = shallowMount(Instructions, { props: { href: '/missing.md' } })
    await flushPromises()

    expect(wrapper.vm.loading).toBe(false)
    expect(wrapper.text()).toContain('Not found')
    expect(wrapper.find('article').exists()).toBe(false)
  })
})
