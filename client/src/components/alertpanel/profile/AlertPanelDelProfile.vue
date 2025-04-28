<script setup lang="ts">
import AlertPanelTemplate, { closeAlertPanel } from '@/components/alertpanel/AlertPanelTemplate.vue'
import Button from '@/components/Button.vue';
import deleteIcon from '@/assets/svg/delete.svg';
import { ref } from 'vue';
import { AsyncRequestWithAuthorization } from '@/classes/request';
import type { AxiosResponse } from 'axios';
import type { BaseResponseMessage } from '@/dto/common';
import { useErrorStore } from '@/stores/error';
import Error from '@/components/Error.vue';
import { useRouter } from 'vue-router';

const errorStore = useErrorStore();
const router = useRouter();

const cancelButton = () => {
  closeAlertPanel();
}

const deleteUser = () => {
  const req = new AsyncRequestWithAuthorization("http://localhost:8000/profile/del", {
    headers: {
      'Content-Type': 'application/json',
    },
    withCredentials: true,
  });
  req.onResponse(async (response: AxiosResponse) => {
    const baseResponse = response.data as BaseResponseMessage
    if (!baseResponse.Ok) {
      errorStore.setText(baseResponse.Error)
    } else {
      sessionStorage.removeItem('profile')
      sessionStorage.removeItem('authJWT')
      router.push('/login');
    }
  })
  req.onError((error: unknown) => {
    errorStore.setText(String(error))
  })
  req.delete();
}

</script>

<template>
  <AlertPanelTemplate>
    <Error />
    <div class="text">Delete profile?</div>
    <div class="buttons">
        <Button @click="deleteUser" type="button" class="_btn" :icon="deleteIcon" text="Delete profile" />
        <Button @click="cancelButton" type="button" class="_btn" :icon="deleteIcon" text="Cancel" />
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
