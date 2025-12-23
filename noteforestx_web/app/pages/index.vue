<script setup lang="ts">
import MyIntroduction from "~/components/MyIntroduction.vue";
import ContactMe from "~/components/MyIntro/ContactMe.vue";
import useThemeStore from "~/store/themeStore";

const tabValue = ref<string>('0');
const themeStore = useThemeStore();

definePageMeta({
  layout: 'main',
  layoutTransition: true,
})

onBeforeMount(() => {
  themeStore.setShowLeftIntroBtn(true)
})

onBeforeUnmount(() => {
  themeStore.setShowLeftIntroBtn(false)
})
</script>

<template>
<div class="">
<!--  <ClientOnly>-->
  <FirstPage />
  <MyDesign />
  <MyIntroduction />

<!--  <AboutMe />-->


</div>
</template>

<style scoped>

</style>