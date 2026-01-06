<script setup lang="ts">
import {useI18n} from "vue-i18n";
import { ref, onMounted, nextTick, onBeforeUnmount } from 'vue'
import { useScrollFadeIn } from '~/composables/useScrollFadeIn'
import murasame_ciallo from '~/assets/imgs/murasame_ciallo.jpg'
import useThemeStore from "~/store/themeStore";
import ProgrammingLangIcons from "~/components/ProgrammingLangIcons.vue";
import VisualPhoto from "~/components/VisualPhoto.vue"
import dayjs from "dayjs";

const {t} = useI18n();
const themeStore = useThemeStore()
const containerRef = ref<HTMLElement | null>(null)
const textRef = ref<HTMLElement | null>(null)
const scrollRef = ref<HTMLElement | null>(null)

const showScroll = ref(false)

let resizeHandler: (() => void) | null = null
let meteorInterval: number | null = null

useScrollFadeIn({
  selector: '.animated-card-firstpage',
  direction: 'up',
  x: 100,
  stagger: 0.1,
  duration: 0.4,
  start: 'top 90%'
})

const METEOR_COUNT = 5
const MIN_LENGTH = 120
const MAX_LENGTH = 480
const MIN_DURATION = 1.5
const MAX_DURATION = 3.5

let W = 0
let H = 0

const onClickWorkBtn = (code: '0' | '1') => {
  themeStore.setWorkTab(code)
  const targetSelector = '#work-section'

  const el = document.querySelector(targetSelector)
  if (el) el.scrollIntoView({ behavior: 'smooth' })
}

onMounted(async () => {
  W = window.innerWidth
  H = window.innerHeight

  showScroll.value = true

  await nextTick()
  const container = containerRef.value
  if (!container) return

  const { gsap } = await import('gsap')

  if (scrollRef.value) {
    gsap.fromTo(
        scrollRef.value,
        { opacity: 0, y: 12 },
        { opacity: 1, y: 0, duration: 1, ease: 'power2.out', delay: 2 }
    )
  }

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

    const rotateDeg =
        (Math.atan2(endY - startY, endX - startX) * 180) / Math.PI - 90

    const duration =
        MIN_DURATION + Math.random() * (MAX_DURATION - MIN_DURATION)

    gsap.set(meteor, {
      x: startX,
      y: startY,
      rotate: rotateDeg,
      transformOrigin: '0 0',
      opacity: 0
    })

    gsap.to(meteor, {
      x: endX,
      y: endY,
      opacity: 0.8,
      duration,
      ease: 'power2.out',
      onComplete: () => {
        gsap.to(meteor, {
          opacity: 0,
          duration: 0.5,
          onComplete: () => meteor.remove()
        })
      }
    })
  }

  const startMeteor = () => {
    if (meteorInterval !== null) return
    meteorInterval = window.setInterval(createMeteor, 800)
  }

  const stopMeteor = () => {
    if (meteorInterval !== null) {
      clearInterval(meteorInterval)
      meteorInterval = null
    }
  }

  // 初始生成
  for (let i = 0; i < METEOR_COUNT; i++) {
    setTimeout(createMeteor, i * 400)
  }

  startMeteor()

  const handleVisibilityChange = () => {
    if (document.hidden) {
      stopMeteor()
    } else {
      startMeteor()
    }
  }

  document.addEventListener('visibilitychange', handleVisibilityChange)

  onBeforeUnmount(() => {
    stopMeteor()
    if (resizeHandler) window.removeEventListener('resize', resizeHandler)
    document.removeEventListener('visibilitychange', handleVisibilityChange)
  })
})
</script>

<template>
  <div
      class="meteor-root relative w-full overflow-hidden index-root pt-[100px]"
  >
    <!-- ================= 背景 ================= -->
    <BackgroundGrid :withFade="true" :withGradient="false" />

    <!-- 流星容器 -->
    <div
        ref="containerRef"
        class="meteor-container absolute inset-0 pointer-events-none"
    />

    <!-- ================= 主内容 ================= -->
    <div
        ref="textRef"
        class="
    relative
    z-20
    flex justify-center
    px-4
  "
    >
      <div
          class="
    w-full max-w-[1200px]
    flex flex-col md:flex-row
    items-center
  "
      >
        <!-- 左侧：文字 -->
        <div
            class="w-full md:flex-1 flex justify-start text-slate-900 dark:text-slate-100 pl-4"
        >
          <div class="w-full max-w-md">
            <!-- 标题 -->
            <div class="opacity-90 animated-card-firstpage">
              <p class="text-2xl sm:text-3xl md:text-4xl">
                hi ! ✌️
              </p>

              <p class="font-bold text-4xl sm:text-5xl md:text-5xl">
                <i18n-t keypath="home.intro_title" tag="span">
                  <template #name>
                    <span class="text-[#1e88a8] underline">kanned1079</span>
                  </template>
                </i18n-t>
              </p>

              <div class="mt-6 text-base sm:text-lg">
                <i18n-t keypath="home.intro_desc" tag="span">
                  <template #role>
                    <span class="font-bold">{{ t('home.fullstack_developer') }}</span>
                  </template>
                </i18n-t>
              </div>

            </div>

            <!-- 按钮区 -->
            <div
                class="mt-8 flex flex-row items-center gap-3 animated-card-firstpage"
            >
              <button
                  class="button h-9 border border-transparent dark:border-gray-500"
                  @click="onClickWorkBtn('1')"
              >
                {{ t('home.get_in_touch') }}
              </button>

              <StarOnGithub link="https://github.com/kanned1079" />
            </div>

            <!-- 技术栈 -->
            <div class="mt-14 animated-card-firstpage">
              <p class="text-sm opacity-80 mb-2">
                {{ t('home.favorite_stack') }}
              </p>
              <ProgrammingLangIcons @click="onClickWorkBtn('0')" />
              <span
                  @click="onClickWorkBtn('0')"
                  class="text-sm opacity-70 mt-2 flex flex-row items-center  hover:underline cursor-pointer">
                {{ t('my_work.main_title') }}
                <span class="pi pi-angle-right font-light"></span>
              </span>
            </div>

          </div>
        </div>

        <!-- 右侧：视觉（桌面端） -->
        <div
            class="
            hidden md:flex
            flex-1
            justify-center items-center
            animated-card-firstpage
          "
        >
          <div
              class="
              aspect-[4/3]
              w-[420px]
              lg:w-[520px]
              xl:w-[620px]
              2xl:w-[700px]
              flex items-center justify-center
            "
          >
            <VisualPhoto class="w-full h-full object-contain" />
          </div>
        </div>
      </div>
    </div>

    <!-- ================= Scroll 提示 ================= -->
    <div
        ref="scrollRef"
        class="mt-20 w-full flex flex-col items-center justify-center gap-2 z-20 transition-all duration-700 ease-out"
    >
<!--      <span class="text-sm md:text-base font-medium opacity-80">-->
<!--        Scroll down to know more-->
<!--      </span>-->
      <span class="text-sm md:text-base font-medium opacity-80">
  {{ t('home.scroll_down') }}
</span>

      <span
          class="
          h-8 w-8
          bg-primary text-primary-contrast
          rounded-full
          inline-flex items-center justify-center
          animate-bounce
        "
      >
        <i class="pi pi-arrow-down" />
      </span>
    </div>
  </div>
</template>

<style>
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


/* From Uiverse.io by zjssun */
.button {
  position: relative;
  border-radius: 6px;
  height: 36px;
  padding: 0 10px;
  border: none;
  color: #fff;
  cursor: pointer;
  background-color: #7d2ae8;
  font-weight: bold;
  transition: all 0.2s ease;
  font-size: 0.75rem;
}

.button:active {
  transform: scale(0.96);
}

.button:before,
.button:after {
  position: absolute;
  content: "";
  width: 150%;
  left: 50%;
  height: 100%;
  transform: translateX(-50%);
  z-index: -1000;
  background-repeat: no-repeat;
}

.button:hover:before {
  top: -70%;
  background-image: radial-gradient(circle, #7d2ae8 20%, transparent 20%),
  radial-gradient(circle, transparent 20%, #7d2ae8 20%, transparent 30%),
  radial-gradient(circle, #7d2ae8 20%, transparent 20%),
  radial-gradient(circle, #7d2ae8 20%, transparent 20%),
  radial-gradient(circle, transparent 10%, #7d2ae8 15%, transparent 20%),
  radial-gradient(circle, #7d2ae8 20%, transparent 20%),
  radial-gradient(circle, #7d2ae8 20%, transparent 20%),
  radial-gradient(circle, #7d2ae8 20%, transparent 20%),
  radial-gradient(circle, #7d2ae8 20%, transparent 20%);
  background-size: 10% 10%, 20% 20%, 15% 15%, 20% 20%, 18% 18%, 10% 10%, 15% 15%,
  10% 10%, 18% 18%;
  background-position: 50% 120%;
  animation: greentopBubbles 0.6s ease;
}

@keyframes greentopBubbles {
  0% {
    background-position: 5% 90%, 10% 90%, 10% 90%, 15% 90%, 25% 90%, 25% 90%,
    40% 90%, 55% 90%, 70% 90%;
  }

  50% {
    background-position: 0% 80%, 0% 20%, 10% 40%, 20% 0%, 30% 30%, 22% 50%,
    50% 50%, 65% 20%, 90% 30%;
  }

  100% {
    background-position: 0% 70%, 0% 10%, 10% 30%, 20% -10%, 30% 20%, 22% 40%,
    50% 40%, 65% 10%, 90% 20%;
    background-size: 0% 0%, 0% 0%, 0% 0%, 0% 0%, 0% 0%, 0% 0%;
  }
}

.button:hover::after {
  bottom: -70%;
  background-image: radial-gradient(circle, #7d2ae8 20%, transparent 20%),
  radial-gradient(circle, #7d2ae8 20%, transparent 20%),
  radial-gradient(circle, transparent 10%, #7d2ae8 15%, transparent 20%),
  radial-gradient(circle, #7d2ae8 20%, transparent 20%),
  radial-gradient(circle, #7d2ae8 20%, transparent 20%),
  radial-gradient(circle, #7d2ae8 20%, transparent 20%),
  radial-gradient(circle, #7d2ae8 20%, transparent 20%);
  background-size: 15% 15%, 20% 20%, 18% 18%, 20% 20%, 15% 15%, 20% 20%, 18% 18%;
  background-position: 50% 0%;
  animation: greenbottomBubbles 0.6s ease;
}

@keyframes greenbottomBubbles {
  0% {
    background-position: 10% -10%, 30% 10%, 55% -10%, 70% -10%, 85% -10%,
    70% -10%, 70% 0%;
  }

  50% {
    background-position: 0% 80%, 20% 80%, 45% 60%, 60% 100%, 75% 70%, 95% 60%,
    105% 0%;
  }

  100% {
    background-position: 0% 90%, 20% 90%, 45% 70%, 60% 110%, 75% 80%, 95% 70%,
    110% 10%;
    background-size: 0% 0%, 0% 0%, 0% 0%, 0% 0%, 0% 0%, 0% 0%;
  }
}

</style>