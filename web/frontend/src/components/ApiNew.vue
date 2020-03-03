<template>
  <v-dialog v-model="dialog" width="900">
    <template v-slot:activator="{ on }">
      <v-btn small icon dark v-on="on">
        <v-icon>mdi-plus-box</v-icon>
      </v-btn>
    </template>
    <v-card>
      <v-card-title>
        Create New API Endpoint
      </v-card-title>
      <v-card-text>
        <v-form ref="form">
          <v-row>
            <v-col>
              <v-select
                v-model="aep.namespace"
                :items="k8s.namespacesInCluster"
                label="Namespace"
              ></v-select>
            </v-col>
            <v-col>
              <v-text-field
                v-model="aep.name"
                label="Name"
                required
              ></v-text-field>
            </v-col>
            <v-col>
              <v-select
                v-model="aep.servicetype"
                :items="serviceTypes"
                label="Service Type"
              ></v-select>
            </v-col>
            <v-col>
              <v-text-field
                v-model="aep.port"
                label="Listen Port"
                required
              ></v-text-field>
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn small color="blue darken-1" text @click="dialog = false"
          >Close</v-btn
        >
        <v-btn small color="blue darken-1" text @click="create()">Create</v-btn>
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
    serviceTypes: ["clusterIP", "nodePort", "loadBalancer"]
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
