<script setup lang="ts">
import projectIcon from '@/assets/svg/project.svg'
import checkBoxIcon from '@/assets/svg/checkbox.svg'

import BaseTemplate from '@/views/BaseTemplate.vue'
import MiddlePanel from '@/views/MiddlePanel.vue'
import PanelTitle from '@/components/PanelTitle.vue'
import InputText from '@/components/input/InputText.vue'
import InputTextarea from '@/components/input/InputTextarea.vue'
import Button from '@/components/Button.vue'
import Separator from '@/components/Separator.vue'
import Error from '@/components/Error.vue'

import { computed } from 'vue'

const props = defineProps<{
  modelValueName: string
  modelValueDescription: string
  onSubmit: () => void
}>()

const emit = defineEmits<{
  (e: 'update:modelValueName', value: string): void
  (e: 'update:modelValueDescription', value: string): void
}>()

const localName = computed({
  get: () => props.modelValueName,
  set: value => emit('update:modelValueName', value)
})

const localDescription = computed({
  get: () => props.modelValueDescription,
  set: value => emit('update:modelValueDescription', value)
})
</script>

<template>
  <BaseTemplate title="New project">
    <MiddlePanel>
      <Error />
      <PanelTitle :icon="projectIcon" text="new project" :sep="false" />
      
      <InputText
        v-model="localName"
        text="Name"
        name="name"
      />
      
      <InputTextarea
        v-model="localDescription"
        text="Description"
        name="description"
      />

      <Separator />
        <slot id="extra"></slot>
      <Button
        @click="onSubmit"
        type="button"
        class="create-btn"
        :icon="checkBoxIcon"
        text="Create"
      />
    </MiddlePanel>
  </BaseTemplate>
</template>

<style scoped lang="scss">
.create-btn {
  margin: 10px;
  width: 200px;
  margin-left: auto;
}
</style>