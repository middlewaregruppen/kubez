const state = {
    cGroup: {},
    httpheaders: {},
    hostname : "",
    updated: null,
    statusCode: -1,
    k8sstats: {},
    requestInfo: {},
    snack: {
      message: '',
      show: false,
      color: 'success'
    }
  }
 
// getters
const getters = {
  cgroup: (state) => {
    return state.info.cGroup 
  },
  status: (state) => {
    return state.statusCode
  },
  updateTime: (state) => {
    return state.updated
  },
  headers: (state) => {
    return state.info.Headers
  },
  requestInfo: (state) => {
    return state.info.requestInfo
  }
}

// actions
const actions = {
  setSnack({ commit }, snack) {
    commit('SET_SNACK', snack)
  },
  clearSnack({ commit }) {
    commit('CLEAR_SNACK')
  }
}

// mutations
const mutations = {
  'SET_INFO' (state, info) {
    state.updated = Date.now()
    state.cGroup = info.cGroup
    state.hostname = info.hostname
    state.httpheaders = info.httpheaders
    state.k8sstats = info.k8sstats
    state.requestInfo = info.requestInfo
  },
  'SET_STATUS' (state, code) {
    state.statusCode = code
  },
  'SET_SNACK' (state, snack) {
    state.snack.message = snack.message
    state.snack.color = snack.color
    state.snack.show = true
  },
  'CLEAR_SNACK' (state) {
    state.snack.message = ''
    state.snack.color = 'success'
    state.snack.show = false
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}