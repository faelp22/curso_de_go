import Products from './Domains/Products'

const modules = {
  Products: Products.store
}

const routes = [
  ...Products.routes
]

export default {
  modules,
  routes
}
