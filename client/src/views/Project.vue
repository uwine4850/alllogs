<script lang="ts">
import projectIcon from '@/assets/svg/project.svg'
import addIcon from '@/assets/svg/add.svg'
import groupIcon from '@/assets/svg/group.svg'
import userIcon from '@/assets/svg/user.svg'
import exportIcon from '@/assets/svg/upload.svg'
import updateIcon from '@/assets/svg/update.svg'
import { useRoute } from 'vue-router'
import { AsyncRequestWithAuthorization } from '@/classes/request'
import type { AxiosResponse } from 'axios'
import type { ProjectLogGroupMessage, ProjectMessage } from '@/dto/project'
import { useErrorStore } from '@/stores/error'
import { onMounted, ref, watch } from 'vue'
</script>

<script setup lang="ts">
import ProjectTemplate from '@/views/ProjectTemplate.vue'
import Button from '../components/Button.vue'
import Separator from '../components/Separator.vue'
import PanelTitle from '../components/PanelTitle.vue'
import Error from '@/components/Error.vue'

const route = useRoute()
const errorStore = useErrorStore()
const projectRef = ref<ProjectMessage | null>(null)
const logGroupsRef = ref<ProjectLogGroupMessage[]>()

const getProject = () => {
  const req = new AsyncRequestWithAuthorization(
    `http://localhost:8000/project/${route.params.id}`,
    {
      withCredentials: true,
    },
  )
  req.onResponse(async (response: AxiosResponse) => {
    const projectMessage = response.data as ProjectMessage
    if (projectMessage.Error != '') {
      errorStore.setText(projectMessage.Error)
    } else {
      projectRef.value = projectMessage
    }
  })
  req.onError((error: unknown) => {
    errorStore.setText(String(error))
  })
  req.get()
}

const getLogGroups = (project_id: number) => {
  const req = new AsyncRequestWithAuthorization(
    `http://localhost:8000/all-log-groups/${project_id}`,
    {
      withCredentials: true,
    },
  )
  req.onResponse(async (response: AxiosResponse) => {
    const projectMessages = response.data as ProjectLogGroupMessage[]
    if (projectMessages[0].Error != '') {
      errorStore.setText(projectMessages[0].Error)
    } else {
      logGroupsRef.value = projectMessages
    }
  })
  req.onError((error: unknown) => {
    errorStore.setText(String(error))
  })
  req.get()
}

onMounted(() => {
  getProject()
})

watch(projectRef, (project) => {
  if (project) {
    getLogGroups(project.Id)
  }
})
</script>

<template>
  <ProjectTemplate title="Project">
    <template #panel-project>
      <Error />
      <div class="proj-base-view">
        <div class="proj-name">{{ projectRef?.Name }}</div>
        <div class="proj-description">
          {{ projectRef?.Description }}
        </div>
      </div>
      <Separator />
      <div class="info-line">
        <Separator class="info-sep" :vertical="true" />
        <router-link class="author-info" :to="`/profile/${projectRef?.Author?.PID}`">
          <div class="ai-avatar">
            <img :src="projectRef?.Author?.Avatar" alt="" />
          </div>
          <div class="ai-username">{{ projectRef?.Author?.Username }}</div>
        </router-link>
      </div>
      <Separator />
      <PanelTitle :icon="groupIcon" text="log groups" />
      <div class="log-group-list">
        <router-link
          v-for="group in logGroupsRef"
          :key="group.Id"
          class="log-group"
          :to="`/project-detail/${projectRef?.Id}/log-group/${group.Id}`"
          >{{ group.Name }}</router-link
        >
      </div>
    </template>
    <template #panel-menu>
      <PanelTitle :icon="projectIcon" text="project management" />
      <div class="pm-wrapper">
        <Button
          class="pm-button"
          :icon="addIcon"
          text="New log group"
          :link="`/project/${route.params.id}/new-log-group`"
        />
        <Button class="pm-button" :icon="addIcon" text="Add group" link="#" />
      </div>
      <Separator />
      <div class="pm-wrapper">
        <Button class="pm-button" :icon="userIcon" text="Users" link="#" />
        <Button class="pm-button" :icon="updateIcon" text="Update" link="#" />
        <Button class="pm-button" :icon="exportIcon" text="Export all as JSON" />
      </div>
    </template>
  </ProjectTemplate>
</template>

<style scoped lang="scss">
@use '../assets/style/global_vars.scss' as vars;
@use '../assets/style/presets.scss' as ps;

.proj-base-view {
  padding: 10px;
  background-color: transparent;
  .proj-name {
    background-color: transparent;
    font-size: 1.25rem;
    margin-bottom: 5px;
  }
  .proj-description {
    background-color: transparent;
    font-size: 1.1rem;
    font-family: vars.$fnt-hint-madurai;
  }
}
.info-line {
  height: 45px;
  display: flex;
  background-color: transparent;
  .info-sep {
    margin-left: auto;
  }
  .author-info {
    display: flex;
    margin: auto 0;
    height: 100%;
    padding: 0 10px;
    background-color: vars.$inner-button;
    text-decoration: none;
    @include ps.inner-shadow-panel;
    &:hover {
      transition: 0.2s;
      filter: brightness(90%);
    }
    .ai-avatar {
      width: 35px;
      height: 35px;
      overflow: hidden;
      border-radius: 50%;
      margin: auto 0;
      margin-right: 5px;
      img {
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
    }
    .ai-username {
      margin: auto 0;
      font-size: 1.25rem;
      background-color: transparent;
    }
  }
}
.log-group-list {
  margin: 10px;
  background-color: transparent;
  display: flex;
  flex-direction: column;
  overflow-y: scroll;
  .log-group {
    border-radius: 3px;
    box-sizing: border-box;
    padding: 10px;
    background-color: vars.$secondary-color;
    margin-bottom: 10px;
    text-decoration: none;
    &:hover {
      transition: 0.2s;
      background-color: vars.$focus-color;
    }
  }
}

.pm-wrapper {
  box-sizing: border-box;
  padding: 10px;
  background-color: transparent;
  display: flex;
  flex-direction: column;
  gap: 10px;
  .pm-button {
    :deep(.btn) {
      padding: 10px 0;
    }
  }
}
</style>
