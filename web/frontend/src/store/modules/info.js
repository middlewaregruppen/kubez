
const state = {
    info: {},
    cGroup: {},
    updated: null,
  }
 



// getters
const getters = {

  cgroup: (state) => {
    
    return state.info.cGroup 
  },
  updated: (state) => {
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
  }


  
}

export default {
    state,
    getters,
    actions,
    mutations
  }