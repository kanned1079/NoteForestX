<script setup lang="ts">
// 无 props，目前直接使用全屏铺满
</script>

<template>
  <!-- 背景容器 -->
  <div class="background-grid absolute inset-0 pointer-events-none z-0">

    <!-- 🌈 漸變層（只在淺色模式） -->
    <div class="background-grid__gradient dark:hidden"></div>

    <!-- 🧩 網格層 -->
    <div class="background-grid__grid"></div>

  </div>
</template>

<style scoped>
.background-grid {
  position: absolute;
  inset: 0;
  pointer-events: none;
  z-index: 0;
}

/* 🌈 漸變層：整片，不 repeat */
.background-grid__gradient {
  position: absolute;
  inset: 0;
  pointer-events: none;

  background: linear-gradient(
      to bottom,
      rgba(171, 201, 255, 0.08) 0%,
      rgba(210, 226, 255, 0.06) 20%,
      rgba(59,130,246,0.0) 30%
  );
  z-index: 0;
}

/* 🧩 網格層：单纯 repeat */
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

  /* 让格子整体渐隐 */
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

/* 🌙 深色模式：灰色网格 */
.dark .background-grid__grid {
  background-image: url("data:image/svg+xml;utf8,\
<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 32 32' fill='none' stroke='rgba(226,232,240,0.12)' stroke-width='1'>\
<path d='M32 0H0V32'/>\
</svg>");
}
</style>