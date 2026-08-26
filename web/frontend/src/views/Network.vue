<template>
  <v-container class="page-shell" fluid>
    <v-card class="page-panel mb-4" variant="flat">
      <PageHeader
        icon="mdi-lan"
        title="Network"
        subtitle="Inspect incoming requests and test outbound connectivity"
      />
    </v-card>

    <v-row align="stretch">
      <v-col cols="12" lg="7">
        <v-card class="tool-card" height="100%" variant="outlined">
          <HTTPHeaders />
        </v-card>
      </v-col>
      <v-col cols="12" lg="5">
        <v-card class="tool-card" height="100%" variant="outlined">
          <ConnectionCheck />
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import HTTPHeaders from '@/components/HTTPHeaders.vue'
import ConnectionCheck from '@/components/ConnectionCheck.vue'
import PageHeader from '@/components/PageHeader.vue'

export default {
  name: 'NetworkView',
  components: {
    HTTPHeaders,
    ConnectionCheck,
    PageHeader
  }
}
</script>

<style scoped>
.page-shell {
  padding: 1rem;
}

.page-panel,
.tool-card {
  border-color: rgba(144, 202, 249, 0.14);
  border-radius: 12px;
}

.page-panel {
  overflow: hidden;
}
</style>
