<template>
  <client-only>
    <div class="meteor-root relative w-full overflow-hidden index-root">
      <BackgroundGrid />

      <!-- 流星容器 -->
      <div
          ref="containerRef"
          class="meteor-container absolute inset-0 pointer-events-none"
      ></div>

      <!-- 中央内容区域（1200px 居中） -->
      <div
          ref="textRef"
          class="absolute inset-0 z-20 flex items-center justify-center px-4"
      >
        <div class="w-full max-w-[1200px] flex items-center">
          <!-- 左侧：文字内容 -->
          <div class="w-full md:w-1/2 text-slate-900 dark:text-slate-100">
            <div class="opacity-90">
              <p class="text-2xl sm:text-3xl md:text-4xl animated-card-firstpage">
                hi !
              </p>

              <p
                  class="font-bold animated-card-firstpage
                       text-4xl sm:text-5xl md:text-6xl lg:text-7xl"
              >
                I'm <span class="text-[#1e88a8] underline">kanned1079</span>,
              </p>

              <div class="mt-6 text-base sm:text-lg animated-card-firstpage">
                a <span class="font-bold">FullStack Developer</span> who loves intuitive,
              </div>
              <div class="text-base sm:text-lg animated-card-firstpage">
                clean and modern UI design.
              </div>
            </div>

            <div class="space-y-2 mt-4 text-sm sm:text-base animated-card-firstpage">
              <div class="mt-10">I never wished for an easy life.</div>
              <div>I am the light itself ✨ —</div>
              <div>seated as the mountain,</div>
              <div>painting the world with the ink of spring.</div>
            </div>
          </div>

          <!-- 右侧：桌面端显示 -->
          <div class="hidden md:flex w-1/2 justify-center items-center">
            <!-- 预留区域：Logo / 插画 / 动画 -->
<!--            <div-->
<!--                class="w-64 h-64 lg:w-80 lg:h-80-->
<!--                     rounded-2xl-->
<!--                     border border-dashed border-slate-400/40-->
<!--                     flex items-center justify-center-->
<!--                     text-slate-400"-->
<!--            >-->
<!--              Logo / Visual-->
<!--            </div>-->

            <div class="hidden md:flex w-1/2 justify-center items-center">
              <div class="w-80 h-80 lg:w-96 lg:h-96">
                <ClientOnly>
                  <RightVisual />

                </ClientOnly>
              </div>
            </div>


          </div>
        </div>
      </div>

      <!-- Scroll Down 提示 -->
      <div
          ref="scrollRef"
          class="absolute bottom-6 left-1/2 transform -translate-x-1/2
               flex flex-col items-center gap-2 z-20 opacity-0"
      >
        <span class="text-xl font-normal">Scroll down to know more.</span>
        <span
            class="h-8 w-8 bg-primary text-primary-contrast
                 rounded-full inline-flex items-center justify-center"
        >
          <i class="mt-2 opacity-70 pi pi-arrow-down" />
        </span>
      </div>
    </div>
  </client-only>
</template>

<script setup lang="ts">
import { ref, onMounted, nextTick, onBeforeUnmount } from 'vue'
import {useScrollFadeIn} from "~/composables/useScrollFadeIn";

const containerRef = ref<HTMLElement | null>(null)
const textRef = ref<HTMLElement | null>(null)
const scrollRef = ref<HTMLElement | null>(null)
let resizeHandler: (() => void) | null = null
let meteorInterval: number | null = null

useScrollFadeIn({
  selector: '.animated-card-firstpage',
  direction: 'up',
  x: 600,
  stagger: 0.2,
  duration: 0.6,
  start: 'top 90%'
})

onMounted(async () => {
  await nextTick()
  const container = containerRef.value
  if (!container) return

  const { gsap } = await import('gsap')

  // ------------------------
  // 文字动画
  // ------------------------
  // if (textRef.value) {
  //   const lines = textRef.value.querySelectorAll('div > div')
  //   gsap.fromTo(
  //       lines,
  //       { y: 20, opacity: 0 },
  //       { y: 0, opacity: 1, duration: 1, ease: 'power2.out', stagger: 0.15 }
  //   )
  // }

  // Scroll Down 延迟出现
  if (scrollRef.value) {
    gsap.to(scrollRef.value, { opacity: 1, y: 0, duration: 1, ease: 'power2.out', delay: 2 })
  }

  // ------------------------
  // 流星动画（循环生成）
  // ------------------------
  const METEOR_COUNT = 5
  const MIN_LENGTH = 120
  const MAX_LENGTH = 480
  const MIN_DURATION = 1.5
  const MAX_DURATION = 3.5

  let W = window.innerWidth
  let H = window.innerHeight
  resizeHandler = () => {
    W = window.innerWidth
    H = window.innerHeight
  }
  window.addEventListener('resize', resizeHandler)

  const createMeteor = () => {
    if (!container) return
    const meteor = document.createElement('div')
    meteor.className = 'meteor'
    const len = MIN_LENGTH + Math.random() * (MAX_LENGTH - MIN_LENGTH)
    meteor.style.height = `${len}px`
    container.appendChild(meteor)

    const startX = Math.random() * W
    const startY = Math.random() * (H * 0.5)
    const angleDeg = 40 + Math.random() * 10
    const rad = (180 - angleDeg) * (Math.PI / 180)
    const travel = W * (0.5 + Math.random() * 0.5)
    const endX = startX + Math.cos(rad) * travel
    const endY = startY + Math.sin(rad) * travel
    const rotateDeg = (Math.atan2(endY - startY, endX - startX) * 180) / Math.PI - 90
    const finalOpacity = 0.4 + Math.random() * 0.6
    const duration = MIN_DURATION + Math.random() * (MAX_DURATION - MIN_DURATION)

    gsap.set(meteor, {
      x: startX,
      y: startY,
      rotate: rotateDeg,
      transformOrigin: '0 0',
      opacity: 0,
      visibility: 'visible'
    })

    // 流星移动动画
    gsap.to(meteor, {
      x: endX,
      y: endY,
      opacity: finalOpacity,
      duration,
      ease: 'power2.out',
      onComplete: () => {
        // 移动结束后淡出动画
        gsap.to(meteor, {
          opacity: 0,
          duration: 0.5,
          ease: 'power1.out',
          onComplete: () => meteor.remove()
        })
      }
    })
  }

  // 初始生成几个流星
  for (let i = 0; i < METEOR_COUNT; i++) {
    setTimeout(createMeteor, i * 400)
  }

  // 循环生成流星
  meteorInterval = window.setInterval(() => {
    createMeteor()
  }, 800) // 每隔 0.8 秒生成一颗
})

onBeforeUnmount(() => {
  if (resizeHandler) window.removeEventListener('resize', resizeHandler)
  if (meteorInterval) clearInterval(meteorInterval)
})
</script>

<style>
.index-root {
  height: 72vh;
}

.meteor {
  position: absolute;
  width: 3px;
  height: 120px;
  background: var(--meteor-color);
  filter: drop-shadow(0 0 8px var(--meteor-glow-color))
  drop-shadow(0 0 20px var(--meteor-glow-color));
  pointer-events: none;
  transform-origin: 0 0;
  border-radius: 2px;
  will-change: transform, opacity;
}

/* 浅色模式 */
:root {
  --meteor-color: linear-gradient(0deg, rgba(59,130,246,0.6), rgba(59,130,246,0));
  --meteor-glow-color: rgba(59,130,246,0.6);
}

/* 深色模式 */
@media (prefers-color-scheme: dark) {
  :root {
    --meteor-color: linear-gradient(0deg, #FFD700, rgba(255,215,0,0));
    --meteor-glow-color: rgba(255,215,0,0.8);
  }
}
</style>