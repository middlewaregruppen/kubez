import Vue from 'vue'
import Vuex from 'vuex'
import info from './modules/info'
//import products from './modules/products'
//import createLogger from '../../../src/plugins/logger'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    info
  },
  strict: true
})