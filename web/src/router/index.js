import Vue from 'vue'
import VueRouter from 'vue-router'
import Dashboard from '../views/dashboard/index.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'dashboard',
    component: Dashboard,
    children: [
      {
        path: 'plugin-list',
        name: '插件列表',
        component: () => import('@/views/dashboard/components/list'),
        meta: { title: '插件列表', icon: 'form' }
      },
      {
        path: 'host-status',
        name: '主机状态',
        component: () => import('@/views/dashboard/components/status'),
        meta: { title: '主机状态', icon: 'form' }
      },
      {
        path: 'plugin-check',
        name: '安全检查',
        component: () => import('@/views/dashboard/components/check'),
        meta: { title: '安全检查', icon: 'form' }
      },
      {
        path: 'about',
        name: '关于',
        component: () => import('@/views/dashboard/components/about'),
        meta: { title: '关于', icon: 'form' }
      }
    ]
  }
  // {
  //   path: '/about',
  //   name: 'About',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  // }
]

const router = new VueRouter({
  routes
})

export default router
