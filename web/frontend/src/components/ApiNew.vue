<template>
  <v-dialog v-model="dialog" width="900">
    <template v-slot:activator="{ props }">
      <v-btn size="small" icon="mdi-plus-box" v-bind="props"></v-btn>
    </template>
    <v-card>
      <v-card-title>Create New API Endpoint</v-card-title>
      <v-card-text>
        <v-form ref="form">
          <v-row>
            <v-col>
              <v-select v-model="aep.namespace" :items="k8sstats.namespacesInCluster" label="Namespace"></v-select>
            </v-col>
            <v-col>
              <v-text-field v-model="aep.name" label="Name" required></v-text-field>
            </v-col>
            <v-col>
              <v-select v-model="aep.servicetype" :items="serviceTypes" label="Service Type"></v-select>
            </v-col>
            <v-col>
              <v-text-field v-model="aep.port" label="Listen Port" required></v-text-field>
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn size="small" color="blue-darken-1" variant="text" @click="dialog = false">Close</v-btn>
        <v-btn size="small" color="blue-darken-1" variant="text" @click="create()">Create</v-btn>
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
