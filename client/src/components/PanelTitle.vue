<script setup lang="ts">
import Separator from '@/components/Separator.vue'
import SvgIcon from './icons/SvgIcon.vue';

const props = defineProps({
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
    <SvgIcon :name="props.icon" class="icon" />
    <p>{{ props.text }}</p>
  </div>
  <Separator v-if="props.sep" />
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;

.content {
  background-color: transparent;
  display: flex;
  padding: 10px;
  display: flex;
  .icon {
    width: 1.25rem;
    height: 1.25rem;
    margin: auto 0;
  }
  p {
    background-color: transparent;
    font-size: 1.25rem;
    margin-left: 5px;
    color: vars.$secondary-font-color;
  }
}
</style>
