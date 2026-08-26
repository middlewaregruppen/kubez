<template>
  <v-container fluid>
    <v-row>
      <v-col>
        <v-card flat>
          <HTTPHeaders />
        </v-card>
      </v-col>
      <v-col>
        <v-card flat>
          <ConnectionCheck />
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import HTTPHeaders from '@/components/HTTPHeaders.vue'
import ConnectionCheck from '@/components/ConnectionCheck.vue'

export default {
  name: 'Network',
  components: {
    HTTPHeaders,
    ConnectionCheck
  }
}
</script>
