<template>
  <div class="tool-content">
    <div class="tool-heading">
      <v-avatar color="blue-lighten-2" rounded="lg" size="38" variant="tonal">
        <v-icon icon="mdi-memory" size="22"></v-icon>
      </v-avatar>
      <div>
        <div class="tool-heading__eyebrow">Linux Kernel</div>
        <div class="tool-heading__title">Control Group</div>
      </div>
      <v-spacer></v-spacer>
      <Instructions href="cgroup.md" />
    </div>
    <p class="tool-description">
        Control groups are used to assign quotas to the containers when they are scheduled on the
        compute node. This information is collected from /sys/fs/cgroup inside the container.
    </p>
      <div class="metric-grid">
        <v-row>
          <v-col>
            <div class="metric-group__title">CPU</div>
            <v-row no-gutters>
              <v-col class="metric-label">CFS quota</v-col>
              <v-col class="metric-value">{{ prettyTime(us2ns(cg.cpuQuota)) }}</v-col>
            </v-row>
            <v-row no-gutters>
              <v-col class="metric-label">CFS period</v-col>
              <v-col class="metric-value">{{ prettyTime(us2ns(cg.cpuPeriod)) }}</v-col>
            </v-row>
            <v-row no-gutters>
              <v-col class="metric-label">Throttled count</v-col>
              <v-col class="metric-value">{{ cg.cpuNumberThrottled }}</v-col>
            </v-row>
            <v-row no-gutters>
              <v-col class="metric-label">Throttled time</v-col>
              <v-col class="metric-value">{{ prettyTime(cg.cpuTimeThrottled) }}</v-col>
            </v-row>
            <v-row no-gutters>
              <v-col class="metric-label">Periods</v-col>
              <v-col class="metric-value">{{ cg.cpuNumberPeriods }}</v-col>
            </v-row>
          </v-col>
          <v-col>
            <div class="metric-group__title">Memory</div>
            <v-row no-gutters>
              <v-col class="metric-label">Limit</v-col>
              <v-col class="metric-value">{{ prettyBytes(cg.memoryLimit) }}</v-col>
            </v-row>
            <v-row no-gutters>
              <v-col class="metric-label">Usage</v-col>
              <v-col class="metric-value">{{ prettyBytes(cg.memoryUsage) }}</v-col>
            </v-row>
          </v-col>
        </v-row>
      </div>
  </div>
</template>

<script>
import { useInfoStore } from '@/stores/info'
import Instructions from '@/components/Instructions.vue'
import { prettyTime, us2ns, prettyBytes } from '@/utils/format'

export default {
  name: 'CGroup',
  components: { Instructions },
  setup() {
    return {
      infoStore: useInfoStore()
    }
  },
  computed: {
    cg() {
      return this.infoStore.cGroup ?? {}
    }
  },
  methods: {
    prettyTime,
    us2ns,
    prettyBytes
  }
}
</script>

<style scoped>
.tool-content {
  padding: 0.25rem;
}

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

.metric-grid {
  padding: 0.9rem;
  border: 1px solid rgba(144, 202, 249, 0.12);
  border-radius: 10px;
  background-color: #2e353b;
}

.metric-group__title {
  margin-bottom: 0.6rem;
  color: #90caf9;
  font-weight: 700;
}

.metric-label,
.metric-value {
  padding-block: 0.25rem;
  font-size: 0.8rem;
}

.metric-label {
  color: rgba(255, 255, 255, 0.62);
}

.metric-value {
  color: #fff;
  font-weight: 600;
}
</style>
