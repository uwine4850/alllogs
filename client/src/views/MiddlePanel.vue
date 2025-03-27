<script setup lang="ts">
import Panel from '../components/Panel.vue'

const props = defineProps({
  single: {
    type: Boolean,
  }
});
</script>

<template>
  <div :class="['wrapper', {'singe-wrapper': single}]">
    <Panel class="middle-panel">
      <slot></slot>
    </Panel>
  </div>
</template>

<style scoped>
.wrapper {
  display: flex;
  width: 100%;
  height: calc(100vh - 58px);
  background-color: transparent;
  box-sizing: border-box;
  padding: 20px;
}
.single-wrapper{
  height: 100vh;
}

.middle-panel {
  width: 600px;
  margin: auto auto;
  max-height: 100%;
}
</style>
