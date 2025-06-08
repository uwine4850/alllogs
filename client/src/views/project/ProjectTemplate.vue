<script setup lang="ts">
import BaseTemplate from '@/views/BaseTemplate.vue'
import Panel from '@/components/Panel.vue'

const props = defineProps({
  title: {
    type: String,
    required: true,
  },
})
</script>

<template>
  <BaseTemplate :title="props.title">
    <div class="content">
      <Panel class="panel-project">
        <slot name="panel-project"></slot>
      </Panel>
      <Panel class="panel-menu">
        <slot name="panel-menu"></slot>
      </Panel>
    </div>
  </BaseTemplate>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;
@use '@/assets/style/presets.scss' as ps;

.content {
  box-sizing: border-box;
  padding: 20px;
  display: flex;
  gap: 20px;
  height: calc(100% - 58px);
  .panel-project {
    flex: 1;
  }
  .panel-menu {
    width: 400px;
  }
}
</style>
