<script setup lang="ts">
import { ref, onBeforeMount, onMounted } from "vue"
import AppHeader from "../components/AppHeader.vue"
import { useRoute } from 'vue-router';
import { useI18n } from "vue-i18n"
const { locale } = useI18n()
const route = useRoute();
const currentPath = route.path;

// const showBanner = ref(false)
//
// onBeforeMount(() => {
//   const closed = localStorage.getItem("bannerClosed")
//   if (closed === "true") {
//     showBanner.value = false
//   }
// })
//
// onMounted(() => {
//   console.log(currentPath, locale.value)
//   if (!localStorage.getItem("bannerClosed") && currentPath === `/${locale.value}`) {
//     setTimeout(() => {
//       showBanner.value = true
//     }, 1000)
//   }
// })
//
// const closeBanner = () => {
//   showBanner.value = false
//   localStorage.setItem("bannerClosed", "true")
// }
</script>

<template>
  <div>
    <div
        class="fixed left-0 right-0 z-50 transition-all duration-500 ease-in-out"

    >
      <AppHeader />
    </div>
    <div class="mx-auto w-full max-w-[1680px] px-4 sm:px-6 md:px-8 lg:px-12">
      <Toast />
      <ConfirmDialog />
      <slot />
    </div>
  </div>
</template>

<style scoped>
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.5s ease;
}

.fade-slide-enter-from {
  opacity: 0;
  transform: translateY(-20px);
}
.fade-slide-enter-to {
  opacity: 1;
  transform: translateY(0);
}

.fade-slide-leave-from {
  opacity: 1;
  transform: translateY(0);
}
.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(-20px);
}
</style>