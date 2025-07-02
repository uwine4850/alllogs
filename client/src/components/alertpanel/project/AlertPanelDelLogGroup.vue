<script setup lang="ts">
import AlertPanelTemplate, { closeAlertPanel } from '@/components/alertpanel/AlertPanelTemplate.vue'
import Button from '@/components/Button.vue'
import { AsyncRequestWithAuthorization } from '@/classes/request'
import type { AxiosError, AxiosResponse } from 'axios'
import type { MsgBaseResponse } from '@/dto/common'
import { useErrorStore } from '@/stores/error'
import Error from '@/components/Error.vue'
import { useRoute, useRouter } from 'vue-router'

const errorStore = useErrorStore()
const router = useRouter()
const route = useRoute()

const cancelButton = () => {
  closeAlertPanel()
}

const deleteProject = () => {
  const req = new AsyncRequestWithAuthorization(`http://localhost:8000/project/${route.params.projId}/log-group/${route.params.logId}`, {
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
      router.push(`/project/${route.params.projId}`)
    }
  })
  req.onError((error: AxiosError) => {
    errorStore.setText("unexpected error: " + error.message)
  }, errorStore, "alertPanelDelLogGroupStoreId")
  req.delete()
}
</script>

<template>
  <AlertPanelTemplate width="600px">
    <Error store-id="alertPanelDelLogGroupStoreId" />
    <div class="text">Delete current log group?</div>
    <div class="buttons">
      <Button
		@click="deleteProject"
        type="button"
        class="_btn"
        icon="delete"
        text="Delete current log group"
      />
      <Button @click="cancelButton" type="button" class="_btn" icon="delete" text="Cancel" />
    </div>
  </AlertPanelTemplate>
</template>

<style scoped>
.text {
  margin: 10px;
  background-color: transparent;
  font-size: 1.2rem;
}
.buttons {
  margin: 10px;
  display: flex;
  gap: 10px;
  justify-content: space-between;
  background-color: transparent;
  ._btn {
    width: 100%;
  }
}
:deep(._btn .btn) {
  padding: 5px;
}
</style>
