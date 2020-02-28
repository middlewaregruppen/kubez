<template>
  <v-dialog v-model="dialog" scrollable max-width="900px" max-height="700px">
    <template v-slot:activator="{ on }">
      <v-btn small icon dark v-on="on">
        <v-icon>mdi-plus-box</v-icon>
      </v-btn>
    </template>
    <v-card fluid>
      <v-card-text class="ma-4" style="height: 200px;">
        Create New Api Endpoint
        <v-form class="mt-4" ref="form">
          <v-row>
            <v-col>
              <v-select v-model="aep.namespace" :items="k8s.namespacesInCluster" label="Namespace"></v-select>
            </v-col>
            <v-col>
              <v-text-field v-model="aep.name" label="Name" required></v-text-field>
            </v-col>
            <v-col>
              <v-select v-model="aep.servicetype" :items="serviceTypes" label="Service Type"></v-select>
            </v-col>
            <v-col>
              <v-text-field v-model="aep.port" label="Listen Port" required></v-text-field>
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>
      <v-divider></v-divider>
      <v-card-actions>
        <v-btn small color="blue darken-1" text @click="create()">Create</v-btn>
        <v-btn small color="blue darken-1" text @click="dialog = false">Close</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { mapState } from "vuex";
export default {
  name: "ApiNew",
  components: {},
  data: () => ({
    dialog: false,
    aep: {
      namespace: "",
      name: "",
      port: "1337"
    },
    serviceTypes: [
      "clusterIP",
      "nodePort",
      "loadBalancer",
    ]
  }),

  methods: {
    create: function() {
      this.$emit("createEndpoint", this.aep);
      this.dialog = false;
    }
  },
  mounted: function() {},

  computed: mapState({
    k8s: state => state.info.k8sstats
  })
};
</script>