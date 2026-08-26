<template>
  <v-container fluid>
    <v-list-item lines="three">
      <template v-slot:default>
        <div class="text-overline mb-4">Kubernetes</div>
        <v-list-item-title class="text-h6 mb-1">Load kubernetes</v-list-item-title>
        <v-list-item-subtitle>Load kubernetes by creating a lot of resources</v-list-item-subtitle>
        <div class="text-body-2 pl-3">
          <v-row class="pt-3" no-gutters>
            <v-col>
              ..<br />
              <b>Cluster:</b><br />
              <b>Namespace:</b>
            </v-col>
            <v-col>
              <b>Namespaces</b><br />
              {{ k8sstats.noNamespaces }}<br />
              {{ k8sstats.namespace }}
            </v-col>
            <v-col>
              <b>Deployments</b><br />
              {{ k8sstats.noDeployments }}<br />
              {{ k8sstats.noDeploymentsinNs }}
            </v-col>
            <v-col>
              <b>Pods</b><br />
              {{ k8sstats.noPods }}<br />
              {{ k8sstats.noPodsInNs }}
            </v-col>
            <v-col>
              <b>Ready Pods</b><br />
              {{ k8sstats.noReadyPods }}<br />
              {{ k8sstats.noReadyPodsInNs }}
            </v-col>
          </v-row>
          <v-row class="pt-3">
            <v-col>
              <v-select v-model="namespace" :items="k8sstats.namespacesInCluster" label="Namespace"></v-select>
            </v-col>
            <v-col>
              <v-text-field v-model="deployments" label="Deployments" hint="Number to create" />
            </v-col>
            <v-col>
              <v-text-field v-model="pods" label="Pods per Deployment" type="number" />
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
              <v-select v-model="profile" :items="profiles" item-title="text" item-value="profile" label="Load"></v-select>
            </v-col>
          </v-row>
          <v-row>
            <v-btn size="small" @click="create">Create {{ deployments * pods }} pods</v-btn>
          </v-row>
        </div>
      </template>
      <template v-slot:append>
        <Instructions href="k8sload.md" />
      </template>
    </v-list-item>
  </v-container>
</template>

<script>
import { useInfoStore } from '@/stores/info'
import Instructions from '@/components/Instructions.vue'
import axios from 'axios'

export default {
  name: 'K8SLoad',
  components: { Instructions },
  setup() {
    return {
      infoStore: useInfoStore()
    }
  },
  computed: {
    k8sstats() {
      return this.infoStore.k8sstats ?? {}
    }
  },
  methods: {
    create() {
      axios.post('/kubez/k8sload', {
        namespace: this.namespace,
        deployments: this.deployments,
        pods: this.pods,
        reqCPU: this.reqCPU,
        reqMem: this.reqMem,
        limCPU: this.limCPU,
        limMem: this.limMem,
        profile: this.profile
      }).catch(error => {
        this.infoStore.setSnack({ message: error.response?.data ?? String(error), color: 'error' })
      })
    }
  },
  data: () => ({
    namespace: '',
    deployments: '1',
    pods: '1',
    reqCPU: '0',
    reqMem: '0',
    limCPU: '0',
    limMem: '0',
    profile: 'none',
    profiles: [
      { profile: 'none', text: 'none' },
      { profile: 'cpu', text: 'cpu 100%' },
      { profile: 'mem100', text: 'mem 100 Mb' },
      { profile: 'mem2000', text: 'mem 2000 Mb' },
      { profile: 'log100ms', text: 'log ~10lns/s' },
      { profile: 'log9ms', text: 'log ~100lns/s' }
    ]
  })
}
</script>
