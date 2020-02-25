<template>
  <v-container fluid>
    <v-list-item three-line>
      <v-list-item-content>
        <div class="headline mb-4">HTTP Request information</div>

        <div class="body-2 pl-3">

          Protocol: {{httprequest.proto}}<br/>
          Major Version: {{httprequest.major}}<br/>
          Minor Version: {{httprequest.minor}}<br/>

          <table class="pt-2">
            <thead>
              <tr><th align="left">Header Name</th><th align="left">Header Value</th></tr>
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
    
    httpheaders() {
      return this.$store.state.info.httpheaders;
    },
    httprequest() {
      return this.$store.state.info.requestInfo;
    }
  },
};
</script>
