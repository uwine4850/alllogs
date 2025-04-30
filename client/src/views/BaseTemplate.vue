<script lang="ts">
import addIcon from '@/assets/svg/add.svg'
import groupIcon from '@/assets/svg/group.svg'
import type { ProfileMessage } from '@/dto/profile'
import { getSocket, MyWebsocket, type SockedMessage } from '@/classes/websocket'
import { onBeforeUnmount, onMounted, ref } from 'vue'
import { NotiicationType } from '@/services/nofications'
</script>

<script setup lang="ts">
import Separator from '@/components/Separator.vue'
import Button from '@/components/Button.vue'

const notificationCount = ref(0)

defineProps({
  title: {
    type: String,
    required: true,
  },
})

let profileData: ProfileMessage
const profileJsonData = sessionStorage.getItem('profile')
if (profileJsonData) {
  profileData = JSON.parse(profileJsonData) as ProfileMessage
}

const socket = new MyWebsocket(
  'notifications',
  `ws://localhost:8000/notifications?authJWT=${sessionStorage.getItem('authJWT')}`,
)
socket.OnOpen(() => {
  console.log('Connected')
})
socket.OnClose(() => {
  console.log('WebSocket closed')
})
socket.OnMessage((event: MessageEvent) => {
  const data = JSON.parse(event.data)
  console.log('Message from server:', data)
  notificationCount.value++
})
socket.Watch()

onBeforeUnmount(() => {
  socket.Close()
})
</script>

<template>
  <div class="container">
    <div class="left-side-wrapper">
      <div class="left-side">
        <router-link class="logo" :to="`/`">ALLLOGS</router-link>
        <Separator />
        <div class="partition">
          <div class="partition-title">menu</div>
          <Button class="partition-button" :icon="addIcon" text="New project" link="/new-project" />
          <Button class="partition-button" :icon="addIcon" text="New group" link="/new-group" />
        </div>
        <Separator />
        <div class="partition">
          <div class="partition-title">projects</div>
          <Button
            class="partition-button"
            :icon="groupIcon"
            text="Group owner"
            link="/my-own-groups"
          />
        </div>
      </div>
    </div>
    <div class="right-side">
      <div class="header-wrapper">
        <div class="header">
          <div class="header-title">{{ title }}</div>
          <Separator :vertical="true" />
          <a class="header-notifications" href="">
            <img src="../assets/svg/bell-green.svg" alt="" class="hn-icon" />
            <div class="hn-count">{{ notificationCount }}</div>
          </a>
          <Separator :vertical="true" />
          <router-link class="header-profile" :to="`/profile/${profileData.Id}`">
            <div class="hp-avatar">
              <img :src="profileData.Avatar" alt="" />
            </div>
            <div class="hp-username">{{ profileData.User.Username }}</div>
          </router-link>
        </div>
      </div>
      <slot></slot>
    </div>
  </div>
  <img class="global-bg-image" src="@/assets/img/ALLLOGS_BG.jpg" />
</template>

<style scoped lang="scss">
@use '../assets/style/global_vars.scss' as vars;
@use '../assets/style/presets.scss' as ps;

.global-bg-image {
  position: absolute;
  z-index: -1;
  width: 100%;
  height: 100%;
  object-fit: cover;
  left: 0;
  top: 0;
  filter: brightness(30%);
}
.container {
  display: flex;
  height: 100vh;
  position: relative;
}
.left-side-wrapper {
  width: 400px;
  background: vars.$primary-color;
  box-sizing: border-box;
  padding-right: 8px;
  -webkit-box-shadow: inset 0px 0px 4px 2px vars.$focus-color;
  -moz-box-shadow: inset 0px 0px 4px 2px vars.$focus-color;
  box-shadow: inset 0px 0px 4px 2px vars.$focus-color;
  .left-side {
    width: 100%;
    height: 100%;
    background: vars.$primary-color;
    @include ps.shadow-panel;
    display: flex;
    flex-direction: column;
    .logo {
      background-color: transparent;
      margin: 10px auto;
      font-family: vars.$fnt-geo;
      font-size: 2.5rem;
      color: vars.$green-color;
      text-decoration: none;
      &:hover {
        cursor: pointer;
        text-shadow: 0px 0px 5px vars.$green-color;
      }
    }
    .partition {
      background-color: transparent;
      margin: 10px 20px;
      .partition-title {
        background-color: transparent;
        font-size: 1.2rem;
        color: vars.$secondary-font-color;
        margin-bottom: 5px;
      }
      .partition-button {
        :deep(.btn) {
          margin-bottom: 10px;
        }
      }
    }
  }
}
.right-side {
  flex: 1;
  position: relative;
  .header-wrapper {
    background: vars.$primary-color;
    height: 58px;
    box-sizing: border-box;
    padding-bottom: 8px;
    -webkit-box-shadow: inset 0px 0px 4px 2px vars.$focus-color;
    -moz-box-shadow: inset 0px 0px 4px 2px vars.$focus-color;
    box-shadow: inset 0px 0px 4px 2px vars.$focus-color;
    .header {
      display: flex;
      background: vars.$primary-color;
      background-color: vars.$secondary-color;
      width: 100%;
      height: 100%;
      @include ps.shadow-panel;
      .header-title {
        width: 100%;
        background-color: transparent;
        font-size: 1.25rem;
        margin: auto 0;
        margin-left: 20px;
      }
      .header-notifications {
        display: flex;
        padding: 0 20px;
        background-color: vars.$inner-button;
        text-decoration: none;
        @include ps.shadow-panel;
        &:hover {
          transition: 0.2s;
          filter: brightness(90%);
        }
        .hn-icon {
          background-color: transparent;
          width: 22px;
          height: 22px;
          margin: auto 0;
          margin-right: 5px;
        }
        .hn-count {
          background-color: transparent;
          margin: auto 0;
          font-size: 1.25rem;
        }
      }
      .header-profile {
        display: flex;
        padding: 0 10px;
        background-color: vars.$inner-button;
        text-decoration: none;
        @include ps.shadow-panel;
        &:hover {
          transition: 0.2s;
          filter: brightness(90%);
        }
        .hp-avatar {
          width: 35px;
          height: 35px;
          overflow: hidden;
          border-radius: 100%;
          margin: auto 0;
          margin-right: 5px;
          img {
            width: 100%;
            height: 100%;
            object-fit: cover;
          }
        }
        .hp-username {
          margin: auto 0;
          font-size: 1.25rem;
          background-color: transparent;
        }
      }
    }
  }
}
</style>
