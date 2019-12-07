<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer" app clipped>
      <v-list dense>
       
          <v-list-item link @click="$router.push('/')">
            <v-list-item-action>
              <v-icon>mdi-view-dashboard</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Pod info</v-list-item-title>
            </v-list-item-content>
          </v-list-item>


                  <v-list-item link @click="$router.push('api')">
            <v-list-item-action>
              <v-icon>mdi-drone</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>API Control Center </v-list-item-title>
            </v-list-item-content>
          </v-list-item>

          <v-list-item link @click="$router.push('compute-stats')">
            <v-list-item-action>
              <v-icon>mdi-chart-line</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Stats &amp; Load Tools</v-list-item-title>
            </v-list-item-content>
          </v-list-item>

          <v-list-item link @click="$router.push('observability')">
            <v-list-item-action>
              <v-icon>mdi-popcorn</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>Observability</v-list-item-title>
            </v-list-item-content>
          </v-list-item>


        

   

        <v-list-item link @click="$router.push('network')">
          <v-list-item-action>
            <v-icon>mdi-web</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Network</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-list-item link @click="$router.push('authentication')">
          <v-list-item-action>
            <v-icon>mdi-key</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>User Authentication</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app clipped-left>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title>Dr. Kubez</v-toolbar-title>
      <v-spacer></v-spacer>
      <span mx-3>
        <v-btn color="secondary" small dark>Kill</v-btn>
        <v-btn color="secondary" small dark>Allocate Ram</v-btn>
        <v-btn color="secondary" small dark>Load CPU</v-btn>
        <v-btn color="secondary" small dark>Fail Healthcheck</v-btn>
      </span>
    </v-app-bar>

    <v-content>
      <router-view></router-view>
    </v-content>

    <v-footer app pl-3>
      <span>Svenska Middlewaregruppen AB</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  props: {
    source: String
  },

  data: () => ({
    drawer: null
  }),

  created() {
    this.$vuetify.theme.dark = true;
  }
};
</script>