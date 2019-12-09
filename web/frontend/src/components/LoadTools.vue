<template>
  <v-container fluid>
    <v-list-item three-line>
      <v-list-item-content>
        <v-row no-gutters>
          <v-col>
            <v-btn  text v-on:click="loadCpu">Load CPU</v-btn>
            {{cpu}}
          </v-col>
          <v-col>
            <v-btn text v-on:click="malloc">Malloc 20 Mb</v-btn>
            {{memory}}
          </v-col>
          <v-col>
            <v-btn text>Write 100 MB</v-btn>
          </v-col>
        </v-row>
      </v-list-item-content>
    </v-list-item>
  </v-container>
</template>
<script>

import axios from 'axios';

export default {
  name: "LoadTools",

  methods: {
      loadCpu: function () {
          this.cpu = "Req ..."
        axios.post("/kubez/action/cpumedium").then( res => this.cpu=res.data)
          
      }, 
      malloc: function () {
          this.memory = "Req ..."
           axios.post("/kubez/action/malloc20mb").then( res => this.memory=res.data)
      }
  } ,
  data: function() {
    return {
      cpu: "click to load",
      memory: "click to allocate"
    };
  }

    
};
</script>
