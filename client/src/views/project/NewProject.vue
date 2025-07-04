<script setup lang="ts">
import type { MsgProject } from '@/dto/project'
import { ref } from 'vue'
import { MutatedAsyncRequest } from '@/common/request'
import { useErrorStore } from '@/stores/error'
import { useRouter } from 'vue-router'
import type { AxiosError, AxiosResponse } from 'axios'
import type { MsgBaseResponse } from '@/dto/common'
import BaseTemplate from '@/views/BaseTemplate.vue'
import MiddlePanel from '@/views/MiddlePanel.vue'
import Error from '@/components/Error.vue'
import PanelTitle from '@/components/PanelTitle.vue'
import Separator from '@/components/Separator.vue'
import Button from '@/components/Button.vue'
import { InputText, InputTextarea } from '@/components/input/index'

const errorStore = useErrorStore()
const router = useRouter()

const formData = ref<MsgProject>({
  Id: 0,
  UserId: 0,
  Name: '',
  Description: '',
  Error: '',
  Author: undefined,
})

const submitForm = () => {
  const req = new MutatedAsyncRequest('http://localhost:8000/new-project', {
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    withCredentials: true,
  })
  req.onResponse(async (response: AxiosResponse) => {
    const baseResponse = response.data as MsgBaseResponse
    if (!baseResponse.Ok) {
      errorStore.setText(baseResponse.Error)
    } else {
      router.push('/')
    }
  })
  req.onError((error: AxiosError) => {
    errorStore.setText('unexpected error: ' + error.message)
  }, errorStore)
  req.setData(formData.value)
  req.post()
}
</script>

<template>
  <BaseTemplate title="New project">
    <MiddlePanel>
      <Error />
      <PanelTitle icon="project" text="new project" :sep="false" />

      <InputText v-model="formData.Name" text="Name" name="name" />

      <InputTextarea v-model="formData.Description" text="Description" name="description" />

      <Separator />
      <Button @click="submitForm" type="button" class="create-btn" icon="checkbox" text="Create" />
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
