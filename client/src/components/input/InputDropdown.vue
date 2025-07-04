<script setup lang="ts">
import InputTemplate from '@/components/input/InputTemplate.vue'
import { ref, onMounted } from 'vue'

type Option = {
  value: string
  name: string
}
const props = defineProps({
  text: {
    type: String,
    required: true,
  },
  inpName: {
    type: String,
    required: true,
  },
  options: {
    type: Array as () => Option[],
    required: true,
  },
  modelValue: {
    type: String,
  },
})

const selectElement = ref<HTMLElement | null>(null)

const emit = defineEmits(['update:modelValue'])

onMounted(() => {
  if (selectElement.value) {
    dropdown(selectElement.value, emit)
  }
})
</script>

<script lang="ts">
function dropdown(
  selectEl: HTMLElement | null,
  emit: (event: 'update:modelValue', value: string) => void,
) {
  const selectButton = selectEl?.querySelector('#select-button') as HTMLElement
  const options = selectEl?.querySelectorAll('.option') as NodeListOf<HTMLElement>
  const selectInput = selectEl?.querySelector('#custom-select-input') as HTMLInputElement

  if (!selectEl || !selectButton || !options || !selectInput) return

  // Opening/closing the list.
  selectButton.onclick = function () {
    selectEl.classList.toggle('open')
  }

  // Processing if option selection
  options.forEach((option) => {
    option.onclick = function () {
      const selectedText = option.textContent?.trim() || 'empty'
      selectButton.textContent = selectedText
      selectInput.value = option.dataset.value || ''
      emit('update:modelValue', selectInput.value)
      selectEl.classList.remove('open')
    }
  })

  document.addEventListener('click', (event) => {
    if (!selectEl.contains(event.target as Node)) {
      selectEl.classList.remove('open')
    }
  })
}
</script>

<template>
  <InputTemplate :text="props.text">
    <div ref="selectElement" class="custom-select">
      <input id="custom-select-input" type="hidden" :name="props.inpName" />
      <button id="select-button" class="select-btn">select...</button>
      <div class="options">
        <button class="option">select...</button>
        <button
          v-for="option in props.options"
          :key="option.value"
          :data-value="option.value"
          class="option"
        >
          {{ option.name }}
        </button>
      </div>
    </div>
  </InputTemplate>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;
@use '@/assets/style/presets.scss' as ps;

.custom-select {
  position: relative;
  width: 100%;

  .select-btn {
    background: vars.$input-color;
    cursor: pointer;
    height: 45px;
    padding: 0 10px;
    display: flex;
    align-items: center;
    position: relative;
    font-size: 1.25rem;
    width: 100%;
    border: none;
    outline: none;
    @include ps.inner-shadow-panel;
  }

  .options {
    display: none;
    position: absolute;
    width: 100%;
    background: vars.$input-color;
    z-index: 10;
  }

  .option {
    padding: 10px;
    cursor: pointer;
    background: vars.$input-color;
    display: flex;
    align-items: center;
    position: relative;
    border: none;
    outline: none;
    font-family: vars.$fnt-hint-madurai;
    font-size: 1.1rem;
    &:hover {
      background: vars.$focus-color;
    }
  }

  &.open .options {
    display: flex;
    flex-direction: column;
  }
}
</style>
