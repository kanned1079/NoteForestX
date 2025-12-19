<script setup lang="ts">
const props = defineProps<{
  keyLabel: string
  size?: 'sm' | 'md' | 'lg'
}>()

// 根据 size 设置 padding 和字体大小
const sizeClasses = {
  sm: 'px-1 py-1 text-xs',
  md: 'px-3 py-2 text-base',
  lg: 'px-4 py-3 text-lg',
}
</script>

<template>
  <div
      class="inline-flex items-center justify-center rounded-md border select-none cursor-default
           transition-all duration-150 ease-in-out
           bg-[#cbd3df] dark:bg-gray-800
           border-gray-500 dark:border-gray-600
           text-gray-700 dark:text-gray-200
           active:translate-y-[1px] active:shadow-inner
           shadow-sm
           "
      :class="sizeClasses[props.size || 'sm']"
  >
    {{ props.keyLabel }}
  </div>
</template>

<style scoped>
/* 鼠标按下的按键视觉效果 */
.active {
  transform: translateY(1px);
  box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.2);
}
</style>