<template>
  <div class="page-header">
    <div class="page-header__identity">
      <v-avatar color="blue-lighten-2" rounded="lg" size="44" variant="tonal">
        <v-icon :icon="icon" size="26"></v-icon>
      </v-avatar>
      <div>
        <h1 class="page-header__title">{{ title }}</h1>
        <p class="page-header__subtitle">{{ subtitle }}</p>
      </div>
    </div>
    <div v-if="$slots.actions" class="page-header__actions">
      <slot name="actions"></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PageHeader',
  props: {
    icon: { type: String, required: true },
    title: { type: String, required: true },
    subtitle: { type: String, required: true }
  }
}
</script>

<style scoped>
.page-header {
  display: flex;
  align-items: center;
  gap: 1rem;
  min-height: 76px;
  padding: 0.9rem 1rem;
  border-bottom: 1px solid #90caf9;
  background-color: #2e353b;
}

.page-header__identity {
  display: flex;
  align-items: center;
  gap: 0.8rem;
  min-width: 0;
}

.page-header__title {
  margin: 0;
  font-size: 1.125rem;
  font-weight: 650;
  line-height: 1.3;
  letter-spacing: -0.01em;
}

.page-header__subtitle {
  margin: 0.1rem 0 0;
  color: rgba(255, 255, 255, 0.62);
  font-size: 0.75rem;
  line-height: 1.3;
}

.page-header__actions {
  display: flex;
  align-items: center;
  gap: 0.4rem;
  margin-left: auto;
}

@media (max-width: 700px) {
  .page-header__subtitle {
    display: none;
  }
}
</style>
