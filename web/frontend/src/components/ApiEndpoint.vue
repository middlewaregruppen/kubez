<template>
  <v-container fluid>
    <v-list-item three-line>
      <v-list-item-content>
        <v-list-item-title class="subtitle mb-1">{{name}}</v-list-item-title> 
        <v-list-item-subtitle v-if="collapsed">
          <v-chip small outlined>namespace: {{namespace}}</v-chip>
          <v-chip small outlined>pods: {{runningpods}}</v-chip>
          <v-chip small outlined>port: {{port}}</v-chip>
          <v-chip small outlined> {{servicetype}}</v-chip>
          
        </v-list-item-subtitle>

        <div class="body-2" v-if="!collapsed">
          <v-form class="mt-4" ref="form">
            <v-row no-gutters>
              <v-col>
                <v-text-field
                  v-model="aep.mindelay"
                  label="Minimum delay in milliseconds"
                  placeholder="0"
                ></v-text-field>
              </v-col>
              <v-col>
                <v-text-field
                  v-model="aep.maxdelay"
                  label="Maximum delay in milliseconds"
                  placeholder="0"
                ></v-text-field>
              </v-col>

              <v-col class="ml-3">
                <v-text-field
                  v-model="aep.requestrate"
                  label="Request data transmission rate in bytes per seconds"
                  hint="0 is no limitation"
                  placeholder="0"
                ></v-text-field>
              </v-col>
              <v-col>
                <v-text-field
                  v-model="aep.responserate"
                  label="Reply data transmission rate in bytes per seconds"
                  hint="0 is no limitation"
                  placeholder="0"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row no-gutters>
              <v-col>
                <v-text-field v-model="aep.failurerate" label="Failure rate in %" placeholder="0"></v-text-field>
              </v-col>
              <v-col class="ml-3">
                <v-select
                  v-model="aep.failurecode"
                  :items="failureCodes"
                  item-text="text"
                  item-value="code"
                  label="Failure Code"
                ></v-select>
              </v-col>
            </v-row>
            <v-row no-gutters></v-row>
            <v-select
              v-model="aep.responsetype"
              :items="responseTypes"
              item-text="text"
              item-value="type"
              label="Response Handler"
            ></v-select>
            <v-textarea
              class="ml-3"
              v-if="aep.responsetype=='static'"
              label="Static Reply Content"
              v-model="aep.staticcontent"
              hint="The payload that should be sent to the client"
            ></v-textarea>
            <v-row no-gutters class="ml-3" v-if="aep.responsetype=='restdb'"></v-row>

            <v-row no-gutters>
              <v-col>
                <v-row no-gutters>
                  <v-col>
                    <v-checkbox v-model="aep.cors" label="CORS from all domains"></v-checkbox>
                  </v-col>
                  <v-col>
                    <v-checkbox v-model="aep.logtoconsole" label="Log requests to console"></v-checkbox>
                  </v-col>
                </v-row>
                <v-row no-gutters></v-row>
              </v-col>
              <v-col>
                <!--v-textarea
                  name="replyHeaders"
                  label="Reply Headers"
                  v-model="replyHeaders"
                  hint="Headers to add to the response. One on each line. eg. Content-Type: application/json "
                ></v-textarea-->
              </v-col>
            </v-row>
          </v-form>
        </div>
      </v-list-item-content>

      <v-list-item-icon v-if="collapsed">
        <v-chip
          small
          v-if="aep.requestrate > 0"
          outlined
        >{{aep.requestrate| prettyBytes}} request rate</v-chip>
        <v-chip small v-if="delayRange != 'no'" outlined>{{delayRange}} delay</v-chip>
        <v-chip small v-if="aep.failurerate > 0" outlined>{{aep.failurerate}}% failure rate</v-chip>
        <v-chip small outlined>{{aep.responsetype}}</v-chip>
        <v-chip
          small
          v-if="aep.responserate > 0"
          outlined
        >{{aep.responserate| prettyBytes}} response rate</v-chip>
        <v-btn icon class="ml-4" @click="collapsed = false">
          <v-icon color="grey lighten-1">mdi-settings</v-icon>
        </v-btn>
      </v-list-item-icon>
    </v-list-item>
    <v-btn
      small
      class="ml-4"
      v-if="!collapsed && !expanded "
      @click="saveAndCollapse()"
    >Apply and Close</v-btn>
    <v-btn small class="ml-4" v-if="expanded" @click="update()">Apply</v-btn>
  </v-container>
</template>
<script>
export default {
  name: "ApiEndpoint",

  props: {
    namespace: String,
    name: String,
    port: String,
    path: String,
    mindelay: String,
    maxdelay: String,
    failurerate: String,
    failurecode: String,
    requestrate: String,
    responserate: String,
    staticcontent: String,
    responsetype: String,
    runningpods: String,
    servicetype: String,
    logtoconsole: Boolean,
    cors:Boolean,


    expanded: Boolean
  },

  mounted: function() {
    if (this.expanded) {
      this.collapsed = false;
    }
  },

  filters: {
    prettyBytes: function(num) {
      num = parseInt(num);
      if (typeof num !== "number" || isNaN(num)) {
        return num;
      }

      if (num == 0) {
        return "no";
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

      return (neg ? "-" : "") + num + " " + unit + "/s";
    }
  },

  methods: {
    saveAndCollapse: function() {
      this.$emit("update", this.aep);
      this.collapsed = true;
    }
    /*httpHeadersToMap(text) {
      var res = {};
      var lines = text.split("\n");
      lines.forEach(element => {
        var kv = element.split(":");
        if (kv[1] != undefined) {
          res[kv[0]] = kv[1].trim();
        }
      });
      return res;
    },
    httpHeadersToText() {
      this.ep.response;
      if (this.ep.response.headers == undefined) return "";

      var res = "";
      for (const [key, value] of Object.entries(this.ep.response.headers)) {
        res = res + key + ": " + value + "\n";
      }
      return res;
    }*/
  },

  computed: {
    delayRange() {
      if (this.aep.mindelay == undefined) {
        return "no";
      }

      let min = parseInt(this.aep.mindelay);
      let max = parseInt(this.aep.maxdelay);

      if (min > max) {
        return min + " ms";
      }

      if (min == max) {
        return max + " ms";
      }

      return min + "-" + max + " ms";
    },
    handlerType() {
      switch (this.aep.responsetype) {
        case "static":
          return "Static Reply";
        case "echo":
          return "Echo back";
        case "restdb":
          return "REST Database";
      }
      return this.aep.responsetype;
    }
  },

  data: function() {
    return {
      ep: {
        failureRate: {},
        response: {}
      },
      aep: {
        namespace: this.namespace,
        name: this.name,
        port: this.port,
        path: this.path,
        mindelay: this.mindelay,
        maxdelay: this.maxdelay,
        failurerate: this.failurerate,
        failurecode: this.failurecode,
        requestrate: this.requestrate,
        responserate: this.responserate,
        responsetype: this.responsetype,
        staticcontent: this.staticcontent,
        runningpods: this.runningpods,
        servicetype: this.servicetype,
        logtoconsole: this.logtoconsole,
        cors: this.cors,
      },
      replyHeaders: "",
      collapsed: true,
      failureCodes: [
        { code: "400", text: "400 - Bad Request" },
        { code: "401", text: "401 - Unauthorized" },
        { code: "403", text: "403 - Forbidden" },
        { code: "404", text: "404 - Not found" },
        { code: "405", text: "405 - Method not allowed" },
        { code: "406", text: "406 - Not acceptable" },
        { code: "408", text: "408 - Request timeout" },
        { code: "409", text: "409 - Conflict" },
        { code: "410", text: "410 - Gone" },
        { code: "411", text: "411 - Lenght required" },
        { code: "412", text: "412 - Precondition failed" },
        { code: "413", text: "413 - Payload too large" },
        { code: "414", text: "414 - URI too long" },
        { code: "415", text: "415 - Unsupported media type" },
        { code: "416", text: "416 - Requested range not satisfiable" },
        { code: "417", text: "417 - Expectation failed" },
        { code: "418", text: "418 - I'm a teapot" },
        { code: "421", text: "421 - Misdirected request" },
        { code: "425", text: "425 - To early" },
        { code: "426", text: "426 - Upgrade requried" },
        { code: "428", text: "428 - Precondition required" },
        { code: "429", text: "429 - Too many requests" },
        { code: "431", text: "431 - Request fields to large" },
        { code: "451", text: "451 - Unavaialable for leagal reasons" },
        { code: "500", text: "500 - Internal server error" },
        { code: "502", text: "502 - Bad gateway" },
        { code: "503", text: "503 - Service unavalable" },
        { code: "504", text: "504 - Gateway timeout" },
        { code: "505", text: "505 - HTTP version not supported" },
        { code: "510", text: "510 - Not extended" },
        { code: "511", text: "511 - Network authenetication required" }
      ],
      responseTypes: [
        { type: "static", text: "Static Response" },
        { type: "echo", text: "Echo Request to Response" }
        // { type: "restdb", text: "REST Database" }
        //  { type: "forward", text: "Forward requests" },
        //  { type: "random", text: "Random generated data" }
      ]
    };
  }
};
</script>
