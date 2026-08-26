<template>
  <v-container fluid>
    <v-list-item lines="three">
      <div class="text-h6 mb-4">HTTP Request information</div>
      <div class="text-body-2 pl-3">
        Protocol: {{ httprequest.proto }}<br />
        Major Version: {{ httprequest.major }}<br />
        Minor Version: {{ httprequest.minor }}<br />

        <table class="mt-2">
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
    </v-list-item>
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
