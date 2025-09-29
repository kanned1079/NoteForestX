<script setup lang="ts">
const props = defineProps<{
  title: string
  subtitle: string
}>()
</script>

<template>
<div class="">
  <div class="opacity-80 font-base text-xl mb-1">{{ props.title }}</div>
  <div v-if="subtitle!==''" class="opacity-60 text-sm">{{ props.subtitle }}</div>
</div>
</template>

<style scoped>

</style>