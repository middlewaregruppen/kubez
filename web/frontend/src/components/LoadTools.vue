<template>
  <v-container fluid>
    <v-row no-gutters>

        <v-card class="pa-2"  tile>
          <v-btn x-small v-on:click="loadCpu">Load CPU</v-btn> <br/>
          <!--v-btn x-small v-on:click="loadCpu">For 1 min</v-btn><v-btn x-small v-on:click="loadCpu">For 5 min</v-btn><v-btn x-small v-on:click="loadCpu">Forever</v-btn-->
          <span class="caption">{{cpu}}</span>
        </v-card>

        <v-card class="pa-2" tile>
          <v-btn x-small v-on:click="malloc">Allocate 20 Mb</v-btn><br/> <span class="caption">{{memory}}</span>
        </v-card>
        
    </v-row>
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
