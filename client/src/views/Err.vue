<script setup lang="ts">
import { useRoute } from 'vue-router'

var route = useRoute()
</script>

<template>
  <h1>{{ route.query.code }} - {{ route.query.text }}</h1>
</template>

<style scoped>
h1 {
  color: black;
}
</style>
