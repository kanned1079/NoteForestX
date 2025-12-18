<script setup lang="ts">

import MyCard from "~/components/RedesignedComponents/MyCard.vue";
import {useScrollFadeIn} from "~/composables/useScrollFadeIn";

const designSteps = [
  {
    id: 0,
    title: "字體設計",
    content: "挑選字體類型、字號大小與字重粗細。"
  },
  {
    id: 1,
    title: "間距配置",
    content: "調整元素位置，並為元素之間設定適當間距。"
  },
  {
    id: 2,
    title: "色彩選用",
    content: "挑選具備充足對比度的色彩調色體系。"
  },
  {
    id: 3,
    title: "視覺效果",
    content: "添加邊框、陰影、圓角等視覺效果。"
  },
]

useScrollFadeIn({
  selector: '.animated-card-my-design',
  direction: 'up',
  x: 300,
  stagger: 0.1,
  duration: 0.6,
  start: 'top 90%'
})

</script>

<template>
  <div class="meteor-root relative w-full overflow-hidden index-root pb-20" id="work-design">
    <!-- 背景下移 20px -->
    <BackgroundGrid :offsetTop="140" />


    <div class="flex justify-center">
      <div class="relative max-w-[1200px] container z-10 px-8 pt-20 text-slate-800 dark:text-slate-100">


        <div class="space-y-3">
          <p class="text-xl font-semibold text-[#3261e4] dark:text-[#4f77e6]">簡潔 & 直觀</p>
          <p class="text-4xl font-semibold">引人注目、現代簡約的設計。</p>
          <p class="text-sm">保持使用者介面簡潔現代，同時不影響使用者體驗。</p>
        </div>

        <div class="flex w-full mt-10 min-h-[400px] animated-card-my-design">
          <!-- 左 -->
          <div class="hidden md:block w-1/2">
            <MyCard
                v-for="i in designSteps"
                :key="i.id"
                rounded="xl"
                :padding="12"
                class="mb-3"
                hoverable
            >
              <p class="text-lg font-semibold mb-2">{{ i.title }}</p>
              <p class="text-sm font-light">{{ i.content }}</p>
            </MyCard>
          </div>

          <!-- 右 -->
          <div class="w-full md:w-1/2 flex justify-center items-center animated-card-my-design">
            <SampleDesignedCard />
          </div>
        </div>

        <div class="space-y-3 animated-card-my-design">
          <p class="text-xl font-semibold text-[#3261e4] dark:text-[#4f77e6]">講究細節</p>
          <p class="text-4xl font-semibold">擁有敏銳的細節洞察能力</p>
          <p class="text-sm">著重操作便捷性、使用者介面一致性，並持續優化使用者體驗</p>
        </div>

        <div class="space-y-3 mt-6 animated-card-my-design">
          <p class="text-xl font-semibold text-[#3261e4] dark:text-[#4f77e6]">美觀且高效</p>
          <p class="text-4xl font-semibold">程式碼簡潔易懂且性能優化</p>
          <p class="text-sm">撰寫乾淨整潔的程式碼為首要原則，同時兼顧最大限度的效能優化</p>
        </div>

      </div>
    </div>


  </div>
</template>

<style scoped>
.p-card-body {

}

--p-card-body-padding {
  padding: 12px !important;
}

</style>