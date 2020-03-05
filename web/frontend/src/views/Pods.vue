<template>
  <v-container fluid>
  <v-data-table
    :headers="headers"
    :items="desserts"
    item-key="name"
    dense
    :search="search"
    single-line
    class="elevation-1"
  >
    <template v-slot:top>

      <v-toolbar flat>
        <v-toolbar-title>Pods</v-toolbar-title>
        <v-spacer></v-spacer>
              <v-text-field
        v-model="search"
        append-icon="mdi-magnify"
        label="Search"
        single-line
        hide-details
      ></v-text-field>
        <!--v-switch v-model="singleExpand" label="Single expand" class="mt-2"></v-switch-->
      </v-toolbar>
    </template>
    <template v-slot:expanded-item="{ headers }">
      <td :colspan="headers.length">Peek-a-boo! woohooo {{calories}} </td>
    </template>
  </v-data-table>
  </v-container>
</template>

<script>
  export default {
    data () {
      return {
        search:'',

        headers: [
          {
            text: 'Name',
            align: 'start',
            sortable: true,
            value: 'name',
          },
          { text: 'Pod phase', value: 'carbs' },
          { text: 'Container Statuses', value: 'protein' },
          { text: 'Namespace', value: 'calories' },
          { text: 'Node', value: 'carbs' },
          { text: 'Conditions', value: 'protein' },
          { text: 'Quality  Of Service', value: 'fat' },
          { text: 'Restarts', value: 'protein' },
          { text: 'Reason', value: 'protein' },
          { text: 'CPU Throttled', value: 'protein' },
          { text: 'Memory Use', value: 'protein' },
        ],
        desserts: [
          {
            name: 'kubez-app-799b5d74bf-4zm9w',
            calories: "Kubez",
            fat: "Guranteed",
            carbs: "Running",
            protein: 4.0,
            iron: '2020-03-04T20:21:52Z',
          },
          {
            name: 'Ice cream sandwich',
            calories: 237,
            fat: 9.0,
            carbs: 37,
            protein: 4.3,
            iron: '1%',
          },
          {
            name: 'Eclair',
            calories: 262,
            fat: 16.0,
            carbs: 23,
            protein: 6.0,
            iron: '7%',
          },
          {
            name: 'Cupcake',
            calories: 305,
            fat: 3.7,
            carbs: 67,
            protein: 4.3,
            iron: '8%',
          },
          {
            name: 'Gingerbread',
            calories: 356,
            fat: 16.0,
            carbs: 49,
            protein: 3.9,
            iron: '16%',
          },
          {
            name: 'Jelly bean',
            calories: 375,
            fat: 0.0,
            carbs: 94,
            protein: 0.0,
            iron: '0%',
          },
          {
            name: 'Lollipop',
            calories: 392,
            fat: 0.2,
            carbs: 98,
            protein: 0,
            iron: '2%',
          },
          {
            name: 'Honeycomb',
            calories: 408,
            fat: 3.2,
            carbs: 87,
            protein: 6.5,
            iron: '45%',
          },
          {
            name: 'Donut',
            calories: 452,
            fat: 25.0,
            carbs: 51,
            protein: 4.9,
            iron: '22%',
          },
          {
            name: 'KitKat',
            calories: 518,
            fat: 26.0,
            carbs: 65,
            protein: 7,
            iron: '6%',
          },
        ],
      }
    },
  }
</script>
<!--  script >
// @ is an alias to /src
//import axios from "axios";
/*import { mapState } from "vuex";
export default {
  name: "Api",
  components: {
    ApiEndpoint,
    Instructions,
    ApiNew
  },

  computed: mapState({
    apis: state => state.apiendpoint.apis
  }),
  methods: {
     createEndpoint: function(api) {
      this.$store.commit("NEW_API", api);
      this.updateList()
    },
    updateEndpoint: function(api) {
      this.$store.commit("UPDATE_API", api);
    },
    updateList: function() {
      this.$store.dispatch("fetchAPIEndpoints");
    }
  },
  mounted: function() {
    this.updateList()
  },
  data: function() {
    return {
      endpoints: {},
      dialog: false
    };
  }
};
</script -->
