import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/api'
  },
  {
    path: '/compute-stats',
    name: 'ComputeStats',

    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "ComputeStats" */ '../views/ComputeStats.vue')
  },
  {
    path: '/network',
    name: 'Network',

    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "ComputeStats" */ '../views/Network.vue')
  },
  {
    path: '/api',
    name: 'API',

    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "ComputeStats" */ '../views/Api.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
