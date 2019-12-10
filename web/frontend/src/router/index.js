import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'



Vue.use(VueRouter)


const routes = [
  {
    path: '/',
    name: 'home',
    component: Home
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
  }
]

const router = new VueRouter({
  routes
})

export default router
