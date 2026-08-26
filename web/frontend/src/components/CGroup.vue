<template>
  <v-list-item lines="three">
    <template v-slot:default>
      <div class="text-overline mb-4">Linux Kernel</div>
      <v-list-item-title class="text-h6 mb-1">Control Group</v-list-item-title>
      <v-list-item-subtitle>
        Control groups are used to assign quotas to the containers when they are scheduled on the
        compute node. This information is collected from /sys/fs/cgroup inside the container.
      </v-list-item-subtitle>
      <div class="text-body-2 pl-3">
        <v-row>
          <v-col>
            <v-row><v-col>CPU</v-col></v-row>
            <v-row no-gutters>
              <v-col>cfs quota</v-col>
              <v-col>{{ prettyTime(us2ns(cg.cpuQuota)) }}</v-col>
            </v-row>
            <v-row no-gutters>
              <v-col>cfs period</v-col>
              <v-col>{{ prettyTime(us2ns(cg.cpuPeriod)) }}</v-col>
            </v-row>
            <v-row no-gutters>
              <v-col>nr_throttled</v-col>
              <v-col>{{ cg.cpuNumberThrottled }}</v-col>
            </v-row>
            <v-row no-gutters>
              <v-col>throttled_time</v-col>
              <v-col>{{ prettyTime(cg.cpuTimeThrottled) }}</v-col>
            </v-row>
            <v-row no-gutters>
              <v-col>nr_periods</v-col>
              <v-col>{{ cg.cpuNumberPeriods }}</v-col>
            </v-row>
          </v-col>
          <v-col>
            <v-row><v-col>Memory</v-col></v-row>
            <v-row no-gutters>
              <v-col>limit</v-col>
              <v-col>{{ prettyBytes(cg.memoryLimit) }}</v-col>
            </v-row>
            <v-row no-gutters>
              <v-col>usage</v-col>
              <v-col>{{ prettyBytes(cg.memoryUsage) }}</v-col>
            </v-row>
          </v-col>
        </v-row>
      </div>
    </template>
    <template v-slot:append>
      <Instructions href="cgroup.md" />
    </template>
  </v-list-item>
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
