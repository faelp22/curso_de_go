import Products from './component/ProductsComponent.vue'
import Product from './component/ProductComponent.vue'

import store from './store'

const routes = [{
  name: 'ListProducts',
  path: '/products',
  component: Products
},
{
  name: 'CreateProduct',
  path: '/create-product',
  component: Product
},
{
  name: 'UpdateProduct',
  path: '/update-product/:id',
  component: Product
}
]

export default {
  namespaced: true,
  routes,
  store
}
