<template>
  <v-dialog v-model="dialog" scrollable max-width="760px">
    <template v-slot:activator="{ props }">
      <v-btn
        aria-label="Open help"
        color="blue-grey-lighten-2"
        icon="mdi-help-circle-outline"
        variant="text"
        v-bind="props"
      ></v-btn>
    </template>
    <v-card class="help-dialog" rounded="xl">
      <v-card-title class="help-dialog__header">
        <v-avatar color="blue-lighten-2" rounded="lg" size="38" variant="tonal">
          <v-icon icon="mdi-book-open-page-variant-outline" size="22"></v-icon>
        </v-avatar>
        <div>
          <div class="help-dialog__eyebrow">Documentation</div>
          <div class="help-dialog__title">About this tool</div>
        </div>
        <v-spacer></v-spacer>
        <v-btn aria-label="Close help" icon="mdi-close" variant="text" @click="dialog = false"></v-btn>
      </v-card-title>
      <v-divider></v-divider>
      <v-card-text class="help-dialog__content">
        <div v-if="loading" class="help-dialog__state">
          <v-progress-circular color="blue-lighten-2" indeterminate></v-progress-circular>
          <span>Loading documentation</span>
        </div>
        <div v-else-if="error" class="help-dialog__state">
          <v-icon color="red-accent-2" icon="mdi-alert-circle-outline" size="30"></v-icon>
          <span>{{ error }}</span>
        </div>
        <!-- Content is loaded from trusted static markdown files bundled with the app. -->
        <article v-else class="markdown-body" v-html="renderedContent"></article>
      </v-card-text>
      <v-divider></v-divider>
      <v-card-actions class="help-dialog__actions">
        <v-spacer></v-spacer>
        <v-btn color="blue-lighten-2" variant="outlined" @click="dialog = false">Close</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { marked } from 'marked'
import axios from 'axios'

export default {
  name: 'InstructionsDialog',
  props: ['href'],
  data: () => ({
    dialog: false,
    content: '',
    loading: true,
    error: ''
  }),
  computed: {
    renderedContent() {
      return this.content ? marked(this.content) : ''
    }
  },
  mounted() {
    axios.get(this.href)
      .then(res => {
        this.content = res.data
      })
      .catch(error => {
        this.error = error.response?.data || error.message || 'Unable to load documentation'
      })
      .finally(() => {
        this.loading = false
      })
  }
}
</script>

<style scoped>
.help-dialog {
  overflow: hidden;
  border: 1px solid rgba(144, 202, 249, 0.2);
}

.help-dialog__header {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  min-height: 72px;
  padding: 0.9rem 1rem;
  background-color: #2e353b;
}

.help-dialog__eyebrow {
  color: #90caf9;
  font-size: 0.65rem;
  font-weight: 700;
  letter-spacing: 0.11em;
  text-transform: uppercase;
}

.help-dialog__title {
  font-size: 1rem;
  font-weight: 650;
}

.help-dialog__content {
  min-height: 260px;
  max-height: 62vh;
  padding: 1.5rem 1.75rem !important;
}

.help-dialog__actions {
  padding: 0.75rem 1rem;
}

.help-dialog__state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
  min-height: 220px;
  color: rgba(255, 255, 255, 0.62);
}

.markdown-body {
  color: rgba(255, 255, 255, 0.78);
  font-size: 0.9rem;
  line-height: 1.65;
}

.markdown-body :deep(h1),
.markdown-body :deep(h2),
.markdown-body :deep(h3),
.markdown-body :deep(h4) {
  color: #fff;
  line-height: 1.3;
}

.markdown-body :deep(h1) {
  margin: 0 0 1rem;
  font-size: 1.45rem;
}

.markdown-body :deep(h2),
.markdown-body :deep(h3) {
  margin: 1.5rem 0 0.5rem;
  color: #90caf9;
  font-size: 1rem;
}

.markdown-body :deep(h4) {
  margin: 1.25rem 0 0.35rem;
  font-size: 0.9rem;
}

.markdown-body :deep(p) {
  margin: 0.5rem 0;
}

.markdown-body :deep(ul) {
  margin: 0.65rem 0;
  padding-left: 1.25rem;
}

.markdown-body :deep(li) {
  margin: 0.3rem 0;
}

.markdown-body :deep(code) {
  padding: 0.15rem 0.35rem;
  border-radius: 4px;
  background-color: #2e353b;
  color: #90caf9;
}

.markdown-body :deep(pre) {
  overflow-x: auto;
  margin: 0.9rem 0;
  padding: 1rem;
  border: 1px solid rgba(144, 202, 249, 0.14);
  border-radius: 8px;
  background-color: #1b2024;
}

.markdown-body :deep(pre code) {
  padding: 0;
  background: transparent;
  color: rgba(255, 255, 255, 0.82);
}
</style>
