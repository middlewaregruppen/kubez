<template>
  <v-container fluid>
    <v-list-item three-line>
      <v-list-item-content>
        <div class="overline mb-4">This is my header</div>
        <v-list-item-title class="headline mb-1">Lolbotz</v-list-item-title>
        <v-list-item-subtitle>This is SPARTAAAA!</v-list-item-subtitle>
        <div class="body-2 pl-3">
                Hello ALLLAN!
              <div>{{ headers }}</div>
        </div>
      </v-list-item-content>

     
    </v-list-item>

  </v-container>
</template>
<script>
import axios from "axios"

export default {
  name: "HTTPHeaders",

  methods: {
    printData: function () {
      this.networkData = 'Requesting network data...'
      //console.log('Getting some headers');
      axios.get('/kubez/info').then(res => (this.networkData = res.data));
      return 'This is Vue!!'
    }
  }
};
</script>
