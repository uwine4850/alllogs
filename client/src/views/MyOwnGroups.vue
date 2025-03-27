<script lang="ts">
import groupIcon from '@/assets/svg/group.svg'
</script>

<script setup lang="ts">
import BaseTemplate from './BaseTemplate.vue';
import MiddlePanel from './MiddlePanel.vue';
import PanelTitle from '@/components/PanelTitle.vue';
import PanelItem from '@/components/PanelItem.vue';
</script>

<template>
    <BaseTemplate title="My own groups">
        <MiddlePanel class="groups">
            <PanelTitle :icon="groupIcon" text="my own groups"/>
            <div class="group-list">
                <PanelItem class="group-item" title="Project name" descr="DESCR" link="/group" />
                <PanelItem class="group-item" title="Project name" descr="DESCR" link="/group" />
                <PanelItem class="group-item" title="Project name" descr="DESCR" link="/group" />
                <PanelItem class="group-item" title="Project name" descr="DESCR" link="/group" />
                <PanelItem class="group-item" title="Project name" descr="DESCR" link="/group" />
                <PanelItem class="group-item" title="Project name" descr="DESCR" link="/group" />
                <PanelItem class="group-item" title="Project name" descr="DESCR" link="/group" />
                <PanelItem class="group-item" title="Project name" descr="DESCR" link="/group" />
                <PanelItem class="group-item" title="Project name" descr="DESCR" link="/group" />
                <PanelItem class="group-item" title="Project name" descr="DESCR" link="/group" />
                <PanelItem class="group-item" title="Project name" descr="DESCR" link="/group" />
                <PanelItem class="group-item" title="Project name" descr="DESCR" link="/group" />
                <PanelItem class="group-item" title="Project name" descr="DESCR" link="/group" />
                <PanelItem class="group-item" title="Project name" descr="DESCR" link="/group" />
                <PanelItem class="group-item" title="Project name" descr="DESCR" link="/group" />
            </div>
        </MiddlePanel>
    </BaseTemplate>
</template>

<style scoped lang="scss">
.groups{
    :deep(.middle-panel){
        height: 100%;
    }
    .group-list{
        background-color: transparent;
        height: 100%;
        overflow-y: scroll;
        .group-item{
            &:last-child{
                margin-bottom: 10px;
            }
        }
    }
}
</style>