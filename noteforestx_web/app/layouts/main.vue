<script setup lang="ts">
import { ref, onMounted } from "vue"
// import AppHeader from "../components/AppHeader.vue"
// import WelcomeBanner from "../components/WelcomeBanner.vue"
// import CollapseTransition from "../components/CollapseTransition.vue" // ✅ 新增导入

const showBanner = ref<boolean>(false)

const expandUpper = () => {
  showBanner.value = !showBanner.value
}
</script>

<template>
  <div class="relative">
    <!-- 使用可复用折叠动画组件 -->
<!--    <ClientOnly>-->

<!--      <CollapseTransition v-model:show="showBanner" :duration="400">-->
<!--&lt;!&ndash;        <WelcomeBanner @close="closeBanner" />&ndash;&gt;-->
<!--        <div class="w-full h-[3vh] bg-cyan-500"></div>-->
<!--      </CollapseTransition>-->
<!--    </ClientOnly>-->


<!--    <CollapseTransition v-model:show="showBanner" :duration="0.5" appear>-->
<!--      <div class="w-full h-[3vh] bg-cyan-500 flex items-center justify-center text-white">-->
<!--        Welcome to Kaneko’s Blog 🚀-->
<!--      </div>-->
<!--    </CollapseTransition>-->

    <WelcomeBanner :expand="expandUpper" />


    <!-- 内容随 banner 展开平滑下移 -->
    <div
        class="transition-all duration-500 ease-in-out"
        :style="{ transform: showBanner ? 'translateY(2.5vh)' : 'translateY(0)' }"
    >
      <AppHeader />
      <Toast />
      <ConfirmDialog />
      <slot />
    </div>
  </div>
</template>

<style scoped>
/* 无需额外样式，过渡动画已在 CollapseTransition 内控制 */
</style>