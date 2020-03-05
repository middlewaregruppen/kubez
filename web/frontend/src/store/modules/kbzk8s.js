import axios from "axios";

const state = {
    podinfo: []
}


// getters
const getters = {

    pods: (state) => {
        return state.podinfo
    }

}


// actions
const actions = {

    fetchPods({ commit }) {
        axios
            .get("/kubez/kbzk8s/pods")
            .then(res => res.data)
            .then(pods => {
                commit ('SET_PODS', pods)
            }).catch(error => {
                throw new Error(`API ${error}`);
              });
    },


}

// mutations
const mutations = {

    'SET_PODS'(state, pods) {
        state.pods = pods
    },

}

export default {
    state,
    getters,
    actions,
    mutations
}