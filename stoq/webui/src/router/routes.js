import productTable from '../components/products/TableComponent.vue'
import productForm from '../components/products/FormComponent.vue'

const routes = [{
  path: '/',
  component: () => import('layouts/LoginLayout.vue'),
  children: [{
    path: '',
    name: 'Login',
    component: () => import('src/pages/LoginPage.vue')
  }]
},

{
  path: '/',
  component: () => import('layouts/MainLayout.vue'),
  children: [{
    path: '',
    name: 'Index',
    component: () => import('pages/IndexPage.vue')
  },
  {
    name: 'ListProducts',
    path: '/products',
    component: productTable
  },
  {
    name: 'CreateProduct',
    path: '/create-product',
    component: productForm
  },
  {
    name: 'UpdateProduct',
    path: '/update-product/:id',
    component: productForm
  }
  ]
},

// Always leave this as last one,
// but you can also remove it
{
  path: '/:catchAll(.*)*',
  component: () => import('pages/ErrorNotFound.vue')
}
]

export default routes
