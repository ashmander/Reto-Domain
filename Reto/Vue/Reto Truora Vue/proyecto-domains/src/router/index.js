import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/domain',
    name: 'Domain',
    component: () => import('../views/DomainSearch.vue')
  },
  {
    path: '/endpoints',
    name: 'Endpoints',
    component: () => import('../views/Enpoints.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
