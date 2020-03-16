import axios from "axios";

const state = {
    podinfo: []
}


// getters
const getters = {

    podInfo: (state) => {

        state.podinfo.forEach( pi => {
                pi.kbzId = pi.namespace+"-"+pi.name
        })
        return state.podinfo

    },
    containerStatuses: (state) => {
        var cs = []
        state.podinfo.forEach( pi => {
            pi.status.containerStatuses.forEach( c => {
                c.kbzPod = pi.name
                c.kbzId = c.namespace+"-"+c.kbzPod+"-"+c.name
                c.kbzState = Object.keys(c['state'])[0] 
                if (c.kbzState == "waiting") {
                    c.kbzReason = c.state.waiting.reason
                    c.kbzReasonMessage = c.state.waiting.message
                }
                 
                c.containerInfo = pi.containerInfo[c.name]
               cs.push(c) 

        })
            
        })
        return cs
    }
    

}


// actions
const actions = {

    fetchPodInfo({ commit }) {
        axios
            .get("/kubez/kbzk8s/podlist")
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
        state.podinfo = pods
    },

}

export default {
    state,
    getters,
    actions,
    mutations
}