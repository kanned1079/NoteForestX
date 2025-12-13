<script setup lang="ts">
import MyWork from "~/components/MyIntro/MyWork.vue"
</script>

<template>
  <section class="relative ">

    <!-- 🌈 漸變層（只在淺色模式） -->
    <div class="my-work__gradient dark:hidden"></div>

    <!-- 🧩 網格層 -->
    <div class="my-work__grid"></div>

    <!-- 📄 內容層 -->
    <div class="relative z-10">
      <MyWork />
    </div>

  </section>
</template>

<style scoped>
/* 🌈 漸變層：整片，不 repeat */
.my-work__gradient {
  position: absolute;
  inset: 0;
  pointer-events: none;
  z-index: 0;

  background: linear-gradient(
      to bottom,
      rgba(171, 201, 255, 0.08) 0%,   /* 原 0.2 -> 0.08 */
      rgba(210, 226, 255, 0.06) 20%,  /* 原 0.12 -> 0.06 */
      rgba(59,130,246,0.0) 30%
  );
}

/* 🧩 網格層：单纯 repeat */
.my-work__grid {
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

  /* ✅ 关键：让格子整体渐隐 */
  /* Safari / WebKit */
  -webkit-mask-image: -webkit-linear-gradient(
      top,
      rgba(0,0,0,1) 0%,
      rgba(0,0,0,1) 30%,
      rgba(0,0,0,0) 70%
  );

  /* 标准 */
  mask-image: linear-gradient(
      to bottom,
      rgba(0,0,0,1) 0%,
      rgba(0,0,0,1) 30%,
      rgba(0,0,0,0) 70%
  );
}

/* 🌙 深色模式：换成灰色网格 */
.dark .my-work__grid {
  background-image: url("data:image/svg+xml;utf8,\
<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 32 32' fill='none' stroke='rgba(226,232,240,0.12)' stroke-width='1'>\
<path d='M32 0H0V32'/>\
</svg>");
}
</style>