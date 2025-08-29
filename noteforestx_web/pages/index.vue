<script setup lang="ts">
import {definePageMeta} from "../.nuxt/imports";

definePageMeta({
  layout: "default"
})

onMounted(() => {
  console.log("index mounted")
})
</script>

<template>
  <div>
    INDEX
  </div>
</template>

<style scoped>

</style>