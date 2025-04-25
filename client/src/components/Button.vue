<script setup lang="ts">
defineProps({
  icon: {
    type: String,
    required: true,
  },
  text: {
    type: String,
    required: true,
  },
  link: {
    type: String,
  },
})
defineEmits(['click'])
</script>

<template>
  <div v-if="link">
    <router-link class="wrapper" :to="link">
      <button class="btn" v-bind="$attrs" @click="$emit('click')">
        <div class="content">
          <img class="icon" :src="icon" alt="icon" />
          <p class="text">{{ text }}</p>
        </div>
      </button>
    </router-link>
  </div>

  <div v-else>
    <span class="wrapper">
      <button class="btn" v-bind="$attrs" @click="$emit('click')">
        <div class="content">
          <img class="icon" :src="icon" alt="icon" />
          <p class="text">{{ text }}</p>
        </div>
      </button>
    </span>
  </div>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;

.wrapper {
  text-decoration: none;
  border-radius: 5px;
  -webkit-box-shadow: 0px 0px 4px 0px vars.$bg-color;
  -moz-box-shadow: 0px 0px 4px 0px vars.$bg-color;
  box-shadow: 0px 0px 4px 0px vars.$bg-color;
  width: 100%;
  .btn {
    background-color: vars.$secondary-color;
    border-radius: 5px;
    width: 100%;
    padding: 15px 0;
    display: flex;
    border: none;
    outline: none;
    &:hover {
      cursor: pointer;
      transition: 0.2s;
      background-color: vars.$focus-color;
    }
    .content {
      background-color: transparent;
      margin: auto 0;
      margin-left: 10px;
      display: flex;
      .icon {
        width: 20px;
        height: 20px;
        margin: auto 0;
        background-color: transparent;
        margin-right: 10px;
      }
      .text {
        background-color: transparent;
        font-size: 1.25rem;
      }
    }
  }
}
</style>
