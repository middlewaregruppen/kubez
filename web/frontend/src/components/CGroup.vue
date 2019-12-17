<template>
  <v-container fluid>
    <v-list-item three-line>
      <v-list-item-content>
        <div class="overline mb-4">Linux Kernel</div>
        <v-list-item-title class="headline mb-1">Control Group</v-list-item-title>
        <v-list-item-subtitle>Control groups are used to assign quotas to the containers when they are scheduled on the compute node. This information is collected from /sys/fs/cgroup inside the container.</v-list-item-subtitle>
        <div class="body-2 pl-3">
          <v-row>
            <v-col>
              <v-row>
                <v-col>CPU</v-col>
              </v-row>
              <v-row no-gutters>
                <v-col>cfs quota</v-col>
                <v-col>{{cg.cpuQuota | us2ns | prettyTime}}</v-col>
              </v-row>
              <v-row no-gutters>
                <v-col>cfs period</v-col>
                <v-col>{{cg.cpuPeriod | us2ns | prettyTime}}</v-col>
              </v-row>
              <v-row no-gutters>
                <v-col>nr_throttled</v-col>
                <v-col>{{cg.cpuNumberThrottled}}</v-col>
              </v-row>

              <v-row no-gutters>
                <v-col>throttled_time</v-col>
                <v-col>{{cg.cpuTimeThrottled | prettyTime}}</v-col>
              </v-row>

              <v-row no-gutters>
                <v-col>nr_periods</v-col>
                <v-col>{{cg.cpuNumberPeriods}}</v-col>
              </v-row>
            </v-col>
            <v-col>
              <v-row>
                <v-col>Memory</v-col>
              </v-row>
              <v-row no-gutters>
                <v-col>
                  limit
                </v-col>

                <v-col>{{ cg.memoryLimit | prettyBytes}}</v-col>
              </v-row>
              <v-row no-gutters>
                <v-col>
                  usage
                </v-col>
                <v-col>{{ cg.memoryUsage | prettyBytes}}</v-col>
              </v-row>
            </v-col>
          </v-row>
        </div>
      </v-list-item-content>
      <v-list-item-icon>
        <Instructions href="cgroup.md" />
      </v-list-item-icon>
    </v-list-item>
  </v-container>
</template>
<script>
import { mapState } from "vuex";
import Instructions from "@/components/Instructions.vue";

export default {
  name: "CGroup",
  components: {
    Instructions
  },

  computed: mapState({
    cg: state => state.info.cGroup
  }),

  filters: {
    prettyTime: function(num) {
      if (typeof num !== "number" || isNaN(num)) {
        return num;
      }
      if (num > Math.pow(10, 9) * 3) {
        return (num / Math.pow(10, 9)).toFixed(1) + " s";
      }
      if (num > Math.pow(10, 6)) {
        return (num / Math.pow(10, 6)).toFixed(0) + " ms";
      }
      if (num > Math.pow(10, 3)) {
        return (num / Math.pow(10, 3)).toFixed(0) + " µs";
      }
      return num + " ns";
    },
    us2ns: function(num) {
      // microseconds to nano seconds
      return num * 1000;
    },
    prettyBytes: function(num) {
      if (typeof num !== "number" || isNaN(num)) {
        return num;
      }

      var exponent;
      var unit;
      var neg = num < 0;
      var units = ["B", "kB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];

      if (neg) {
        num = -num;
      }

      if (num < 1) {
        return (neg ? "-" : "") + num + " B";
      }

      exponent = Math.min(
        Math.floor(Math.log(num) / Math.log(1000)),
        units.length - 1
      );
      num = (num / Math.pow(1000, exponent)).toFixed(2) * 1;
      unit = units[exponent];

      return (neg ? "-" : "") + num + " " + unit;
    }
  }
};
</script>
