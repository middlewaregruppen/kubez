
const state = {
  cGroup: {},
  httpheaders: {},
  hostname: "",
  updated: null,
  statusCode: -1,
  k8sstats: {},
  requestInfo: {}
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

}

// mutations
const mutations = {

  'SET_INFO'(state, info) {
    state.updated = Date.now()

    state.cGroup = info.cGroup
    state.hostname = info.hostname
    state.httpheaders = info.httpheaders
    state.k8sstats = info.k8sstats
    state.requestInfo = info.requestInfo
  },
  'SET_STATUS'(state, code) {
    state.statusCode = code
  }



}

export default {
  state,
  getters,
  actions,
  mutations
}