<script setup lang="ts">
import { onMounted } from 'vue'
import Panel from '@/components/Panel.vue'

const props = defineProps({
  width: {
    type: String,
  },
});

onMounted(() => {
  const alertPanel = document.getElementById('alert-panel')
  if (alertPanel) {
    alertPanel.onclick = function () {
      alertPanel.classList.add('hide')
    }
  }
})
</script>

<script lang="ts">
export function openAlertPanel() {
  const alertPanel = document.getElementById('alert-panel')
  if (alertPanel) {
    alertPanel.classList.remove('hide')
  }
}

export function closeAlertPanel() {
  const alertPanel = document.getElementById('alert-panel')
  if (alertPanel) {
    alertPanel.classList.add('hide')
  }
}
</script>

<template>
  <div v-if="props.width">
    <div id="alert-panel" class="hide">
      <Panel :style="{width: props.width}" class="alert-inner-panel" @click.stop>
        <slot></slot>
      </Panel>
    </div>
  </div>
  <div v-else>
    <div id="alert-panel" class="hide">
      <Panel class="alert-inner-panel" @click.stop>
        <slot></slot>
      </Panel>
    </div>
  </div>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;

#alert-panel {
  top: 0;
  left: 0;
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
.hide {
  display: none !important;
}
</style>
