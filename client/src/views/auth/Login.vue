<script lang="ts">
import checkboxIcon from '@/assets/svg/checkbox.svg'
import { useErrorStore } from '@/stores/error'
import { useRouter } from 'vue-router'
import type { Login } from '@/dto/auth'
import MiddlePanel from '@/views/MiddlePanel.vue'
import InputTemplate from '@/components/input/InputTemplate.vue'
import PasswordInp from '@/components/input/InputPassword.vue'
import Separator from '@/components/Separator.vue'
import Button from '@/components/Button.vue'
import { ref } from 'vue'
import type { LoginResponse } from '@/dto/auth'
import axios from 'axios'
import Error from '@/components/Error.vue'
</script>

<script setup lang="ts">
const errorStore = useErrorStore();
const router = useRouter();

const formData = ref<Login>({
  Username: "",
  Password: "",
});

const submitForm = async () => {
  try{
    const response = await axios.post("http://localhost:8000/login", formData.value, {
      headers: {
        "Content-Type": "application/json",
      },
      withCredentials: true,
    });
    const loginResponse = response.data as LoginResponse;
    if (loginResponse.Error != ""){
      errorStore.setText(loginResponse.Error);
    } else {
      sessionStorage.setItem("authJWT", loginResponse.JWT)
      router.push("/")
    }
  } catch (error){
    errorStore.setText(String(error));
  }
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

form{
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
