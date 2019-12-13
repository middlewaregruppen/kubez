<template>
  <div class="home">

      <v-row no-gutters>
        <v-col>
          <v-card>
            
          </v-card>
        </v-col>
      </v-row>

      <v-row v-for="e in endpoints" v-bind:key="e.name">
        <v-col>
          <v-card>
            <ApiEndpoint :name="e.name" :href="e.self" />
          </v-card>
        </v-col>
      </v-row>

  </div>
</template>

<script>
// @ is an alias to /src
import ApiEndpoint from "@/components/ApiEndpoint.vue";
import axios from "axios";
export default {
  name: "Api",
  components: {
    ApiEndpoint
  },

  mounted: function() {
    axios.get("/kubez/apicc/").then(res => (this.endpoints = res.data));
  },
  data: function() {
    return {
      endpoints: {}
    };
  }
};
</script>
