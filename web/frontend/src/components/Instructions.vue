<template>
  <v-dialog v-model="dialog" scrollable max-width="800px">
    <template v-slot:activator="{ on }">
      <v-btn  small icon dark v-on="on">
       <v-icon >mdi-help-box</v-icon>
      </v-btn >
    </template>
    <v-card>
      <v-card-text  style="height: 300px;">
        <VueMarkdown class="mt-4" >{{content}}</VueMarkdown>
      </v-card-text>
      <v-divider></v-divider>
      <v-card-actions>
        <v-btn small color="blue darken-1" text @click="dialog = false">Close</v-btn>
      </v-card-actions>
      
    </v-card>
  </v-dialog>
</template>

<script>
import VueMarkdown from "vue-markdown";
import axios from "axios";

export default {
  name: "Instructions",
  props: [ "href"],
  components: {
    VueMarkdown
  },
  data: () => ({
    dialogm1: "",
    dialog: false,
    content: "## ddd "
  }),
  mounted: function() {
     axios.get(this.href).then(res => (this.content = res.data));
  },

};
</script>