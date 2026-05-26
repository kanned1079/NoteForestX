<script setup lang="ts">
import { useI18n } from "vue-i18n";
import toolslist from "~/assets/imgs/toolslist.png"
import {useScrollFadeIn} from "~/composables/useScrollFadeIn";

const { t } = useI18n()

useScrollFadeIn({
  selector: '.animate-card-sponsor',
  y: 60,
  stagger: 0.15
})
</script>

<template>
  <div class="mt-8">

    <div class="container mx-auto max-w-[980px] px-4 space-y-6">

      <!-- 页面标题 -->
      <!--      <div class="opacity-85">關於本網站的一些技術細節</div>-->

      <!-- 背景與目標 -->
      <section class="space-y-2 animate-card-sponsor">
        <h2 class="text-xl font-bold opacity-90">{{ t('details.background_title') }}</h2>
        <p>
          {{ t('details.background_desc') }}
        </p>
      </section>

      <section class="space-y-2 animate-card-sponsor">
        <h2 class="text-xl font-bold opacity-90">{{ t('layout.sponsor') }}</h2>
        <p class="font-mono">0xaa53b6eb01103118a909a018a8c24d480e47bf7e <span class="font-bold">(Ethereum)</span></p>
        <p class="font-mono">0xaa53b6eb01103118a909a018a8c24d480e47bf7e <span class="font-bold">(Arbitrum)</span></p>
      </section>

    </div>
  </div>
</template>

<style scoped>

</style>