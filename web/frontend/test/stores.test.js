import { flushPromises } from '@vue/test-utils'
import { createPinia, setActivePinia } from 'pinia'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'
import axios from 'axios'
import { useApiEndpointStore } from '@/stores/apiendpoint'
import { useInfoStore } from '@/stores/info'
import { useKbzK8sStore } from '@/stores/kbzk8s'

vi.mock('axios', () => ({
  default: {
    get: vi.fn(),
    post: vi.fn(),
    put: vi.fn()
  }
}))

describe('Pinia stores', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
    vi.clearAllMocks()
  })

  afterEach(() => {
    vi.useRealTimers()
  })

  it('sets info fields, getters, and defaults for missing fields', () => {
    const store = useInfoStore()
    store.setInfo({ cGroup: { cpu: 2 }, hostname: 'pod-a' })

    expect(store.cgroup).toEqual({ cpu: 2 })
    expect(store.hostname).toBe('pod-a')
    expect(store.headers).toEqual({})
    expect(store.requestInfo).toEqual({})
    expect(store.updateTime).toEqual(expect.any(Number))
  })

  it('sets and clears snack notifications', () => {
    const store = useInfoStore()
    store.setSnack({ message: 'failed', color: 'error' })
    expect(store.snack).toEqual({ message: 'failed', color: 'error', show: true })

    store.clearSnack()
    expect(store.snack).toEqual({ message: '', color: 'success', show: false })
  })

  it('polls info immediately and once per second', async () => {
    vi.useFakeTimers()
    axios.get.mockResolvedValue({ data: { hostname: 'pod-a' }, status: 200 })
    const store = useInfoStore()

    store.startPolling()
    await flushPromises()
    expect(axios.get).toHaveBeenCalledTimes(1)
    expect(store.status).toBe(200)

    await vi.advanceTimersByTimeAsync(1000)
    expect(axios.get).toHaveBeenCalledTimes(2)
  })

  it('derives pod IDs without mutating API state', () => {
    const store = useKbzK8sStore()
    store.podinfo = [{ namespace: 'tools', name: 'kubez-1' }]

    expect(store.podInfo).toEqual([{ namespace: 'tools', name: 'kubez-1', kbzId: 'tools-kubez-1' }])
    expect(store.podinfo[0]).not.toHaveProperty('kbzId')
  })

  it('flattens container states and tolerates missing statuses', () => {
    const store = useKbzK8sStore()
    store.podinfo = [
      { namespace: 'tools', name: 'empty' },
      {
        namespace: 'tools',
        name: 'kubez-1',
        containerInfo: { app: { image: 'kubez:latest' } },
        status: {
          containerStatuses: [{ name: 'app', namespace: 'tools', state: { waiting: { reason: 'Starting', message: 'soon' } } }]
        }
      }
    ]

    expect(store.containerStatuses).toHaveLength(1)
    expect(store.containerStatuses[0]).toMatchObject({
      kbzPod: 'kubez-1',
      kbzId: 'tools-kubez-1-app',
      kbzState: 'waiting',
      kbzReason: 'Starting',
      kbzReasonMessage: 'soon',
      containerInfo: { image: 'kubez:latest' }
    })
  })

  it('fetches pod information', async () => {
    const pods = [{ namespace: 'tools', name: 'kubez-1' }]
    axios.get.mockResolvedValue({ data: pods })
    const store = useKbzK8sStore()

    await store.fetchPodInfo()
    expect(axios.get).toHaveBeenCalledWith('/kubez/kbzk8s/podlist')
    expect(store.podinfo).toEqual(pods)
  })

  it('surfaces pod request failures with API context', async () => {
    axios.get.mockRejectedValue(new Error('offline'))
    const store = useKbzK8sStore()

    await expect(store.fetchPodInfo()).rejects.toThrow('API Error: offline')
  })

  it('tracks endpoint loading and stores a successful response', async () => {
    let resolveRequest
    axios.get.mockReturnValue(new Promise(resolve => { resolveRequest = resolve }))
    const store = useApiEndpointStore()

    const request = store.fetchAPIEndpoints()
    expect(store.loading).toBe(true)
    resolveRequest({ data: [{ name: 'echo' }] })
    await request

    expect(store.list).toEqual([{ name: 'echo' }])
    expect(store.loading).toBe(false)
    expect(store.error).toBe('')
  })

  it('clears endpoints and records API response failures', async () => {
    axios.get.mockRejectedValue({ response: { data: 'service unavailable' } })
    const store = useApiEndpointStore()
    store.apis = [{ name: 'old' }]

    await store.fetchAPIEndpoints()
    expect(store.apis).toEqual([])
    expect(store.error).toBe('service unavailable')
    expect(store.loading).toBe(false)
  })

  it('reports endpoint create and update failures through the info snack', async () => {
    axios.post.mockRejectedValue(new Error('create failed'))
    axios.put.mockRejectedValue(new Error('update failed'))
    const store = useApiEndpointStore()
    const infoStore = useInfoStore()

    await store.createNewAPI({ name: 'echo' })
    expect(axios.post).toHaveBeenCalledWith('/kubez/apicc/', { name: 'echo' })
    expect(infoStore.snack).toMatchObject({ message: 'Error: create failed', color: 'error', show: true })

    await store.updateAPI({ name: 'echo' })
    expect(axios.put).toHaveBeenCalledWith('/kubez/apicc/echo', { name: 'echo' })
    expect(infoStore.snack.message).toBe('Error: update failed')
  })
})
