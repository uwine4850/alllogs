<script setup lang="ts">
	import { onMounted, ref } from 'vue';
	import { useErrorStore } from '@/stores/error';
	import Separator from './Separator.vue';
	
	const errorStore = useErrorStore();
	const errorEl = ref<HTMLElement>();

	onMounted(() => {
		if(errorEl.value){
			errorStore.setErrorElement(errorEl.value);
		}
	});
</script>

<template>
	<div id="wrapper" ref="errorEl">
		<span id="text">ewfew</span>
		<Separator class="sep" />
	</div>
</template>

<style scoped lang="scss">
	@use '../assets/style/global_vars.scss' as vars;
	@use '../assets/style/presets.scss' as ps;

	#wrapper{
		background-color: vars.$secondary-color;
		width: 100%;
		display: none;
		flex-direction: column;
		#text{
			font-size: 1.1rem;
			padding: 10px;
			background-color: transparent;
			box-sizing: border-box;
			color: red;
			@include ps.inner-shadow-panel;
		}
	}
</style>