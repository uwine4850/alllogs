<script setup lang="ts">
import { AsyncRequestWithAuthorization } from '@/classes/request';
import ProjectForm from '@/components/project/ProjectForm.vue';
import type { BaseResponseMessage } from '@/dto/common';
import type { ProfileMessage } from '@/dto/profile';
import { type ProjectMessage } from '@/dto/project';
import { getProject } from '@/services/project';
import { useErrorStore } from '@/stores/error';
import type { AxiosError, AxiosResponse } from 'axios';
import { onMounted, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

var route = useRoute();
var router = useRouter();
var errorStore = useErrorStore();

const formData = ref<ProjectMessage>({
  Id: 0,
  UserId: 0,
  Name: '',
  Description: '',
  Error: '',
  Author: undefined,
})

const projectRef = ref<ProjectMessage | null>(null)
onMounted(() => {
  getProject(route.params.id, projectRef, errorStore);
})
watch(projectRef, (project) => {
  if(project){
    formData.value.Id = project.Id;
    formData.value.Name = project.Name;
    formData.value.Description = project.Description;
  }
});


const submitForm = () => {
  const req = new AsyncRequestWithAuthorization('http://localhost:8000/project', {
    withCredentials: true,
  })
  req.onResponse(async (response: AxiosResponse) => {
    const baseResponse = response.data as BaseResponseMessage
    if (!baseResponse.Ok) {
      errorStore.setText(baseResponse.Error)
    } else {
      router.push("/project/" + route.params.id)
    }
  })
  req.onError((error: AxiosError) => {
    errorStore.setText("unexpected error: " + error.message)
  }, errorStore)
  req.setData(formData.value);
  req.patch();
}
</script>

<template>
  <ProjectForm 
  v-model:model-value-name="formData.Name"
  v-model:model-value-description="formData.Description"
  :onSubmit="submitForm"
  />
</template>