<template>
  <v-container fluid>
    <v-row no-gutters>
      <v-col>
        <v-card flat>
            <ApiNew href="/kubez/apicc/" @createEndpoint="createEndpoint" />
            <!-- Reload button -->
            <v-btn small icon dark v-on:click="updateList()">
              <v-icon>mdi-refresh</v-icon>
            </v-btn>

            <Instructions href="api-control-center.md" />
        </v-card>
      </v-col>
    </v-row>

    <v-row v-for="e in apis" v-bind:key="e.name">
      <v-col>
        <v-card>
          <ApiEndpoint
            :namespace="e.namespace"
            :name="e.name"
            :port="e.port"
            :path="e.path"
            :mindelay="e.mindelay"
            :maxdelay="e.maxdelay"
            :failurerate="e.failurerate"
            :failurecode="e.failurecode"
            :requestrate="e.requestrate"
            :responserate="e.responserate"
            :responsetype="e.responsetype"
            :staticcontent="e.staticcontent"
            :runningpods="e.status.runningpods"
            :servicetype="e.servicetype"
            :logtoconsole="e.logtoconsole"
            :cors="e.cors"
            @update="updateEndpoint"
          />
        </v-card>
        <v-card></v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
// @ is an alias to /src
import ApiEndpoint from "@/components/ApiEndpoint.vue";
import ApiNew from "@/components/ApiNew.vue";
import Instructions from "@/components/Instructions.vue";
//import axios from "axios";
import { mapState } from "vuex";
export default {
  name: "Api",
  components: {
    ApiEndpoint,
    Instructions,
    ApiNew
  },

  computed: mapState({
    apis: state => state.apiendpoint.apis
  }),
  methods: {
     createEndpoint: function(api) {
      this.$store.dispatch("createNewAPI", api);
      this.updateList()
    },
    updateEndpoint: function(api) {
      this.$store.commit("UPDATE_API", api);
    },
    updateList: function() {
      this.$store.dispatch("fetchAPIEndpoints");
    }
  },
  mounted: function() {
    this.updateList()
  },
  data: function() {
    return {
      endpoints: {},
      dialog: false
    };
  }
};
</script>
