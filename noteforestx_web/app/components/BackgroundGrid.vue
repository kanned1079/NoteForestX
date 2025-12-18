<script setup lang="ts">
import { computed } from 'vue'
import type { CSSProperties } from 'vue'

const props = withDefaults(
    defineProps<{
      withFade?: boolean
      withGradient?: boolean
      gradientColor?: string
      offsetTop?: number
    }>(),
    {
      withFade: true,
      withGradient: true,
      gradientColor: '',
      offsetTop: 0
    }
)

/**
 * 🌈 渐变背景样式（浅色模式）
 * - 默认：沿用原有渐变
 * - 自定义：使用 gradientColor
 */
const gradientStyle = computed(() => {
  const startOffset = props.offsetTop || 0;
  return {
    background: props.gradientColor
        ? `linear-gradient(
          to bottom,
          ${props.gradientColor} ${startOffset}px,
          ${props.gradientColor} calc(${startOffset}px + 20%),
          transparent calc(${startOffset}px + 30%)
        )`
        : `linear-gradient(
          to bottom,
          rgba(171, 201, 255, 0.08) ${startOffset}px,
          rgba(210, 226, 255, 0.06) calc(${startOffset}px + 20%),
          rgba(59,130,246,0.0) calc(${startOffset}px + 30%)
        )`,
    position: 'absolute',
    left: '0',
    right: '0',
    bottom: '0',
    top: startOffset + 'px'
  }
})
</script>

<template>
  <div class="background-grid absolute inset-0 pointer-events-none z-0" :style="{ top: offsetTop + 'px' }">
    <!-- 渐变层 -->
    <div
        v-if="withGradient"
        class="background-grid__gradient dark:hidden"
        :style="gradientStyle"
    />

    <!-- 网格层 -->
    <div
        class="background-grid__grid"
        :class="{ 'background-grid__grid--fade': withFade }"
        :style="{ top: offsetTop + 'px' }"
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