<template>
  <v-container fluid>
    <v-list-item three-line>
      <v-list-item-content>
        <div class="headline mb-4">HTTP HEADERS FROM INCOMING REQUESTS</div>
        <div class="body-2 pl-3">
          <table>
            <thead>
              <tr><th align="left">NAME</th><th align="left">VALUE</th></tr>
            </thead>
            <tbody>
              <tr v-for="(values,name) in httpheaders" v-bind:key="name">
                <td width="150">{{name}}</td>
                <td v-for="(value,index) in values" v-bind:key="index">{{value}}</td>
              </tr>
            </tbody>
          </table>  
          <!-- <div v-for="(values,name) in httpheaders" v-bind:key="name">{{name}} - <span v-for="(value,index) in values" v-bind:key="index">{{value}}</span></div> -->
        </div>
      </v-list-item-content>
    </v-list-item>
  </v-container>
</template>
<script>
export default {
  name: "HTTPHeaders",

  computed: {
    birds() {
      return "Fly away to Thailand "
    },
    httpheaders() {
      return this.$store.state.info.httpheaders;
    },
  },
};
</script>
