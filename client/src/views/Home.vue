<script setup lang="ts">
import { AsyncRequestWithAuthorization } from '@/classes/request'
import type { AxiosError, AxiosResponse } from 'axios'
import type { ProjectMessage } from '@/dto/project'
import { useErrorStore } from '@/stores/error'
import { ref } from 'vue'
import BaseTemplate from '../views/BaseTemplate.vue'
import Panel from '../components/Panel.vue'
import PanelTitle from '../components/PanelTitle.vue'
import PanelItem from '@/components/PanelItem.vue'
import type { ProfileMessage } from '@/dto/profile'

const projectsRef = ref<ProjectMessage[]>()
const errorStore = useErrorStore()

let profileData: ProfileMessage
const profileJsonData = sessionStorage.getItem('profile')
if (profileJsonData) {
  profileData = JSON.parse(profileJsonData) as ProfileMessage
}

const req = new AsyncRequestWithAuthorization(`http://localhost:8000/all-projects/${profileData!.UserId}`, {
  withCredentials: true,
})
req.onResponse(async (response: AxiosResponse) => {
  const projectMessages = response.data as ProjectMessage[]
  if (projectMessages[0].Error != '') {
    console.log('Error: ', projectMessages[0].Error)
  } else {
    projectsRef.value = projectMessages
  }
})
req.onError((error: AxiosError) => {
  errorStore.setText("unexpected error: " + error.message)
}, errorStore)
req.get()
</script>

<template>
  <BaseTemplate title="Home">
    <div class="home-content">
      <Panel class="panel">
        <PanelTitle icon="project" text="my projects" />
        <div class="panel-content-wrapper">
          <PanelItem
            v-for="project in projectsRef"
            :key="project.Id"
            :title="project.Name"
            :descr="project.Description"
            :link="`/project/${project.Id}`"
          />
        </div>
      </Panel>
      <Panel class="panel">
        <PanelTitle icon="group" text="my groups" />
        <div class="panel-content-wrapper">
          <PanelItem title="Group name" descr="DESCR" link="#" />
        </div>
      </Panel>
    </div>
  </BaseTemplate>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;

.home-content {
  width: 100%;
  height: calc(100vh - 58px);
  box-sizing: border-box;
  padding: 20px;
  display: flex;
  gap: 20px;
  .panel {
    flex: 1;
    display: flex;
    flex-direction: column;
    :deep(.panel) {
      display: flex;
      flex-direction: column;
    }
    .panel-content-wrapper {
      background-color: transparent;
      flex-grow: 1;
      overflow-y: scroll;
    }
  }
}
</style>
