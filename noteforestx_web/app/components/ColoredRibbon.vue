<template>
  <!--  1-->
  <canvas ref="canvasRef" class="confetti-canvas"></canvas>
</template>

<script lang="ts" setup>
import {onMounted, onUnmounted, ref, watch} from 'vue';
import * as confetti from 'canvas-confetti';

// Props
const props = defineProps({
  show: {
    type: Boolean,
    default: false,
  },
});

// Refs
const canvasRef = ref<HTMLCanvasElement | null>(null);
let confettiInstance: confetti.CreateTypes | null = null;

// 初始化 confetti 实例
const initConfetti = () => {
  if (canvasRef.value) {
    confettiInstance = confetti.create(canvasRef.value, {
      resize: true,
      useWorker: true,
    });
  }
};

// Fire confetti animation
const fireConfetti = () => {
  if (!confettiInstance) return;

  const count = 600; // 基础粒子数量
  const defaults = {
    origin: {y: 1.4}, // 从页面底部发射
    gravity: 3, // 模拟自然重力
    scalar: 1.5,
    startVelocity: 3,
  };

  const fire = (particleRatio: number, opts: confetti.Options) => {
    confettiInstance?.({
      ...defaults,
      ...opts,
      particleCount: Math.floor(count * particleRatio),
    });
  };

  // 不同的动画设置
  fire(0.25, {
    spread: 90,
    startVelocity: 120,
  });
  fire(0.2, {
    spread: 140,
    startVelocity: 100,
  });
  fire(0.35, {
    spread: 150,
    decay: 0.91,
    scalar: 0.8,
    startVelocity: 95,
  });
  fire(0.1, {
    spread: 180,
    startVelocity: 130,
    decay: 0.92,
    scalar: 1.2,
  });
  fire(0.1, {
    spread: 120,
    startVelocity: 70,
  });
  fire(0.1, {
    spread: 180,          // 彩条将以120度的范围从中心发射
    startVelocity: 130,   // 彩条的起始速度为130，意味着它会飞得较高较远
    decay: 0.92,          // 彩条会逐渐减速，但速度衰减的速度比较平缓
    scalar: 1.1,          // 彩条会稍微大一些
  });
};

// Watch `show` prop to trigger animation
watch(
    () => props.show,
    (newValue) => {
      if (newValue) {
        fireConfetti();
      }
    },
);

// Lifecycle hooks
onMounted(() => {
  initConfetti();
});

onUnmounted(() => {
  confettiInstance = null;
});
</script>

<script lang="ts">
export default {
  name: 'ColoredRibbon',
}
</script>

<style scoped>
.confetti-canvas {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 9999;
  pointer-events: none;
}
</style>