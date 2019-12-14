<template>
  <v-container fluid>
    <v-list-item three-line>
      <v-list-item-content>
        <v-list-item-title class="subtitle mb-1">{{name}}</v-list-item-title>
        <v-list-item-subtitle v-if="collapsed">{{ep.path}}</v-list-item-subtitle>

        <div class="body-2" v-if="!collapsed">
          <v-form class="mt-4" ref="form">
            <v-text-field v-model="ep.path" label="Path" required></v-text-field>
            <v-row no-gutters>
              <v-col>
                <v-text-field
                  v-model="ep.delay.minTime"
                  label="Minimum delay in milliseconds"
                  placeholder="0"
                ></v-text-field>
              </v-col>
              <v-col>
                <v-text-field
                  v-model="ep.delay.maxTime"
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
            <v-select
              v-model="ep.response.type"
              :items="responseTypes"
              item-text="text"
              item-value="type"
              label="Response Handler"
            ></v-select>
            <v-textarea
              v-if="ep.response.type=='static'"
              name="staticReplyContent"
              label="Static Reply Content"
              v-model="ep.response.static"
              hint="The payload that should be sent to the client"
            ></v-textarea>

            <v-row no-gutters>
              <v-col>
                <v-switch v-model="ep.cors" class="ml-2" label="Promiscuous Cross-Origin Resource Sharing (CORS)"></v-switch>
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
        <v-chip small outlined>{{delayRange}} delay</v-chip>
        <v-chip small outlined>{{ep.failureRate.rate}}% failure rate</v-chip>
        <v-chip small outlined>{{handlerType}}</v-chip>
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
    >Save and Close</v-btn>
    <v-btn small class="ml-4" v-if="expanded" @click="update()">Save</v-btn>
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

  methods: {
    fetch: function() {
      axios.get(this.href).then(res => {
        this.ep = res.data;
        this.replyHeaders = this.httpHeadersToText();
      });
    },
    update: function() {
      this.ep.delay.minTime = parseInt(this.ep.delay.minTime);
      this.ep.delay.maxTime = parseInt(this.ep.delay.maxTime);
      this.ep.failureRate.rate = parseInt(this.ep.failureRate.rate);
      this.ep.response.headers = this.httpHeadersToMap(this.replyHeaders);
      window.console.log(this.replyHeaders);
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
      if (this.ep.delay.minTime == 0) {
        return "No";
      }

      if (this.ep.delay.minTime > this.ep.delay.maxTime) {
        return this.ep.delay.minTime + " ms";
      }

      if (this.ep.delay.minTime == this.ep.delay.maxTime) {
        return this.ep.delay.maxTime + " ms";
      }

      return this.ep.delay.minTime + "-" + this.ep.delay.maxTime + " ms";
    },
    handlerType() {
      switch (this.ep.response.type){
        case "static":
          return "Static Reply"
        case "random":
          return "Random Data"
      }
    return this.ep.response.type
    }

  },

  data: function() {
    return {
      ep: {
        delay: {},
        failureRate: {}
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
        { type: "static", text: "Static Response" }
        //  { type: "proxy", text: "Proxy requests" },
        //  { type: "cruddb", text: "CRUD database" },
        //   { type: "random", text: "Random generated data" }
      ]
    };
  }
};
</script>
