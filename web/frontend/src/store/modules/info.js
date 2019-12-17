
const state = {
    info: {},
    cGroup: {},
    Headers: {},
    updated: null,
  }
 



// getters
const getters = {

  cgroup: (state) => {
    
    return state.info.cGroup 
  },
  updated: (state) => {
    return state.updated
  },

  headers: (state) => {
    return state.info.Headers
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
  }


  
}

export default {
    state,
    getters,
    actions,
    mutations
  }