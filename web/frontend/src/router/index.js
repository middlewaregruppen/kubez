import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/api'
  },
  {
    path: '/compute-stats',
    name: 'ComputeStats',
    component: () => import('../views/ComputeStats.vue')
  },
  {
    path: '/network',
    name: 'Network',
    component: () => import('../views/Network.vue')
  },
  {
    path: '/api',
    name: 'API',
    component: () => import('../views/Api.vue')
  },
  {
    path: '/pods',
    name: 'Pods',
    component: () => import('../views/Pods.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
