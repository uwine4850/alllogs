<script setup lang="ts">
import { AsyncRequestWithAuthorization } from '@/classes/request';
import type { MsgBaseResponse } from '@/dto/common';
import { type MsgProject } from '@/dto/project';
import { getProject } from '@/services/project';
import { useErrorStore } from '@/stores/error';
import type { AxiosError, AxiosResponse } from 'axios';
import { onMounted, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import BaseTemplate from '@/views/BaseTemplate.vue';
import MiddlePanel from '@/views/MiddlePanel.vue';
import Error from '@/components/Error.vue';
import PanelTitle from '@/components/PanelTitle.vue';
import Separator from '@/components/Separator.vue';
import Button from '@/components/Button.vue';
import { InputText, InputTextarea } from '@/components/input/index';
import AlertPanelDelProject from '@/components/alertpanel/project/AlertPanelDelProject.vue';
import { openAlertPanel } from '@/components/alertpanel/AlertPanelTemplate.vue';
import type { MsgProfile } from '@/dto/profile';

var route = useRoute();
var router = useRouter();
var errorStore = useErrorStore();

const formData = ref<MsgProject>({
  Id: 0,
  UserId: 0,
  Name: '',
  Description: '',
  Error: '',
  Author: undefined,
})

const projectRef = ref<MsgProject | null>(null)
onMounted(() => {
  getProject(route.params.id, projectRef, errorStore);
})
watch(projectRef, (project) => {
  if(project){
    let profileData: MsgProfile
    const profileJsonData = sessionStorage.getItem('profile')
    if (profileJsonData) {
      profileData = JSON.parse(profileJsonData) as MsgProfile
      if (projectRef.value?.UserId != profileData.UserId){
        router.replace("/error?code=403 Forbidden&text=no access for project updates")
        return
      }
    }

    formData.value.Id = project.Id;
    formData.value.Name = project.Name;
    formData.value.Description = project.Description;
  }
});

const submitForm = () => {
  const req = new AsyncRequestWithAuthorization('http://localhost:8000/project', {
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
  <BaseTemplate :title="`Update project - ${formData.Name}`">
    <AlertPanelDelProject />
    <MiddlePanel>
      <Error />
      <PanelTitle icon="project" :text="`Update project - ${formData.Name}`" :sep="false" />
      
      <InputText
        v-model="formData.Name"
        text="Name"
        name="name"
      />
      
      <InputTextarea
        v-model="formData.Description"
        text="Description"
        name="description"
      />

      <Separator />
      <div class="buttons">
        <Button
          @click="openAlertPanel"
          type="button"
          class="lbutton delete-btn"
          icon="delete"
          text="Delete"
        />
        <Button
          @click="submitForm"
          type="button"
          class="lbutton"
          icon="checkbox"
          text="Update"
        />
      </div>
    </MiddlePanel>
  </BaseTemplate>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;

.buttons{
  display: flex;
  padding: 10px;
  gap: 10px;
}
.lbutton {
  width: 100%;
  margin-left: auto;
}
.delete-btn{
  :deep(.btn){
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
