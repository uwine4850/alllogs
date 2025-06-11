import { AsyncRequestWithAuthorization } from "@/classes/request"
import type { ProjectMessage } from "@/dto/project"
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
