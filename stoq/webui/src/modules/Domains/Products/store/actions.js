import config from '../../../../config';
import Http from '../../../Http';


const setProductAction = async (store, obj) => {
    try {

        if(obj.hasOwnProperty('usuario')){
            store.commit('SET_PRODUCT_MUTATION', { dados: obj.usuario });
            return;
        }

        let url = config.apiPath + 'api/v1/product/' + obj.id;
        const { data } = await Http.getData(url);

        if(data.hasOwnProperty('product')){
            store.commit('SET_PRODUCT_MUTATION', { dados: data['product'] });
        }

    } catch (e) {
        console.log('setProductAction', e);
    }
};

const setListProductsAction = async (store, obj) => {
    try {

        let url = config.apiPath + 'api/v1/products';
        const { data } = await Http.getData(url);

        if(data.hasOwnProperty('list')){
            store.commit('SET_LIST_PRODUCTS_MUTATION', { lista: data['list'] });
        }
    } catch (e) {
        console.log('setListProductsAction', e);
    }
};

const setProductsInListAction = async (store, obj) => {
    try {
        if(obj.hasOwnProperty('product')){
            store.commit('SET_PRODUCTS_IN_LIST_MUTATION', { dados: obj.product });
            return;
        }

        let url = config.apiPath + 'api/v1/product/' + obj.id;
        const { data } = await Http.getData(url);

        if(data.hasOwnProperty('product')){
            store.commit('SET_PRODUCTS_IN_LIST_MUTATION', { dados: data['usuario'] });
        }
    } catch (e) {
        console.log('setProductsInListAction', e);
    }
};


const deleteProductAction = async (store, obj) => {
    try {

      let url = config.apiPath + 'api/v1/product/' + store.state.product.id;
        await Http.deleteData(url).then(function () {
            store.commit('SET_PRODUCT_MUTATION', { dados: {} });
        }).catch(function (error) {
            console.log(error);
        });
    } catch (e) {
        console.log('deleteProductAction', e);
    }
};

const iniciarProductAction = async (store, obj) => {
  let hash = randomHash(64)
  let web_service = {
    hash: hash,
    path: "http://localhost:8000/",
    name: "",
    image: "",
    created: "2022-03-05 19:57:30",
    count: 0,
    send: false,
    found: false,
    jwt_check: false,
    db_auto: false,
  };
  store.commit('SET_PRODUCT_MUTATION', {dados: web_service});
};

export default {
    setProductAction,
    setListProductsAction,
    setProductsInListAction,
    deleteProductAction,
    iniciarProductAction,
};
