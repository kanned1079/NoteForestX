<script setup lang="ts">
import Ulli from "~/components/Ulli.vue"
import SkillRepresent from "~/components/SkillRepresent.vue"
import { useScrollFadeIn } from "~/composables/useScrollFadeIn"

useScrollFadeIn({
  selector: '.animate-card-mywork',
  direction: 'up',      // 默认其实可以不写
  y: 50,
  duration: 1,
  stagger: 0.1,
  start: 'top 85%',
  toggleActions: 'play none none reverse'
})
</script>

<template>
  <div class="flex justify-center">
    <div class="relative max-w-[1000px] container z-10 px-4 pt-20 text-slate-800 dark:text-slate-100">

<!--      <Button variant="text" class="mt-4 mb-4" label="技能 & 工具" icon="pi pi-wrench" />-->

<!--      <p class="mb-4 underline">轉到 <span class="font-bold">技能 & 工具</span> <span class="pi pi-wrench"></span> </p>-->

      <!-- 🌟 作品标题 -->
      <div class="mb-6">
        <p class="text-2xl font-bold text-[#3261e4] dark:text-[#4f77e6]">作品</p>
        <p class="text-7xl font-bold mt-5 mb-4">技能與工具</p>
        <p class="text-xl">瀏覽我所使用過的程式語言、函式庫以及工具。</p>
      </div>

<!--      <Divider class="mt-6 mb-6" />-->

      <!-- 简介 -->
      <div class="animate-card-mywork mb-6">
        <p>我從事網頁開發已有約九年，嘗試過多種程式語言與技術堆疊，涵蓋後端與前端開發。雖然網頁開發領域廣泛，但我對前端開發與 UI/UX 設計特別感興趣並專注其中。</p>
      </div>

      <!-- App Developer -->
      <div class="animate-card-mywork mb-6">
        <p class="text-3xl font-bold mt-5 mb-4">應用程式開發者</p>
        <p class="mt-2 mb-2">我喜歡創造乾淨、極簡、美觀且具有價值的作品，當然也要易於使用。</p>
        <p class="mt-2 mb-2">我特別享受製作富有 UI 元件的網站，包括：</p>

        <Ulli value="Web 應用程式" bold />
        <Ulli value="文件與說明頁面" bold />
        <Ulli value="CMS 內容佈局" bold />
        <Ulli value="後台儀表板布局" bold />
        <Ulli value="以及其他" />
      </div>

      <div class="animate-card-mywork mb-6">
        <p class="mt-2 mb-2">但我仍然喜歡創建簡單的網頁，例如登陸頁面。在網站製作過程中，我最熟悉使用的工具有哪些呢？</p>
      </div>

      <!-- 技術棧 -->
      <div class="animate-card-mywork mb-6">
        <p class="text-2xl font-bold mt-5 mb-4">Vue 與 Nuxt.js 搭配 TypeScript</p>
        <p class="mt-2 mb-2 text-slate-800 dark:text-slate-400">
          單頁網站是我的專長，而 <span class="font-bold">Vue.js</span> 就像我的魔法棒。我已經使用它多年，創作的作品越來越好看。
        </p>
        <p class="mt-2 mb-2">
          首先是程式語言，我習慣使用 TypeScript，雖然 PHP、Python 也可以用於網頁開發。TypeScript 已經陪伴我約七年。
        </p>
      </div>

      <div class="animate-card-mywork mb-6">
        <p class="text-2xl font-bold mt-5 mb-4">Golang 與 Nest.js</p>
        <p class="mt-2 mb-2 text-slate-800 dark:text-slate-400">
          我用它們構建可擴展、高效能的後端服務，從 API 設計到資料操作皆涵蓋。起初專注於後端開發，但後來我意識到 CLI 無法滿足所有需求，因此開始學習 GUI 開發。
        </p>
      </div>

      <div class="animate-card-mywork mb-6">
        <p class="text-3xl font-bold mt-5 mb-4">運維經驗</p>
        <p class="mt-2 mb-2">
          最初源於對 IT 行業的濃厚興趣，很早就接觸 Linux，經過幾年的學習，目前已經使用虛擬化與 Linux 部署了多個服務，包括但不限於邊緣計算服務、檔案目錄程式、個人 Blog、Zabbix 監控、SoftEtherVPN 等。
        </p>
      </div>

<!--      <Divider class="mt-10 mb-10 animate-card" />-->

      <!-- Detail and Summary -->
      <div class="animate-card-mywork mb-6">
        <p class="text-3xl font-bold mt-5 mb-4">細節與摘要</p>
        <p class="mt-2 mb-2">我使用標籤呈現所有資料以便閱讀。底部橫線指示使用頻率，例如：</p>
        <div class="flex flex-wrap gap-2 justify-start">
          <SkillRepresent value="經常使用" :percent="90" />
          <SkillRepresent value="偶爾使用" :percent="23" />
        </div>
      </div>

      <!-- Programming Languages -->
      <div class="animate-card-mywork mb-6">
        <p class="text-2xl font-bold mt-5 mb-4">程式語言</p>
        <div class="flex flex-wrap gap-2 justify-start">
          <SkillRepresent value="TypeScript" :percent="100" />
          <SkillRepresent value="JavaScript" :percent="90" />
          <SkillRepresent value="Golang" :percent="100" />
          <SkillRepresent value="Java" :percent="30" />
          <SkillRepresent value="Shell" :percent="50" />
        </div>
      </div>

      <!-- Technologies -->
      <div class="animate-card-mywork mb-6">
        <p class="text-2xl font-bold mt-5 mb-4">技術與工具</p>
        <div class="flex flex-wrap gap-2 justify-start">
          <SkillRepresent value="Tailwind CSS" :percent="90" />
          <SkillRepresent value="NaiveUI" :percent="50" />
          <SkillRepresent value="PrimeVue" :percent="60" />
          <SkillRepresent value="MaterialUI" :percent="30" />
          <SkillRepresent value="Vue" :percent="90" />
          <SkillRepresent value="Nuxt.js" :percent="80" />
          <SkillRepresent value="Electron" :percent="10" />
          <SkillRepresent value="pinia" :percent="90" />
          <SkillRepresent value="React" :percent="30" />
          <SkillRepresent value="GIN" :percent="100" />
          <SkillRepresent value="gORM" :percent="100" />
          <SkillRepresent value="gRPC" :percent="20" />
          <SkillRepresent value="Redis" :percent="80" />
          <SkillRepresent value="MySQL/MariaDB" :percent="80" />
          <SkillRepresent value="MongoDB" :percent="10" />
          <SkillRepresent value="Docker" :percent="95" />
          <SkillRepresent value="Nginx" :percent="80" />
        </div>
      </div>

      <!-- Softwares and Tools -->
      <div class="animate-card-mywork mb-6">
        <p class="text-2xl font-bold mt-5 mb-4">軟體與工具</p>
        <div class="flex flex-wrap gap-2 justify-start">
          <SkillRepresent value="JetBrains IntelliJ IDEA" :percent="95" />
          <SkillRepresent value="Visual Studio Code" :percent="80" />
          <SkillRepresent value="Vim" :percent="40" />
          <SkillRepresent value="Git" :percent="80" />
          <SkillRepresent value="MySQL Workbench" :percent="70" />
          <SkillRepresent value="Adobe PhotoShop" :percent="80" />
          <SkillRepresent value="Adobe Illustrator" :percent="30" />
        </div>
      </div>

    </div>
  </div>
</template>


<style scoped>

</style>