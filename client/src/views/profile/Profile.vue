<script lang="ts">
import updateIcon from '@/assets/svg/update.svg'
import logoutIcon from '@/assets/svg/log_out.svg'
import refreshIcon from '@/assets/svg/refresh.svg'
import deleteIcon from '@/assets/svg/delete.svg'
import { onMounted } from 'vue'
import type { ProfileMessage } from '@/dto/profile'
import { deleteToken, generateTokenForm, getProfileData } from '@/services/profile'
import { logout } from '@/services/auth'
</script>

<script setup lang="ts">
import MiddlePanel from '@/views/MiddlePanel.vue'
import Button from '@/components/Button.vue'
import Separator from '@/components/Separator.vue'
import InputPassword from '@/components/input/InputPassword.vue'
import BaseTemplate from '@/views/BaseTemplate.vue'
import { useRoute } from 'vue-router'
import { ref, type Ref } from 'vue'
import Error from '@/components/Error.vue'
import { useErrorStore } from '@/stores/error'

const errorStore = useErrorStore()
const route = useRoute()

const profileDataRef = ref<ProfileMessage | null>(null)
const tokenRef = ref('')

const handleGetProfileData = () => {
  getProfileData(profileDataRef, tokenRef, route.params.id as string, errorStore)
}

const handleGenerateTokenForm = () => {
  generateTokenForm(tokenRef, profileDataRef.value, errorStore)
}

const handleDeleteToken = () => {
  deleteToken(tokenRef, profileDataRef.value, errorStore)
}

onMounted(async () => {
  handleGetProfileData()
  const logoutBtn = document.getElementById('logout-btn')
  if (logoutBtn) {
    logoutBtn.onclick = function () {
      if(profileDataRef.value?.User?.Id){
        try{
          logout(profileDataRef.value.User.Id)
        } catch (e) {
          errorStore.setText(String(e))
        }
      }
    }
  }
})
</script>

<template>
  <BaseTemplate title="Profile">
    <MiddlePanel class="middle-panel">
      <Error />
      <div class="base-info">
        <div class="profile-info">
          <div class="avatar">
            <img :src="profileDataRef?.Avatar" alt="" />
          </div>
          <div class="description">
            <div class="username">{{ profileDataRef?.User?.Username }}</div>
            <div class="desc-text">
              {{ profileDataRef?.Description }}
            </div>
          </div>
        </div>
        <div class="profile-btns">
          <Button
            class="pbtn"
            :icon="updateIcon"
            text="Update"
            :link="`/profile/update/${profileDataRef?.UserId}`"
          />
          <Button id="logout-btn" class="pbtn" :icon="logoutIcon" text="Log out" />
        </div>
      </div>
      <InputPassword text="Token" name="token" :readonly="true" v-model="tokenRef" />
      <div class="token-btns">
        <Button
          @click="handleGenerateTokenForm"
          class="tbtn"
          :icon="refreshIcon"
          text="Regenerate token"
        />
        <Separator class="tsep" :vertical="true" />
        <Button
          @click="handleDeleteToken"
          class="tbtn tbtn-delete"
          :icon="deleteIcon"
          text="Delete token"
        />
      </div>
    </MiddlePanel>
  </BaseTemplate>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;
@use '@/assets/style/presets.scss' as ps;

.middle-panel {
  height: calc(100vh - 58px);
}

.base-info {
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  padding: 20px;
  padding-bottom: 10px;
  background-color: transparent;
  .profile-info {
    background-color: transparent;
    display: flex;
    .avatar {
      width: 150px;
      min-width: 150px;
      height: 150px;
      min-height: 150px;
      overflow: hidden;
      border-radius: 5px;
      img {
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
    }
    .description {
      background-color: transparent;
      display: flex;
      flex-direction: column;
      margin-left: 10px;
      .username {
        font-size: 1.5rem;
        margin-bottom: 5px;
        background-color: transparent;
      }
      .desc-text {
        font-size: 1.1rem;
        font-family: vars.$fnt-hint-madurai;
        background-color: transparent;
      }
    }
  }
  .profile-btns {
    width: 100%;
    display: flex;
    gap: 20px;
    margin-top: 10px;
    background-color: transparent;
    .pbtn {
      width: 100%;
      :deep(.btn) {
        padding: 10px 0;
      }
    }
  }
}
.token-btns {
  display: flex;
  background-color: transparent;
  height: fit-content;
  height: 50px;
  .tbtn {
    width: 100%;
    border-radius: 0;
    :deep(.btn) {
      padding: 0;
      border-radius: 0;
      height: 50px;
      background-color: vars.$inner-button;
      @include ps.inner-shadow-panel;
      &:hover {
        cursor: pointer;
        transition: 0.2s;
        filter: brightness(90%);
      }
    }
  }
  .tbtn-delete {
    :deep(.btn) {
      background-color: vars.$color-red;
    }
  }
}
</style>
