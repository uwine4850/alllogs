<script lang="ts">
import checkboxIcon from '@/assets/svg/checkbox.svg'
</script>

<script setup lang="ts">
import MiddlePanel from './MiddlePanel.vue'
import InputTemplate from '@/components/input/InputTemplate.vue'
import PasswordInp from '@/components/input/InputPassword.vue'
import Separator from '@/components/Separator.vue'
import Button from '@/components/Button.vue'
</script>

<template>
  <MiddlePanel :single="true">
    <div class="title">Log in</div>
    <InputTemplate text="Username" class="inp">
      <input type="text" name="username" />
    </InputTemplate>
    <PasswordInp text="Password" name="password" />
    <Separator />
    <router-link class="link" to="/register">Register</router-link>
    <Button class="button" text="Log in" :icon="checkboxIcon" />
  </MiddlePanel>
</template>

<style scoped lang="scss">
@use '../assets/style/global_vars.scss' as vars;
@use '../assets/style/presets.scss' as ps;

.title {
  font-size: 1.25rem;
  background-color: transparent;
  padding: 10px;
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
