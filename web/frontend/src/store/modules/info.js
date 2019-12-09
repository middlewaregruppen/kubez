
const state = {
    info: {},
    cGroup: {},
  }
 



// getters
const getters = {

  cgrssoup: (state) => {
    
    return state.info.cGroup 
  }

  }


  // actions
const actions = {

}

// mutations
const mutations = {
 
  'SET_INFO' (state, info) {

      state.cGroup = info.cGroup
  }


  
}

export default {
    state,
    getters,
    actions,
    mutations
  }