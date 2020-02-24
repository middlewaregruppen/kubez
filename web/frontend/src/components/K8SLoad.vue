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
              ..
              <br />
              <b>Cluster:</b>
              <br />
              <b>Namespace:</b>
            </v-col>
            <v-col>
              <b>Namespaces</b>
              <br />
              {{ k8s.noNamespaces }}
              <br />
              {{ k8s.namespace }}
            </v-col>
            <v-col>
              <b>Deployments</b>
              <br />
              {{ k8s.noDeployments }}
              <br />
              {{ k8s.noDeploymentsinNs }}
            </v-col>
            <v-col>
              <b>Pods</b>
              <br />
              {{ k8s.noPods }}
              <br />
              {{ k8s.noPodsInNs}}
            </v-col>
            <v-col>
              <b>Ready Pods</b>
              <br />
              {{k8s.noReadyPods}}
              <br />
              {{k8s.noReadyPodsInNs}}
            </v-col>
          </v-row>
          <v-row class="pt-3">
            <v-col>
              <v-text-field
                v-model="namespaces"
                label="Namespaces"
                hint=" "
                disabled
                type="number"
              />
            </v-col>
            <v-col>
              <v-text-field v-model="deployments" label="Deployments" hint="Number to create" />
            </v-col>
            <v-col>
              <v-text-field v-model="pods" label="Pods per Deployments" type="number" />
            </v-col>
          </v-row>
          <v-row no-gutters>
            <v-col>
              <v-text-field v-model="reqCPU" label="Req. CPU" hint="ex. 200m" />
            </v-col>
            <v-col>
              <v-text-field v-model="reqMem" label="Req. Memory" hint="ex. 120Mi" />
            </v-col>
            <v-col>
              <v-text-field v-model="limCPU" label="Limit CPU" hint="ex. 500m" />
            </v-col>
            <v-col>
              <v-text-field v-model="limMem" label="Limit Memory" hint="ex. 150Mi" />
            </v-col>
            <v-col>
              <v-select
                v-model="profile"
                :items="profiles"
                item-text="text"
                item-value="profile"
                label="Load"
                
              ></v-select>
            </v-col>
          </v-row>
          <v-row>
            <v-btn x-small v-on:click="create">Create {{ deployments * pods}} pods</v-btn>
          </v-row>
        </div>
      </v-list-item-content>
      <v-list-item-icon>
        <Instructions href="k8sload.md" />
      </v-list-item-icon>
    </v-list-item>
  </v-container>
</template>

<style>
  .v-select input {
    font-size: 1.2em;
  }
</style>
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
        .post("/kubez/k8sload", {
          namespaces: this.namespaces,
          deployments: this.deployments,
          pods: this.pods,
          reqCPU: this.reqCPU,
          reqMem: this.reqMem,
          limCPU: this.limCPU,
          limMem: this.limMem,
          profile: this.profile
        })
        .then(res => (this.cpu = res.data));
    }
  },
  data: () => ({
    namespaces: "0",
    deployments: "1",
    pods: "1",
    reqCPU: "0",
    reqMem: "0",
    limCPU: "0",
    limMem: "0",
    profile: "none",
    profiles: [
      { profile: "none", text: "none" },
      { profile: "cpu", text: "cpu 100%" },
      { profile: "mem100", text: "mem 100 Mb" },
      { profile: "mem2000", text: "mem 2000 Mb" }
    ]
  }),
  computed: mapState({
    k8s: state => state.info.k8sstats
  })
};
</script>
