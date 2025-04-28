<script lang="ts">
import filterIcon from '@/assets/svg/filter.svg'
import checkboxIcon from '@/assets/svg/checkbox.svg'
</script>

<script setup lang="ts">
import AlertPanelTemplate from '../alertpanel/AlertPanelTemplate.vue'
import PanelTitle from '../PanelTitle.vue'
import InputDropdown from '../input/InputDropdown.vue'
import InputText from '../input/InputText.vue'
import InputDateTime from '../input/InputDateTime.vue'
import Button from '../Button.vue'
import Separator from '../Separator.vue'

type Option = {
  value: string
  name: string
}

const options: Option[] = [
  { value: 'warn', name: 'warn' },
  { value: 'error', name: 'error' },
  { value: 'info', name: 'info' },
]
</script>

<template>
  <AlertPanelTemplate class="a-panel">
    <PanelTitle :icon="filterIcon" text="Filter" :sep="false" />
    <InputDropdown text="type" inp-name="type" :options="options" />
    <InputText text="tag" name="tag" />
    <InputDateTime text="Date Time" name="dateTime" />
    <Separator />
    <Button
      id="save-btn"
      class="save-btn"
      :icon="checkboxIcon"
      text="Save"
    />
  </AlertPanelTemplate>
</template>

<style scoped lang="scss">
.a-panel {
  background: #000;
  :deep(.alert-inner-panel) {
    width: 700px !important;
  }
  .save-btn {
    margin: 10px 0;
    width: 200px;
    margin-left: auto;
    margin-right: 10px;
  }
}
</style>
