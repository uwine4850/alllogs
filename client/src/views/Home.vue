<script lang="ts">
import projectIcon from '@/assets/svg/project.svg'
import groupIcon from '@/assets/svg/group.svg'
</script>

<script setup lang="ts">
import BaseTemplate from '../views/BaseTemplate.vue'
import Panel from '../components/Panel.vue'
import PanelTitle from '../components/PanelTitle.vue'
import PanelItem from '@/components/PanelItem.vue'
</script>

<template>
  <BaseTemplate title="Home">
    <div class="home-content">
      <Panel class="panel">
        <PanelTitle :icon="projectIcon" text="my projects" />
        <div class="panel-content-wrapper">
          <PanelItem title="Project name" descr="DESCR" link="/project" />
        </div>
      </Panel>
      <Panel class="panel">
        <PanelTitle :icon="groupIcon" text="my groups" />
        <div class="panel-content-wrapper">
          <PanelItem title="Group name" descr="DESCR" link="#" />
        </div>
      </Panel>
    </div>
  </BaseTemplate>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;

.home-content {
  width: 100%;
  height: calc(100vh - 58px);
  box-sizing: border-box;
  padding: 20px;
  display: flex;
  gap: 20px;
  .panel {
    flex: 1;
    display: flex;
    flex-direction: column;
    :deep(.panel) {
      display: flex;
      flex-direction: column;
    }
    .panel-content-wrapper {
      background-color: transparent;
      flex-grow: 1;
      overflow-y: scroll;
    }
  }
}
</style>
