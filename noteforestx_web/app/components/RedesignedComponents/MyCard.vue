<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(
    defineProps<{
      title?: string
      padding?: number          // ✅ px
      rounded?: 'none' | 'sm' | 'md' | 'lg' | 'xl'
      bordered?: boolean
      hoverable?: boolean
    }>(),
    {
      padding: 20,
      rounded: 'lg',
      bordered: true,
      hoverable: false
    }
)

/* 圆角映射 */
const roundedClassMap = {
  none: 'rounded-none',
  sm: 'rounded-sm',
  md: 'rounded-md',
  lg: 'rounded-lg',
  xl: 'rounded-xl'
}

/* class 组合 */
const cardClass = computed(() => [
  'bg-white dark:bg-[#141414]',
  props.bordered && 'border border-slate-200 dark:border-slate-700',
  roundedClassMap[props.rounded],
  props.hoverable &&
  'transition-all duration-300 hover:-translate-y-1 hover:shadow-lg'
])

/* padding style */
const cardStyle = computed(() => ({
  padding: `${props.padding}px`
}))
</script>

<template>
  <div :class="cardClass" :style="cardStyle">
    <!-- 标题 -->
    <div
        v-if="title"
        class="mb-3 text-lg font-semibold text-slate-800 dark:text-slate-100"
    >
      {{ title }}
    </div>

    <!-- 内容 -->
    <div class="text-slate-700 dark:text-slate-300">
      <slot />
    </div>
  </div>
</template>