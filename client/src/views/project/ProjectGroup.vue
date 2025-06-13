<script setup lang="ts">
import { AsyncRequestWithAuthorization } from '@/classes/request'
import type { AxiosError, AxiosResponse } from 'axios'
import { useErrorStore } from '@/stores/error'
import { useRoute } from 'vue-router'
import { onMounted, ref } from 'vue'
import type { ProjectLogGroupMessage, ProjectMessage } from '@/dto/project'
import ProjectTemplate from '@/views/project/ProjectTemplate.vue'
import Separator from '@/components/Separator.vue'
import Button from '@/components/Button.vue'
import PanelTitle from '@/components/PanelTitle.vue'
import AlertFilter from '@/components/project_group/AlertFilter.vue'
import { addComponent } from '@/utils/component'
import Error from '@/components/Error.vue'
import SvgIcon from '@/components/icons/SvgIcon.vue'
import { getProjectLogGroup } from '@/services/project'

const route = useRoute()
const errorStore = useErrorStore()
const projectRef = ref<ProjectMessage | null>(null)
const logRef = ref<ProjectLogGroupMessage | null>(null)

onMounted(() => {
  getProjectLogGroup(route.params.projID, route.params.logID, logRef, projectRef, errorStore);
})
</script>

<template>
  <ProjectTemplate title="Project group">
    <template #panel-project>
      <Error />
      <div class="base-info-view">
        <div class="biv-name">{{ logRef?.Name }}</div>
        <div class="biv-description">{{ logRef?.Description }}</div>
      </div>
      <Separator />
      <div class="info-line">
        <form class="search-form">
          <input type="text" placeholder="Search..." />
          <button>Search</button>
        </form>
        <Separator :vertical="true" />
        <button
          class="il-button"
          id="filter-btn1"
          @click="addComponent('alert-container', AlertFilter)"
        >
          <SvgIcon name="filter" class="icon" />
          <p>filters</p>
        </button>
        <Separator :vertical="true" />
        <router-link class="il-button" :to="`/project/${projectRef?.Id}`">
          <SvgIcon name="project" class="icon" />
          <p>{{ projectRef?.Name }}</p>
        </router-link>
      </div>
      <Separator class="row-header-sep" />

      <div class="row row-header">
        <div class="cell c-type cell-header">
          <SvgIcon name="info" />
          <p>type</p>
        </div>
        <Separator :vertical="true" />
        <div class="cell c-tag cell-header">
          <SvgIcon name="tag" />
          <p>tag</p>
        </div>
        <Separator :vertical="true" />
        <div class="cell c-time cell-header">
          <SvgIcon name="calendar" />
          <p>time</p>
        </div>
        <Separator :vertical="true" />
        <div class="cell c-text cell-header">
          <SvgIcon name="text" />
          <p>text</p>
        </div>
      </div>
      <Separator />
      <div class="table">
        <div class="row">
          <div class="cell c-type c-type-warn">WARN</div>
          <Separator :vertical="true" />
          <div class="cell c-tag">tag</div>
          <Separator :vertical="true" />
          <div class="cell c-time">time</div>
          <Separator :vertical="true" />
          <div class="cell c-text">text</div>
        </div>
        <Separator />
        <div class="row">
          <div class="cell c-type c-type-error">ERROR</div>
          <Separator :vertical="true" />
          <div class="cell c-tag">tag</div>
          <Separator :vertical="true" />
          <div class="cell c-time">time</div>
          <Separator :vertical="true" />
          <div class="cell c-text">text</div>
        </div>
        <Separator />
        <div class="row">
          <div class="cell c-type c-type-info">INFO</div>
          <Separator :vertical="true" />
          <div class="cell c-tag">tag</div>
          <Separator :vertical="true" />
          <div class="cell c-time">time</div>
          <Separator :vertical="true" />
          <div class="cell c-text">text</div>
        </div>
        <Separator />
      </div>
    </template>
    <template #panel-menu>
      <PanelTitle icon="project" text="log group management" />
      <div class="pm-wrapper">
        <Button class="pm-button" icon="upload" text="Export as JSON" link="" />
        <Button class="pm-button" icon="update" text="Update" :link="`/project/${projectRef?.Id}/log-group/${logRef?.Id}/update`" />
        <Button class="pm-button" icon="clear" text="Clear" link="" />
        <Button class="pm-button" icon="bell" text="Sutup notfications" link="" />
      </div>
    </template>
  </ProjectTemplate>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;
@use '@/assets/style/presets.scss' as ps;

.base-info-view {
  background-color: transparent;
  box-sizing: border-box;
  padding: 10px;
  .biv-name {
    background-color: transparent;
    font-size: 1.25rem;
    margin-bottom: 5px;
  }
  .biv-description {
    background-color: transparent;
    font-size: 1.1rem;
    font-family: vars.$fnt-hint-madurai;
  }
}
.info-line {
  height: 45px;
  min-height: 45px;
  display: flex;
  background-color: transparent;
  .search-form {
    display: flex;
    width: 100%;
    input {
      border: none;
      outline: none;
      background-color: vars.$input-color;
      padding: 5px 10px;
      font-family: vars.$fnt-hint-madurai;
      font-size: 1.1rem;
      @include ps.shadow-panel;
      width: 100%;
    }
    button {
      padding: 0 10px;
      font-size: 1.1rem;
      border: none;
      outline: none;
      background-color: vars.$inner-button;
      @include ps.inner-shadow-panel;
      &:hover {
        transition: 0.2s;
        cursor: pointer;
        filter: brightness(90%);
      }
    }
  }
  .il-button {
    display: flex;
    margin: auto 0;
    height: 100%;
    gap: 5px;
    padding: 0 10px;
    background-color: vars.$inner-button;
    border: none;
    outline: none;
    text-decoration: none;
    position: relative;
    @include ps.inner-shadow-panel;
    &:hover {
      transition: 0.2s;
      cursor: pointer;
      filter: brightness(90%);
    }
    .icon {
      width: 20px;
      margin: auto 0;
      background-color: transparent;
    }
    p {
      background-color: transparent;
      margin: auto 0;
      font-size: 1.1rem;
    }
  }
}

.row-header-sep {
  z-index: 1;
}

.table {
  width: 100%;
  height: 100%;
  overflow-y: scroll;
  background-color: transparent;
  ::-webkit-scrollbar {
    width: 0px;
    height: 0px;
  }
  scrollbar-width: none;
}

.row-header {
  font-size: 1.1rem;
  font-family: vars.$fnt-gabarito;
  .cell-header {
    font-size: 1.25rem;
    background-color: vars.$focus-color !important;
    font-family: vars.$fnt-gabarito !important;
    display: flex;
    img {
      background-color: transparent;
      margin: auto 0;
    }
    p {
      background-color: transparent;
      margin-left: 5px;
    }
  }
}

.row {
  background-color: transparent;
  display: flex;
  height: 40px;
  .cell {
    background-color: vars.$secondary-color;
    padding: 10px;
    text-align: left;
    font-size: 1.1rem;
    font-family: vars.$fnt-hint-madurai;
    @include ps.inner-shadow-panel;
  }
  .c-type {
    width: 80px;
    font-family: vars.$fnt-gabarito;
  }
  .c-type-error {
    color: #d80000;
  }
  .c-type-warn {
    color: #dddd00;
  }
  .c-type-info {
    color: #8294d3;
  }
  .c-tag {
    width: 100px;
  }
  .c-time {
    width: 100px;
  }
  .c-text {
    flex: 1;
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
