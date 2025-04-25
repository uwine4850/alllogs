<script setup lang="ts">
import AlertPanelTemplate from '@/components/alertpanel/AlertPanelTemplate.vue'

const props = defineProps({
  text: {
    type: String,
    required: true,
  },
  hide: {
    type: Boolean,
  },
})
</script>

<template>
  <AlertPanelTemplate :hide="hide">
    <div class="text">
      {{ props.text }}
    </div>
  </AlertPanelTemplate>
</template>

<style scoped>
.text {
  margin: 10px;
  background-color: transparent;
}
</style>
