<template>
  <v-container fluid>
    <v-list-item three-line>
      <v-list-item-content>
        <v-list-item-title class="subtitle mb-1">{{name}}</v-list-item-title>
        <v-list-item-subtitle v-if="collapsed">http://*:{{ep.port}}{{ep.path}}</v-list-item-subtitle>

        <div class="body-2" v-if="!collapsed">
          <v-form class="mt-4" ref="form">
            <v-row no-gutters>
              <v-col>
                <v-text-field
                  v-model="ep.port"
                  label="Port"
                  hint="Port to listen to"
                  placeholder
                  required
                ></v-text-field>
              </v-col>
              <v-col cols="10">
                <v-text-field v-model="ep.path" label="Path" required></v-text-field>
              </v-col>
            </v-row>
            <v-row no-gutters>
              <v-col>
                <v-text-field
                  v-model="ep.delayMin"
                  label="Minimum delay in milliseconds"
                  placeholder="0"
                ></v-text-field>
              </v-col>
              <v-col>
                <v-text-field
                  v-model="ep.delayMax"
                  label="Maximum delay in milliseconds"
                  placeholder="0"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row no-gutters>
              <v-col>
                <v-text-field
                  v-model="ep.failureRate.rate"
                  label="Failure rate in %"
                  placeholder="0"
                ></v-text-field>
              </v-col>
              <v-col>
                <v-select
                  v-model="ep.failureRate.responseCodes"
                  :items="failureCodes"
                  item-text="text"
                  item-value="code"
                  label="Failure Code"
                ></v-select>
              </v-col>
            </v-row>
            <v-row no-gutters>
              <v-col>
                <v-text-field
                  v-model="ep.requestRate"
                  label="Request data transmission rate in bytes per seconds"
                  hint="0 means no limitation"
                  placeholder="0"
                ></v-text-field>
              </v-col>
              <v-col>
                <v-text-field
                  v-model="ep.responseRate"
                  label="Reply data transmission rate in bytes per seconds"
                  hint="0 means no limitation"
                  placeholder="0"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-select
              v-model="ep.response.type"
              :items="responseTypes"
              item-text="text"
              item-value="type"
              label="Response Handler"
            ></v-select>
            <v-textarea
              class="ml-3"
              v-if="ep.response.type=='static'"
              name="staticReplyContent"
              label="Static Reply Content"
              v-model="ep.response.static"
              hint="The payload that should be sent to the client"
            ></v-textarea>
            <v-row no-gutters class="ml-3" v-if="ep.response.type=='restdb'">
              Crud Contents..
              Database file path.
            </v-row>

            <v-row no-gutters>
              <v-col>
                <v-row no-gutters>
                  <v-col>
                    <v-checkbox v-model="ep.cors" label="CORS from all domains"></v-checkbox>
                  </v-col>
                  <v-col>
                    <v-checkbox v-model="ep.logToConsole" label="Log requests to console"></v-checkbox>
                  </v-col>
                </v-row>
                <v-row no-gutters></v-row>
              </v-col>
              <v-col>
                <v-textarea
                  name="replyHeaders"
                  label="Reply Headers"
                  v-model="replyHeaders"
                  hint="Headers to add to the response. One on each line. eg. Content-Type: application/json "
                ></v-textarea>
              </v-col>
            </v-row>
            <!--v-select v-model="ep.methods" :items="methods" label="Accepted Methods" multiple></v-select-->
          </v-form>
        </div>
      </v-list-item-content>

      <v-list-item-icon v-if="collapsed">
        <v-chip
          small
          v-if="ep.requestRate > 0"
          outlined
        >{{ep.requestRate| prettyBytes}} request rate</v-chip>
        <v-chip small v-if="delayRange != 'no'" outlined>{{delayRange}} delay</v-chip>
        <v-chip small v-if="ep.failureRate.rate > 0" outlined>{{ep.failureRate.rate}}% failure rate</v-chip>
        <v-chip small outlined>{{handlerType}}</v-chip>
        <v-chip
          small
          v-if="ep.responseRate > 0"
          outlined
        >{{ep.responseRate| prettyBytes}} response rate </v-chip>
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
import axios from "axios";
export default {
  name: "ApiEndpoint",

  props: {
    name: String,
    href: String,
    expanded: Boolean
  },

  mounted: function() {
    this.fetch();
    if (this.expanded) {
      this.collapsed = false;
    }
  },

  filters: {
    prettyBytes: function(num) {
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
    fetch: function() {
      axios.get(this.href).then(res => {
        this.ep = res.data;
        this.replyHeaders = this.httpHeadersToText();
      });
    },
    update: function() {
      this.ep.delayMin = parseInt(this.ep.delayMin);
      this.ep.delayMax = parseInt(this.ep.delayMax);
      this.ep.failureRate.rate = parseInt(this.ep.failureRate.rate);
      this.ep.responseRate = parseInt(this.ep.responseRate);
      this.ep.requestRate = parseInt(this.ep.requestRate);
      this.ep.port = parseInt(this.ep.port);
      this.ep.response.headers = this.httpHeadersToMap(this.replyHeaders);
      axios.put(this.href, this.ep).then(res => (this.ep = res.data));
    },
    saveAndCollapse: function() {
      this.update();
      this.collapsed = true;
    },
    httpHeadersToMap(text) {
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
    }
  },

  computed: {
    delayRange() {
      if (this.ep.delayMin == undefined) {
        return "no";
      }

      if (this.ep.delayMin > this.ep.delayMax) {
        return this.ep.delayMin + " ms";
      }

      if (this.ep.delayMin == this.ep.delayMax) {
        return this.ep.delayMax + " ms";
      }

      return this.ep.delayMin + "-" + this.ep.delayMax + " ms";
    },
    handlerType() {
      switch (this.ep.response.type) {
        case "static":
          return "Static Reply";
        case "echo":
          return "Echo back";
        case "restdb":
          return "REST Database";
      }
      return this.ep.response.type;
    }
  },

  data: function() {
    return {
      ep: {
        failureRate: {},
        response: {}
      },
      replyHeaders: "",
      collapsed: true,
      methods: [
        "get",
        "post",
        "put",
        "delete",
        "patch",
        "head",
        "options",
        "trace"
      ],
      failureCodes: [
        { code: 400, text: "400 - Bad Request" },
        { code: 401, text: "401 - Unauthorized" },
        { code: 403, text: "403 - Forbidden" },
        { code: 404, text: "404 - Not found" },
        { code: 405, text: "405 - Method not allowed" },
        { code: 406, text: "406 - Not acceptable" },
        { code: 408, text: "408 - Request timeout" },
        { code: 409, text: "409 - Conflict" },
        { code: 410, text: "410 - Gone" },
        { code: 411, text: "411 - Lenght required" },
        { code: 412, text: "412 - Precondition failed" },
        { code: 413, text: "413 - Payload too large" },
        { code: 414, text: "414 - URI too long" },
        { code: 415, text: "415 - Unsupported media type" },
        { code: 416, text: "416 - Requested range not satisfiable" },
        { code: 417, text: "417 - Expectation failed" },
        { code: 418, text: "418 - I'm a teapot" },
        { code: 421, text: "421 - Misdirected request" },
        { code: 425, text: "425 - To early" },
        { code: 426, text: "426 - Upgrade requried" },
        { code: 428, text: "428 - Precondition required" },
        { code: 429, text: "429 - Too many requests" },
        { code: 431, text: "431 - Request fields to large" },
        { code: 451, text: "451 - Unavaialable for leagal reasons" },
        { code: 500, text: "500 - Internal server error" },
        { code: 502, text: "502 - Bad gateway" },
        { code: 503, text: "503 - Service unavalable" },
        { code: 504, text: "504 - Gateway timeout" },
        { code: 505, text: "505 - HTTP version not supported" },
        { code: 510, text: "510 - Not extended" },
        { code: 511, text: "511 - Network authenetication required" }
      ],
      responseTypes: [
        { type: "static", text: "Static Response" },
        { type: "echo", text: "Echo Request to Response" },
       // { type: "restdb", text: "REST Database" }
        //  { type: "forward", text: "Forward requests" },
        //  { type: "random", text: "Random generated data" }
      ]
    };
  }
};
</script>
