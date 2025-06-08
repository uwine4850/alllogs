<script setup lang="ts">
import ProjectForm from '@/components/project/ProjectForm.vue';
import { type ProjectMessage } from '@/dto/project';
import { ref } from 'vue';
import { useRoute } from 'vue-router';

var route = useRoute();

const formData = ref<ProjectMessage>({
  Id: 0,
  UserId: 0,
  Name: '',
  Description: '',
  Error: '',
  Author: undefined,
})

const submitForm = () => {
  // formData.value.UserId = parseInt(route.params.id.toString());
}
</script>

<template>
  <ProjectForm 
  v-model:model-value-name="formData.Name"
  v-model:model-value-description="formData.Description"
  :onSubmit="submitForm"
  />
</template>