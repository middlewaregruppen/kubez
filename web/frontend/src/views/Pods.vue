<template>
  <v-container fluid>
    <v-data-table
      :headers="headers"
      :items="pods"
      item-key="name"
      dense
      :search="search"
      single-line
      class="elevation-1"
    >
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>Pods</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-text-field
            v-model="search"
            append-icon="mdi-magnify"
            label="Search"
            single-line
            hide-details
          ></v-text-field>
          <!--v-switch v-model="singleExpand" label="Single expand" class="mt-2"></v-switch-->
        </v-toolbar>
      </template>
    </v-data-table>
  </v-container>
</template>

<script>
import { mapState } from "vuex";
export default {
  mounted: function() {
    this.updateList();
  },
  computed: {
    ...mapState({
    pods: state => state.kbzk8s.podinfo
  }),
  },

  
  methods: {
    updateList: function() {
      this.$store.dispatch("fetchPodInfo");
    }
  },

  data() {
    return {
      search: "",

      headers: [
        { text: "Namespace", value: "namespace" },
        {
          text: "Pod name",
          align: "start",
          sortable: true,
          value: "name"
        },
        { text: "Pod phase", value: "status.phase" },
        { text: "Container Status", value: "protein" },
        { text: "Restarts", value: "" },
        { text: "Conditions", value: "protein" },
        { text: "Quality  Of Service", value: "status.qosClass" },
        { text: "Reason", value: "protein" },
        { text: "CPU Throttled", value: "protein" },
        { text: "Memory Use", value: "protein" },
        { text: "Node", value: "status.hostIP" },
      ],
      
    };
  }
};
</script>
<!--  script >
// @ is an alias to /src
//import axios from "axios";
/*import { mapState } from "vuex";
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
      this.$store.commit("NEW_API", api);
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
</script -->
