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
                <v-col>
                  <v-tooltip right>
                    <template v-slot:activator="{ on }">
                      <span v-on="on">cfs quota</span>
                    </template>
                    <span>
                      cfs quota: Maximum time in microseconds during each cpu cfs period in which
                      the current group will be allowed to run.
                    </span>
                  </v-tooltip>
                </v-col>
                <v-col>{{cg.cpuQuota | us2ns | prettyTime}} </v-col>
              </v-row>
              <v-row no-gutters>
                <v-col>
                  <v-tooltip right>
                    <template v-slot:activator="{ on }">
                      <span v-on="on">cfs period</span>
                    </template>
                    <span>cfs period: Duration in microseconds of each scheduler period, for bandwidth decisions.</span>
                  </v-tooltip>
                </v-col>
                <v-col>{{cg.cpuPeriod | us2ns | prettyTime}}</v-col>
              </v-row>
              <v-row no-gutters>
                <v-col>
                  <v-tooltip right>
                    <template v-slot:activator="{ on }">
                      <span v-on="on">nr_throttled</span>
                    </template>
                    <span>nr_throttled: Bandwidth statistics: Number of times we exausted the full allowed bandwidth.</span>
                  </v-tooltip>
                </v-col>
                <v-col>{{cg.cpuNumberThrottled}}</v-col>
              </v-row>

              <v-row no-gutters>
                <v-col>
                  <v-tooltip right>
                    <template v-slot:activator="{ on }">
                      <span v-on="on">throttled_time</span>
                    </template>
                    <span>throttled_time: Bandwidth statistics: Total time the tasks were not run due to being over quota.</span>
                  </v-tooltip>
                </v-col>
                <v-col>{{cg.cpuTimeThrottled | prettyTime}}</v-col>
              </v-row>

              <v-row no-gutters>
                <v-col>
                  <v-tooltip right>
                    <template v-slot:activator="{ on }">
                      <span v-on="on">nr_periods</span>
                    </template>
                    <span>nr_periods: Bandwidth statistics: How many full periods have been elapsed</span>
                  </v-tooltip>
                </v-col>
                <v-col>{{cg.cpuNumberPeriods}}</v-col>
              </v-row>
            </v-col>
            <v-col>
              <v-row>
                <v-col>Memory</v-col>
              </v-row>
              <v-row no-gutters>
                <v-col>
                  <v-tooltip right>
                    <template v-slot:activator="{ on }">
                      <span v-on="on">limit</span>
                    </template>
                    <span>limit: Maximum memory that the group is allowed to use</span>
                  </v-tooltip>
                </v-col>

                <v-col>{{ cg.memoryLimit | prettyBytes}}</v-col>
              </v-row>
              <v-row no-gutters>
                <v-col>
                  <v-tooltip right>
                    <template v-slot:activator="{ on }">
                      <span v-on="on">usage</span>
                    </template>
                    <span>usage: Current memory usage</span>
                  </v-tooltip>
                </v-col>
                <v-col>{{ cg.memoryUsage | prettyBytes}}</v-col>
              </v-row>
            </v-col>
          </v-row>
        </div>
      </v-list-item-content>

     
    </v-list-item>

  </v-container>
</template>
<script>
import { mapState } from 'vuex'
  
export default {
  name: "CGroup",

  computed: mapState({
    cg: state => state.info.cGroup,
  }),
  
  filters: {
    prettyTime: function(num) {
      if (typeof num !== "number" || isNaN(num)) {
        return num
      }
      if (num > Math.pow(10, 9) * 3 ) {
        return (num / Math.pow(10, 9)).toFixed(1)+" s"
      }
      if (num > Math.pow(10, 6)) {
        return (num / Math.pow(10, 6)).toFixed(0)+ " ms"
      }
      if (num > Math.pow(10, 3)) {
        return (num / Math.pow(10, 3)).toFixed(0) + " µs"
      }
      return num + " ns"
    },
    us2ns:function(num) {
      // microseconds to nano seconds
      return num*1000

    },
    prettyBytes: function(num) {
      if (typeof num !== "number" || isNaN(num)) {
        return num
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
