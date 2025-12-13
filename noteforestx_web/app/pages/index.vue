<script setup lang="ts">
import MyIntroduction from "~/components/MyIntroduction.vue";

definePageMeta({
  layout: 'main',
  layoutTransition: true,
})
</script>

<template>
<div class="">
<!--  <ClientOnly>-->
    <FirstPage />
  <MyIntroduction />
    <AboutMe />

<!--  </ClientOnly>-->


</div>
</template>

<style scoped>

</style>