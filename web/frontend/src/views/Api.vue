<template>
  <v-container fluid>
    <v-row no-gutters>
      <v-col>
        <v-card flat>
          <ApiNew @createEndpoint="createEndpoint" />
          <v-btn size="small" icon="mdi-refresh" @click="updateList()"></v-btn>
          <Instructions href="api-control-center.md" />
        </v-card>
      </v-col>
    </v-row>

    <v-row v-for="e in apis" :key="e.name">
      <v-col>
        <v-card>
          <ApiEndpoint
            :namespace="e.namespace"
            :name="e.name"
            :port="e.port"
            :path="e.path"
            :mindelay="e.mindelay"
            :maxdelay="e.maxdelay"
            :failurerate="e.failurerate"
            :failurecode="e.failurecode"
            :requestrate="e.requestrate"
            :responserate="e.responserate"
            :responsetype="e.responsetype"
            :staticcontent="e.staticcontent"
            :runningpods="e.status.runningpods"
            :servicetype="e.servicetype"
            :logtoconsole="e.logtoconsole"
            :cors="e.cors"
            @update="updateEndpoint"
          />
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import ApiEndpoint from '@/components/ApiEndpoint.vue'
import ApiNew from '@/components/ApiNew.vue'
import Instructions from '@/components/Instructions.vue'
import { useApiEndpointStore } from '@/stores/apiendpoint'

export default {
  name: 'Api',
  components: {
    ApiEndpoint,
    Instructions,
    ApiNew
  },
  setup() {
    return {
      apiStore: useApiEndpointStore()
    }
  },
  computed: {
    apis() {
      return this.apiStore.apis
    }
  },
  methods: {
    createEndpoint(api) {
      this.apiStore.createNewAPI(api).then(() => this.updateList())
    },
    updateEndpoint(api) {
      this.apiStore.updateAPI(api)
    },
    updateList() {
      this.apiStore.fetchAPIEndpoints()
    }
  },
  mounted() {
    this.updateList()
  }
}
</script>
