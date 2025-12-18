<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(
    defineProps<{
      withFade?: boolean
      withGradient?: boolean
      gradientColor?: string
    }>(),
    {
      withFade: true,
      withGradient: true,
      gradientColor: ''
    }
)

/**
 * 🌈 渐变背景样式（浅色模式）
 * - 默认：沿用原有渐变
 * - 自定义：使用 gradientColor
 */
const gradientStyle = computed(() => {
  if (!props.gradientColor) {
    return {
      background: `linear-gradient(
        to bottom,
        rgba(171, 201, 255, 0.08) 0%,
        rgba(210, 226, 255, 0.06) 20%,
        rgba(59,130,246,0.0) 30%
      )`
    }
  }

  return {
    background: `linear-gradient(
      to bottom,
      ${props.gradientColor} 0%,
      ${props.gradientColor} 20%,
      transparent 30%
    )`
  }
})
</script>

<template>
  <!-- 背景容器 -->
  <div class="background-grid absolute inset-0 pointer-events-none z-0">

    <!-- 🌈 漸變層（可选 + 浅色模式） -->
    <div
        v-if="withGradient"
        class="background-grid__gradient dark:hidden"
        :style="gradientStyle"
    />

    <!-- 🧩 網格層 -->
    <div
        class="background-grid__grid"
        :class="{
        'background-grid__grid--fade': withFade
      }"
    />
  </div>
</template>

<style scoped>
.background-grid {
  position: absolute;
  inset: 0;
  pointer-events: none;
  z-index: 0;
}

/* 🌈 漸變層 */
.background-grid__gradient {
  position: absolute;
  inset: 0;
  pointer-events: none;
  z-index: 0;
}

/* 🧩 網格層（基础） */
.background-grid__grid {
  position: absolute;
  inset: 0;
  pointer-events: none;
  z-index: 1;

  background-image: url("data:image/svg+xml;utf8,\
<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 32 32' fill='none' stroke='rgba(100,116,139,0.22)' stroke-width='1'>\
<path d='M32 0H0V32'/>\
</svg>");
  background-repeat: repeat;
  background-size: 32px 32px;
}

/* ✅ 网格渐隐（可选） */
.background-grid__grid--fade {
  -webkit-mask-image: -webkit-linear-gradient(
      top,
      rgba(0,0,0,1) 0%,
      rgba(0,0,0,1) 30%,
      rgba(0,0,0,0) 70%
  );
  mask-image: linear-gradient(
      to bottom,
      rgba(0,0,0,1) 0%,
      rgba(0,0,0,1) 30%,
      rgba(0,0,0,0) 70%
  );
}

/* 🌙 深色模式 */
.dark .background-grid__grid {
  background-image: url("data:image/svg+xml;utf8,\
<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 32 32' fill='none' stroke='rgba(226,232,240,0.12)' stroke-width='1'>\
<path d='M32 0H0V32'/>\
</svg>");
}
</style>