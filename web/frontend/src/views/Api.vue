<template>
  <div class="home">
    <v-row no-gutters>
      <v-col>
        <v-card>
          <v-row class="mt-2 ml-1">
            <NewApi href="/kubez/apicc/" />

            <Instructions href="api-control-center.md" />
          </v-row>
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
import NewApi from "@/components/NewApi.vue";
import Instructions from "@/components/Instructions.vue";
import axios from "axios";
export default {
  name: "Api",
  components: {
    ApiEndpoint,
    Instructions, 
    NewApi
  },

  mounted: function() {
    axios.get("/kubez/apicc/").then(res => (this.endpoints = res.data));
  },
  data: function() {
    return {
      endpoints: {},
       dialog: false,
    };
  }
};
</script>
