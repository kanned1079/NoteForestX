<template>
  <client-only>
    <div class="meteor-root relative w-full overflow-hidden index-root">
      <div ref="containerRef" class="meteor-container absolute inset-0 pointer-events-none"></div>

      <!-- 中间文字，略偏上 -->

      <Transition >
        <div v-if="false">

        </div>
      </Transition>

      <div
          class="absolute inset-0 flex flex-col justify-center items-center z-10 px-6 leading-relaxed opacity-90"
      >
        <div class="text-start space-y-2 transform -translate-y-1/6 text-xl md:text-2xl font-mono">
          <div>I never wished for an easy life.</div>
          <div>I am the light itself ✨ —</div>
          <div>seated as the mountain,</div>
          <div>painting the world with the ink of spring.</div>
        </div>
      </div>

      <!-- Scroll Down 固定在底部 -->
      <div class="absolute bottom-6 left-1/2 transform -translate-x-1/2 flex flex-col items-center gap-2 z-10">
        <span class="text-xl font-normal">Scroll down to know more.</span>
        <span class="animate-bounce h-8 w-8 bg-primary text-primary-contrast rounded-full inline-flex items-center justify-center">
          <i class="mt-2 opacity-70 pi pi-arrow-down" />
        </span>
      </div>
    </div>
  </client-only>
</template>

<script setup lang="ts">
import { ref, onMounted, nextTick, onBeforeUnmount } from 'vue'

const containerRef = ref<HTMLElement | null>(null)
let resizeHandler: (() => void) | null = null


onMounted(async () => {
  await nextTick()
  const container = containerRef.value
  if (!container) return

  const { gsap } = await import('gsap')

  const METEOR_COUNT = 5
  const MIN_DURATION = 1.8
  const MAX_DURATION = 3
  const MIN_LENGTH = 120
  const MAX_LENGTH = 480

  let W = window.innerWidth
  let H = window.innerHeight
  resizeHandler = () => {
    W = window.innerWidth
    H = window.innerHeight
  }
  window.addEventListener('resize', resizeHandler)

  const createMeteor = (index: number) => {
    const meteor = document.createElement('div')
    meteor.className = 'meteor'
    const len = MIN_LENGTH + Math.random() * (MAX_LENGTH - MIN_LENGTH)
    meteor.style.height = `${len}px`
    meteor.style.visibility = 'hidden'
    // meteor.style.background = getMeteorColor() // 设置颜色
    container.appendChild(meteor)

    // 起点：屏幕上方随机位置
    const startX = Math.random() * W
    const startY = -100 - Math.random() * 200  // -100 ~ -300px

    // 角度：保持45度 ±5°
    const angleDeg = 40 + Math.random() * 10
    const rad = (180 - angleDeg) * (Math.PI / 180)

    // 运动距离：屏幕宽度的一半左右随机
    const travel = W/2 * (0.5 + Math.random() * 0.4)

    const endX = startX + Math.cos(rad) * travel
    const endY = startY + Math.sin(rad) * travel

    // 旋转角度
    const rotateDeg = (Math.atan2(endY - startY, endX - startX) * 180) / Math.PI - 90

    gsap.set(meteor, {
      x: startX,
      y: startY,
      rotate: rotateDeg,
      transformOrigin: '0 0',
      opacity: 1,
      visibility: 'visible'
    })

    const duration = MIN_DURATION + Math.random() * (MAX_DURATION - MIN_DURATION)
    gsap.to(meteor, {
      x: endX,
      y: endY,
      opacity: 1,
      duration,
      ease: 'power2.out'
      // 不删除元素，停留在屏幕上
    })
  }

  for (let i = 0; i < METEOR_COUNT; i++) {
    setTimeout(() => createMeteor(i), i * 400 + Math.random() * 600)
  }
})

onBeforeUnmount(() => {
  if (resizeHandler) window.removeEventListener('resize', resizeHandler)
})
</script>

<style>

.index-root {
  height: 80vh;
}

.meteor {
  position: absolute;
  width: 3px;
  height: 120px;
  background: linear-gradient(0deg, rgba(255, 255, 255, 1), rgba(255, 255, 255, 0));
  filter: drop-shadow(0 0 10px rgba(255, 255, 255, 0.8))
  drop-shadow(0 0 20px rgba(180, 200, 255, 0.6));
  pointer-events: none;
  transform-origin: 0 0;
  border-radius: 2px;
  will-change: transform, opacity;
}

:root {
  --meteor-color: linear-gradient(0deg, #555555, rgba(85, 85, 85, 0));
}

@media (prefers-color-scheme: dark) {
  :root {
    --meteor-color: linear-gradient(0deg, #FFD700, rgba(255, 215, 0, 0));
  }
}

.meteor {
  position: absolute;
  width: 3px;
  height: 120px;
  background: var(--meteor-color);
  filter: drop-shadow(0 0 10px rgba(255, 255, 255, 0.8))
  drop-shadow(0 0 30px rgba(180, 200, 255, 0.2));
  pointer-events: none;
  transform-origin: 0 0;
  border-radius: 2px;
  will-change: transform, opacity;
}

</style>