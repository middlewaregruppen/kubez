<template>
  <div class="home">
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
      </v-col>
    </v-row>
  </div>
</template>

<script>
// @ is an alias to /src
import HTTPHeaders from "@/components/HTTPHeaders.vue";
import ConnectionCheck from "@/components/ConnectionCheck.vue";

export default {
  name: "Network",
  components: {
    HTTPHeaders,
    ConnectionCheck
  }
};
</script>
