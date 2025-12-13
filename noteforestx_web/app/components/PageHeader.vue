<script setup lang="ts">
const props = defineProps<{
  title: string
  subtitle?: string
}>()
</script>

<template>
<div class="mb-4 mt-4">
  <div class="font-base text-2xl mb-1">{{ props.title }}</div>
  <div class="opacity-70" v-if="props.subtitle">{{ props.subtitle }}</div>
  <div class="opacity-70" v-if="!props.subtitle"> <slot name="subtitle"></slot> </div>
</div>
</template>

<style scoped>

</style>