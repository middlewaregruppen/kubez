<template>
  <v-dialog v-model="dialog" scrollable max-width="800px">
    <template v-slot:activator="{ props }">
      <v-btn size="small" icon="mdi-help-box" v-bind="props"></v-btn>
    </template>
    <v-card>
      <v-card-text style="height: 300px;">
        <!-- Rendered from public markdown file fetched via axios -->
        <div class="mt-4 markdown-body" v-html="renderedContent"></div>
      </v-card-text>
      <v-divider></v-divider>
      <v-card-actions>
        <v-btn size="small" color="blue-darken-1" variant="text" @click="dialog = false">Close</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { marked } from 'marked'
import axios from 'axios'

export default {
  name: 'Instructions',
  props: ['href'],
  data: () => ({
    dialog: false,
    content: ''
  }),
  computed: {
    renderedContent() {
      return this.content ? marked(this.content) : ''
    }
  },
  mounted() {
    axios.get(this.href).then(res => {
      this.content = res.data
    })
  }
}
</script>
