import { describe, expect, it } from 'vitest'
import router from '@/router'

describe('router', () => {
  it('registers all primary application routes', () => {
    const paths = router.getRoutes().map(route => route.path)
    expect(paths).toEqual(expect.arrayContaining(['/api', '/compute-stats', '/network', '/pods']))
  })

  it('redirects the root route to the API page', () => {
    const root = router.getRoutes().find(route => route.path === '/')
    expect(root.redirect).toBe('/api')
  })

  it('resolves named routes to hash URLs', () => {
    expect(router.resolve({ name: 'API' }).href).toBe('#/api')
    expect(router.resolve({ name: 'ComputeStats' }).href).toBe('#/compute-stats')
    expect(router.resolve({ name: 'Network' }).href).toBe('#/network')
    expect(router.resolve({ name: 'Pods' }).href).toBe('#/pods')
  })
})
