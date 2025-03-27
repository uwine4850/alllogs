<script lang="ts">
import projectIcon from '@/assets/svg/project.svg'
import addIcon from '@/assets/svg/add.svg'
import groupIcon from '@/assets/svg/group.svg'
import userIcon from '@/assets/svg/user.svg'
import exportIcon from '@/assets/svg/upload.svg'
import updateIcon from '@/assets/svg/update.svg'
</script>

<script setup lang="ts">
import ProjectTemplate from '@/views/ProjectTemplate.vue'
import Button from '../components/Button.vue'
import Separator from '../components/Separator.vue'
import PanelTitle from '../components/PanelTitle.vue'
</script>

<template>
  <ProjectTemplate title="Project">
    <template #panel-project>
      <div class="proj-base-view">
        <div class="proj-name">Project name</div>
        <div class="proj-description">
          It is a long established fact that a reader will be distracted by the readable content of
          a page when looking at its layout. The point of using Lorem Ipsum is that it has
        </div>
      </div>
      <Separator />
      <div class="info-line">
        <Separator class="info-sep" :vertical="true" />
        <a class="author-info" href="">
          <div class="ai-avatar">
            <img src="../assets/tmp/1.jpg" alt="" />
          </div>
          <div class="ai-username">renxob</div>
        </a>
      </div>
      <Separator />
      <PanelTitle :icon="groupIcon" text="log groups" />
      <div class="log-group-list">
        <router-link class="log-group" :to="`/project-group`"> Auth </router-link>
        <router-link class="log-group" :to="`/project-group`"> Database </router-link>
        <router-link class="log-group" :to="`/project-group`"> Router </router-link>
      </div>
    </template>
    <template #panel-menu>
      <PanelTitle :icon="projectIcon" text="project management" />
      <div class="pm-wrapper">
        <Button class="pm-button" :icon="addIcon" text="New log group" link="#" />
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
