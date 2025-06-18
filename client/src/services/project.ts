import { AsyncRequestWithAuthorization } from "@/classes/request"
import type { LogItemMessage, LogItemPayload, LogItemsFilterMessage, ProjectLogGroupMessage, ProjectMessage } from "@/dto/project"
import type { AxiosError, AxiosResponse } from "axios"
import { useErrorStore } from '@/stores/error'
import { type Ref } from "vue"

export function getProject(id: any, projectRef: Ref<ProjectMessage | null>, errorStore: ReturnType<typeof useErrorStore>){
  const req = new AsyncRequestWithAuthorization(
    `http://localhost:8000/project/${id}`,
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
  req.onError((error: AxiosError) => {
    errorStore.setText("unexpected error: " + error.message)
  }, errorStore)
  req.get()
}

export function getProjectLogGroup(
  projectId: any,
  groupId: any,
  groupRef: Ref<ProjectLogGroupMessage | null>,
  projectRef: Ref<ProjectMessage | null>,
  errorStore: ReturnType<typeof useErrorStore>
  ){
  const req = new AsyncRequestWithAuthorization(
    `http://localhost:8000/project-detail/${projectId}/log-group/${groupId}`,
    {
      withCredentials: true,
    },
  )
  req.onResponse(async (response: AxiosResponse) => {
    const _project = response.data['project']
    const _log = response.data['log']
    if (_log) {
      const log = _log as ProjectLogGroupMessage
      if (log.Error != '') {
        errorStore.setText(log.Error)
      } else {
        if (_project) {
          projectRef.value = _project as ProjectMessage
        }
        groupRef.value = log
      }
    }
  })
  req.onError((error: AxiosError) => {
    errorStore.setText("unexpected error: " + error.message)
  }, errorStore)
  req.get()
}

export function getLogGroupItems(
  logGroupId: any,
  startId: any,
  count: any,
  logItemRef: Ref<LogItemPayload[] | null>,
  filterRef: Ref<LogItemsFilterMessage | null>,
  isLastLogs: Ref<boolean>,
  errorStore: ReturnType<typeof useErrorStore>
  ){
  let queryParams = ""
  if (filterRef.value){
    let isFirst = true
    for (const [key, value] of Object.entries(filterRef.value)) {
      if (!value) continue
      if (isFirst) {
        isFirst = false
        queryParams += `${key}=${value}`
      } else {
        queryParams += `&${key}=${value}`
      }
    }
  }
  const req = new AsyncRequestWithAuthorization(
    `http://localhost:8000/log-items/${logGroupId}/${startId}/${count}?${queryParams}`,
    {
      withCredentials: true,
    },
  )
  req.onResponse(async (response: AxiosResponse) => {
    const logs = response.data as LogItemPayload[]
    if (!logs){
      return
    }
    if (logs.length != count){
      isLastLogs.value = true
    }
    if (logItemRef.value){
      logItemRef.value.push(...logs)
    } else {
      logItemRef.value = logs
    }
  })
  req.onError((error: AxiosError) => {
    errorStore.setText("unexpected error: " + error.message)
  }, errorStore)
  req.get()
}