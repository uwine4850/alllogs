<script setup lang="ts">
import Separator from '@/components/Separator.vue'

defineProps({
  text: {
    type: String,
    required: true,
  },
})
</script>

<template>
  <Separator />
  <div class="text">
    {{ text }}
  </div>
  <Separator />
  <div class="inp">
    <slot></slot>
  </div>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;
@use '@/assets/style/presets.scss' as ps;

.text {
  font-size: 1.25rem;
  display: flex;
  align-items: center;
  padding-left: 10px;
  height: 45px;
  background-color: vars.$focus-color;
  @include ps.inner-shadow-panel;
}
.inp {
  display: flex;
  background-color: transparent;
  min-height: 45px;
}
</style>
