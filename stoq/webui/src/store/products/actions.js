import config from '../../config'
import Http from '../../services/Http'

const setProductAction = async (store, obj) => {
  try {
    if (Object.prototype.hasOwnProperty.call(obj, 'product')) {
      store.commit('SET_PRODUCT_MUTATION', { dados: obj.product })
      return
    }

    const url = config.apiPath + 'api/v1/product/' + obj.id
    const { data } = await Http.getData(url)

    if (Object.prototype.hasOwnProperty.call(data, 'id')) {
      store.commit('SET_PRODUCT_MUTATION', { dados: data })
    }
  } catch (e) {
    console.log('setProductAction', e)
  }
}

const setListProductsAction = async (store, obj) => {
  try {
    const url = config.apiPath + 'api/v1/products'
    const { data } = await Http.getData(url)

    if (Object.prototype.hasOwnProperty.call(data, 'list')) {
      store.commit('SET_LIST_PRODUCTS_MUTATION', { lista: data.list })
    }
  } catch (e) {
    console.log('setListProductsAction', e)
  }
}

const setProductsInListAction = async (store, obj) => {
  try {
    if (Object.prototype.hasOwnProperty.call(obj, 'name') && obj.name) {
      if (Object.prototype.hasOwnProperty.call(obj, 'id') && obj.id) {
        const url = config.apiPath + 'api/v1/product/' + obj.id
        await Http.putData(url, obj)
      } else {
        const url = config.apiPath + 'api/v1/product'
        await Http.postData(url, obj)
      }
    }
  } catch (e) {
    console.log('setProductsInListAction', e)
  }
}

const deleteProductAction = async (store, obj) => {
  try {
    const url = config.apiPath + 'api/v1/product/' + store.state.product.id
    await Http.deleteData(url).then(() => {
      store.commit('SET_PRODUCT_MUTATION', { dados: {} })
    }).catch((error) => {
      console.log(error)
    })
  } catch (e) {
    console.log('deleteProductAction', e)
  }
}

const iniciarProductAction = async (store, obj) => {
  const product = {
    name: '',
    code: '',
    price: 0
  }
  store.commit('SET_PRODUCT_MUTATION', { dados: product })
}

export default {
  setProductAction,
  setListProductsAction,
  setProductsInListAction,
  deleteProductAction,
  iniciarProductAction
}
