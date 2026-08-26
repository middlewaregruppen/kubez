<template>
  <v-container fluid>
    <div class="load-tools-heading">Container load actions</div>
    <div class="load-actions">
      <div class="load-action">
        <v-btn color="blue-lighten-2" prepend-icon="mdi-cpu-64-bit" variant="outlined" @click="loadCpu">Load CPU</v-btn>
        <span v-if="cpu !== '-'" class="load-action__result">{{ cpu }}</span>
      </div>
      <div class="load-action">
        <v-btn color="blue-lighten-2" prepend-icon="mdi-memory" variant="outlined" @click="malloc">Allocate 20 MB</v-btn>
        <span v-if="memory !== '-'" class="load-action__result">{{ memory }}</span>
      </div>
    </div>
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
.load-tools-heading {
  margin-bottom: 0.75rem;
  color: rgba(255, 255, 255, 0.62);
  font-size: 0.75rem;
  font-weight: 700;
  letter-spacing: 0.06em;
  text-transform: uppercase;
}

.load-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
}

.load-action {
  display: flex;
  align-items: center;
  gap: 0.6rem;
}

.load-action__result {
  color: rgba(255, 255, 255, 0.7);
  font-size: 0.75rem;
}
</style>
