import { createVuetify } from 'vuetify'
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// vite-plugin-vuetify handles component/directive auto-import when autoImport: true
export default createVuetify({
  icons: {
    defaultSet: 'mdi'
  },
  theme: {
    defaultTheme: 'dark'
  }
})
