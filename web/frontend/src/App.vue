<template>
  <v-app id="inspire">
    <v-snackbar v-model="snack.show" :color="snack.color" :timeout="-1">
      {{ snack.message }}
      <template v-slot:actions>
        <v-btn variant="text" icon="mdi-close" @click="hideSnack"></v-btn>
      </template>
    </v-snackbar>

    <v-navigation-drawer v-model="drawer" app>
      <v-list density="compact">
        <v-list-item
          prepend-icon="mdi-drone"
          title="API Control Center"
          @click="$router.push('/api')"
        />
        <v-list-item
          prepend-icon="mdi-chart-line"
          title="Stats &amp; Load Tools"
          @click="$router.push('/compute-stats')"
        />
        <v-list-item
          prepend-icon="mdi-popcorn"
          title="Pods and Containers"
          @click="$router.push('/pods')"
        />
        <v-list-item
          prepend-icon="mdi-web"
          title="Network"
          @click="$router.push('/network')"
        />
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title>Dr. Kubez</v-toolbar-title>
      <v-spacer></v-spacer>
      {{ hostname }} in {{ namespace }}
      <v-spacer></v-spacer>
      <v-chip :color="connectionStatus.colour" size="x-small">{{ connectionStatus.code }}</v-chip>
    </v-app-bar>

    <v-main>
      <router-view></router-view>
    </v-main>

    <v-footer app class="pl-3">
      <span>Svenska Middlewaregruppen AB</span>
      <v-spacer></v-spacer>
      <span>
        <a href="https://middleware.se">middleware.se</a>
      </span>
      <v-spacer></v-spacer>
      <span>
        <v-icon>mdi-github</v-icon>
        <a href="https://github.com/middlewaregruppen/kubez">kubez</a>
      </span>
    </v-footer>
  </v-app>
</template>

<script>
import { useInfoStore } from './stores/info'

export default {
  setup() {
    return {
      infoStore: useInfoStore()
    }
  },
  methods: {
    hideSnack() {
      this.infoStore.clearSnack()
    }
  },
  computed: {
    hostname() {
      return this.infoStore.hostname
    },
    namespace() {
      return this.infoStore.k8sstats?.namespace
    },
    snack() {
      return this.infoStore.snack
    },
    connectionStatus() {
      const c = { colour: '', code: '' }
      switch (this.infoStore.status) {
        case 200:
          c.colour = 'green-darken-4'
          c.code = 'connected'
          break
        case -1:
          c.code = ''
          return c
        default:
          c.colour = 'red'
          c.code = String(this.infoStore.status)
          break
      }
      return c
    }
  },
  data: () => ({
    drawer: null
  })
}
</script>
