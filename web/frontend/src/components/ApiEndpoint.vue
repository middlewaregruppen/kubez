<template>
  <v-container fluid>
    <v-list-item three-line>
      <v-list-item-content>
        <v-list-item-title class="subtitle mb-1">{{name}}</v-list-item-title>
        <v-list-item-subtitle>{{ep.path}}</v-list-item-subtitle>

        <div class="body-2" v-if="!collapsed">
          <v-form class="mt-4" ref="form">
            <v-text-field v-model="ep.path" label="Path" required></v-text-field>
            <v-row no-gutters>
              <v-col>
                <v-text-field
                  v-model="ep.delay.minTime"
                  label="Minimum delay in milliseconds"
                  placeholder="0"
                  :rules="delayMinRules"
                ></v-text-field>
              </v-col>
              <v-col>
                <v-text-field
                  v-model="ep.delay.maxTime"
                  label="Maximum delay in milliseconds"
                  placeholder="0"
                  :rules="delayRules"
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
            <!--v-select v-model="ep.methods" :items="methods" label="Accepted Methods" multiple></v-select-->
          </v-form>
        </div>
      </v-list-item-content>

      <v-list-item-icon  v-if="collapsed">
        <v-chip small outlined>{{delayRange}} delay</v-chip>
        <v-chip small outlined>{{ep.failureRate.rate}}% failure rate</v-chip>
        <v-chip small outlined>Static Relpy</v-chip>
            <v-btn icon class="ml-4" @click="collapsed = false">
      <v-icon  color="grey lighten-1">mdi-settings</v-icon>
    </v-btn>
      </v-list-item-icon>
    </v-list-item>
    <v-btn small class="ml-4 " v-if="!collapsed" @click="saveAndClose()">Save and Close</v-btn>
  </v-container>
</template>
<script>
import axios from "axios";
export default {
  name: "ApiEndpoint",
  props: ["name", "href"],
  mounted: function() {
   this.fetch()
  },

  methods: {
    fetch: function () {
      axios.get(this.href).then(res => (this.ep = res.data));
    },  
    update: function () {
      this.ep.delay.minTime = parseInt(this.ep.delay.minTime)
      this.ep.delay.maxTime = parseInt(this.ep.delay.maxTime)
      this.ep.failureRate.rate = parseInt(this.ep.failureRate.rate)

      axios.put(this.href, this.ep).then(res => (this.ep = res.data));
    },
    saveAndClose: function () {
      this.update();
      this.collapsed = true;
    },
    

  },
  

  computed: {
    delayRange() {
      if (this.ep.delay.maxTime == 0) {
        return "No";
      }

      if (this.ep.delay.minTime == this.ep.delay.maxTime) {
        return this.ep.delay.maxTime + " ms";
      }

      return this.ep.delay.minTime + "-" + this.ep.delay.maxTime + " ms";
    }
  },


  data: function() {
    return {
      ep: {
        delay: {},
        failureRate: {}
      },
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
        { type: "proxy", text: "Proxy requests to a server" },
        { type: "cruddb", text: "Database with CRUD operations" },
        { type: "random", text: "Random generated data" }
      ]
    };
  }
};
</script>
