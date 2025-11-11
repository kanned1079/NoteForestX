<template>
<!--  <client-only>-->
    <div class="p-10 flex flex-col items-center gap-2 h-screen">

      <!-- 顶部提示 -->
<!--      <div class="flex flex-col items-center gap-2">-->
<!--        <span class="text-xl font-medium">Scroll Down</span>-->
<!--        <span class="animate-bounce h-8 w-8 bg-primary text-primary-contrast rounded-full inline-flex items-center justify-center">-->
<!--          <i class="pi pi-arrow-down" />-->
<!--        </span>-->
<!--      </div>-->

      <!-- 第一排卡片 -->
      <div class="flex flex-wrap justify-center gap-8">
        <div v-for="(item, i) in firstRow" :key="i" ref="cards" class="card flex flex-col border border-surface shadow-lg justify-center items-center max-w-80 rounded-lg p-8 gap-4 opacity-0  transition-transform duration-500 ease-out hover:-translate-y-[10px]">
          <div class="rounded-full bg-primary text-primary-contrast w-12 h-12 flex items-center justify-center">
            <i :class="item.icon + ' !text-2xl'"></i>
          </div>
          <span class="text-2xl font-bold">{{ item.title }}</span>
          <span class="text-muted-color text-center">{{ item.description }}</span>
        </div>
      </div>

      <!-- 第二排卡片 -->
      <div class="flex flex-wrap justify-center gap-8">
        <div v-for="(item, i) in secondRow" :key="i" ref="cards" class="card flex flex-col justify-center items-center max-w-80 rounded-2xl p-8 gap-4 opacity-0">
          <img v-if="item.avatar" :src="item.avatar" class="w-24 h-24 rounded-full object-cover"/>
          <span class="text-2xl font-medium">{{ item.title }}</span>
          <span class="text-muted-color text-center">{{ item.description }}</span>
        </div>
      </div>

      <!-- 第三排统计卡片 -->
      <div class="flex flex-wrap justify-center gap-8">
        <div v-for="(item, i) in thirdRow" :key="i" ref="cards" class="card flex flex-col bg-primary text-primary-contrast border-primary shadow-lg justify-center items-center max-w-80 rounded-2xl p-8 gap-4 opacity-0">
          <span class="bg-white/20 text-xl font-medium rounded-xl px-4 py-2">{{ item.value }}</span>
          <span class="text-2xl font-bold">{{ item.title }}</span>
          <span class="text-center">{{ item.description }}</span>
        </div>
      </div>

    </div>
<!--  </client-only>-->
</template>

<script setup lang="ts">
import { ref, onMounted, nextTick } from 'vue'
import { gsap } from 'gsap'
import { ScrollTrigger } from 'gsap/ScrollTrigger'

gsap.registerPlugin(ScrollTrigger)

const cards = ref<HTMLElement[]>([])

const firstRow = [
  { title: 'Vue / Nuxt.js', description: '使用Vue系列前端框架构建用户界面，合理运用Nuxt进行SSR以提高搜索引擎索引。', icon: 'pi pi-user' },
  { title: 'Golang / Nest.js', description: 'Lorem ipsum dolor sit amet consectetur.', icon: 'pi pi-users' },
  { title: 'Final Cut Pro X', description: 'Lorem ipsum dolor sit amet consectetur.', icon: 'pi pi-building' }
]

const secondRow = [
  { title: 'Jenna Thompson', description: 'Lorem ipsum dolor sit amet consectetur.', avatar: 'https://primefaces.org/cdn/primevue/images/avatar/amyelsner.png' },
  { title: 'Isabel Garcia', description: 'Lorem ipsum dolor sit amet consectetur.', avatar: 'https://primefaces.org/cdn/primevue/images/avatar/asiyajavayant.png' },
  { title: 'Xavier Mason', description: 'Lorem ipsum dolor sit amet consectetur.', avatar: 'https://primefaces.org/cdn/primevue/images/avatar/onyamalimba.png' }
]

const thirdRow = [
  { title: 'Customers', value: '850K', description: 'Lorem ipsum dolor sit amet consectetur.' },
  { title: 'Revenue', value: '$1.5M', description: 'Lorem ipsum dolor sit amet consectetur.' },
  { title: 'Sales', value: '140K', description: 'Lorem ipsum dolor sit amet consectetur.' }
]

onMounted(async () => {
  await nextTick()
  const elements = document.querySelectorAll<HTMLElement>('.card')

  elements.forEach((el, i) => {
    gsap.fromTo(
        el,
        { opacity: 0, y: 100 },
        {
          opacity: 1,
          y: 0,
          duration: 1,
          ease: 'power2.out',
          scrollTrigger: {
            trigger: el,
            start: 'top 85%',
            toggleActions: 'play none none reverse'
          },
          delay: i * 0.2
        }
    )
  })
})
</script>

<style scoped>
.card {
  /* 初始隐藏通过 gsap 控制 */
}
</style>