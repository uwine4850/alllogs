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
import axios, { AxiosError, type AxiosResponse } from 'axios'
import Error from '@/components/Error.vue'
import { AsyncRequest, catchClientError, catchServerError } from '@/classes/request'
import type { ProfileMessage } from '@/dto/profile'
import { isClientErrorMessage, isServerErrorMessage, type ClientErrorMessage, type ServerErrorMessage } from '@/dto/common'
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
      'Content-Type': 'application/x-www-form-urlencoded',
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
      req.onError((error: AxiosError) => {
        if(error.response?.data, isClientErrorMessage(error.response?.data)){
          catchClientError(error.response?.data as ClientErrorMessage, errorStore)
        } else if (error.response?.data, isServerErrorMessage(error.response?.data)){
          catchServerError(error.response?.data as ServerErrorMessage, errorStore)
        } else {
          errorStore.setText("unexpected error: " + error.message)
        }
      })
      req.get()
    }
  })
  loginReq.onError((error: AxiosError) => {
    if(error.response?.data, isClientErrorMessage(error.response?.data)){
      catchClientError(error.response?.data as ClientErrorMessage, errorStore)
    } else if (error.response?.data, isServerErrorMessage(error.response?.data)){
      catchServerError(error.response?.data as ServerErrorMessage, errorStore)
    } else {
      errorStore.setText("unexpected error: " + error.message)
    }
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
