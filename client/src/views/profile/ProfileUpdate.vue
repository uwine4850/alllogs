<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import type { MsgProfile, MsgProfileUpdate } from '@/dto/profile'
import { getProfileData } from '@/services/profile'
import { useRoute, useRouter } from 'vue-router'
import { useErrorStore } from '@/stores/error'
import { AsyncRequestWithAuthorization } from '@/classes/request'
import type { AxiosError, AxiosResponse } from 'axios'
import type { MsgBaseResponse } from '@/dto/common'
import { openAlertPanel } from '@/components/alertpanel/AlertPanelTemplate.vue'
import MiddlePanel from '@/views/MiddlePanel.vue'
import BaseTemplate from '@/views/BaseTemplate.vue'
import PanelTitle from '@/components/PanelTitle.vue'
import InputTextarea from '@/components/input/InputTextarea.vue'
import InputFile from '@/components/input/InputFile.vue'
import InputCheckbox from '@/components/input/InputCheckbox.vue'
import Button from '@/components/Button.vue'
import Separator from '@/components/Separator.vue'
import Error from '@/components/Error.vue'
import AlertPanelDelProfile from '@/components/alertpanel/profile/AlertPanelDelProfile.vue'

const errorStore = useErrorStore()
const router = useRouter()
const route = useRoute()
const profileDataRef = ref<MsgProfile | null>(null)

const formData = ref<MsgProfileUpdate>({
  UID: 0,
  Description: '',
  Avatar: null,
  OldAvatarPath: '',
  DelAvatar: false,
})
const saveChanges = async () => {
  const req = new AsyncRequestWithAuthorization('http://localhost:8000/profile/update', {
    withCredentials: true,
  })
  req.onResponse(async (response: AxiosResponse) => {
    const baseResponse = response.data as MsgBaseResponse
    if (!baseResponse.Ok) {
      errorStore.setText(baseResponse.Error)
    } else {
      const newProfileDataRef = ref<MsgProfile | null>(null)
      await getProfileData(newProfileDataRef, null, route.params.id as string, errorStore)
      sessionStorage.setItem('profile', JSON.stringify(newProfileDataRef.value))
      router.push('/')
    }
  })
  req.onError((error: AxiosError) => {
    errorStore.setText("unexpected error: " + error.message)
  }, errorStore)

  const data = new FormData()
  data.append('UID', String(formData.value.UID))
  data.append('Description', formData.value.Description)
  if (formData.value.Avatar) {
    data.append('Avatar', formData.value.Avatar)
  }
  data.append('OldAvatarPath', formData.value.OldAvatarPath)
  data.append('DelAvatar', formData.value.DelAvatar ? 'true' : 'false')

  req.setData(data)
  req.patch()
}

const showDeleteAlert = () => {
  openAlertPanel()
}

onMounted(() => {
  getProfileData(profileDataRef, null, route.params.id as string, errorStore)
})

watch(profileDataRef, (profile) => {
  if (profile) {
    let profileData: MsgProfile
    const profileJsonData = sessionStorage.getItem('profile')
    if (profileJsonData) {
      profileData = JSON.parse(profileJsonData) as MsgProfile
      if (profileDataRef.value?.UserId != profileData.UserId){
        router.replace("/error?code=403 Forbidden&text=no access for user profile updates")
        return
      }
    }

    formData.value.UID = profile.UserId
    formData.value.Description = profile.Description || ''
    formData.value.DelAvatar = false
    formData.value.OldAvatarPath = profile.Avatar
  }
})
</script>

<template>
  <BaseTemplate title="Profile update">
    <AlertPanelDelProfile />
    <MiddlePanel>
      <Error />
      <PanelTitle icon="user" text="Profile update" :sep="false" />
      <form @submit.prevent="saveChanges">
        <InputTextarea v-model="formData.Description" text="Description" name="description" />
        <InputFile v-model="formData.Avatar" text="Avatar" />
        <InputCheckbox v-model="formData.DelAvatar" text="Delete avatar" inptext="delete" />
        <Separator />
        <div class="update-btns">
          <Button
            @click="showDeleteAlert"
            type="button"
            class="btn btn-delete"
            icon="delete"
            text="Delete user"
          />
          <Button type="submit" class="btn" icon="checkbox" text="Save" />
        </div>
      </form>
    </MiddlePanel>
  </BaseTemplate>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;

.update-btns {
  margin: 10px;
  display: flex;
  gap: 10px;
  justify-content: space-between;
  background-color: transparent;
  .btn {
    width: 100%;
    padding: 0;
  }
  .btn-delete {
    :deep(.btn) {
      background-color: vars.$color-red;
      width: 100%;
      &:hover {
        cursor: pointer;
        transition: 0.2s;
        filter: brightness(90%);
      }
    }
  }
}
</style>
