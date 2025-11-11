<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { X } from 'lucide-vue-next'
import gsap from 'gsap'

const props = defineProps<{
  expand?: () => void
}>()

const showBanner = ref(false)
const bannerRef = ref<HTMLElement | null>(null)

function closeBanner() {
  if (!bannerRef.value) return

  // 隐藏横幅动画
  props.expand?.()
  gsap.to(bannerRef.value, {
    y: -80,
    opacity: 0,
    duration: 0.6,
    ease: 'power2.inOut',
    onComplete: () => {
      showBanner.value = false
      localStorage.setItem('bannerClosed', 'true') // 用户关闭过就记录
    }
  })

}

onMounted(() => {
  const closed = localStorage.getItem('bannerClosed')
  if (closed === 'true') return // 用户之前关闭过，直接不显示

  showBanner.value = true // 元素显示

  const el = bannerRef.value
  if (!el) return

  props.expand?.() // 横幅出现，下方内容往下移

  // 初次加载，从上方滑入
  gsap.fromTo(
      el,
      { y: -80, opacity: 0 },
      { y: 0, opacity: 1, duration: 0.8, ease: 'power3.out' }
  )
})
</script>

<template>
  <div
      ref="bannerRef"
      :style="{ display: showBanner ? 'flex' : 'none' }"
      class="fixed top-0 left-0 w-full z-[9999] p-1 flex flex-row justify-between bg-[#1e50a2] text-white items-center shadow-lg"
  >
    <div></div>
    <div class="text-sm">👋 Welcome to Kaneko's Blog.</div>
    <div
        class="text-sm opacity-60 pr-2 hover:opacity-95 cursor-pointer"
        @click="closeBanner"
    >
      <X class="w-5" />
    </div>
  </div>
</template>

<style scoped>
/* 全部动画用 GSAP 控制，不需要 transition */
</style>