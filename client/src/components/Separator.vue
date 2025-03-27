<script setup lang="ts">
const props = defineProps({
  vertical: {
    type: Boolean,
  },
})
</script>

<template>
  <hr v-if="props.vertical" class="vl" />
  <div v-else class="hl"></div>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;

.hl {
  height: 4px;
  min-height: 4px;
  border: none;
  background-color: vars.$focus-color;
  -webkit-box-shadow: 0px 0px 4px 0px vars.$bg-color;
  -moz-box-shadow: 0px 0px 4px 0px vars.$bg-color;
  box-shadow: 0px 0px 4px 0px vars.$bg-color;
}
.vl {
  height: 100%;
  width: 4px;
  min-width: 4px;
  border: none;
  background-color: vars.$focus-color;
  -webkit-box-shadow: 0px 0px 4px 0px vars.$bg-color;
  -moz-box-shadow: 0px 0px 4px 0px vars.$bg-color;
  box-shadow: 0px 0px 4px 0px vars.$bg-color;
}
</style>
