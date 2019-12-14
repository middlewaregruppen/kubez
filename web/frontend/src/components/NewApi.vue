<template>
  <v-dialog v-model="dialog" scrollable max-width="900px" max-height="500px">
    <template v-slot:activator="{ on }">
      <v-btn small icon dark v-on="on">
        <v-icon>mdi-plus-box</v-icon>
      </v-btn>
    </template>
    <v-card v-if="!apiConfigDialog">
      <v-card-text class="ma-4" style="height: 300px;">
        Create New Api Endpoint
        <v-form class="mt-4" ref="form">
          <v-text-field v-model="ep.name" label="Unique name for API Endpoint" required></v-text-field>
        </v-form>
      </v-card-text>
      <v-divider></v-divider>
      <v-card-actions>
        <v-btn small color="blue darken-1" text @click="next()">Next</v-btn>
        <v-btn small color="blue darken-1" text @click="dialog = false">Close</v-btn>
      </v-card-actions>
    </v-card>
    <v-card v-if="apiConfigDialog">
      <v-card-text class="ma-4" style="height: 500px;">
        API Endpoint created Successfully. Configure it.
        <ApiEndpoint expanded :href="href+ep.name" :name="ep.name" />
      </v-card-text>
      <v-card-actions>
        <v-btn small color="blue darken-1" text @click="dialog = false">Close</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import axios from "axios";
import ApiEndpoint from "@/components/ApiEndpoint.vue";

export default {
  name: "NewApi",
  props: ["href"],
  components: {
    ApiEndpoint
  },
  data: () => ({
    dialog: false,
    ep: {
      name: ""
    },
    apiConfigDialog: false
  }),

  methods: {
    next: function() {
      axios.post(this.href, this.ep).then((this.apiConfigDialog = true));
    }
  },
  mounted: function() {}
};
</script>