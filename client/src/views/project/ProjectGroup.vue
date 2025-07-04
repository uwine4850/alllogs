<script setup lang="ts">
import { useErrorStore } from '@/stores/error'
import { useRoute } from 'vue-router'
import { onBeforeUnmount, onMounted, ref, watch } from 'vue'
import {
  isMsgLogItem,
  type MsgLogItem,
  type MsgLogItemPayload,
  type MsgLogItemsFilter,
  type MsgProjectLogGroup,
  type MsgProject,
} from '@/dto/project'
import ProjectTemplate from '@/views/project/ProjectTemplate.vue'
import Separator from '@/components/Separator.vue'
import Button from '@/components/Button.vue'
import PanelTitle from '@/components/PanelTitle.vue'
import AlertFilter from '@/components/alertpanel/project/AlertFilter.vue'
import Error from '@/components/Error.vue'
import SvgIcon from '@/components/icons/SvgIcon.vue'
import { getLogGroupItems, getProjectLogGroup } from '@/services/project'
import { MyWebsocket } from '@/common/websocket'
import { WrappedObserver } from '@/common/observer'
import { openAlertPanel } from '@/components/alertpanel/AlertPanelTemplate.vue'
import { MutatedAsyncRequest } from '@/common/request'
import type { AxiosError, AxiosResponse } from 'axios'
import AlertClearLogs from '@/components/alertpanel/project/AlertClearLogs.vue'

const route = useRoute()
const errorStore = useErrorStore()
const projectRef = ref<MsgProject | null>(null)
const logRef = ref<MsgProjectLogGroup | null>(null)
const logItemsPayloadRef = ref<MsgLogItemPayload[] | null>(null)
const isLastLogs = ref<boolean>(false)

enum LogMessageType {
  TYPE_ERROR = 0,
  TYPE_LOGITEM = 1,
}

let filterFormRef = ref<MsgLogItemsFilter>({
  Text: '',
  Type: '',
  Tag: '',
  DateTime: '',
})

let wrappedObserver: WrappedObserver
onMounted(() => {
  getProjectLogGroup(route.params.projID, route.params.logID, logRef, projectRef, errorStore)
  getLogGroupItems(
    route.params.logID,
    -1,
    10,
    logItemsPayloadRef,
    filterFormRef,
    isLastLogs,
    errorStore,
  )
  const CLASS_NAME = 'lastlog'
  wrappedObserver = new WrappedObserver(CLASS_NAME)
  wrappedObserver.onTrigger((el: HTMLElement) => {
    if (isLastLogs.value) {
      wrappedObserver.intersectionObserver?.disconnect()
      return
    }
    el.classList.remove(CLASS_NAME)
    const lastLogID = el.getAttribute('data-logID')
    getLogGroupItems(
      route.params.logID,
      lastLogID,
      10,
      logItemsPayloadRef,
      filterFormRef,
      isLastLogs,
      errorStore,
    )
  })
  wrappedObserver.watch()

  const initialEl = document.querySelector<HTMLElement>(`.${CLASS_NAME}`)
  if (initialEl) {
    wrappedObserver.observeNow(initialEl)
  }
})

var socket: MyWebsocket
var items = ref<MsgLogItemPayload[]>([])

watch(logRef, (log) => {
  socket = new MyWebsocket<MsgLogItem>(
    'log_item',
    `ws://localhost:8000/logitem?token=${log?.AuthorToken}`,
  )
  socket.OnOpen(() => {
    console.log('CLIENT log connected')
  })
  socket.OnClose(() => {
    console.log('CLIENT log webSocket closed')
  })

  socket.OnMessage((event: MessageEvent) => {
    const data = JSON.parse(event.data)
    if (data && isMsgLogItem(data)) {
      if (data.Type == LogMessageType.TYPE_ERROR) {
        errorStore.setText(data.Error)
        return
      }
      if (data.Payload && data.Type == LogMessageType.TYPE_LOGITEM) {
        items.value.unshift(data.Payload)
      }
    }
  })

  socket.Watch()
})

onBeforeUnmount(() => {
  socket?.Close()
  wrappedObserver.mutationObserver?.disconnect()
  wrappedObserver.intersectionObserver?.disconnect()
})

function exportJson() {
  const req = new MutatedAsyncRequest(
    `http://localhost:8000/logs-export-json/${logRef.value?.Id}`,
    {
      withCredentials: true,
      responseType: 'blob',
    },
  )
  req.onResponse(async (response: AxiosResponse) => {
    const blob = new Blob([response.data], { type: 'application/json' })
    const link = document.createElement('a')
    link.href = window.URL.createObjectURL(blob)
    link.download = `PJ-${projectRef.value?.Name}-LG-${logRef.value?.Name}-logs.json`
    link.click()
    window.URL.revokeObjectURL(link.href)
  })
  req.onError((error: AxiosError) => {
    errorStore.setText('unexpected error: ' + error.message)
  }, errorStore)
  req.get()
}

function searchLogs() {
  logItemsPayloadRef.value = null
  isLastLogs.value = false
  getLogGroupItems(
    route.params.logID,
    -1,
    10,
    logItemsPayloadRef,
    filterFormRef,
    isLastLogs,
    errorStore,
  )
}
defineExpose({ searchLogs })
</script>

<template>
  <AlertFilter
    v-on:search="searchLogs"
    v-model:model-value-type="filterFormRef.Type"
    v-model:model-value-tag="filterFormRef.Tag"
    v-model:model-value-date-time="filterFormRef.DateTime"
  />
  <AlertClearLogs custom-id="alert-clear-logs" logs-container-class="table" />
  <ProjectTemplate title="Project group">
    <template #panel-project>
      <Error />
      <div class="base-info-view">
        <div class="biv-name">{{ logRef?.Name }}</div>
        <div class="biv-description">{{ logRef?.Description }}</div>
      </div>
      <Separator />
      <div class="log-group-id">
        <span>Log group id: </span>
        <div>{{ logRef?.Id }}</div>
      </div>
      <Separator />
      <div class="info-line">
        <form class="search-form">
          <input v-model="filterFormRef.Text" type="text" placeholder="Search..." />
          <button type="button" @click="searchLogs">Search</button>
        </form>
        <Separator :vertical="true" />
        <button class="il-button" id="filter-btn1" @click="openAlertPanel()">
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
        <div v-for="item in items" :key="item.Id">
          <div class="row">
            <div class="cell c-type" :class="`c-type-${item.Type}`">
              {{ item.Type }}
            </div>
            <Separator :vertical="true" />
            <div class="cell c-tag" :title="item.Tag">{{ item.Tag }}</div>
            <Separator :vertical="true" />
            <div class="cell c-time">{{ item.Datetime }}</div>
            <Separator :vertical="true" />
            <div class="cell c-text" :title="item.Text">{{ item.Text }}</div>
          </div>
          <Separator />
        </div>
        <div v-for="(log, index) in logItemsPayloadRef" :key="log.Id">
          <div
            class="row"
            :class="{
              lastlog: index == Object.keys(logItemsPayloadRef!).length - 1 && !isLastLogs,
            }"
            :data-logID="log.Id"
          >
            <div class="cell c-type" :class="`c-type-${log.Type}`">
              {{ log.Type }}
            </div>
            <Separator :vertical="true" />
            <div class="cell c-tag" :title="log.Tag">{{ log.Tag }}</div>
            <Separator :vertical="true" />
            <div class="cell c-time">{{ log.Datetime }}</div>
            <Separator :vertical="true" />
            <div class="cell c-text" :title="log.Text">{{ log.Text }}</div>
          </div>
          <Separator />
        </div>
      </div>
    </template>
    <template #panel-menu>
      <PanelTitle icon="project" text="log group management" />
      <div class="pm-wrapper">
        <Button @click="exportJson" class="pm-button" icon="upload" text="Export as JSON" link="" />
        <Button
          class="pm-button"
          icon="update"
          text="Update"
          :link="`/project/${projectRef?.Id}/log-group/${logRef?.Id}/update`"
        />
        <Button
          class="pm-button"
          icon="clear"
          text="Clear"
          @click="openAlertPanel('alert-clear-logs')"
        />
        <Button class="pm-button" icon="bell" text="Setup notfications" link="" />
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
.log-group-id {
  display: flex;
  font-size: 1.1rem;
  padding: 5px 10px;
  span {
    margin-right: 5px;
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
    text-wrap: nowrap;
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
  .c-type-ERROR {
    color: #d80000;
  }
  .c-type-WARN {
    color: #dddd00;
  }
  .c-type-INFO {
    color: #8294d3;
  }
  .c-tag {
    width: 100px;
    text-wrap: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }
  .c-time {
    width: 150px;
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
