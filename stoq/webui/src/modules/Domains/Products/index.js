import Products from './component/ProductsComponent.vue'

import store from './store';

const routes = [{
        name: 'ListProducts',
        path: '/products',
        component: Products
    },
    {
        name: 'CreateProduct',
        path: '/create-product',
        component: Products
    },
    {
        name: 'UpdateProduct',
        path: '/update-product/:id',
        component: Products
    },
];

export default {
    namespaced: true,
    routes,
    store
};
