<template>
  <v-container class="page-shell" fluid>
    <v-card class="page-panel" variant="flat">
      <PageHeader
        icon="mdi-api"
        title="API Control Center"
        subtitle="Create and configure simulated API endpoints"
      >
        <template v-slot:actions>
          <ApiNew @createEndpoint="createEndpoint" />
          <v-btn
            aria-label="Refresh API endpoints"
            color="blue-grey-lighten-2"
            icon="mdi-refresh"
            variant="text"
            @click="updateList()"
          ></v-btn>
          <Instructions href="api-control-center.md" />
        </template>
      </PageHeader>

      <div class="endpoint-list pa-4">
        <v-card v-for="e in apis" :key="e.name" class="endpoint-card" variant="outlined">
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
            :runningpods="e.status?.runningpods ?? '0'"
            :servicetype="e.servicetype"
            :logtoconsole="e.logtoconsole"
            :cors="e.cors"
            @update="updateEndpoint"
          />
        </v-card>
        <div v-if="apiStore.loading" class="empty-state">
          <v-progress-circular color="blue-lighten-2" indeterminate></v-progress-circular>
          <span>Loading API endpoints</span>
        </div>
        <div v-else-if="apiStore.error" class="empty-state">
          <v-icon color="red-accent-2" icon="mdi-alert-circle-outline" size="34"></v-icon>
          <span>Unable to load API endpoints</span>
          <span class="empty-state__detail">{{ apiStore.error }}</span>
          <v-btn color="blue-lighten-2" prepend-icon="mdi-refresh" variant="outlined" @click="updateList()">Try Again</v-btn>
        </div>
        <div v-else-if="apis.length === 0" class="empty-state">
          <v-icon color="blue-grey-lighten-1" icon="mdi-api-off" size="34"></v-icon>
          <span>No API endpoints found</span>
        </div>
      </div>
    </v-card>
  </v-container>
</template>

<script>
import ApiEndpoint from '@/components/ApiEndpoint.vue'
import ApiNew from '@/components/ApiNew.vue'
import Instructions from '@/components/Instructions.vue'
import PageHeader from '@/components/PageHeader.vue'
import { useApiEndpointStore } from '@/stores/apiendpoint'

export default {
  name: 'ApiView',
  components: {
    ApiEndpoint,
    Instructions,
    ApiNew,
    PageHeader
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

<style scoped>
.page-shell {
  padding: 1rem;
}

.page-panel {
  overflow: hidden;
  border: 1px solid rgba(144, 202, 249, 0.14);
  border-radius: 12px;
}

.endpoint-list {
  display: grid;
  gap: 0.75rem;
}

.endpoint-card {
  border-color: rgba(144, 202, 249, 0.16);
  border-radius: 10px;
}

.endpoint-card:hover {
  border-color: rgba(144, 202, 249, 0.38);
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.6rem;
  padding: 3rem;
  color: rgba(255, 255, 255, 0.55);
  font-size: 0.875rem;
  text-align: center;
}

.empty-state__detail {
  max-width: 48rem;
  color: rgba(255, 255, 255, 0.42);
  font-size: 0.75rem;
}
</style>
