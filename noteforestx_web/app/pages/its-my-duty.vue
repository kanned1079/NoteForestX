<script setup lang="ts">
import { useScrollFadeIn } from '~/composables/useScrollFadeIn'

useScrollFadeIn({
  selector: '.animate-card-its-my-duty',
  y: 60,
  stagger: 0.15
})
</script>

<template>
  <!-- 外层：负责居中 & 呼吸空间 -->

  <div class="flex flex-col justify-center items-center ">
    <!-- 纸张本体 -->


    <div class="space-y-3 mt-10 mb-10 text-center">
      <p class="text-3xl font-semibold animate-card-its-my-duty">「 為什麼？ 」</p>
      <p class="opacity-80 animate-card-its-my-duty">Why ?</p>
      <p class="text-3xl font-bold mt-6 animate-card-its-my-duty">因為這是我的職責</p>
      <p class="opacity-80 animate-card-its-my-duty">Because this is my duty.</p>
    </div>
    <div class="a4-paper animate-card-its-my-duty">
      <slot />
    </div>

    <p class="mt-10 animate-card-its-my-duty">A4 Plain Paper</p>

  </div>
</template>

<style scoped>
/* A4 横向：29.7 : 21 ≈ 1.414 */
.a4-paper {
  width: min(100%, 1000px);
  aspect-ratio: 29.7 / 21;

  /* ✅ 永远纯白 */
  background-color: #f6f6f6;

  /* 轻微边缘（像真实纸张） */
  border-radius: 2px;

  /* 🧠 关键：多层阴影，适配深浅色背景 */
  box-shadow:
    /* 近处阴影（边缘厚度感） */
      0 2px 4px rgba(0, 0, 0, 0.06),
        /* 中层悬浮 */
      0 12px 24px rgba(0, 0, 0, 0.12),
        /* 远处环境阴影 */
      0 32px 64px rgba(0, 0, 0, 0.18);

  /* 性能 & 视觉稳定 */
  will-change: transform, box-shadow;
}

.a4-paper:hover {
  transform: translateY(-10px);
}

</style>