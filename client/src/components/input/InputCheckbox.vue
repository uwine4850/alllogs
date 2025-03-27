<script setup lang="ts">
import InputTemplate from './InputTemplate.vue';

const props = defineProps({
    text:{
        type: String,
        required: true,
    },
    inptext:{
        type: String,
    }
});
</script>

<template>
    <InputTemplate :text="props.text">
        <span>
            <label for="ch">{{ props.inptext }}</label>
            <input id="ch" type="checkbox" class="inp">
        </span>
    </InputTemplate>
</template>

<style scoped lang="scss">
    @use '@/assets/style/global_vars.scss' as vars;
    @use '@/assets/style/presets.scss' as ps;
    span{
        background-color: vars.$input-color;
        display: flex;
        width: 100%;
        @include ps.inner-shadow-panel;
        label{
            user-select: none;
            -webkit-user-select: none;
            -moz-user-select: none;
            -ms-user-select: none; 
            margin: auto 0;
            margin-left: 10px;
            font-size: 1.1rem;
            font-family: vars.$fnt-hint-madurai;
            background-color: transparent;
        }
        .inp{
            margin: auto 0;
            margin-left: 10px;
            background-color: transparent;
            font-size: 1rem;
            font-family: vars.$fnt-hint-madurai;
        }
    }
</style>