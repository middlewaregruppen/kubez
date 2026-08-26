import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vuetify from 'vite-plugin-vuetify'
import { fileURLToPath, URL } from 'node:url'

export default defineConfig({
  plugins: [
    vue(),
    vuetify({ autoImport: true }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    proxy: {
      '/kubez': {
        // Set BACKEND_TARGET env var to point to your backend, e.g.:
        //   BACKEND_TARGET=http://192.168.64.12:31503 npm run dev
        target: process.env.BACKEND_TARGET || 'http://localhost:3000/',
        ws: true,
        changeOrigin: true
      }
    }
  }
})
