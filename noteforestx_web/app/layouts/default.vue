<script setup lang="ts">
import AppHeader from "../components/AppHeader.vue";

onMounted(() => {
  console.log("aaa")
})
</script>

<template>
  <div>
    <AppHeader />

    <p>Some default layout content shared across all pages</p>
    <slot />
  </div>
</template>

<style scoped>

</style>