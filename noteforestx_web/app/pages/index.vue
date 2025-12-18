<script setup lang="ts">
import MyIntroduction from "~/components/MyIntroduction.vue";
import ContactMe from "~/components/MyIntro/ContactMe.vue";

const tabValue = ref<string>('0');

definePageMeta({
  layout: 'main',
  layoutTransition: true,
})
</script>

<template>
<div class="">
<!--  <ClientOnly>-->
  <FirstPage />
  <MyDesign />
  <MyIntroduction />

  <AboutMe />


</div>
</template>

<style scoped>

</style>