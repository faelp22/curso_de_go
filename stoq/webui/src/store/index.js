import { store } from 'quasar/wrappers'
import { createStore } from 'vuex'

import getters from './getters'
import mutations from './mutations'
import actions from './actions'
import state from './state'

import Modules from '../modules'

/*
 * If not building with SSR mode, you can
 * directly export the Store instantiation;
 *
 * The function below can be async too; either use
 * async/await or return a Promise which resolves
 * with the Store instance.
 */

export default store(function (/* { ssrContext } */) {
  const Store = createStore({
    modules: {
      ...Modules.modules
    },

    getters, mutations, actions, state,

    // enable strict mode (adds overhead!)
    // for dev mode and --debug builds only
    strict: process.env.DEBUGGING
  })

  return Store
})
