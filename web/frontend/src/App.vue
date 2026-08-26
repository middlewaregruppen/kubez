<template>
  <v-app id="inspire">
    <v-snackbar v-model="snack.show" :color="snack.color" :timeout="-1">
      {{ snack.message }}
      <template v-slot:actions>
        <v-btn variant="text" icon="mdi-close" @click="hideSnack"></v-btn>
      </template>
    </v-snackbar>

    <v-navigation-drawer v-model="drawer" class="app-drawer" width="280">
      <div class="drawer-header">
        <v-avatar color="blue-lighten-2" rounded="lg" size="44" variant="tonal">
          <v-icon icon="mdi-kubernetes" size="26"></v-icon>
        </v-avatar>
        <div>
          <div class="drawer-header__title">Dr. Kubez</div>
          <div class="drawer-header__subtitle">Cluster diagnostics</div>
        </div>
      </div>

      <v-divider></v-divider>

      <div class="drawer-section-label">Tools</div>
      <v-list class="navigation-list" density="comfortable" nav>
        <v-list-item
          class="navigation-item"
          color="blue-lighten-2"
          rounded="lg"
          title="API Control Center"
          to="/api"
        >
          <template v-slot:prepend>
            <span class="navigation-item__icon">
              <v-icon icon="mdi-api" size="21"></v-icon>
            </span>
          </template>
        </v-list-item>
        <v-list-item
          class="navigation-item"
          color="blue-lighten-2"
          rounded="lg"
          title="Stats &amp; Load Tools"
          to="/compute-stats"
        >
          <template v-slot:prepend>
            <span class="navigation-item__icon">
              <v-icon icon="mdi-chart-timeline-variant" size="21"></v-icon>
            </span>
          </template>
        </v-list-item>
        <v-list-item
          class="navigation-item"
          color="blue-lighten-2"
          rounded="lg"
          title="Pods and Containers"
          to="/pods"
        >
          <template v-slot:prepend>
            <span class="navigation-item__icon">
              <v-icon icon="mdi-cube-outline" size="21"></v-icon>
            </span>
          </template>
        </v-list-item>
        <v-list-item
          class="navigation-item"
          color="blue-lighten-2"
          rounded="lg"
          title="Network"
          to="/network"
        >
          <template v-slot:prepend>
            <span class="navigation-item__icon">
              <v-icon icon="mdi-lan" size="21"></v-icon>
            </span>
          </template>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar class="app-header" height="84">
      <v-app-bar-nav-icon
        :aria-label="drawer ? 'Hide navigation' : 'Show navigation'"
        :title="drawer ? 'Hide navigation' : 'Show navigation'"
        @click.stop="drawer = !drawer"
      />
      <div v-if="!drawer" class="app-bar-brand">
        <v-avatar color="blue-lighten-2" rounded="lg" size="44" variant="tonal">
          <v-icon icon="mdi-kubernetes" size="26"></v-icon>
        </v-avatar>
        <div>
          <div class="drawer-header__title">Dr. Kubez</div>
          <div class="drawer-header__subtitle">Cluster diagnostics</div>
        </div>
      </div>
      <v-spacer></v-spacer>
      <div class="cluster-context">
        <v-chip
          prepend-icon="mdi-server"
          class="context-chip"
          color="blue-grey-lighten-2"
          size="small"
          variant="tonal"
        >
          <span class="context-chip__label">Host</span>
          {{ hostname || 'unknown' }}
        </v-chip>
        <v-chip
          prepend-icon="mdi-kubernetes"
          class="context-chip"
          color="blue-lighten-2"
          size="small"
          variant="tonal"
        >
          <span class="context-chip__label">Namespace</span>
          {{ namespace || 'unknown' }}
        </v-chip>
      </div>
      <v-spacer></v-spacer>
      <v-chip
        v-if="connectionStatus.label"
        :color="connectionStatus.colour"
        :title="connectionStatus.detail"
        class="connection-status mr-4"
        size="small"
        variant="outlined"
      >
        <span class="connection-status__dot" aria-hidden="true"></span>
        {{ connectionStatus.label }}
      </v-chip>
    </v-app-bar>

    <v-main>
      <router-view></router-view>
    </v-main>

    <v-footer app class="app-footer px-4">
      <div class="footer-brand">
        <span class="footer-brand__mark"></span>
        <span>Svenska Middlewaregruppen AB</span>
      </div>
      <v-spacer></v-spacer>
      <nav class="footer-links" aria-label="External links">
        <a
          class="footer-link"
          href="https://middlewaregruppen.se/"
          rel="noopener noreferrer"
          target="_blank"
        >
          <v-icon icon="mdi-web" size="17"></v-icon>
          middlewaregruppen.se
          <v-icon class="footer-link__external" icon="mdi-open-in-new" size="12"></v-icon>
        </a>
        <span class="footer-divider" aria-hidden="true"></span>
        <a
          class="footer-link"
          href="https://github.com/middlewaregruppen/kubez"
          rel="noopener noreferrer"
          target="_blank"
        >
          <v-icon icon="mdi-github" size="18"></v-icon>
          kubez
          <v-icon class="footer-link__external" icon="mdi-open-in-new" size="12"></v-icon>
        </a>
      </nav>
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
      const c = { colour: '', label: '', detail: '' }
      switch (this.infoStore.status) {
        case 200:
          c.colour = 'green-accent-3'
          c.label = 'connected'
          break
        case -1:
          return c
        default:
          c.colour = 'red-accent-2'
          c.label = 'disconnected'
          c.detail = String(this.infoStore.status)
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

<style scoped>
.app-drawer {
  border-right-color: rgba(144, 202, 249, 0.16) !important;
}

.app-header {
  border-bottom: 1px solid rgba(144, 202, 249, 0.1) !important;
}

.drawer-header {
  display: flex;
  align-items: center;
  gap: 0.8rem;
  min-height: 84px;
  padding: 1rem 1.1rem;
}

.app-bar-brand {
  display: flex;
  align-items: center;
  gap: 0.8rem;
  margin-left: 0.35rem;
}

.drawer-header__title {
  font-size: 1rem;
  font-weight: 700;
  line-height: 1.3;
  letter-spacing: -0.01em;
}

.drawer-header__subtitle {
  margin-top: 0.1rem;
  color: rgba(255, 255, 255, 0.58);
  font-size: 0.75rem;
}

.drawer-section-label {
  padding: 1.1rem 1.25rem 0.45rem;
  color: #90caf9;
  font-size: 0.7rem;
  font-weight: 700;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.navigation-list {
  padding-inline: 0.65rem;
}

.navigation-item {
  min-height: 48px;
  margin-bottom: 0.3rem;
  font-size: 0.875rem;
  font-weight: 600;
  letter-spacing: 0.01em;
}

:deep(.navigation-item .v-list-item__prepend) {
  margin-inline-end: 0.8rem;
}

.navigation-item__icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 34px;
  height: 34px;
  border-radius: 9px;
  color: #90caf9;
  background-color: #2e353b;
}

:deep(.navigation-item.v-list-item--active) {
  background-color: #2e353b;
  box-shadow: inset 3px 0 0 #90caf9;
}

:deep(.navigation-item.v-list-item--active .navigation-item__icon) {
  color: #fff;
  background-color: #35617f;
}

.cluster-context {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.context-chip {
  font-size: 0.875rem;
  font-weight: 600;
  letter-spacing: 0.01em;
}

.context-chip__label {
  margin-right: 0.35rem;
  font-weight: 700;
}

.connection-status {
  font-size: 0.875rem;
  font-weight: 600;
  letter-spacing: 0.01em;
}

.connection-status__dot {
  width: 0.65rem;
  height: 0.65rem;
  margin-right: 0.45rem;
  border-radius: 50%;
  background-color: currentColor;
  box-shadow: 0 0 0.35rem currentColor;
}

.app-footer {
  min-height: 44px;
  border-top: 1px solid rgba(144, 202, 249, 0.18) !important;
  background-color: #22282d !important;
  color: rgba(255, 255, 255, 0.68);
  font-size: 0.75rem;
}

.footer-brand,
.footer-links,
.footer-link {
  display: flex;
  align-items: center;
}

.footer-brand {
  gap: 0.55rem;
  font-weight: 600;
}

.footer-brand__mark {
  width: 0.5rem;
  height: 0.5rem;
  border-radius: 50%;
  background-color: #90caf9;
  box-shadow: 0 0 0.3rem rgba(144, 202, 249, 0.65);
}

.footer-links {
  gap: 0.8rem;
}

.footer-link {
  gap: 0.35rem;
  color: #90caf9;
  font-weight: 600;
  text-decoration: none;
  transition: color 150ms ease;
}

.footer-link:hover {
  color: #fff;
}

.footer-link__external {
  opacity: 0.6;
}

.footer-divider {
  width: 1px;
  height: 18px;
  background-color: rgba(144, 202, 249, 0.22);
}

@media (max-width: 700px) {
  .cluster-context {
    display: none;
  }

  .footer-brand span:last-child {
    display: none;
  }
}
</style>
