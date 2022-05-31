import Modules from '../modules'

const routes = [

  {
    path: '/',
    component: () => import('layouts/LoginLayout.vue'),
    children: [{
      path: '',
      name: 'Login',
      component: () => import('pages/Login.vue')
    }]
  },

  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [{
        path: '',
        name: 'Index',
        component: () => import('pages/Index.vue')
      },
      ...Modules.routes,
    ]
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/Error404.vue')
  }
]

export default routes
