<template>
  <v-container fluid>
    <div class="home">

      <v-row>
        <v-col>
          <v-card>
            <CGroup />
            <LoadTools />
          </v-card>
        </v-col>
        <v-col>
          <v-card>
            <K8SLoad />
          </v-card>
        </v-col>
      </v-row>
    </div>
  </v-container>
</template>

<script>
// @ is an alias to /src
import CGroup from "@/components/CGroup.vue";
import LoadTools from "@/components/LoadTools.vue";
import K8SLoad from "@/components/K8SLoad.vue";

export default {
  name: "ComputeStats",
  components: {
    CGroup,
    LoadTools,
    K8SLoad
  }
};
</script>
