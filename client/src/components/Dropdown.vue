<script setup lang="ts">
import { computed, onMounted, reactive, ref, watch } from 'vue'
import { activeState } from '@/states/dropdown'

const props = defineProps({
  targetID: {
    type: String,
    required: true,
  },
})

const isHidden = ref(true)
const uniqueId = ref('')

const isActive = computed(() => {
  return activeState.activeId === uniqueId.value
})

watch(isActive, (newVal) => {
  if (!newVal) {
    isHidden.value = true
  }
})

const activateComponent = () => {
  activeState.activeId = uniqueId.value
  isHidden.value = !isHidden.value
}

onMounted(() => {
  uniqueId.value = crypto.randomUUID()
  const target = document.getElementById(props.targetID)
  if (target) {
    target.onclick = activateComponent
  }
})
</script>

<template>
  <div :class="['dropdown', { hide: isHidden }]">
    {{ isActive }}
    <slot></slot>
  </div>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;
@use '@/assets/style/presets.scss' as ps;

.dropdown {
  z-index: 1000;
  position: absolute;
  right: 0;
  top: 100%;
  box-sizing: border-box;
  padding: 10px;
  font-size: 1rem;
  font-family: vars.$fnt-hint-madurai;
  background-color: vars.$secondary-color;
  border-radius: 3px;
  border: 1px solid vars.$focus-color;
  @include ps.inner-shadow-panel;
}
.hide {
  display: none;
}
</style>
