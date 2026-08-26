import { defineStore } from 'pinia'
import axios from 'axios'

export const useKbzK8sStore = defineStore('kbzk8s', {
  state: () => ({
    podinfo: []
  }),

  getters: {
    // Fixed: old Vuex getter was mutating objects on state.podinfo directly,
    // which violated strict-mode and caused unpredictable reactivity.
    // Now returns new derived objects without touching state.
    podInfo: (state) => {
      return state.podinfo.map(pi => ({
        ...pi,
        kbzId: `${pi.namespace}-${pi.name}`
      }))
    },
    containerStatuses: (state) => {
      const cs = []
      state.podinfo.forEach(pi => {
        if (!pi.status?.containerStatuses) return
        pi.status.containerStatuses.forEach(c => {
          const kbzState = c.state ? Object.keys(c.state)[0] : 'unknown'
          const entry = {
            ...c,
            kbzPod: pi.name,
            kbzId: `${c.namespace}-${pi.name}-${c.name}`,
            kbzState,
            containerInfo: pi.containerInfo?.[c.name] ?? {}
          }
          if (kbzState === 'waiting') {
            entry.kbzReason = c.state.waiting?.reason
            entry.kbzReasonMessage = c.state.waiting?.message
          }
          cs.push(entry)
        })
      })
      return cs
    }
  },

  actions: {
    async fetchPodInfo() {
      try {
        const res = await axios.get('/kubez/kbzk8s/podlist')
        this.podinfo = res.data
      } catch (error) {
        throw new Error(`API ${error}`)
      }
    }
  }
})
