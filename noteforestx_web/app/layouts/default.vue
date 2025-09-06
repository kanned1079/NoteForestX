<script setup lang="ts">
import AppHeader from "../components/AppHeader.vue";

onMounted(() => {
  console.log("aaa")
})
</script>

<template>
  <AppHeader />
  <div class="mx-auto w-full max-w-[1680px] px-4 sm:px-6 md:px-8 lg:px-12">
    <slot />
  </div>
</template>

<style scoped>

</style>