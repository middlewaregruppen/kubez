<template>
  <v-container fluid>
    <v-list-item three-line>
      <v-list-item-content>
        <div class="overline mb-4">Linux Kernel</div>
        <v-list-item-title class="headline mb-1">Control Group</v-list-item-title>
        <v-list-item-subtitle>Control groups are used to assign quotas to the containers when they are scheduled on the compute node. This information is collected from /sys/fs/cgroup inside the container.</v-list-item-subtitle>
        <div class="body-2 pl-3">
          <div class="mt-2">CPU</div>
          <v-row no-gutters>
            <v-col>cfs quota</v-col>
            <v-col>{{cgroup.cpuCfs_quota_us}} µs</v-col>
            <v-col
              cols="9"
            >Maximum time in microseconds during each cpu cfs period in which the current group will be allowed to run.</v-col>
          </v-row>
          <v-row no-gutters>
            <v-col>cfs period</v-col>
            <v-col>{{cgroup.cpuCfs_period_us}} µs</v-col>
            <v-col
              cols="9"
            >Duration in microseconds of each scheduler period, for bandwidth decisions.</v-col>
          </v-row>
          <v-row no-gutters>
            <v-col>nr_throttled</v-col>
            <v-col>{{cgroup.cpuStatNr_throttled}}</v-col>
            <v-col
              cols="9"
            >Bandwidth statistics: Number of times we exausted the full allowed bandwidth</v-col>
          </v-row>

          <v-row no-gutters>
            <v-col>throttled_time</v-col>
            <v-col>{{cgroup.cpuStatThrottled_time}}</v-col>
            <v-col
              cols="9"
            >Bandwidth statistics: Total time the tasks were not run due to being overquota</v-col>
          </v-row>

          <v-row no-gutters>
            <v-col>nr_periods</v-col>
            <v-col>{{cgroup.cpuStatNr_periods}}</v-col>
            <v-col cols="9">Bandwidth statistics: How many full periods have been elapsed</v-col>
          </v-row>
          <div class="mt-2">Memory</div>
          <v-row no-gutters>
            <v-col>limit</v-col>
            <v-col>{{ cgroup.memoryLimit_in_bytes | prettyBytes}}</v-col>
            <v-col cols="9">Maximum memory that the group is allowed to use</v-col>
          </v-row>
          <v-row no-gutters>
            <v-col>usage</v-col>
            <v-col>{{ cgroup.memoryUsage_in_bytes | prettyBytes}}</v-col>
            <v-col cols="9">Current memory usage</v-col>
          </v-row>
        </div>
      </v-list-item-content>

      <v-list-item-avatar tile size="80" color="grey">
        <v-icon></v-icon>
      </v-list-item-avatar>
    </v-list-item>

    <v-card-actions>
      <!--v-btn text>Refresh</v-btn>
      <v-btn text>Show all</v-btn-->
    </v-card-actions>
  </v-container>
</template>
<script>
export default {
  name: "CGroup",

  data: () => ({ cgroup: {} }),

  mounted: function() {
    this.fetchData();
          setInterval(
        function() {
          this.fetchData();
        }.bind(this),
        1000
      );
  },

  filters: {
    prettyBytes: function(num) {
      if (typeof num !== "number" || isNaN(num)) {
        throw new TypeError("Expected a number");
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
  },

  methods: {
    fetchData: function() {
      var xhr = new XMLHttpRequest();
      var self = this;
      xhr.open("GET", "/api/cgroup");
      xhr.onload = function() {
        self.cgroup = JSON.parse(xhr.responseText);
      };
      xhr.send();
    }
  }
};
</script>
