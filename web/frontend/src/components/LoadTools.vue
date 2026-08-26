<template>
  <v-container fluid>
    <v-row align-content="space-around" no-gutters>
      <v-col cols="2">
        <v-btn size="small" @click="loadCpu">Load CPU</v-btn>
        <br />
        <span v-if="cpu !== '-'" class="text-caption">{{ cpu }}</span>
      </v-col>
      <v-col cols="2">
        <v-btn size="small" @click="malloc">Allocate 20 Mb</v-btn>
        <br />
        <span v-if="memory !== '-'" class="text-caption">{{ memory }}</span>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from 'axios'

export default {
  name: 'LoadTools',
  methods: {
    loadCpu() {
      this.cpu = 'Requesting ...'
      axios.post('/kubez/action/cpumedium').then(res => {
        this.cpu = res.data
      })
    },
    malloc() {
      this.memory = 'Requesting ...'
      axios.post('/kubez/action/malloc20mb').then(res => {
        this.memory = res.data
      })
    }
  },
  data() {
    return {
      cpu: '-',
      memory: '-'
    }
  }
}
</script>

<style scoped>
.text-caption {
  padding-left: 2px;
}
</style>
