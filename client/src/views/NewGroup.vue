<script lang="ts">
import groupIcon from '@/assets/svg/group.svg'
import checkBoxIcon from '@/assets/svg/checkbox.svg'
</script>

<script setup lang="ts">
import BaseTemplate from './BaseTemplate.vue'
import MiddlePanel from './MiddlePanel.vue'
import PanelTitle from '@/components/PanelTitle.vue'
import InputText from '@/components/input/InputText.vue'
import InputTextarea from '@/components/input/InputTextarea.vue'
import Button from '@/components/Button.vue'
import Separator from '@/components/Separator.vue'
</script>

<template>
  <BaseTemplate title="New group">
    <MiddlePanel>
      <PanelTitle icon="group" text="new group" :sep="false" />
      <InputText text="Name" name="name" />
      <InputTextarea text="Description" name="description" />
      <Separator />
      <Button class="create-btn" icon="checkbox" text="Create" />
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
