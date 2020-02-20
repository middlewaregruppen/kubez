<template>
  <v-container fluid>
    <v-list-item three-line>
      <v-list-item-content>
        <div class="overline mb-4">Kubernetes</div>
        <v-list-item-title class="headline mb-1">Load kubernetes</v-list-item-title>
        <v-list-item-subtitle>Load kubernetes by creating a lot of resources</v-list-item-subtitle>
        <div class="body-2 pl-3">
          <v-row class="pt-3" no-gutters>
            <v-col>
              Namespaces
              <br />
              {{ k8s.noNamespaces }}
            </v-col>
            <v-col>
              Deployments
              <br />
              {{ k8s.noDeployments }}
            </v-col>
            <v-col>
              Pods
              <br />
              {{ k8s.noPods }}
            </v-col>
            <v-col>
              Ready Pods
              <br />
              {{k8s.noReadyPods}}
            </v-col>
          </v-row>
          <v-row class="pt-3">
            <v-col>
              <v-text-field v-model="namespaces" label="Namespaces" value="1" type="number" />
            </v-col>
            <v-col>
              <v-text-field
                v-model="deployments"
                label="Deployments per NS"
                value="1"
                type="number"
              />
            </v-col>
            <v-col>
              <v-text-field v-model="pods" label="Pods per Deployments" type="number" />
            </v-col>
          </v-row>
          <v-row no-gutters>
            <v-btn x-small v-on:click="create">Create {{namespaces * deployments * pods}} pods</v-btn>
          </v-row>
        </div>
      </v-list-item-content>
      <v-list-item-icon>
        <Instructions href="k8sload.md" />
      </v-list-item-icon>
    </v-list-item>
  </v-container>
</template>
<script>
import { mapState } from "vuex";
import Instructions from "@/components/Instructions.vue";
import axios from "axios";

export default {
  name: "K8SLoad",
  components: {
    Instructions
  },

  methods: {
    create: function() {
      axios
        .post("/kubez/action/k8sload", {
          namespaces: this.namespaces,
          deployments: this.deployments,
          pods: this.pods
        })
        .then(res => (this.cpu = res.data));
    }
  },
  data: () => ({
    namespaces: "1",
    deployments: "1",
    pods: "1"
  }),
  computed: mapState({
    k8s: state => state.info.k8sinfo
  })
};
</script>
