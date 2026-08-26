import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router'
import { useInfoStore } from './stores/info'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(vuetify)
app.use(router)

app.mount('#app')

// Start background polling for server info after stores are initialised
const infoStore = useInfoStore()
infoStore.startPolling()
