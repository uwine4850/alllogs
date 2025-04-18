<script lang="ts">
import updateIcon from '@/assets/svg/update.svg'
import logoutIcon from '@/assets/svg/log_out.svg'
import refreshIcon from '@/assets/svg/refresh.svg'
import deleteIcon from '@/assets/svg/delete.svg'
import { onMounted } from 'vue'
import type { GenTokenMessage, ProfileMessage, TokenResponse } from '@/dto/profile'
import { AsyncRequest } from '@/classes/request'
</script>

<script setup lang="ts">
import MiddlePanel from '@/views/MiddlePanel.vue'
import Button from '@/components/Button.vue'
import Separator from '@/components/Separator.vue'
import InputPassword from '@/components/input/InputPassword.vue'
import BaseTemplate from '@/views/BaseTemplate.vue'
import { useRouter, useRoute } from 'vue-router'
import axios, { type AxiosResponse } from 'axios'
import { ref } from 'vue'
import Error from '@/components/Error.vue'
import { useErrorStore } from '@/stores/error'

const errorStore = useErrorStore();
const router = useRouter()
const route = useRoute()

const profileData = ref<ProfileMessage | null>(null)
const token = ref("");
const getProfileData = async () => {
  const req = new AsyncRequest("http://localhost:8000/profile/" + route.params.id, {
    headers: {
      "Content-Type": "application/json",
    },
    withCredentials: true,
  });
  req.onResponse((response: AxiosResponse) => {
    const profileResponse = response.data as ProfileMessage;
    if (profileResponse.Error !== "") {
      errorStore.setText(profileResponse.Error);
    } else {
      profileData.value = profileResponse
      token.value = profileData.value.Token
    }
  });
  req.onError((error: unknown) => {
    errorStore.setText(String(error));
  });
  await req.get()
}


const generateTokenForm = async () => {
  if (!profileData.value){
    return
  }
  const tokenFormData = ref<GenTokenMessage>({
    UserId: profileData.value.Id
  });

  const req = new AsyncRequest("http://localhost:8000/gen-token",  {
    headers: {
      "Content-Type": "application/json",
    },
    withCredentials: true,
  });
  req.onResponse((response: AxiosResponse) => {
    const tokenResponse = response.data as TokenResponse;
    if (tokenResponse.Error !== "") {
      errorStore.setText(tokenResponse.Error);
    } else {
      token.value = tokenResponse.Token;
    }
  });
  req.onError((error: unknown) => {
    errorStore.setText(String(error));
  });
  req.post(tokenFormData.value);
}

onMounted(() => {
  getProfileData()
  const logoutBtn = document.getElementById("logout-btn")
  if (logoutBtn){
    logoutBtn.onclick = function(){
      sessionStorage.removeItem("authJWT");
      router.go(0);
    }
  }
});
</script>

<template>
  <BaseTemplate title="Profile">
    <MiddlePanel class="middle-panel">
      <Error />
      <div class="base-info">
        <div class="profile-info">
          <div class="avatar">
            <img :src="profileData?.Avatar" alt="" />
          </div>
          <div class="description">
            <div class="username">{{ profileData?.User.Username }}</div>
            <div class="desc-text">
             {{ profileData?.Description }}
            </div>
          </div>
        </div>
        <div class="profile-btns">
          <Button class="pbtn" :icon="updateIcon" text="Update" link="/profile/update" />
          <Button id="logout-btn" class="pbtn" :icon="logoutIcon" text="Log out" />
        </div>
      </div>
      <InputPassword text="Token" name="token" :readonly="true" v-model="token" />
      <div class="token-btns">
        <Button @click="generateTokenForm" class="tbtn" :icon="refreshIcon" text="Regenerate" />
        <Separator class="tsep" :vertical="true" />
        <Button class="tbtn tbtn-delete" :icon="deleteIcon" text="Delete" />
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
