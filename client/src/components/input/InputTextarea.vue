<script setup lang="ts">
import InputTemplate from '@/components/input/InputTemplate.vue';

const props = defineProps({
    text: {
        type: String,
        required: true,
    },
    name: {
        type: String,
        required: true,
    },
});
</script>

<template>
<InputTemplate :text="props.text">
    <textarea :name="props.name" class="inp"></textarea>
</InputTemplate>
</template>

<style scoped lang="scss">
@use '@/assets/style/global_vars.scss' as vars;
@use '@/assets/style/presets.scss' as ps;

.inp{
    width: 100%;
    border: 0;
    outline: none;
    background-color: vars.$input-color;
    @include ps.inner-shadow-panel;
    font-size: 1.1rem;
    padding: 10px;
    height: 100px;
    resize: none;
}
</style>