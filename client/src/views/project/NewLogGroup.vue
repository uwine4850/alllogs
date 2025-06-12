<script setup lang="ts">
import type { ProjectLogGroupMessage, ProjectMessage } from '@/dto/project'
import { ref } from 'vue'
import { AsyncRequestWithAuthorization } from '@/classes/request'
import { useErrorStore } from '@/stores/error'
import { useRoute, useRouter } from 'vue-router'
import type { AxiosError, AxiosResponse } from 'axios'
import type { BaseResponseMessage } from '@/dto/common'
import BaseTemplate from '@/views/BaseTemplate.vue'
import MiddlePanel from '@/views/MiddlePanel.vue'
import PanelTitle from '@/components/PanelTitle.vue'
import InputText from '@/components/input/InputText.vue'
import InputTextarea from '@/components/input/InputTextarea.vue'
import Button from '@/components/Button.vue'
import Separator from '@/components/Separator.vue'
import Error from '@/components/Error.vue'

const errorStore = useErrorStore()
const router = useRouter()
const route = useRoute()

const formData = ref<ProjectLogGroupMessage>({
  Id: 0,
  ProjectId: parseInt(String(route.params.id)),
  Name: '',
  Description: '',
  Error: '',
})

const submitForm = () => {
  const req = new AsyncRequestWithAuthorization('http://localhost:8000/new-log-group', {
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
  console.log(formData.value)
  req.setData(formData.value)
  req.post()
}
</script>

<template>
  <BaseTemplate title="New log group">
    <MiddlePanel>
      <Error />
      <PanelTitle icon="project" text="new log group" :sep="false" />
      <InputText v-model="formData.Name" text="Name" name="name" />
      <InputTextarea v-model="formData.Description" text="Description" name="description" />
      <Separator />
      <Button
        @click="submitForm"
        type="button"
        class="create-btn"
        icon="checkbox"
        text="Create"
      />
    </MiddlePanel>
  </BaseTemplate>
</template>

<style scoped lang="scss">
.create-btn {
  margin: 10px;
  width: 200px;
  margin-left: auto;
}
</style>
