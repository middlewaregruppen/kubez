import { defineStore } from 'pinia'
import axios from 'axios'
import { useInfoStore } from './info'

export const useApiEndpointStore = defineStore('apiendpoint', {
  state: () => ({
    apis: []
  }),

  getters: {
    list: (state) => state.apis
  },

  actions: {
    async fetchAPIEndpoints() {
      try {
        const res = await axios.get('/kubez/apicc/')
        this.apis = res.data
      } catch (error) {
        throw new Error(`API ${error}`)
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
