<script setup lang="ts">
const props = defineProps<{
  title: string
  subtitle: string
}>()
</script>

<template>
<div class="mb-4">
  <Button class="mb-2 pl-0" size="small" type="button" label="返回" icon="pi pi-angle-left" variant="link"/>
  <div class="font-bold text-xl mb-1">{{ props.title }}</div>
  <div class="opacity-70">{{ props.subtitle }}</div>
</div>
</template>

<style scoped>

</style>