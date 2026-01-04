<script setup lang="ts">
import { useI18n } from "vue-i18n";
import MyCard from "~/components/RedesignedComponents/MyCard.vue";
import { useScrollFadeIn } from "~/composables/useScrollFadeIn";
import ColorSelector from "~/components/MyIntro/ColorSelector.vue";
import { ref } from "vue"; // 仅保留必要的 ref 导入

const { t } = useI18n();

const currentCardId = ref<number>(0);

// 恢复你熟悉的原始数组写法，将中文替换为 t() 翻译函数调用
const designSteps = [
  {
    id: 0,
    title: t("my_design.step0_title"), // 对应国际化配置中的键
    content: t("my_design.step0_content")
  },
  {
    id: 1,
    title: t("my_design.step1_title"),
    content: t("my_design.step1_content")
  },
  {
    id: 2,
    title: t("my_design.step2_title"),
    content: t("my_design.step2_content")
  },
  {
    id: 3,
    title: t("my_design.step3_title"),
    content: t("my_design.step3_content")
  },
];

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
      <div class="relative max-w-[1200px] container px-4 z-10 pt-20 text-slate-800 dark:text-slate-100">
        <div class="space-y-3">
          <!-- 其他文案也用 t() 调用 -->
          <p class="text-xl font-semibold text-[#3261e4] dark:text-[#4f77e6]">
            {{ t("my_design.section1_tag") }}
          </p>
          <p class="text-4xl font-semibold">
            {{ t("my_design.section1_title") }}
          </p>
          <p class="text-sm">
            {{ t("my_design.section1_desc") }}
          </p>
        </div>

<!--        <div class="flex w-full mt-10 min-h-[400px] animated-card-my-design">-->
<!--          &lt;!&ndash; 左：沿用原有的 v-for 遍历 &ndash;&gt;-->
<!--          <div class="hidden md:block w-1/2">-->
<!--            <MyCard-->
<!--                v-for="i in designSteps"-->
<!--                :key="i.id"-->
<!--                rounded="xl"-->
<!--                :padding="12"-->
<!--                class="mb-3"-->
<!--                hoverable-->
<!--                @click="currentCardId = i.id"-->
<!--            >-->
<!--              <p class="text-lg font-semibold mb-2">{{ i.title }}</p>-->
<!--              <p class="text-sm font-light">{{ i.content }}</p>-->
<!--            </MyCard>-->
<!--          </div>-->

<!--          &lt;!&ndash; 右 &ndash;&gt;-->
<!--          <div class="w-full md:w-1/2 flex flex-col justify-center items-center">-->
<!--            <SampleDesignedCard class="animated-card-my-design" />-->
<!--            &lt;!&ndash;            <ColorSelector class="animated-card-my-design" />&ndash;&gt;-->
<!--          </div>-->
<!--        </div>-->

        <div class="flex flex-col md:flex-row w-full mt-10 min-h-[400px] animated-card-my-design">
          <!-- 左：移动端 + 桌面端都显示 -->
          <div class="w-full md:w-1/2">

<!--            <div-->
<!--                class="mb-10"-->
<!--                v-for="i in designSteps"-->
<!--                :key="i.id"-->
<!--            >-->
<!--&lt;!&ndash;              <p class="text-lg font-semibold mb-2">{{ i.title }}</p>&ndash;&gt;-->
<!--              <Ulli :value="i.title" bold />-->
<!--              <p class="text-sm font-light">{{ i.content }}</p>-->
<!--            </div>-->

            <MyCard
                v-for="i in designSteps"
                :key="i.id"
                rounded="xl"
                :padding="12"
                class="mb-6"
                hoverable
                @click="currentCardId = i.id"
            >
              <div class="flex flex-row items-center space-x-4">
                <div class="text-4xl font-bold opacity-90 font-mono pl-3 pr-1 text-[#3261e4] dark:text-[#4f77e6]">
                  {{ i.id+1 }}
                </div>
                <div>
                  <p class="text-lg font-semibold mb-1">{{ i.title }}</p>
                  <p class="text-sm font-light">{{ i.content }}</p>
                </div>

              </div>


            </MyCard>
          </div>

          <!-- 右：仅桌面端显示 -->
          <div class="hidden md:flex w-full md:w-1/2 flex-col items-center justify-center">
            <SampleDesignedCard class="animated-card-my-design" />
<!--            <p class="text-2xl font-bold mt-4">Just Sample</p>-->
<!--            <TicketSample class="animated-card-my-design" />-->
          </div>
        </div>

        <div class="space-y-3 mt-12 animated-card-my-design"> <!-- 修正笔误 mt-8q -> mt-8 -->
          <p class="text-xl font-semibold text-[#3261e4] dark:text-[#4f77e6]">
            {{ t("my_design.section2_tag") }}
          </p>
          <p class="text-4xl font-semibold">
            {{ t("my_design.section2_title") }}
          </p>
          <p class="text-sm">
            {{ t("my_design.section2_desc") }}
          </p>
        </div>

        <div class="space-y-3 mt-12 animated-card-my-design">
          <p class="text-xl font-semibold text-[#3261e4] dark:text-[#4f77e6]">
            {{ t("my_design.section3_tag") }}
          </p>
          <p class="text-4xl font-semibold">
            {{ t("my_design.section3_title") }}
          </p>
          <p class="text-sm">
            {{ t("my_design.section3_desc") }}
          </p>
        </div>

      </div>
    </div>
  </div>
</template>

<style scoped>
.p-card-body {

}

:deep(.p-card-body) { /* 修正样式穿透，确保生效 */
  padding: 12px !important;
}
</style>