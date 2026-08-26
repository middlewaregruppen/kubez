<template>
  <v-container fluid>
    <v-data-table
      :headers="headers"
      :items="tdata"
      item-value="kbzId"
      density="compact"
      :search="search"
      :items-per-page="10000"
      hide-default-footer
      class="elevation-1"
    >
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>Pods and Containers</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-text-field
            v-model="search"
            append-inner-icon="mdi-magnify"
            label="Search"
            single-line
            hide-details
          ></v-text-field>
        </v-toolbar>
        <div class="d-flex align-center ml-4 mb-2">
          <v-btn-toggle v-model="type" mandatory>
            <v-btn size="x-small" variant="text" value="pods">pods</v-btn>
            <v-btn size="x-small" variant="text" value="containers">container overview</v-btn>
          </v-btn-toggle>
          <v-btn size="x-small" variant="text" class="ml-3" @click="updateList()">refresh</v-btn>
        </div>
      </template>

      <template v-slot:[`item.containerReason`]="{ item }">
        <span v-for="e in item.status.containerStatuses" :key="e.name">
          <span :style="{ color: statusColour(e) }">&#9673;</span>
        </span>
      </template>

      <template v-slot:[`item.kbzState`]="{ item }">
        <span v-if="item.kbzState === 'running' && item.ready">
          <v-chip size="x-small" variant="outlined" color="green">ok</v-chip>
        </span>

        <v-tooltip location="bottom" max-width="600" v-if="item.kbzReason">
          <template v-slot:activator="{ props }">
            <span v-bind="props">
              <v-chip size="x-small" variant="outlined" color="orange">{{ item.kbzReason }}</v-chip>
              <v-chip size="x-small" variant="outlined" color="orange" v-if="terminationReason(item)">{{ terminationReason(item) }}</v-chip>
            </span>
          </template>
          <span>Pod is in {{ item.kbzState }} state.<br />{{ item.kbzReasonMessage }}<br />--<br /></span>
          <span v-if="terminationReason(item)">Reason container terminated: {{ item.lastState.terminated.reason }}<br />{{ item.lastState.terminated.message }}</span>
        </v-tooltip>

        <v-tooltip location="bottom" max-width="600" v-else-if="item.kbzState === 'running' && !item.ready">
          <template v-slot:activator="{ props }">
            <span v-bind="props">
              <v-chip size="x-small" variant="outlined" color="orange">not ready</v-chip>
            </span>
          </template>
          <span>Container is running but is not ready to receive traffic</span>
        </v-tooltip>

        <v-tooltip
          location="bottom"
          max-width="600"
          v-else-if="item.kbzState === 'running' && item.ready && recentlyRestarted(item)"
        >
          <template v-slot:activator="{ props }">
            <span v-bind="props">
              <v-chip size="x-small" variant="outlined" color="orange">restarted due to {{ item.lastState.terminated.reason }}</v-chip>
            </span>
          </template>
          <span>The container was restarted less than 15 minutes ago<br />--<br />{{ item.lastState.terminated.message }}</span>
        </v-tooltip>
      </template>

      <template v-slot:[`item.ready`]="{ item }">
        <v-tooltip location="bottom" max-width="600">
          <template v-slot:activator="{ props }">
            <span v-bind="props" v-if="item.containerInfo && item.containerInfo.redynessProbeConfig">
              <v-chip size="x-small" variant="outlined" :color="probeChipColor(item.ready)">
                {{ probeType(item.containerInfo.redynessProbeConfig) }} probe
              </v-chip>
            </span>
            <span v-else-if="!item.ready">
              <v-chip size="x-small" variant="outlined" :color="probeChipColor(item.ready)">not ready</v-chip>
            </span>
          </template>
          <span>{{ item.containerInfo && item.containerInfo.redynessProbeConfig }}</span>
        </v-tooltip>
      </template>

      <template v-slot:[`item.livenessprobe`]="{ item }">
        <v-tooltip location="bottom" max-width="600">
          <template v-slot:activator="{ props }">
            <span v-bind="props" v-if="item.containerInfo && item.containerInfo.livenessProbeConfig">
              <v-chip variant="outlined" size="x-small" color="black">
                {{ probeType(item.containerInfo.livenessProbeConfig) }} probe
              </v-chip>
            </span>
          </template>
          <span>{{ item.containerInfo && item.containerInfo.livenessProbeConfig }}</span>
        </v-tooltip>
      </template>

      <template v-slot:[`item.terminationReason`]="{ item }">
        <v-tooltip location="bottom" max-width="600">
          <template v-slot:activator="{ props }">
            <span v-bind="props" v-if="item.restartCount > 0">{{ item.lastState.terminated.reason }}</span>
          </template>
          <span>{{ item.lastState }}</span>
        </v-tooltip>
      </template>
    </v-data-table>
  </v-container>
</template>

<script>
import { useKbzK8sStore } from '@/stores/kbzk8s'

export default {
  setup() {
    return {
      k8sStore: useKbzK8sStore()
    }
  },
  mounted() {
    this.updateList()
  },
  computed: {
    containers() {
      return this.k8sStore.containerStatuses
    },
    pods() {
      return this.k8sStore.podInfo
    },
    headers() {
      return this.type === 'containers' ? this.containerheaders : this.podheaders
    },
    tdata() {
      return this.type === 'containers' ? this.containers : this.pods
    }
  },
  methods: {
    updateList() {
      this.k8sStore.fetchPodInfo()
    },
    statusColour(cs) {
      return cs.ready ? 'green' : 'red'
    },
    probeChipColor(ready) {
      return ready ? 'green' : 'orange'
    },
    probeType(cfg) {
      for (const key of Object.keys(cfg)) {
        switch (key) {
          case 'httpGet': return 'http'
          case 'exec': return 'exec'
          case 'tcpSocket': return 'socket'
        }
      }
      return 'probed'
    },
    recentlyRestarted(status) {
      if (!status.lastState?.terminated) return false
      const rt = new Date(status.lastState.terminated.finishedAt)
      const MIN_15 = 15 * 60 * 1000
      return new Date() - rt < MIN_15
    },
    terminationReason(status) {
      return status.lastState?.terminated?.reason ?? ''
    }
  },
  data() {
    return {
      search: '',
      type: 'pods',
      podheaders: [
        { title: 'Namespace', value: 'namespace' },
        { title: 'Pod name', align: 'start', sortable: true, value: 'name' },
        { title: 'Pod phase', value: 'status.phase' },
        { title: 'Conditions', value: 'condition' },
        { title: 'Container Status', value: 'containerReason' },
        { title: 'Restarts', value: 'containerRestarts' },
        { title: 'Quality Of Service', value: 'status.qosClass' },
        { title: 'Reason', value: 'reason' },
        { title: 'Node', value: 'status.hostIP' }
      ],
      containerheaders: [
        { title: 'Container name', align: 'start', sortable: true, value: 'name' },
        { title: 'State', value: 'kbzState' },
        { title: 'Ready', value: 'ready', sortable: true },
        { title: 'Liveness', value: 'livenessprobe', sortable: false },
        { title: 'Restarts', value: 'restartCount' },
        { title: 'Image Name', value: 'image' },
        { title: 'Pod', value: 'kbzPod' }
      ]
    }
  }
}
</script>
