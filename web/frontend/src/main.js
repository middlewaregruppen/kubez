import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import VueResource from 'vue-resource'
import router from './router'

Vue.config.productionTip = false

new Vue({
  vuetify,
  router,
  VueResource,
  render: h => h(App)
}).$mount('#app')
