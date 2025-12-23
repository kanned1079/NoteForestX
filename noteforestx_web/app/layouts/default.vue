<script setup lang="ts">
import { useRoute } from 'vue-router'
import { useI18n } from "vue-i18n"
import AppHeader from "../components/AppHeader.vue"
import AppFooter from "~/components/AppFooter.vue"

const { locale } = useI18n()
const route = useRoute()
</script>

<template>
  <div class="min-h-screen bg-surface-ground relative">

    <!-- Header（固定） -->
    <div class="fixed left-0 right-0 z-50">
      <AppHeader />
    </div>

    <!-- 🌈 背景层 -->
    <BackgroundGrid
        class="absolute inset-0 z-0"
        :withFade="true"
        :withGradient="true"
    />

    <!-- 页面主体（相对定位，隔离背景） -->
    <div
        class="
        relative
        mx-auto
        w-full
        max-w-[1600px]
        px-4 sm:px-6 md:px-8 lg:px-12
      "
    >


      <!-- 🧱 内容层（关键） -->
      <div class="relative z-10 pt-[60px]">
        <Toast />
        <ConfirmDialog />
        <slot />
      </div>
    </div>

    <!-- <AppFooter /> -->
  </div>
</template>