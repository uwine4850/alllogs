<script setup lang="ts">
import type { ProjectMessage } from '@/dto/project'
import { ref } from 'vue'
import { AsyncRequestWithAuthorization } from '@/classes/request'
import { useErrorStore } from '@/stores/error'
import { useRouter } from 'vue-router'
import type { AxiosError, AxiosResponse } from 'axios'
import type { BaseResponseMessage } from '@/dto/common'
import ProjectForm from '@/components/project/ProjectForm.vue'

const errorStore = useErrorStore()
const router = useRouter()

const formData = ref<ProjectMessage>({
  Id: 0,
  UserId: 0,
  Name: '',
  Description: '',
  Error: '',
  Author: undefined,
})

const submitForm = () => {
  const req = new AsyncRequestWithAuthorization('http://localhost:8000/new-project', {
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    withCredentials: true,
  })
  req.onResponse(async (response: AxiosResponse) => {
    const baseResponse = response.data as BaseResponseMessage
    if (!baseResponse.Ok) {
      errorStore.setText(baseResponse.Error)
    } else {
      router.push('/')
    }
  })
  req.onError((error: AxiosError) => {
    errorStore.setText("unexpected error: " + error.message)
  }, errorStore)
  req.setData(formData.value)
  req.post()
}
</script>

<template>
  <ProjectForm 
  v-model:model-value-name="formData.Name"
  v-model:model-value-description="formData.Description"
  :onSubmit="submitForm"
  />
</template>