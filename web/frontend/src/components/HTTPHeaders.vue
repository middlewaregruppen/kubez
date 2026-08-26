<template>
  <v-container fluid>
    <div class="tool-heading">
      <v-avatar color="blue-lighten-2" rounded="lg" size="38" variant="tonal">
        <v-icon icon="mdi-text-box-search-outline" size="22"></v-icon>
      </v-avatar>
      <div>
        <div class="tool-heading__title">HTTP Request Information</div>
        <div class="tool-heading__subtitle">Headers received by this kubez instance</div>
      </div>
    </div>

      <div class="request-summary">
        <v-chip color="blue-grey-lighten-2" size="small" variant="tonal">Protocol {{ httprequest.proto || 'unknown' }}</v-chip>
        <v-chip color="blue-grey-lighten-2" size="small" variant="tonal">Version {{ httprequest.major || '0' }}.{{ httprequest.minor || '0' }}</v-chip>

        <table class="headers-table">
          <thead>
            <tr>
              <th align="left">Header Name</th>
              <th align="left">Header Value</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(values, name) in httpheaders" :key="name">
              <td width="150">{{ name }}</td>
              <td v-for="(value, index) in values" :key="index">{{ value }}</td>
            </tr>
          </tbody>
        </table>
      </div>
  </v-container>
</template>

<script>
import { useInfoStore } from '@/stores/info'

export default {
  name: 'HTTPHeaders',
  setup() {
    return {
      infoStore: useInfoStore()
    }
  },
  computed: {
    httpheaders() {
      return this.infoStore.httpheaders ?? {}
    },
    httprequest() {
      return this.infoStore.requestInfo ?? {}
    }
  }
}
</script>

<style scoped>
.tool-heading {
  display: flex;
  align-items: center;
  gap: 0.7rem;
  margin-bottom: 1rem;
}

.tool-heading__title {
  font-size: 1rem;
  font-weight: 650;
}

.tool-heading__subtitle {
  color: rgba(255, 255, 255, 0.62);
  font-size: 0.75rem;
}

.request-summary {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.headers-table {
  width: 100%;
  margin-top: 0.75rem;
  overflow: hidden;
  border: 1px solid rgba(144, 202, 249, 0.14);
  border-collapse: separate;
  border-spacing: 0;
  border-radius: 8px;
}

.headers-table th {
  padding: 0.65rem 0.75rem;
  border-bottom: 1px solid #90caf9;
  background-color: #2e353b;
  font-size: 0.75rem;
}

.headers-table td {
  padding: 0.55rem 0.75rem;
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
  font-size: 0.8rem;
}
</style>
