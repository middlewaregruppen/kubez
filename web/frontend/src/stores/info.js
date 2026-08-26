import { defineStore } from 'pinia'
import axios from 'axios'

export const useInfoStore = defineStore('info', {
  state: () => ({
    cGroup: {},
    httpheaders: {},
    hostname: '',
    updated: null,
    statusCode: -1,
    k8sstats: {},
    requestInfo: {},
    snack: {
      message: '',
      show: false,
      color: 'success'
    }
  }),

  getters: {
    // Fixed: old Vuex module had `state.info.cGroup` which was always undefined
    // inside the module. Now accessing state directly.
    cgroup: (state) => state.cGroup,
    status: (state) => state.statusCode,
    updateTime: (state) => state.updated,
    headers: (state) => state.httpheaders
  },

  actions: {
    setInfo(info) {
      this.updated = Date.now()
      this.cGroup = info.cGroup ?? {}
      this.hostname = info.hostname ?? ''
      this.httpheaders = info.httpheaders ?? {}
      this.k8sstats = info.k8sstats ?? {}
      this.requestInfo = info.requestInfo ?? {}
    },
    setStatus(code) {
      this.statusCode = code
    },
    setSnack(snack) {
      this.snack.message = snack.message
      this.snack.color = snack.color
      this.snack.show = true
    },
    clearSnack() {
      this.snack.message = ''
      this.snack.color = 'success'
      this.snack.show = false
    },
    startPolling() {
      const poll = () => {
        axios.get('/kubez/info')
          .then(res => {
            this.setInfo(res.data)
            this.setStatus(res.status)
          })
          .catch(err => {
            this.setStatus(err)
          })
      }
      poll()
      setInterval(poll, 1000)
    }
  }
})
