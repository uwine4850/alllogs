<script setup lang="ts">
import ProjectTemplate from './project/ProjectTemplate.vue'
import Separator from '@/components/Separator.vue'
import PanelTitle from '@/components/PanelTitle.vue'
import Button from '@/components/Button.vue'
</script>

<template>
  <ProjectTemplate title="Group">
    <template #panel-project>
      <div class="group-view">
        <div class="group-name">group name</div>
        <div class="group-description">
          It is a long established fact that a reader will be distracted by the readable content of
          a page when looking at its layout. The point of using Lorem Ipsum is that it has
        </div>
      </div>
      <Separator />
      <div class="info-line">
        <form class="search-form">
          <input type="text" class="search-input" />
          <button class="search-btn">Search</button>
        </form>
      </div>
      <Separator />
      <div class="group-user-list">
        <div class="group-user">
          <div class="gu-username">renxob</div>
          <button class="gu-remove-btn">remove</button>
        </div>
        <div class="group-user">
          <div class="gu-username">renxob</div>
          <button class="gu-remove-btn">remove</button>
        </div>
        <div class="group-user">
          <div class="gu-username">renxob</div>
          <button class="gu-remove-btn">remove</button>
        </div>
        <div class="group-user">
          <div class="gu-username">renxob</div>
          <button class="gu-remove-btn">remove</button>
        </div>
        <div class="group-user">
          <div class="gu-username">renxob</div>
          <button class="gu-remove-btn">remove</button>
        </div>
        <div class="group-user">
          <div class="gu-username">renxob</div>
          <button class="gu-remove-btn">remove</button>
        </div>
        <div class="group-user">
          <div class="gu-username">renxob</div>
          <button class="gu-remove-btn">remove</button>
        </div>
        <div class="group-user">
          <div class="gu-username">renxob</div>
          <button class="gu-remove-btn">remove</button>
        </div>
        <div class="group-user">
          <div class="gu-username">renxob</div>
          <button class="gu-remove-btn">remove</button>
        </div>
        <div class="group-user">
          <div class="gu-username">renxob</div>
          <button class="gu-remove-btn">remove</button>
        </div>
        <div class="group-user">
          <div class="gu-username">renxob</div>
          <button class="gu-remove-btn">remove</button>
        </div>
        <div class="group-user">
          <div class="gu-username">renxob</div>
          <button class="gu-remove-btn">remove</button>
        </div>
      </div>
    </template>
    <template #panel-menu>
      <PanelTitle icon="group" text="Group management" />
      <div class="group-management-btns">
        <Button icon="update" text="Update" />
        <Button icon="add-user" text="Add user" />
      </div>
    </template>
  </ProjectTemplate>
</template>

<style scoped lang="scss">
@use '../assets/style/global_vars.scss' as vars;
@use '../assets/style/presets.scss' as ps;

.group-view {
  display: flex;
  flex-direction: column;
  padding: 10px;
  background-color: transparent;
  .group-name {
    background-color: transparent;
    font-size: 1.25rem;
    margin-bottom: 5px;
  }
  .group-description {
    background-color: transparent;
    font-family: vars.$fnt-hint-madurai;
    font-size: 1.1rem;
  }
}
.info-line {
  min-height: 45px;
  display: flex;
  .search-form {
    display: flex;
    width: 100%;
    .search-input {
      width: 100%;
      border: none;
      outline: none;
      padding: 0 10px;
      font-size: 1.1rem;
      background-color: vars.$input-color;
      @include ps.shadow-panel;
    }
    .search-btn {
      width: 100px;
      border: none;
      outline: none;
      background-color: vars.$inner-button;
      @include ps.inner-shadow-panel;
      font-size: 1.1rem;
      &:hover {
        transition: 0.2s;
        cursor: pointer;
        filter: brightness(90%);
      }
    }
  }
}
.group-user-list {
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  padding: 10px;
  background-color: transparent;
  overflow-y: scroll;
  .group-user {
    display: flex;
    justify-content: space-between;
    margin-bottom: 10px;
    box-sizing: border-box;
    padding: 10px;
    border-radius: 5px;
    background-color: vars.$secondary-color;
    &:last-child {
      margin-bottom: 0;
    }
    .gu-username {
      font-size: 1.25rem;
      background-color: transparent;
      margin: auto 0;
    }
    .gu-remove-btn {
      padding: 5px;
      background-color: vars.$color-red;
      border: none;
      outline: none;
      border-radius: 3px;
      font-size: 1.1rem;
      &:hover {
        cursor: pointer;
        transition: 0.2s;
        filter: brightness(90%);
      }
    }
  }
}
.group-management-btns {
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  padding: 10px;
  background-color: transparent;
  gap: 10px;
}
</style>
