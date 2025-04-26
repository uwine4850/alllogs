<script setup lang="ts">
import InputTemplate from './InputTemplate.vue'

const props = defineProps({
  text: {
    type: String,
    required: true,
  },
})
const emit = defineEmits(['update:modelValue'])

const updateValue = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files.length > 0) {
    const file = target.files[0];
    emit('update:modelValue', file);
  } else {
    emit('update:modelValue', null);
  }
}
</script>

<template>
  <InputTemplate :text="props.text">
    <span>
      <input type="file" class="inp" @change="updateValue" />
    </span>
  </InputTemplate>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;
@use '@/assets/style/presets.scss' as ps;
span {
  background-color: vars.$input-color;
  width: 100%;
  display: flex;
  @include ps.inner-shadow-panel;
  .inp {
    margin: auto 0;
    margin-left: 10px;
    font-family: vars.$fnt-hint-madurai;
    font-size: 1rem;
    background-color: transparent;
  }
}
</style>
