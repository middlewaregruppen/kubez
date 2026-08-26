<template>
  <v-container fluid>
    <v-data-table
      :headers="headers"
      :items="tdata"
      item-value="kbzId"
      density="compact"
      fixed-header
      height="calc(100vh - 132px)"
      :search="search"
      :items-per-page="10000"
      hide-default-footer
      class="workloads-table elevation-1"
    >
      <template v-slot:top>
        <v-toolbar class="table-header px-4" flat>
          <div class="section-heading">
            <v-avatar color="blue-lighten-2" rounded="lg" size="40" variant="tonal">
              <v-icon icon="mdi-cube-outline" size="24"></v-icon>
            </v-avatar>
            <div>
              <div class="section-heading__title">Pods and Containers</div>
              <div class="section-heading__subtitle">Workload health and runtime status</div>
            </div>
          </div>
          <v-spacer></v-spacer>
          <v-btn-toggle
            v-model="type"
            class="view-selector"
            color="blue-lighten-2"
            density="comfortable"
            divided
            mandatory
            rounded="lg"
            variant="outlined"
          >
            <v-btn class="view-selector__button" value="pods">Pods</v-btn>
            <v-btn class="view-selector__button" value="containers">Container Overview</v-btn>
          </v-btn-toggle>
          <v-text-field
            v-model="search"
            class="table-search ml-3"
            append-inner-icon="mdi-magnify"
            density="compact"
            label="Search workloads"
            single-line
            hide-details
            variant="outlined"
          ></v-text-field>
          <v-btn
            class="refresh-button ml-2"
            color="blue-grey-lighten-2"
            icon="mdi-refresh"
            rounded="lg"
            variant="text"
            aria-label="Refresh workloads"
            @click="updateList()"
          ></v-btn>
        </v-toolbar>
      </template>

      <template v-slot:[`item.containerReason`]="{ item }">
        <span v-for="e in item.status.containerStatuses" :key="e.name">
          <span :style="{ color: statusColour(e) }">&#9673;</span>
        </span>
      </template>

      <template v-slot:[`item.kbzState`]="{ item }">
        <span v-if="item.kbzState === 'running' && item.ready">
          <v-chip size="x-small" variant="outlined" color="green-accent-3">ok</v-chip>
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
            <span v-bind="props" v-else-if="!item.ready">
              <v-chip size="x-small" variant="outlined" :color="probeChipColor(item.ready)">not ready</v-chip>
            </span>
          </template>
          <div class="probe-tooltip__status">
            {{ readinessStatus(item) }}
          </div>
          <pre v-if="item.containerInfo && item.containerInfo.redynessProbeConfig" class="probe-tooltip__config">{{ formatProbeConfig(item.containerInfo.redynessProbeConfig) }}</pre>
        </v-tooltip>
      </template>

      <template v-slot:[`item.livenessprobe`]="{ item }">
        <v-tooltip location="bottom" max-width="600">
          <template v-slot:activator="{ props }">
            <span v-bind="props" v-if="item.containerInfo && item.containerInfo.livenessProbeConfig">
              <v-chip variant="outlined" size="x-small" color="grey-lighten-1">
                {{ probeType(item.containerInfo.livenessProbeConfig) }} probe
              </v-chip>
            </span>
          </template>
          <div class="probe-tooltip__status">
            Liveness probe configured
          </div>
          <pre class="probe-tooltip__config">{{ formatProbeConfig(item.containerInfo && item.containerInfo.livenessProbeConfig) }}</pre>
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
  name: 'PodsView',
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
      return cs.ready ? '#76ff03' : 'red'
    },
    probeChipColor(ready) {
      return ready ? 'green-accent-3' : 'orange'
    },
    readinessStatus(item) {
      const configured = item.containerInfo?.redynessProbeConfig
      if (!configured) return 'No readiness probe configured; container is not ready'
      return item.ready ? 'Container is ready' : 'Container is not ready'
    },
    formatProbeConfig(config) {
      return JSON.stringify(config, null, 2)
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

<style scoped>
.table-header {
  min-height: 72px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.section-heading {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.section-heading__title {
  font-size: 1.125rem;
  font-weight: 650;
  line-height: 1.3;
  letter-spacing: -0.01em;
}

.section-heading__subtitle {
  margin-top: 0.1rem;
  color: rgba(255, 255, 255, 0.62);
  font-size: 0.75rem;
  line-height: 1.3;
}

.table-search {
  max-width: 280px;
}

:deep(.workloads-table .v-table__wrapper > table > thead > tr > th) {
  height: 44px;
  background-color: #2e353b !important;
  border-top: 1px solid rgba(144, 202, 249, 0.22);
  border-bottom: 1px solid #90caf9 !important;
  color: #fff !important;
  font-size: 0.75rem;
  font-weight: 700 !important;
  letter-spacing: 0.025em;
}

:deep(.workloads-table .v-table__wrapper > table > thead > tr > th:not(:last-child)) {
  border-right: 1px solid rgba(144, 202, 249, 0.1);
}

.view-selector__button,
.refresh-button {
  text-transform: none;
  letter-spacing: 0.01em;
}

.view-selector__button {
  font-weight: 600;
}

.probe-tooltip__status {
  font-weight: 600;
}

.probe-tooltip__config {
  margin-top: 0.5rem;
  padding-top: 0.5rem;
  border-top: 1px solid rgba(255, 255, 255, 0.25);
  white-space: pre-wrap;
  font-size: 0.75rem;
}

@media (max-width: 700px) {
  .section-heading__subtitle {
    display: none;
  }

  .view-selector__button {
    padding-inline: 0.5rem;
  }

  .table-search {
    display: none;
  }
}
</style>
