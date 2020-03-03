<template>
  <v-container fluid>
    <v-row dense>
      <v-col>
        <v-card height="100%" flat>
          <v-card-text>
            <CGroup />
          </v-card-text>
          <v-card-actions>
            <LoadTools />
          </v-card-actions>
        </v-card>
      </v-col>
      <v-col>
        <v-card height="100%" flat>
          <K8SLoad />
        </v-card>
      </v-col>
    </v-row>
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
