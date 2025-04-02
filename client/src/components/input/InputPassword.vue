<script setup lang="ts">
import { onMounted, ref } from 'vue'
import InputTemplate from '@/components/input/InputTemplate.vue'
import eyeIcon from '@/assets/svg/eye.svg'

const props = defineProps({
  text: {
    type: String,
    required: true,
  },
  name: {
    type: String,
    required: true,
  },
  readonly: {
    type: Boolean,
  },
  value: {
    type: String,
  },
  modelValue: {
    type: String,
  },
})

const emit = defineEmits(['update:modelValue'])

const passwordInput = ref<HTMLInputElement | null>(null)
const togglePasswordButton = ref<HTMLButtonElement | null>(null)

onMounted(() => {
  if (props.readonly) {
    const passwordInput = document.getElementById('password') as HTMLInputElement
    if (passwordInput && props.value) {
      passwordInput.value = props.value
    }
  }
  if(togglePasswordButton.value && passwordInput.value){
    toggle(togglePasswordButton.value, passwordInput.value);
  }
})

const updateValue = (event: Event) => {
  const target = event.target as HTMLInputElement
  emit('update:modelValue', target.value)
}

</script>

<template>
  <InputTemplate :text="props.text">
    <input ref="passwordInput" :readonly="props.readonly" type="password" class="p-input" :value="modelValue" @input="updateValue" />
    <button ref="togglePasswordButton" type="button" class="p-toggle">
      <img :src="eyeIcon" />
    </button>
  </InputTemplate>
</template>

<script lang="ts">
function toggle(togglePasswordButton: HTMLButtonElement, passwordInput: HTMLInputElement){
  if (passwordInput && togglePasswordButton) {
    togglePasswordButton.addEventListener('click', () => {
      passwordInput.type = passwordInput.type === 'password' ? 'text' : 'password'
    })
  }
}
</script>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;
@use '@/assets/style/presets.scss' as ps;

.p-input {
  width: 100%;
  font-size: 1.1rem;
  font-family: vars.$fnt-hint-madurai;
  background-color: vars.$input-color;
  height: 45px;
  padding: 0 10px;
  border: none;
  outline: none;
  @include ps.inner-shadow-panel;
}
.p-toggle {
  width: 45px;
  height: 45px;
  display: flex;
  border: none;
  outline: none;
  background-color: vars.$inner-button;
  @include ps.inner-shadow-panel;
  &:hover {
    transition: 0.2s;
    cursor: pointer;
    filter: brightness(90%);
  }
  img {
    background-color: transparent;
    margin: auto auto;
  }
}
</style>
