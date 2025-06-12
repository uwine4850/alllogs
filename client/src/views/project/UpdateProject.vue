<script setup lang="ts">
import { AsyncRequestWithAuthorization } from '@/classes/request';
import type { BaseResponseMessage } from '@/dto/common';
import projectIcon from '@/assets/svg/project.svg'
import checkBoxIcon from '@/assets/svg/checkbox.svg'
import { type ProjectMessage } from '@/dto/project';
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
  <BaseTemplate :title="`Update project - ${formData.Name}`">
    <MiddlePanel>
      <Error />
      <PanelTitle :icon="projectIcon" :text="`Update project - ${formData.Name}`" :sep="false" />
      
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
        <slot id="extra"></slot>
      <Button
        @click="submitForm"
        type="button"
        class="create-btn"
        :icon="checkBoxIcon"
        text="Update"
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
