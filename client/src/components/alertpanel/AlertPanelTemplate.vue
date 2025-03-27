<script setup lang="ts">
import { onMounted } from 'vue';
import { clearHTMLElement } from '@/utils/component';
import Panel from '@/components/Panel.vue';

defineProps({
  hide: {
    type: Boolean,
  },
})

onMounted(() => {
  const alertPanel = document.getElementById('alert-panel');
  if (alertPanel) {
    alertPanel.onclick = function () {
      closeAlertPanel();
    }
  }
})
</script>

<script lang="ts">
  export function closeAlertPanel(){
    clearHTMLElement(document.getElementById('alert-container') as HTMLElement);
  }
</script>

<template>
  <div id="alert-panel" :class="{ hidden: hide }">
    <Panel class="alert-inner-panel" @click.stop>
      <slot></slot>
    </Panel>
  </div>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;

#alert-panel {
  position: absolute;
  z-index: 10;
  width: 100%;
  height: 100%;
  background-color: #0a0a0a84;
  display: flex;
  .alert-inner-panel {
    width: 400px;
    margin: auto auto;
  }
}
.hidden {
  display: none;
}
</style>
