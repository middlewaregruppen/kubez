<template>
  <v-container fluid>
    <v-list-item three-line>
      <v-list-item-content>
        <div class="headline mb-4">{{title}}</div>
        <v-list-item-subtitle>{{subtitle}}</v-list-item-subtitle>

        <div class="body-2">
          <v-row no-gutters>
            <v-col>
              <v-select
                v-model="recordType"
                :items="recordTypes"
                item-text="type"
                item-value="type"
                label="Type"
              ></v-select>
            </v-col>
            <v-col cols="8" class="ml-2">
              <v-text-field label="Name" hint="e.g middleware.se" v-model="recordName"></v-text-field>
            </v-col>
          </v-row>
          <v-row no-gutters>
            <v-col>
              <v-btn color="blue darken-1" @click="check()" text>Check</v-btn>
            </v-col>
            <v-col>
              <v-chip color="gray" small v-if="inProgress">Connecting ...</v-chip>
              <v-chip
                color="green darken-4"
                small
                v-if="res.success && !inProgress && hasChecked"
              >{{record}}</v-chip>
              <v-chip
                small
                color="red darken-4"
                v-if="!res.success && !inProgress && hasChecked"
              >{{res.error}}</v-chip>
            </v-col>
          </v-row>
        </div>
      </v-list-item-content>
    </v-list-item>
  </v-container>
</template>
<script>
import axios from "axios";
export default {
  name: "DNSLookup",
  props: {
    title: { type: String, default: "DNS Lookup" },
    subtitle: {
      type: String,
      default: "Look up DNS records"
    },
    recordType: { type: String, default: "" },
    recordName: { type: String, default: "" }
  },
  methods: {
    check: function() {
      this.inProgress = true;
      axios
        .get("/kubez/dnslookup/" + this.recordType + "/" + this.recordName)
        .then(res => {
          this.res = res.data;
          this.inProgress = false;
          this.hasChecked = true;
        });
    },
    responsesIdentical: function(serverResponses) {
      if (serverResponses) {
        //TODOY
      }
      // TODO!
      return true;
    }
  },

  computed: {
    record() {
      if (this.res.serverResponses[0].msg.Answer == null) {
        return "Not found";
      }
      if (!this.responsesIdentical) {
        return "";
      }

      switch (this.res.type) {
        // A-Records
        case "A":
          var isCname = false;
          var target = "";
          var as = [];
          for (var i in this.res.serverResponses[0].msg.Answer) {
            var a = this.res.serverResponses[0].msg.Answer[i];
            if (a["Target"] != null) {
              isCname = true;
              target = a.Target;
            }
            if (a["A"] != null) {
              as.push(a["A"]);
            }
          }
          // Single A record
          if (as.length == 1 && !isCname) {
            return "Resolves to " + as[0];
          }
          // CName -> no A:s?
          if (as.length == 0 && isCname) {
            return "Is a CNAME (" + target + ") that has no IP's ";
          }
          // CName -> one A?
          if (as.length == 1 && isCname) {
            return "Is a CNAME (" + target + ") resolves to " + as[0];
          }
          // Cname -> multiple A?
          if (as.length > 1 && isCname) {
            return (
              "Is a CNAME (" +
              target +
              ") that resolves to" +
              as.length +
              " IP's"
            );
          }
          // Multiple A:s
          if (as.length > 1 && !isCname) {
            return "Resolves to " + as.length + " ip addresses";
          }
          break;
        // SRV-Records
        case "SRV":
          // One SRV Response 
          if (this.res.serverResponses[0].msg.Answer.length == 1) {
            var s = this.res.serverResponses[0].msg.Answer[0];
            return  "prio: "+ s.Priority  + " weight: "+ s.Weight + " port: "+ s.Port +  " target: "+s.Target
          }
          // Multiple SRV Responses
          return "Resolves to "+ this.res.serverResponses[0].msg.Answer.length +" SRV records "
        

          
      }
      return "x";
    }
  },

  data: () => ({
    inProgress: false,
    hasChecked: false,
    res: { serverResponses: [] },
    recordTypes: [
      { type: "A", desc: "Ipv4 Host" },
      //  { type: "AAAA", desc: "Ipv6 Host" },
      { type: "SRV", desc: "Service Location" }
      /*{ type: "CNAME", desc: "Canonical name for an alias" },
      { type: "TXT", desc: "Descriptive Text" },
      { type: "MX", desc: "Mail eXchange" },
      { type: "NS", desc: "Name Server" },
      { type: "PTR", desc: "Pointer" },
      { type: "SOA", desc: "Start Of Authority" },
      { type: "ALIAS", desc: "Auto Resolved Alias" }*/
    ]
  })
};
</script>
