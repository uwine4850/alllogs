<script lang="ts">
import projectIcon from '@/assets/svg/project.svg'
import checkBoxIcon from '@/assets/svg/checkbox.svg'
import type { ProjectMessage } from '@/dto/project'
import { ref } from 'vue'
import { AsyncRequestWithAuthorization } from '@/classes/request'
import { useErrorStore } from '@/stores/error'
import { useRouter } from 'vue-router'
import type { AxiosResponse } from 'axios'
import type { BaseResponseMessage } from '@/dto/common'
</script>

<script setup lang="ts">
import BaseTemplate from './BaseTemplate.vue'
import MiddlePanel from './MiddlePanel.vue'
import PanelTitle from '@/components/PanelTitle.vue'
import InputText from '@/components/input/InputText.vue'
import InputTextarea from '@/components/input/InputTextarea.vue'
import Button from '@/components/Button.vue'
import Separator from '@/components/Separator.vue'
import Error from '@/components/Error.vue'

const errorStore = useErrorStore()
const router = useRouter()

const formData = ref<ProjectMessage>({
  UserId: 0,
  Name: '',
  Description: '',
  Error: '',
  Author: undefined
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
  req.onError((error: unknown) => {
    errorStore.setText(String(error))
  })
  console.log(formData.value)
  req.setData(formData.value)
  req.post()
}
</script>

<template>
  <BaseTemplate title="New project">
    <MiddlePanel>
      <Error />
      <PanelTitle :icon="projectIcon" text="new project" :sep="false" />
      <InputText v-model="formData.Name" text="Name" name="name" />
      <InputTextarea v-model="formData.Description" text="Description" name="description" />
      <Separator />
      <Button
        @click="submitForm"
        type="button"
        class="create-btn"
        :icon="checkBoxIcon"
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
