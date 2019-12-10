<template>
  <v-container fluid>
    <v-list-item three-line>
      <v-list-item-content>
        <v-row no-gutters>
          <v-col body-2>
            <v-card class="pa-2" outlined tile>
              <span class="caption">Load CPU</span> <v-btn x-small v-on:click="loadCpu">once</v-btn><!--v-btn x-small v-on:click="loadCpu">For 1 min</v-btn><v-btn x-small v-on:click="loadCpu">For 5 min</v-btn><v-btn x-small v-on:click="loadCpu">Forever</v-btn-->
            </v-card>
            <v-card class="pa-2" outlined tile><span class="caption" >{{cpu}}</span></v-card>
          </v-col>

            <v-col >
            <v-card class="pa-2" outlined tile>
              <v-btn x-small v-on:click="malloc">Malloc 20 Mb</v-btn>
            </v-card>
            <v-card body-2 class="pa-2" outlined tile><span class="caption" >{{memory}}</span></v-card>
          </v-col>
         
        </v-row>
      </v-list-item-content>
    </v-list-item>
  </v-container>
</template>
<script>
import axios from "axios";

export default {
  name: "LoadTools",

  methods: {
    loadCpu: function() {
      this.cpu = "Requesting ...";
      axios.post("/kubez/action/cpumedium").then(res => (this.cpu = res.data));
    },
    malloc: function() {
      this.memory = "Requesting ...";
      axios
        .post("/kubez/action/malloc20mb")
        .then(res => (this.memory = res.data));
    }
  },
  data: function() {
    return {
      cpu: "-",
      memory: "-"
    };
  }
};
</script>
