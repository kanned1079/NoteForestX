<script setup lang="ts">
import { useRoute } from 'vue-router'
import { useI18n } from "vue-i18n"
import AppHeader from "../components/AppHeader.vue"
import AppFooter from "~/components/AppFooter.vue"

const { locale } = useI18n()
const route = useRoute()
</script>

<template>
  <div class="min-h-screen bg-surface-ground">

    <!-- Header 占位区（关键！） -->
    <div class="h-[60px]">
      <div class="fixed left-0 right-0 top-0 z-50 h-[60px]">
        <AppHeader />
      </div>
    </div>

    <!-- 背景层（不影响滚动） -->
    <BackgroundGrid
        class="fixed inset-0 -z-10"
        :withFade="true"
        :withGradient="true"
    />

    <!-- 页面主体 -->
    <div
        class="
        mx-auto
        w-full
        max-w-[1600px]
        px-4 sm:px-6 md:px-8 lg:px-12
        overflow-visible
      "
    >
      <Toast />
      <ConfirmDialog />
      <slot />
    </div>

    <!-- <AppFooter /> -->
  </div>
</template>