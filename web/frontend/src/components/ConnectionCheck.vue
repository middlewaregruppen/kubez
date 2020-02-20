<template>
  <v-container fluid>
    <v-list-item three-line>
      <v-list-item-content>
        <div class="headline mb-4">{{title}}</div>
        <v-list-item-subtitle>{{subtitle}}</v-list-item-subtitle>

        <div class="body-2">
          <v-text-field
            label="Host and port"
            hint="e.g middleware.se:443 or 10.5.12.12:22"
            v-model="target"
          ></v-text-field>
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
              >Connection to {{res.address}} successful</v-chip>
              <v-chip
                color="red darken-4"
                small
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
  name: "ConnectionCheck",
  props: {
    title: { type: String, default: "Connection Checker" },
    subtitle: {
      type: String,
      default: "Test TCP/IP connectivity to other systems from the container"
    },
    target: { type: String, default: "" }
  },
  methods: {
    check: function() {
      this.inProgress = true;
      axios.get("/kubez/connectioncheck/" + this.target).then(res => {
        this.res = res.data;
        this.inProgress = false;
        this.hasChecked = true;
      });
    }
  },
  data: () => ({
    inProgress: false,
    hasChecked: false,
    res: {}
  })
};
</script>
