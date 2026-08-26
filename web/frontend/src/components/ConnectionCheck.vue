<template>
  <v-container fluid>
    <v-list-item lines="three">
      <div class="text-h6 mb-4">{{ title }}</div>
      <v-list-item-subtitle>{{ subtitle }}</v-list-item-subtitle>
      <div class="text-body-2">
        <v-text-field
          label="Host and port"
          hint="e.g middleware.se:443 or 10.5.12.12:22"
          v-model="targetValue"
        ></v-text-field>
        <v-row no-gutters>
          <v-col>
            <v-btn color="blue-darken-1" @click="check()" variant="text">Check</v-btn>
          </v-col>
          <v-col>
            <v-chip color="grey" size="small" v-if="inProgress">Connecting ...</v-chip>
            <v-chip color="green-darken-4" size="small" v-if="res.success && !inProgress && hasChecked">
              Connection to {{ res.address }} successful
            </v-chip>
            <v-chip color="red-darken-4" size="small" v-if="!res.success && !inProgress && hasChecked">
              {{ res.error }}
            </v-chip>
          </v-col>
        </v-row>
      </div>
    </v-list-item>
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
