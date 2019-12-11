import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import VueResource from 'vue-resource'
import router from './router'
import store from './store'
import axios from 'axios';


Vue.config.productionTip = false
setInterval(function () {
  store.commit('SET_STATUS', -1)
  axios.get('/kubez/info').then(res => {
    store.commit('SET_INFO', res.data)
    store.commit('SET_STATUS', res.status)
  }).catch(
    err =>{
      store.commit('SET_STATUS', err)
    }
  )
}, 1000)

new Vue({
  vuetify,
  router,
  store,
  VueResource,
  render: h => h(App)
}).$mount('#app')
