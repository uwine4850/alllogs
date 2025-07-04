<script setup lang="ts">
import type { MsgProjectLogGroup, MsgProject } from '@/dto/project'
import { onMounted, ref, watch } from 'vue'
import { MutatedAsyncRequest } from '@/common/request'
import { useErrorStore } from '@/stores/error'
import { useRoute, useRouter } from 'vue-router'
import type { AxiosError, AxiosResponse } from 'axios'
import type { MsgBaseResponse } from '@/dto/common'
import BaseTemplate from '@/views/BaseTemplate.vue'
import MiddlePanel from '@/views/MiddlePanel.vue'
import PanelTitle from '@/components/PanelTitle.vue'
import InputText from '@/components/input/InputText.vue'
import InputTextarea from '@/components/input/InputTextarea.vue'
import Button from '@/components/Button.vue'
import Separator from '@/components/Separator.vue'
import Error from '@/components/Error.vue'
import { getProjectLogGroup } from '@/services/project'
import { openAlertPanel } from '@/components/alertpanel/AlertPanelTemplate.vue'
import AlertPanelDelLogGroup from '@/components/alertpanel/project/AlertPanelDelLogGroup.vue'
import type { MsgProfile } from '@/dto/profile'

const errorStore = useErrorStore()
const router = useRouter()
const route = useRoute()
const projectRef = ref<MsgProject | null>(null)
const logRef = ref<MsgProjectLogGroup | null>(null)

onMounted(() => {
  getProjectLogGroup(route.params.projId, route.params.logId, logRef, projectRef, errorStore)
})

const formData = ref<MsgProjectLogGroup>({
  Id: parseInt(String(route.params.logId)),
  ProjectId: parseInt(String(route.params.projId)),
  Name: '',
  Description: '',
  Error: '',
  AuthorToken: '',
})

watch(logRef, (log) => {
  if (log) {
    let profileData: MsgProfile
    const profileJsonData = sessionStorage.getItem('profile')
    if (profileJsonData) {
      profileData = JSON.parse(profileJsonData) as MsgProfile
      if (projectRef.value?.UserId != profileData.UserId) {
        router.replace('/error?code=403 Forbidden&text=no access for log group updates')
        return
      }
    }

    formData.value.Name = log.Name
    formData.value.Description = log.Description
  }
})

const submitForm = () => {
  const req = new MutatedAsyncRequest('http://localhost:8000/log-group', {
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
      router.push(`/project-detail/${projectRef.value?.Id}/log-group/${logRef.value?.Id}`)
    }
  })
  req.onError((error: AxiosError) => {
    errorStore.setText('unexpected error: ' + error.message)
  }, errorStore)
  console.log(formData.value)
  req.setData(formData.value)
  req.patch()
}
</script>

<template>
  <BaseTemplate title="Update log group">
    <AlertPanelDelLogGroup />
    <MiddlePanel>
      <Error />
      <PanelTitle icon="project" text="update log group" :sep="false" />
      <InputText v-model="formData.Name" text="Name" name="name" />
      <InputTextarea v-model="formData.Description" text="Description" name="description" />
      <Separator />
      <div class="buttons">
        <Button
          @click="openAlertPanel"
          type="button"
          class="lbutton delete-btn"
          icon="delete"
          text="Delete"
        />
        <Button @click="submitForm" type="button" class="lbutton" icon="checkbox" text="Update" />
      </div>
    </MiddlePanel>
  </BaseTemplate>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;

.buttons {
  display: flex;
  padding: 10px;
  gap: 10px;
}
.lbutton {
  width: 100%;
  margin-left: auto;
}
.delete-btn {
  :deep(.btn) {
    background-color: vars.$color-red;
    &:hover {
      background-color: vars.$color-red;
      cursor: pointer;
      transition: 0.2s;
      filter: brightness(90%);
    }
  }
}
</style>
