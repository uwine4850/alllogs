<script setup lang="ts">
import AlertPanelTemplate, { closeAlertPanel } from '@/components/alertpanel/AlertPanelTemplate.vue'
import PanelTitle from '@/components/PanelTitle.vue'
import InputDropdown from '@/components/input/InputDropdown.vue'
import InputText from '@/components/input/InputText.vue'
import InputDateTime from '@/components/input/InputDateTime.vue'
import Button from '@/components/Button.vue'
import Separator from '@/components/Separator.vue'
import { computed } from 'vue'

const props = defineProps<{
  onSearch: () => void
  modelValueType: string
  modelValueTag: string
  modelValueDateTime: string
}>()

const emit = defineEmits<{
  (e: 'update:modelValueType', value: string): void
  (e: 'update:modelValueTag', value: string): void
  (e: 'update:modelValueDateTime', value: string): void
}>()

const formType = computed({
  get: () => props.modelValueType,
  set: value => emit('update:modelValueType', value),
})
const formTag = computed({
  get: () => props.modelValueTag,
  set: value => emit('update:modelValueTag', value),
})
const formDateTime = computed({
  get: () => props.modelValueDateTime,
  set: value => emit('update:modelValueDateTime', value),
})

type Option = {
  value: string
  name: string
}

const options: Option[] = [
  { value: 'warn', name: 'warn' },
  { value: 'error', name: 'error' },
  { value: 'info', name: 'info' },
]
function search(){
  closeAlertPanel()
  props.onSearch()
}
</script>

<template>
  <AlertPanelTemplate class="a-panel">
    <PanelTitle icon="filter" text="Filter" :sep="false" />
    <InputDropdown text="type" inp-name="type" :options="options" v-model="formType" />
    <InputText text="tag" name="tag" v-model="formTag" />
    <InputDateTime text="Date Time" name="dateTime" v-model="formDateTime" />
    <Separator />
    <Button id="save-btn" class="save-btn" icon="checkbox" text="Search" @click="search" />
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
