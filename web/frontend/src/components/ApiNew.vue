<template>
  <v-dialog v-model="dialog" width="900">
    <template v-slot:activator="{ props }">
      <v-btn
        aria-label="Create API endpoint"
        color="blue-lighten-2"
        icon="mdi-plus-circle-outline"
        variant="text"
        v-bind="props"
      ></v-btn>
    </template>
    <v-card>
      <v-card-title class="dialog-title">
        <v-avatar color="blue-lighten-2" rounded="lg" size="36" variant="tonal">
          <v-icon icon="mdi-api" size="21"></v-icon>
        </v-avatar>
        Create New API Endpoint
      </v-card-title>
      <v-card-text class="pt-4">
        <v-form ref="form">
          <v-row>
            <v-col>
              <v-select v-model="aep.namespace" :items="k8sstats.namespacesInCluster" density="comfortable" label="Namespace" variant="outlined"></v-select>
            </v-col>
            <v-col>
              <v-text-field v-model="aep.name" density="comfortable" label="Name" required variant="outlined"></v-text-field>
            </v-col>
            <v-col>
              <v-select v-model="aep.servicetype" :items="serviceTypes" density="comfortable" label="Service Type" variant="outlined"></v-select>
            </v-col>
            <v-col>
              <v-text-field v-model="aep.port" density="comfortable" label="Listen Port" required variant="outlined"></v-text-field>
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue-grey-lighten-2" variant="text" @click="dialog = false">Cancel</v-btn>
        <v-btn color="blue-lighten-2" prepend-icon="mdi-plus" variant="outlined" @click="create()">Create Endpoint</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { useInfoStore } from '@/stores/info'

export default {
  name: 'ApiNew',
  setup() {
    return {
      infoStore: useInfoStore()
    }
  },
  data: () => ({
    dialog: false,
    aep: {
      namespace: '',
      name: '',
      port: '1337'
    },
    serviceTypes: ['clusterIP', 'nodePort', 'loadBalancer']
  }),
  computed: {
    k8sstats() {
      return this.infoStore.k8sstats ?? {}
    }
  },
  methods: {
    create() {
      this.$emit('createEndpoint', this.aep)
      this.dialog = false
    }
  }
}
</script>

<style scoped>
.dialog-title {
  display: flex;
  align-items: center;
  gap: 0.7rem;
  border-bottom: 1px solid rgba(144, 202, 249, 0.14);
}
</style>
