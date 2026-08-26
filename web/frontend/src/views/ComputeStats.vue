<template>
  <v-container class="page-shell" fluid>
    <v-card class="page-panel mb-4" variant="flat">
      <PageHeader
        icon="mdi-chart-timeline-variant"
        title="Stats &amp; Load Tools"
        subtitle="Inspect container resources and generate controlled workloads"
      />
    </v-card>

    <v-row align="stretch">
      <v-col cols="12" lg="6">
        <v-card class="tool-card" height="100%" variant="outlined">
          <v-card-text>
            <CGroup />
          </v-card-text>
          <v-divider></v-divider>
          <v-card-actions class="pa-4">
            <LoadTools />
          </v-card-actions>
        </v-card>
      </v-col>
      <v-col cols="12" lg="6">
        <v-card class="tool-card" height="100%" variant="outlined">
          <K8SLoad />
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import CGroup from '@/components/CGroup.vue'
import LoadTools from '@/components/LoadTools.vue'
import K8SLoad from '@/components/K8SLoad.vue'
import PageHeader from '@/components/PageHeader.vue'

export default {
  name: 'ComputeStats',
  components: {
    CGroup,
    LoadTools,
    K8SLoad,
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
