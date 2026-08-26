<template>
  <v-container fluid>
    <v-list-item class="endpoint-summary" lines="three">
      <v-list-item-title class="text-subtitle-1 mb-1">{{ name }}</v-list-item-title>
      <v-list-item-subtitle v-if="collapsed">
        <v-chip class="mr-1 mt-1" color="blue-grey-lighten-2" size="small" variant="tonal">namespace: {{ namespace }}</v-chip>
        <v-chip class="mr-1 mt-1" color="blue-grey-lighten-2" size="small" variant="tonal">pods: {{ runningpods }}</v-chip>
        <v-chip class="mr-1 mt-1" color="blue-grey-lighten-2" size="small" variant="tonal">port: {{ port }}</v-chip>
        <v-chip class="mr-1 mt-1" color="blue-grey-lighten-2" size="small" variant="tonal">{{ servicetype }}</v-chip>
      </v-list-item-subtitle>

      <div class="text-body-2" v-if="!collapsed">
        <v-form class="mt-4" ref="form">
          <v-row no-gutters>
            <v-col>
              <v-text-field v-model="aep.mindelay" label="Minimum delay in milliseconds" placeholder="0"></v-text-field>
            </v-col>
            <v-col>
              <v-text-field v-model="aep.maxdelay" label="Maximum delay in milliseconds" placeholder="0"></v-text-field>
            </v-col>
            <v-col class="ml-3">
              <v-text-field v-model="aep.requestrate" label="Request data transmission rate in bytes per seconds" hint="0 is no limitation" placeholder="0"></v-text-field>
            </v-col>
            <v-col>
              <v-text-field v-model="aep.responserate" label="Reply data transmission rate in bytes per seconds" hint="0 is no limitation" placeholder="0"></v-text-field>
            </v-col>
          </v-row>
          <v-row no-gutters>
            <v-col>
              <v-text-field v-model="aep.failurerate" label="Failure rate in %" placeholder="0"></v-text-field>
            </v-col>
            <v-col class="ml-3">
              <v-select v-model="aep.failurecode" :items="failureCodes" item-title="text" item-value="code" label="Failure Code"></v-select>
            </v-col>
          </v-row>
          <v-select v-model="aep.responsetype" :items="responseTypes" item-title="text" item-value="type" label="Response Handler"></v-select>
          <v-textarea
            class="ml-3"
            v-if="aep.responsetype === 'static'"
            label="Static Reply Content"
            v-model="aep.staticcontent"
            hint="The payload that should be sent to the client"
          ></v-textarea>
          <v-row no-gutters>
            <v-col>
              <v-row no-gutters>
                <v-col>
                  <v-checkbox v-model="aep.cors" label="CORS from all domains"></v-checkbox>
                </v-col>
                <v-col>
                  <v-checkbox v-model="aep.logtoconsole" label="Log requests to console"></v-checkbox>
                </v-col>
              </v-row>
            </v-col>
          </v-row>
        </v-form>
      </div>

      <template v-slot:append v-if="collapsed">
        <div class="d-flex flex-column align-end ga-1">
          <v-chip size="small" v-if="aep.requestrate > 0" variant="outlined">{{ prettyBytesRate(aep.requestrate) }} request rate</v-chip>
          <v-chip size="small" v-if="delayRange !== 'no'" variant="outlined">{{ delayRange }} delay</v-chip>
          <v-chip size="small" v-if="aep.failurerate > 0" variant="outlined">{{ aep.failurerate }}% failure rate</v-chip>
          <v-chip size="small" variant="outlined">{{ aep.responsetype }}</v-chip>
          <v-chip size="small" v-if="aep.responserate > 0" variant="outlined">{{ prettyBytesRate(aep.responserate) }} response rate</v-chip>
          <v-btn color="blue-lighten-2" icon="mdi-tune-variant" variant="text" class="mt-1" @click="collapsed = false"></v-btn>
        </div>
      </template>
    </v-list-item>

    <v-btn color="blue-lighten-2" size="small" variant="outlined" class="ml-4 mb-3" v-if="!collapsed && !expanded" @click="saveAndCollapse()">Apply and Close</v-btn>
    <v-btn color="blue-lighten-2" size="small" variant="outlined" class="ml-4 mb-3" v-if="expanded" @click="update()">Apply</v-btn>
  </v-container>
</template>

<script>
import { prettyBytesRate } from '@/utils/format'

export default {
  name: 'ApiEndpoint',
  props: {
    namespace: String,
    name: String,
    port: String,
    path: String,
    mindelay: String,
    maxdelay: String,
    failurerate: String,
    failurecode: String,
    requestrate: String,
    responserate: String,
    staticcontent: String,
    responsetype: String,
    runningpods: String,
    servicetype: String,
    logtoconsole: Boolean,
    cors: Boolean,
    expanded: Boolean
  },
  mounted() {
    if (this.expanded) {
      this.collapsed = false
    }
  },
  methods: {
    prettyBytesRate,
    saveAndCollapse() {
      this.$emit('update', this.aep)
      this.collapsed = true
    }
  },
  computed: {
    delayRange() {
      if (this.aep.mindelay == undefined) return 'no'
      const min = parseInt(this.aep.mindelay)
      const max = parseInt(this.aep.maxdelay)
      if (min > max) return min + ' ms'
      if (min === max) return max + ' ms'
      return min + '-' + max + ' ms'
    }
  },
  data() {
    return {
      aep: {
        namespace: this.namespace,
        name: this.name,
        port: this.port,
        path: this.path,
        mindelay: this.mindelay,
        maxdelay: this.maxdelay,
        failurerate: this.failurerate,
        failurecode: this.failurecode,
        requestrate: this.requestrate,
        responserate: this.responserate,
        responsetype: this.responsetype,
        staticcontent: this.staticcontent,
        runningpods: this.runningpods,
        servicetype: this.servicetype,
        logtoconsole: this.logtoconsole,
        cors: this.cors
      },
      collapsed: true,
      failureCodes: [
        { code: '400', text: '400 - Bad Request' },
        { code: '401', text: '401 - Unauthorized' },
        { code: '403', text: '403 - Forbidden' },
        { code: '404', text: '404 - Not found' },
        { code: '405', text: '405 - Method not allowed' },
        { code: '406', text: '406 - Not acceptable' },
        { code: '408', text: '408 - Request timeout' },
        { code: '409', text: '409 - Conflict' },
        { code: '410', text: '410 - Gone' },
        { code: '411', text: '411 - Length required' },
        { code: '412', text: '412 - Precondition failed' },
        { code: '413', text: '413 - Payload too large' },
        { code: '414', text: '414 - URI too long' },
        { code: '415', text: '415 - Unsupported media type' },
        { code: '416', text: '416 - Requested range not satisfiable' },
        { code: '417', text: '417 - Expectation failed' },
        { code: '418', text: "418 - I'm a teapot" },
        { code: '421', text: '421 - Misdirected request' },
        { code: '425', text: '425 - Too early' },
        { code: '426', text: '426 - Upgrade required' },
        { code: '428', text: '428 - Precondition required' },
        { code: '429', text: '429 - Too many requests' },
        { code: '431', text: '431 - Request fields too large' },
        { code: '451', text: '451 - Unavailable for legal reasons' },
        { code: '500', text: '500 - Internal server error' },
        { code: '502', text: '502 - Bad gateway' },
        { code: '503', text: '503 - Service unavailable' },
        { code: '504', text: '504 - Gateway timeout' },
        { code: '505', text: '505 - HTTP version not supported' },
        { code: '510', text: '510 - Not extended' },
        { code: '511', text: '511 - Network authentication required' }
      ],
      responseTypes: [
        { type: 'static', text: 'Static Response' },
        { type: 'echo', text: 'Echo Request to Response' }
      ]
    }
  }
}
</script>

<style scoped>
.endpoint-summary {
  padding: 0.75rem 1rem;
}
</style>
