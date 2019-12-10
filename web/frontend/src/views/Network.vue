<template>
  <div class="home">
    <v-row>
         <v-col>
        <v-card>
          <HTTPHeaders />
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
// @ is an alias to /src
import HTTPHeaders from "@/components/HTTPHeaders.vue";

export default {
  name: "Network",
  components: {
    HTTPHeaders
  }
};
</script>
