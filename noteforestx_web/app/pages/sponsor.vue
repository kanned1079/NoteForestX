<script setup lang="ts">
import toolslist from "~/assets/imgs/toolslist.png"
import {useScrollFadeIn} from "~/composables/useScrollFadeIn";
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
        <h2 class="text-xl font-bold opacity-90">背景與目標</h2>
        <p>
          本博客的目的是記錄日常學習、技術分享與項目經驗。我希望透過這個項目實踐全棧開發、
          系統設計與 DevOps 部署中涉及的技術。在該博客的舊版本中，部分功能設計較為簡陋，
          因此在全新的 <span class="font-mono">NoteForest X</span> 版本中，我修復了舊版存在的 Bug，
          並擴展了一些新的功能。
        </p>
      </section>

      <section class="space-y-2 animate-card-sponsor">
        <h2 class="text-xl font-bold opacity-90">虛擬貨幣</h2>
        <p class="font-mono">0xaa53b6eb01103118a909a018a8c24d480e47bf7e <span class="font-bold">(Ethereum)</span></p>
        <p class="font-mono">0xaa53b6eb01103118a909a018a8c24d480e47bf7e <span class="font-bold">(Arbitrum)</span></p>
      </section>

    </div>
  </div>
</template>

<style scoped>

</style>