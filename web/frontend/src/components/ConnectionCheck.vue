<template>
  <v-container fluid>
    <div class="tool-heading">
      <v-avatar color="blue-lighten-2" rounded="lg" size="38" variant="tonal">
        <v-icon icon="mdi-connection" size="22"></v-icon>
      </v-avatar>
      <div>
        <div class="tool-heading__title">{{ title }}</div>
        <div class="tool-heading__subtitle">{{ subtitle }}</div>
      </div>
    </div>
      <div class="connection-form">
        <v-text-field
          density="comfortable"
          label="Host and port"
          hint="e.g middleware.se:443 or 10.5.12.12:22"
          prepend-inner-icon="mdi-server-network"
          variant="outlined"
          v-model="targetValue"
        ></v-text-field>
        <v-row no-gutters>
          <v-col>
            <v-btn color="blue-lighten-2" prepend-icon="mdi-lan-connect" @click="check()" variant="outlined">Check Connection</v-btn>
          </v-col>
          <v-col>
            <v-chip color="grey" size="small" v-if="inProgress">Connecting ...</v-chip>
            <v-chip color="green-accent-3" size="small" variant="outlined" v-if="res.success && !inProgress && hasChecked">
              Connection to {{ res.address }} successful
            </v-chip>
            <v-chip color="red-accent-2" size="small" variant="outlined" v-if="!res.success && !inProgress && hasChecked">
              {{ res.error }}
            </v-chip>
          </v-col>
        </v-row>
      </div>
  </v-container>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ConnectionCheck',
  props: {
    title: { type: String, default: 'Connection Checker' },
    subtitle: {
      type: String,
      default: 'Test TCP/IP connectivity to other systems from the container'
    },
    target: { type: String, default: '' }
  },
  methods: {
    check() {
      this.inProgress = true
      axios.get('/kubez/connectioncheck/' + this.targetValue).then(res => {
        this.res = res.data
        this.inProgress = false
        this.hasChecked = true
      })
    }
  },
  data() {
    return {
      inProgress: false,
      hasChecked: false,
      res: {},
      targetValue: this.target
    }
  }
}
</script>

<style scoped>
.tool-heading {
  display: flex;
  align-items: center;
  gap: 0.7rem;
  margin-bottom: 1.25rem;
}

.tool-heading__title {
  font-size: 1rem;
  font-weight: 650;
}

.tool-heading__subtitle {
  color: rgba(255, 255, 255, 0.62);
  font-size: 0.75rem;
}

.connection-form {
  padding: 1rem;
  border: 1px solid rgba(144, 202, 249, 0.12);
  border-radius: 10px;
  background-color: #2e353b;
}
</style>
