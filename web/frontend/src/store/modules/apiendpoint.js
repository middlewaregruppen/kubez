import axios from "axios";

const state = {
    apis: []
}


// getters
const getters = {

    list: (state) => {

        return state.apis
    }

}


// actions
const actions = {

    fetchAPIEndpoints({ commit }) {
        axios
            .get("/kubez/apicc/")
            .then(res => res.data)
            .then(apis => {
                commit ('SET_APIS', apis)
            }).catch(error => {
                throw new Error(`API ${error}`);
              });
    },
    createNewAPI({ dispatch}, api) {
        axios.post("/kubez/apicc/", api)
        .catch(error => {
            dispatch('setSnack', { message: error, color: 'error'})
        })
    }

}

// mutations
const mutations = {

    'SET_APIS'(state, apis) {
        state.apis = apis
    },
    'UPDATE_API'(state, api){
        let uri = "/kubez/apicc/"+api.name
        axios.put(uri, api)
       .catch(error => {
            throw new Error(`API ${error}`);
          });
    }
}

export default {
    state,
    getters,
    actions,
    mutations
}