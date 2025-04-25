<script setup lang="ts">
import Separator from '@/components/Separator.vue'

defineProps({
  icon: {
    type: String,
    required: true,
  },
  text: {
    type: String,
    required: true,
  },
  sep: {
    type: Boolean,
    default: true,
  },
})
</script>

<template>
  <div class="content">
    <img :src="icon" alt="icon" />
    <p>{{ text }}</p>
  </div>
  <Separator v-if="sep" />
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;

.content {
  background-color: transparent;
  display: flex;
  padding: 10px;
  img {
    background-color: transparent;
  }
  p {
    background-color: transparent;
    font-size: 1.25rem;
    margin-left: 5px;
    color: vars.$secondary-font-color;
  }
}
</style>
