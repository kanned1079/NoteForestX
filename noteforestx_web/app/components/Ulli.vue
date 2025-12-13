<script setup lang="ts">
const props = defineProps<{
  value: string
  bold?: boolean
}>()
</script>

<template>
<div class="flex flex-row justify-start items-center mt-1.5 mb-1.5">
  <div class="w-3 h-3 bg-[#667388] mr-4" style="border-radius: 4px"></div>
  <div class="font-normal opacity-90" :style="props.bold?'font-weight: bold':''">{{ props.value }}</div>
</div>
</template>

<style scoped>

</style>