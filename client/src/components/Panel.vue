<template>
  <div class="panel-border">
    <div class="panel">
      <slot></slot>
    </div>
  </div>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;
@use '@/assets/style/presets.scss' as ps;

.panel-border {
  background-color: vars.$primary-color;
  box-sizing: border-box;
  padding: 8px;
  border-radius: 5px;
  -webkit-box-shadow: inset 0px 0px 4px 2px vars.$focus-color;
  -moz-box-shadow: inset 0px 0px 4px 2px vars.$focus-color;
  box-shadow: inset 0px 0px 4px 2px vars.$focus-color;
  .panel {
    display: flex;
    flex-direction: column;
    box-sizing: border-box;
    background-color: vars.$primary-color;
    border-radius: 5px;
    width: 100%;
    height: 100%;
    @include ps.shadow-panel;
  }
}
</style>
