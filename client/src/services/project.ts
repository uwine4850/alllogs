import { AsyncRequestWithAuthorization } from "@/classes/request"
import type { ProjectLogGroupMessage, ProjectMessage } from "@/dto/project"
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