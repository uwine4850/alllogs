<script setup lang="ts">
    const props = defineProps({
        title: {
            type: String,
            required: true,
        },
        descr: {
            type: String,
            required: true,
        },
        link: {
            type: String,
            default: "#",
        }
    });
</script>

<template>
  <router-link class="panel-content" :to="link">
    <div class="pc-name">{{ props.title }}</div>
    <div class="pc-description">
      {{ props.descr }}
    </div>
  </router-link>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;

.panel-content {
        text-decoration: none;
        box-sizing: border-box;
        padding: 10px;
        display: flex;
        flex-direction: column;
        margin: 10px 10px 0 10px;
        border-radius: 5px;
        background-color: vars.$secondary-color;
        &:hover {
          transition: 0.2s;
          background-color: vars.$focus-color;
        }
        .pc-name {
          background-color: transparent;
          font-size: 1.25rem;
          margin-bottom: 5px;
        }
        .pc-description {
          background-color: transparent;
          font-size: 1.1rem;
          font-family: vars.$fnt-hint-madurai;
        }
      }
</style>