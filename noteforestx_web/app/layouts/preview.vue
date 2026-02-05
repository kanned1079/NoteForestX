<script setup lang="ts">
import { useRoute } from 'vue-router'
import { useI18n } from "vue-i18n"
import AppHeader from "../components/AppHeader.vue"
import AppFooter from "~/components/AppFooter.vue"

const { locale } = useI18n()
const route = useRoute()
</script>

<!--<template>-->
<!--  <div class="bg-surface-ground">-->

<!--    &lt;!&ndash; Header 占位区（关键！） &ndash;&gt;-->
<!--    <div class="h-[60px]">-->
<!--      <div class="fixed left-0 right-0 top-0 z-50 h-[60px]">-->

<!--        <ClientOnly>-->
<!--          <AppHeader />-->
<!--        </ClientOnly>-->

<!--      </div>-->
<!--    </div>-->

<!--    &lt;!&ndash; 背景层（不影响滚动） &ndash;&gt;-->
<!--    <BackgroundGrid-->
<!--        class="fixed inset-0 -z-10"-->
<!--        :withFade="true"-->
<!--        :withGradient="true"-->
<!--    />-->

<!--    &lt;!&ndash; 页面主体 &ndash;&gt;-->

<!--    <div-->
<!--        class="-->
<!--    mx-auto-->
<!--    w-full-->
<!--    max-w-[1600px]-->
<!--    px-3-->
<!--    sm:px-6-->
<!--    md:px-8-->
<!--    lg:px-12-->
<!--    overflow-visible-->
<!--  "-->
<!--    >-->
<!--        <Toast />-->
<!--        <ConfirmDialog />-->
<!--        <slot />-->
<!--      </div>-->


<!--     <AppFooter />-->
<!--  </div>-->
<!--</template>-->

<template>
  <!-- 必须是 flex flex-col 且最小高度全屏 -->
  <div class="min-h-screen bg-surface-ground flex flex-col relative">

    <!-- Header 占位区 -->
    <div class="h-[60px] flex-none">
      <div class="fixed left-0 right-0 top-0 z-50 h-[60px] bg-inherit">
        <ClientOnly>
          <AppHeader />
        </ClientOnly>
      </div>
    </div>

    <!-- 背景层 (fixed 不占空间) -->
    <BackgroundGrid class="fixed inset-0 -z-10 max-h-[100vh]" :withFade="true" :withGradient="true" />

    <!-- 页面主体：使用 flex-1 撑开高度 -->
    <div class="mx-auto w-full max-w-[1600px] px-3 sm:px-6 md:px-8 lg:px-12 flex-1 flex flex-col">
      <Toast />
      <ConfirmDialog />
      <!-- 这里是关键：slot 渲染的内容 -->
      <slot />
    </div>

<!--    <div>-->
<!--      Comment-->
<!--    </div>-->

    <!-- Footer：确保它在文档流的最后 -->
<!--    <AppFooter class="flex-none mt-auto" />-->
  </div>
</template>