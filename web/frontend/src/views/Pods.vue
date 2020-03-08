<template>
  <v-container fluid>
    <v-data-table
      :headers="headers"
      :items="tdata"
      item-key="kbzId"
      dense
      :search="search"
      single-line
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
            append-icon="mdi-magnify"
            label="Search"
            single-line
            hide-details
          ></v-text-field>
          <!--v-switch v-model="singleExpand" label="Single expand" class="mt-2"></v-switch-->
        </v-toolbar>
        <v-btn-toggle v-model="type" mandatory class="ml-4">
          <v-btn x-small text value="pods">pods</v-btn>
          <v-btn x-small text value="containers">container overview</v-btn>
        </v-btn-toggle>
        <v-btn x-small text class="ml-3" @click="updateList()">refresh</v-btn>
      </template>

      <template v-slot:item.containerReason="{ item }">
        <span v-for="e in item.status.containerStatuses" v-bind:key="e.name">
          <font :color="statusColour(e)">&#9673;</font>
        </span>
      </template>

        
      <template v-slot:item.kbzState="{ item }">
        
        <span v-if="item.kbzState ===  'running' && item.ready">
          <v-chip x-small text-color="white"  outlined color="green">ok</v-chip>
        </span>

        <v-tooltip bottom max-width="600" v-if="item.kbzReason">
          <template v-slot:activator="{ on }">
            <span v-on="on">
              <v-chip x-small outlined text-color="white" color="orange">{{ item.kbzReason }}</v-chip>
              <v-chip x-small outlined text-color="white" color="orange"  v-if="terminationReason(item)">{{terminationReason(item)}} </v-chip>
            </span>
          </template>
          <span>Pod is in {{item.kbzState}} state. <br/> {{item.kbzReasonMessage}} <br/> -- <br/> </span> 
          <span v-if="terminationReason(item)"> Reason container terminated: {{item.lastState.terminated.reason}} <br/> {{item.lastState.terminated.message}} </span>
        </v-tooltip>


        <v-tooltip bottom max-width="600" v-else-if="item.kbzState ===  'running' && !item.ready">
          <template v-slot:activator="{ on }">
            <span v-on="on">
              <v-chip x-small text-color="white" outlined color="orange">not ready</v-chip>
            </span>
          </template>
          <span>Container is running but is not ready to receive traffic</span>
        </v-tooltip>
        
        

        <v-tooltip
          bottom
          max-width="600"
          v-else-if="item.kbzState ===  'running' && item.ready && recentlyRestarted(item)"
        >
          <template v-slot:activator="{ on }">
            <span v-on="on">
              <v-chip
                x-small
                outlined
                text-color="white" 
                color="orange"
              >restarted due to {{item.lastState.terminated.reason}}</v-chip>
            </span>
          </template>
          <span>The container was restarted less than 15 minutes ago <br/> -- <br/> {{item.lastState.terminated.message}} </span>
        </v-tooltip>
     
      </template>

      <template v-slot:item.ready="{ item }">
        <v-tooltip bottom max-width="600">
          <template v-slot:activator="{ on }">
            <span v-on="on" v-if="item.containerInfo.redynessProbeConfig">
              <v-chip
                x-small
                outlined
                :color="probeChipColor(item.ready)" text-color="white" 
              >{{probeType(item.containerInfo.redynessProbeConfig)}} probe</v-chip>
            </span>
            <span v-else-if="!item.ready">
              <v-chip x-small outlined text-color="white" :color="probeChipColor(item.ready)">not ready</v-chip>
            </span>
          </template>
          <span>{{ item.containerInfo.redynessProbeConfig }}</span>
        </v-tooltip>
      </template>

      <template v-slot:item.livenessprobe="{ item }">
        <v-tooltip bottom max-width="600">
          <template v-slot:activator="{ on }">
            <span v-on="on" v-if="item.containerInfo.livenessProbeConfig">
              <v-chip outlined x-small text-color="white"  color="black ">{{probeType(item.containerInfo.livenessProbeConfig)}} probe</v-chip>
            </span>
          </template>
          <span>{{ item.containerInfo.livenessProbeConfig }}</span>
        </v-tooltip>
      </template>

      <template v-slot:item.terminationReason="{ item }">
        <v-tooltip bottom max-width="600">
          <template v-slot:activator="{ on }">
            <span v-on="on" v-if="item.restartCount > 0">{{item.lastState.terminated.reason}}</span>
          </template>
          <span>{{ item.lastState}}</span>
        </v-tooltip>
      </template>
    </v-data-table>
  </v-container>
</template>

<script>
import { mapState } from "vuex";
import { mapGetters } from "vuex";
export default {
  mounted: function() {
    this.updateList();
  },
  computed: {
    ...mapGetters({
      containers: "containerStatuses"
    }),
    ...mapState({
      pods: state => state.kbzk8s.podinfo
    }),
    headers: function() {
      if (this.type == "containers") {
        return this.containerheaders;
      }
      return this.podheaders;
    },
    tdata: function() {
      if (this.type == "containers") {
        return this.containers;
      }
      return this.pods;
    }
  },

  methods: {
    updateList: function() {
      this.$store.dispatch("fetchPodInfo");
    },
    statusColour: function(cs) {
      if (!cs.ready) {
        return "red";
      }
    },
    probeChipColor: function(ready) {
      if (ready) {
        return "green";
      }
      return "orange";
    },
    probeType: function(cfg) {
      for (var k in Object.keys(cfg)) {
        switch (Object.keys(cfg)[k]) {
          case "httpGet":
            return "http";
          case "exec":
            return "exec";
          case "tcpSocket":
            return "socket";
        }
      }
      return "probed";
    },
    recentlyRestarted: function(status) {
      if (status.lastState == null) {
        return false;
      }
      if (status.lastState.terminated == null) {
        return false;
      }
      var rt = new Date(status.lastState.terminated.finishedAt);
      var MIN_15 = 15 * 60 * 1000; /* ms */
      if (new Date() - rt < MIN_15) {
        return true;
      }

      return false;
    },
    terminationReason: function (status)  {
      
      if (status.lastState == null) {
        return "";
      }
      if (status.lastState.terminated == null) {
        return "";
      }
      
      return status.lastState.terminated.reason
      

    }
    
  },

  data() {
    return {
      search: "",
      type: "pods",
      podheaders: [
        { text: "Namespace", value: "namespace" },
        {
          text: "Pod name",
          align: "start",
          sortable: true,
          value: "name"
        },
        { text: "Pod phase", value: "status.phase" },
        { text: "Conditions", value: "condition" },
        { text: "Container Status", value: "containerReason" },
        { text: "Restarts", value: "containerRestarts" },
        { text: "Quality  Of Service", value: "status.qosClass" },
        { text: "Reason", value: "reason" },
        { text: "Node", value: "status.hostIP" }
      ],
      containerheaders: [

        {
          text: "Container name",
          align: "start",
          sortable: true,
          value: "name"
        },
        { text: "State", value: "kbzState" },

        { text: "Ready", value: "ready", sortable: true },
        { text: "Liveness", value: "livenessprobe", sortable: false },
        //{ text: "Reason", value: "reason" },
        { text: "Restarts", value: "restartCount" },
        /*{
          text: "Termination reason",
          value: "terminationReason",
          sortable: false
        },*/
      { text: "Image Name", value: "image" },
      { text: "Pod", value: "kbzPod" },
      ]
    };
  }
};
</script>
<!--  script >
// @ is an alias to /src
//import axios from "axios";
/*import { mapState } from "vuex";
export default {
  name: "Api",
  components: {
    ApiEndpoint,
    Instructions,
    ApiNew
  },

  computed: mapState({
    apis: state => state.apiendpoint.apis
  }),
  methods: {
     createEndpoint: function(api) {
      this.$store.commit("NEW_API", api);
      this.updateList()
    },
    updateEndpoint: function(api) {
      this.$store.commit("UPDATE_API", api);
    },
    updateList: function() {
      this.$store.dispatch("fetchAPIEndpoints");
    }
  },
  mounted: function() {
    this.updateList()
  },
  data: function() {
    return {
      endpoints: {},
      dialog: false
    };
  }
};
</script -->
