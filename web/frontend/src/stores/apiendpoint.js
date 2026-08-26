import { defineStore } from 'pinia'
import axios from 'axios'
import { useInfoStore } from './info'

export const useApiEndpointStore = defineStore('apiendpoint', {
  state: () => ({
    apis: [],
    loading: false,
    error: ''
  }),

  getters: {
    list: (state) => state.apis
  },

  actions: {
    async fetchAPIEndpoints() {
      this.loading = true
      this.error = ''
      try {
        const res = await axios.get('/kubez/apicc/')
        this.apis = Array.isArray(res.data) ? res.data : []
      } catch (error) {
        this.apis = []
        this.error = error.response?.data || error.message || String(error)
      } finally {
        this.loading = false
      }
    },
    async createNewAPI(api) {
      const infoStore = useInfoStore()
      try {
        await axios.post('/kubez/apicc/', api)
      } catch (error) {
        infoStore.setSnack({ message: String(error), color: 'error' })
      }
    },
    // Fixed: was previously an axios call inside a Vuex mutation (illegal side effect).
    // Now a proper async action.
    async updateAPI(api) {
      const infoStore = useInfoStore()
      try {
        await axios.put(`/kubez/apicc/${api.name}`, api)
      } catch (error) {
        infoStore.setSnack({ message: String(error), color: 'error' })
      }
    }
  }
})
