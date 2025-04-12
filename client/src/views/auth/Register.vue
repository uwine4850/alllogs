<script lang="ts">
import checkboxIcon from '@/assets/svg/checkbox.svg'
import { useRouter } from 'vue-router'
</script>

<script setup lang="ts">
import MiddlePanel from '@/views/MiddlePanel.vue'
import InputTemplate from '@/components/input/InputTemplate.vue'
import PasswordInp from '@/components/input/InputPassword.vue'
import Separator from '@/components/Separator.vue'
import Button from '@/components/Button.vue'
import Error from '@/components/Error.vue'
import { useErrorStore } from '@/stores/error'
import { ref } from 'vue'
import type { Register } from '@/dto/auth'
import type { BaseResponse } from '@/dto/common'
import axios from 'axios';

const errorStore = useErrorStore();
const router = useRouter();

const formData = ref<Register>({
  Username: "",
  Password: "",
  RepeatPassword: "",
})

const submitForm = async () => {
  try{
    const response = await axios.post("http://localhost:8000/register", formData.value, {
      headers: {
        "Content-Type": "application/json",
      },
    });
    const baseResponse = response.data as BaseResponse;
    if (!baseResponse.Ok){
      errorStore.setText(baseResponse.Error);
    } else {
      router.push("/login")
    }
  } catch (error){
    errorStore.setText(String(error));
  }
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
      <PasswordInp v-model="formData.RepeatPassword" text="Repeat password" name="repeat_password" />
      <Separator />
      <router-link class="link" to="/login">Log in</router-link>
      <Button type="submit" class="button" text="Register" :icon="checkboxIcon" />
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
