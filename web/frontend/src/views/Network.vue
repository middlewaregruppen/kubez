<template>
  <v-container fluid>
    <v-row>
      <v-col>
        <v-card>
          <HTTPHeaders />
        </v-card>
      </v-col>
            <v-col>
        <v-card>
          <ConnectionCheck />
        </v-card>
   
         <v-card class="mt-3">
          <!--DNSLookup /-->
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
// @ is an alias to /src
import HTTPHeaders from "@/components/HTTPHeaders.vue";
import ConnectionCheck from "@/components/ConnectionCheck.vue";
//import DNSLookup from "@/components/DNSLookup.vue";

export default {
  name: "Network",
  components: {
    HTTPHeaders,
    ConnectionCheck, 
   // DNSLookup
  }
};
</script>
