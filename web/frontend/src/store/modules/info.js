
const state = {
    cGroup: {},
    hostname : "",
    updated: null,
    statusCode: -1
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
  }

  }


  // actions
const actions = {

}

// mutations
const mutations = {
 
  'SET_INFO' (state, info) {
    state.updated = Date.now()

      state.cGroup = info.cGroup
      state.hostname = info.hostname
  },
  'SET_STATUS' (state, code) {
    state.statusCode = code
  }


  
}

export default {
    state,
    getters,
    actions,
    mutations
  }