<template>
  <div class="home">
    <v-row>
         <v-col>
        <v-card>
          <LoadTools />
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-card>
          <CGroup />
        </v-card>
      </v-col>
      <v-col>
        <v-card>
          <LoadTools />
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
// @ is an alias to /src
import CGroup from "@/components/CGroup.vue";
import LoadTools from "@/components/LoadTools.vue";

export default {
  name: "ComputeStats",
  components: {
    CGroup,
    LoadTools
  }
};
</script>
