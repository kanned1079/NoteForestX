<script setup lang="ts">
import { ref, onMounted } from "vue"
import AppFooter from "~/components/AppFooter.vue";
// import AppHeader from "../components/AppHeader.vue"
// import WelcomeBanner from "../components/WelcomeBanner.vue"
// import CollapseTransition from "../components/CollapseTransition.vue" // ✅ 新增导入

const showBanner = ref<boolean>(false)

const bannerHeight = 32   // WelcomeBanner 展开高度
const headerHeight = 0   // AppHeader 高度

const expandUpper = () => {
  showBanner.value = !showBanner.value
  // headerHeight.value = headerRef.value?.offsetHeight
  // bannerHeight.value = bannerRef.value?.offsetHeight
}
</script>

<template>
  <div class="relative">
    <!-- Banner -->
    <WelcomeBanner :expand="expandUpper" />

    <!-- 固定 Header，但根据 Banner 高度下移 -->
    <div
        class="fixed left-0 right-0 z-50 transition-all duration-500 ease-in-out"
        :style="{ top: showBanner ? bannerHeight + 'px' : '0px' }"
    >
      <AppHeader />
    </div>

    <!-- 页面内容，同样下移，避免被 Header 覆盖 -->
    <div
        class="transition-all duration-500 ease-in-out"
        :style="{
        paddingTop: showBanner
          ? bannerHeight + headerHeight + 'px'
          : headerHeight + 'px'
      }"
    >
      <Toast />
      <ConfirmDialog />
      <slot />
    </div>

    <AppFooter />

  </div>


</template>

<!--<template>-->
<!--  <div class="relative">-->

<!--    <WelcomeBanner :expand="expandUpper" />-->

<!--    <div-->
<!--        class="transition-all duration-500 ease-in-out"-->
<!--        :style="{ transform: showBanner ? 'translateY(2.5vh)' : 'translateY(0)' }"-->
<!--    >-->
<!--      <AppHeader />-->
<!--      <Toast />-->
<!--      <ConfirmDialog />-->
<!--      <slot />-->
<!--    </div>-->
<!--  </div>-->
<!--</template>-->

<style scoped>
/* 无需额外样式，过渡动画已在 CollapseTransition 内控制 */
</style>