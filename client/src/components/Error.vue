<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useErrorStore } from '@/stores/error'
import Separator from './Separator.vue'

const errorStore = useErrorStore()
const errorEl = ref<HTMLElement>()

const props = defineProps({
  storeId: {
    type: String,
  },
})

onMounted(() => {
  if (errorEl.value) {
    errorStore.setErrorElement(errorEl.value, props.storeId)
  }
})
</script>

<template>
  <div class="error-wrapper-c" id="wrapper" ref="errorEl">
    <span class="error-text-c"></span>
    <Separator class="sep" />
  </div>
</template>

<style scoped lang="scss">
@use '../assets/style/global_vars.scss' as vars;
@use '../assets/style/presets.scss' as ps;

.error-wrapper-c {
  background-color: vars.$secondary-color;
  width: 100%;
  display: none;
  flex-direction: column;
  .error-text-c {
    font-size: 1.1rem;
    padding: 10px;
    background-color: transparent;
    box-sizing: border-box;
    color: red;
    @include ps.inner-shadow-panel;
  }
}
</style>
