<script setup lang="ts">
const props = defineProps<{
  value: string
  percent: number // 0-100
  flagUrl?: string
}>()
</script>

<template>
  <Tag
      class="relative overflow-hidden w-auto h-6"
      style="border: 2px solid var(--border-color);"
  >
    <div class="flex items-center gap-2 px-1">
      <img v-if="props.flagUrl" :src="props.flagUrl" alt="Flag" class="w-4 h-auto" />
      <span class="text-sm font-medium">{{ props.value }}</span>
    </div>

    <!-- 横条在 Tag 内部 -->
    <div
        class="absolute bottom-0 left-0 h-[2px] bg-blue-500"
        :style="{ width: props.percent + '%' }"
    ></div>
  </Tag>
</template>

<style scoped>
.absolute {
  transition: width 0.3s ease;
}
</style>