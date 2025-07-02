<script setup lang="ts">
import { useRouter } from 'vue-router'
import { AsyncRequest, catchClientError, catchServerError } from '@/classes/request'
import MiddlePanel from '@/views/MiddlePanel.vue'
import InputTemplate from '@/components/input/InputTemplate.vue'
import PasswordInp from '@/components/input/InputPassword.vue'
import Separator from '@/components/Separator.vue'
import Button from '@/components/Button.vue'
import Error from '@/components/Error.vue'
import { useErrorStore } from '@/stores/error'
import { ref } from 'vue'
import type { MsgRegister } from '@/dto/auth'
import { isMsgClientError, isMsgServerError, type MsgBaseResponse, type MsgClientError, type MsgServerError } from '@/dto/common'
import { AxiosError, type AxiosResponse } from 'axios'

const errorStore = useErrorStore()
const router = useRouter()

const formData = ref<MsgRegister>({
  Username: '',
  Password: '',
  RepeatPassword: '',
})

const submitForm = async () => {
  const req = new AsyncRequest('http://localhost:8000/register', {
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
  })
  req.onResponse((response: AxiosResponse) => {
    const baseResponse = response.data as MsgBaseResponse
    if (!baseResponse.Ok) {
      errorStore.setText(baseResponse.Error)
    } else {
      router.push('/login')
    }
  })
  req.onError((error: AxiosError) => {
    if(error.response?.data, isMsgClientError(error.response?.data)){
      catchClientError(error.response?.data as MsgClientError, errorStore)
    } else if (error.response?.data, isMsgServerError(error.response?.data)){
      catchServerError(error.response?.data as MsgServerError, errorStore)
    } else {
      errorStore.setText("unexpected error: " + error.message)
    }
  })
  req.setData(formData.value)
  req.post()
}
</script>

<template>
  <MiddlePanel :single="true">
    <Error />
    <div class="title">Register</div>
    <form @submit.prevent="submitForm">
      <InputTemplate text="Username" class="inp">
        <input v-model="formData.Username" type="text" name="Username" required />
      </InputTemplate>
      <PasswordInp v-model="formData.Password" text="Password" name="Password" />
      <PasswordInp
        v-model="formData.RepeatPassword"
        text="Repeat password"
        name="repeat_password"
      />
      <Separator />
      <router-link class="link" to="/login">Log in</router-link>
      <Button type="submit" class="button" text="Register" icon="checkbox" />
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
