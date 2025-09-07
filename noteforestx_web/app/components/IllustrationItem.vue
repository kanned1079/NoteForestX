<script setup lang="ts">
const props = defineProps<{
  name?: string
  author?: string
  img_url: string
}>()
</script>

<template>
  <div class="flex flex-col mb-5">
    <!-- 上半部分：正方形 + 背景 + 圆角 -->
    <div class="aspect-square w-full rounded-md overflow-hidden bg-gray-100 dark:bg-gray-800 flex items-center justify-center">
      <img
          :src="img_url"
          alt="illustration"
          class="w-full h-full object-contain"
          loading="lazy"
      />
    </div>

    <!-- 下半部分：纯文字 -->
    <div class="mt-2 px-1">
      <div class="font-medium text-base truncate">
        {{ name || 'Untitled Illustration' }}
      </div>
      <div class="text-sm text-gray-600 dark:text-gray-300 truncate">
        {{ author || 'Unknown Author' }}
      </div>
    </div>
  </div>
</template>