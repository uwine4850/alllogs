<script setup lang="ts">
import AlertPanelTemplate, { closeAlertPanel } from '@/components/alertpanel/AlertPanelTemplate.vue'
import Button from '@/components/Button.vue'
import { AsyncRequestWithAuthorization } from '@/classes/request'
import type { AxiosError, AxiosResponse } from 'axios'
import type { MsgBaseResponse } from '@/dto/common'
import { useErrorStore } from '@/stores/error'
import Error from '@/components/Error.vue'
import { useRoute } from 'vue-router'

const props = defineProps({
  customId: {
    type: String,
  },
  logsContainerClass: {
    type: String,
    required: true,
  },
})

const errorStore = useErrorStore()
const route = useRoute()

const cancelButton = () => {
  closeAlertPanel(props.customId)
}

const clearLogs = () => {
  const req = new AsyncRequestWithAuthorization(`http://localhost:8000/log-items/${route.params.logID}`, {
    withCredentials: true,
  })
  req.onResponse(async (response: AxiosResponse) => {
    const baseResponse = response.data as MsgBaseResponse
    if (!baseResponse.Ok) {
      errorStore.setText(baseResponse.Error)
    } else {
      const logsContainer = document.getElementsByClassName(props.logsContainerClass)
      if (logsContainer.length > 0){
        logsContainer[0].innerHTML = ""
        closeAlertPanel(props.customId)
      }
    }
  })
  req.onError((error: AxiosError) => {
    errorStore.setText("unexpected error: " + error.message)
  }, errorStore, "alertClearLogsStoreId")
  req.delete()
}
</script>

<template>
  <AlertPanelTemplate width="600px" :custom-id="customId">
    <Error store-id="alertClearLogsStoreId" />
    <div class="text">Clear logs?</div>
    <div class="buttons">
      <Button
		  @click="clearLogs"
        type="button"
        class="_btn"
        icon="delete"
        text="Clear logs"
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
