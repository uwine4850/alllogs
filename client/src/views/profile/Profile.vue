<script lang="ts">
import updateIcon from '@/assets/svg/update.svg'
import logoutIcon from '@/assets/svg/log_out.svg'
import refreshIcon from '@/assets/svg/refresh.svg'
import deleteIcon from '@/assets/svg/delete.svg'
import { onMounted } from 'vue'
</script>

<script setup lang="ts">
import MiddlePanel from '@/views/MiddlePanel.vue'
import Button from '@/components/Button.vue'
import Separator from '@/components/Separator.vue'
import InputPassword from '@/components/input/InputPassword.vue'
import BaseTemplate from '@/views/BaseTemplate.vue'
import { useRouter } from 'vue-router'

const router = useRouter()

onMounted(() => {
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
      <div class="base-info">
        <div class="profile-info">
          <div class="avatar">
            <img src="../assets/tmp/1.jpg" alt="" />
          </div>
          <div class="description">
            <div class="username">renxob</div>
            <div class="desc-text">
              It is a long established fact that a reader will be distracted by the readable content
              of a page when looking at its layout. The point of using Lorem Ipsum is that it has
            </div>
          </div>
        </div>
        <div class="profile-btns">
          <Button class="pbtn" :icon="updateIcon" text="Update" link="/profile/update" />
          <Button id="logout-btn" class="pbtn" :icon="logoutIcon" text="Log out" />
        </div>
      </div>
      <InputPassword text="Token" name="token" :readonly="true" value="TOKEN" />
      <div class="token-btns">
        <Button class="tbtn" :icon="refreshIcon" text="Regenerate" />
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
