<script setup lang="ts">
import InputTemplate from '@/components/input/InputTemplate.vue'

const props = defineProps({
  text: {
    type: String,
    required: true,
  },
  name: {
    type: String,
    required: true,
  },
  modelValue: {
    type: String,
  },
})
const emit = defineEmits(['update:modelValue'])
const updateValue = (event: Event) => {
  const target = event.target as HTMLInputElement
  const date = new Date(target.value)
  const formattedDate = formatDate(date)
  emit('update:modelValue', formattedDate)
}

function formatDate(date: Date): string {
  if(isNaN(date.getTime())){
    return ""
  }
  const yyyy = date.getFullYear()
  const mm = String(date.getMonth() + 1).padStart(2, '0')
  const dd = String(date.getDate()).padStart(2, '0')


  return `${yyyy}-${mm}-${dd}`
}
</script>

<template>
  <InputTemplate :text="props.text">
    <input class="text-input" type="date" @input="updateValue" />
  </InputTemplate>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;
@use '@/assets/style/presets.scss' as ps;

.text-input {
  width: 100%;
  border: 0;
  outline: none;
  background-color: vars.$input-color;
  @include ps.inner-shadow-panel;
  font-size: 1.1rem;
  padding: 0 10px;
}
</style>
