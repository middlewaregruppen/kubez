<template>
  <v-container fluid>
    <div class="tool-heading">
      <v-avatar color="blue-lighten-2" rounded="lg" size="38" variant="tonal">
        <v-icon icon="mdi-kubernetes" size="22"></v-icon>
      </v-avatar>
      <div>
        <div class="tool-heading__eyebrow">Kubernetes</div>
        <div class="tool-heading__title">Workload Generator</div>
      </div>
      <v-spacer></v-spacer>
      <Instructions href="k8sload.md" />
    </div>
    <p class="tool-description">Create controlled deployments to exercise cluster scheduling and resource limits.</p>

    <div class="stats-grid">
      <div class="stat-tile">
        <span>Namespaces</span>
        <strong>{{ k8sstats.noNamespaces ?? 0 }}</strong>
      </div>
      <div class="stat-tile">
        <span>Deployments</span>
        <strong>{{ k8sstats.noDeployments ?? 0 }}</strong>
      </div>
      <div class="stat-tile">
        <span>Pods</span>
        <strong>{{ k8sstats.noPods ?? 0 }}</strong>
      </div>
      <div class="stat-tile">
        <span>Ready Pods</span>
        <strong>{{ k8sstats.noReadyPods ?? 0 }}</strong>
      </div>
    </div>

    <div class="load-form mt-4">
          <v-row>
            <v-col cols="12" md="4">
              <v-select v-model="namespace" :items="k8sstats.namespacesInCluster" density="comfortable" label="Namespace" variant="outlined"></v-select>
            </v-col>
            <v-col cols="6" md="4">
              <v-text-field v-model="deployments" density="comfortable" label="Deployments" hint="Number to create" variant="outlined" />
            </v-col>
            <v-col cols="6" md="4">
              <v-text-field v-model="pods" density="comfortable" label="Pods per Deployment" type="number" variant="outlined" />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="6" md="4">
              <v-text-field v-model="reqCPU" density="comfortable" label="Requested CPU" hint="e.g. 200m" variant="outlined" />
            </v-col>
            <v-col cols="6" md="4">
              <v-text-field v-model="reqMem" density="comfortable" label="Requested Memory" hint="e.g. 120Mi" variant="outlined" />
            </v-col>
            <v-col cols="6" md="4">
              <v-text-field v-model="limCPU" density="comfortable" label="CPU Limit" hint="e.g. 500m" variant="outlined" />
            </v-col>
            <v-col cols="6" md="4">
              <v-text-field v-model="limMem" density="comfortable" label="Memory Limit" hint="e.g. 150Mi" variant="outlined" />
            </v-col>
            <v-col cols="12" md="8">
              <v-select v-model="profile" :items="profiles" density="comfortable" item-title="text" item-value="profile" label="Load Profile" variant="outlined"></v-select>
            </v-col>
          </v-row>
          <v-btn color="blue-lighten-2" prepend-icon="mdi-plus-circle-outline" variant="outlined" @click="create">
            Create {{ deployments * pods }} pods
          </v-btn>
    </div>
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

<style scoped>
.tool-heading {
  display: flex;
  align-items: center;
  gap: 0.7rem;
}

.tool-heading__eyebrow {
  color: #90caf9;
  font-size: 0.68rem;
  font-weight: 700;
  letter-spacing: 0.11em;
  text-transform: uppercase;
}

.tool-heading__title {
  font-size: 1rem;
  font-weight: 650;
}

.tool-description {
  margin: 1rem 0;
  color: rgba(255, 255, 255, 0.62);
  font-size: 0.82rem;
  line-height: 1.5;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 0.5rem;
}

.stat-tile {
  display: flex;
  flex-direction: column;
  gap: 0.2rem;
  padding: 0.75rem;
  border: 1px solid rgba(144, 202, 249, 0.12);
  border-radius: 9px;
  background-color: #2e353b;
}

.stat-tile span {
  color: rgba(255, 255, 255, 0.58);
  font-size: 0.68rem;
}

.stat-tile strong {
  color: #90caf9;
  font-size: 1.05rem;
}

.load-form {
  padding: 1rem;
  border: 1px solid rgba(144, 202, 249, 0.12);
  border-radius: 10px;
}

@media (max-width: 700px) {
  .stats-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}
</style>
