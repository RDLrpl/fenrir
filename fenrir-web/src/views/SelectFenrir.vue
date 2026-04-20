<style>
    @import "@/res/css/Select.css";
</style>

<script lang="ts" src="../res/scripts/select.ts"></script>

<template>
    <div ref="routeElement" class="select-route">
        <button class="settings-button" @click="toSettings">⚙️</button>
        <div class="select-menu">
            <span class="menu-title">Fenrir@Utility</span>
            <button class="menu-button" @click="toTG">Telegram</button>
        </div>
  </div>
</template>

