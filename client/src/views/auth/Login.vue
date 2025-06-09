<script lang="ts">
import checkboxIcon from '@/assets/svg/checkbox.svg'
import { useErrorStore } from '@/stores/error'
import { useRouter } from 'vue-router'
import type { LoginMessage, LoginResponseMessage } from '@/dto/auth'
import MiddlePanel from '@/views/MiddlePanel.vue'
import InputTemplate from '@/components/input/InputTemplate.vue'
import PasswordInp from '@/components/input/InputPassword.vue'
import Separator from '@/components/Separator.vue'
import Button from '@/components/Button.vue'
import { ref } from 'vue'
import axios, { type AxiosResponse } from 'axios'
import Error from '@/components/Error.vue'
import { AsyncRequest } from '@/classes/request'
import type { ProfileMessage } from '@/dto/profile'
</script>

<script setup lang="ts">
const errorStore = useErrorStore()
const router = useRouter()

const formData = ref<LoginMessage>({
  Username: '',
  Password: '',
})

const submitForm = async () => {
  const loginReq = new AsyncRequest('http://localhost:8000/login', {
    headers: {
      'Content-Type': 'application/json',
    },
    withCredentials: true,
  })
  loginReq.onResponse((response: AxiosResponse) => {
    const loginResponse = response.data as LoginResponseMessage
    if (loginResponse.Error != '') {
      errorStore.setText(loginResponse.Error)
    } else {
      sessionStorage.setItem('authJWT', loginResponse.JWT)
      // Get profile data.
      const req = new AsyncRequest('http://localhost:8000/profile/' + loginResponse.UID, {
        headers: {
          'Content-Type': 'application/json',
        },
        withCredentials: true,
      })
      req.onResponse(function (response: AxiosResponse) {
        const profileResponse = response.data as ProfileMessage
        if (profileResponse.Error != '') {
          console.log(profileResponse.Error)
          errorStore.setText(profileResponse.Error)
          sessionStorage.remove('authJWT');
        } else {
          sessionStorage.setItem('profile', JSON.stringify(profileResponse))
          router.push('/')
        }
      })
      req.onError((error: unknown) => {
        errorStore.setText(loginResponse.Error)
      })
      req.get()
    }
  })
  loginReq.onError((error: unknown) => {
    errorStore.setText(String(error))
  })
  loginReq.setData(formData.value)
  loginReq.post()
}
</script>

<template>
  <MiddlePanel :single="true">
    <Error />
    <div class="title">Log in</div>
    <form @submit.prevent="submitForm">
      <InputTemplate text="Username" class="inp">
        <input v-model="formData.Username" type="text" name="username" />
      </InputTemplate>
      <PasswordInp v-model="formData.Password" text="Password" name="password" />
      <Separator />
      <router-link class="link" to="/register">Register</router-link>
      <Button type="submit" class="button" text="Log in" :icon="checkboxIcon" />
    </form>
  </MiddlePanel>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;
@use '@/assets/style/presets.scss' as ps;

.title {
  font-size: 1.25rem;
  background-color: transparent;
  padding: 10px;
}

form {
  display: flex;
  flex-direction: column;
  background-color: transparent;
}

.inp {
  input {
    border: none;
    outline: none;
    font-size: 1.1rem;
    background-color: vars.$input-color;
    width: 100%;
    height: 45px;
    padding: 0 10px;
    @include ps.inner-shadow-panel;
  }
}
.link {
  background-color: transparent;
  font-size: 1.25rem;
  margin: 10px;
  text-decoration: none;
  color: vars.$secondary-font-color;
  &:hover {
    transition: 0.2s;
    filter: brightness(130%);
  }
}
.button {
  margin: 10px;
  width: 50%;
  margin-left: auto;
}
</style>
