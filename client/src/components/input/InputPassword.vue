<!-- <script lang="ts">
</script> -->

<script setup lang="ts">
import { onMounted } from 'vue'
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
})

onMounted(() => {
  if (props.readonly) {
    const passwordInput = document.getElementById('password') as HTMLInputElement
    if (passwordInput && props.value) {
      passwordInput.value = props.value
    }
  }
})
</script>

<template>
  <InputTemplate :text="props.text">
    <input v-if="props.readonly" readonly type="password" id="password" class="p-input" />
    <input v-else type="password" :name="props.name" id="password" class="p-input" />
    <button id="toggle-password" class="p-toggle">
      <img :src="eyeIcon" />
    </button>
  </InputTemplate>
</template>

<script lang="ts">
document.addEventListener('DOMContentLoaded', () => {
  const passwordInput = document.getElementById('password') as HTMLInputElement
  const toggleButton = document.getElementById('toggle-password') as HTMLElement

  if (passwordInput && toggleButton) {
    toggleButton.addEventListener('click', () => {
      passwordInput.type = passwordInput.type === 'password' ? 'text' : 'password'
    })
  }
})
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
