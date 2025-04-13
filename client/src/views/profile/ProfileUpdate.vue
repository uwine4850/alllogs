<script lang="ts">
import deleteIcon from '@/assets/svg/delete.svg'
import checkboxIcon from '@/assets/svg/checkbox.svg'
import profileIcon from '@/assets/svg/user.svg'
</script>

<script setup lang="ts">
import MiddlePanel from '@/views/MiddlePanel.vue';
import BaseTemplate from '@/views/BaseTemplate.vue';
import PanelTitle from '@/components/PanelTitle.vue';
import InputTextarea from '@/components/input/InputTextarea.vue';
import InputFile from '@/components/input/InputFile.vue';
import InputCheckbox from '@/components/input/InputCheckbox.vue';
import Button from '@/components/Button.vue';
import Separator from '@/components/Separator.vue';
</script>

<template>
	<BaseTemplate title="Profile update">
		<MiddlePanel>
			<PanelTitle :icon="profileIcon" text="Profile update" :sep="false" />
			<InputTextarea text="Description" name="description" />
			<InputFile text="Avatar"/>
			<InputCheckbox text="Delete avatar" inptext="delete"/>
			<Separator />
			<div class="update-btns">
				<Button class="btn btn-delete" :icon="deleteIcon" text="Delete user" />
				<Button class="btn" :icon="checkboxIcon" text="Save" />
			</div>
		</MiddlePanel>
	</BaseTemplate>
</template>

<style scoped lang="scss">
	@use '@/assets/style/global_vars.scss' as vars;

	.update-btns{
		margin: 10px;
		display: flex;
		gap: 10px;
		justify-content: space-between;
		background-color: transparent;
		.btn{
			width: 100%;
		}
		.btn-delete{
			:deep(.btn){
				background-color: vars.$color-red;
				&:hover {
					cursor: pointer;
					transition: 0.2s;
					filter: brightness(90%);
				}
			}
		}
	}
</style>