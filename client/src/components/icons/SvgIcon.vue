<script setup lang="ts">
import { computed } from 'vue';
    const props = defineProps({
        name:{
            type: String,
            required: true,
        },
    })
    const iconName = computed(() => `#icon-${props.name}`)
</script>

<template>
  <svg
    aria-hidden="true"
    class="svg-icon"
    xmlns="http://www.w3.org/2000/svg"
    height="100%"
    viewBox="0 0 24 24"
    fill="currentColor"
  >
    <use :xlink:href="iconName" />
  </svg>
</template>