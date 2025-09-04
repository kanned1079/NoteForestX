<script setup lang="ts">
import useThemeStore from "~/store/themeStore";
import CommonHeaderLeft from "~/components/CommonHeaderLeft.vue";
import CommonHeaderRight from "~/components/CommonHeaderRight.vue";
const themeStore = useThemeStore()
const { t } = useI18n()
</script>

<template>
  <div class="flex flex-col min-h-screen bg-gray-50 dark:bg-gray-900">
    <!-- 顶部导航栏 -->
    <header
        class="flex items-center justify-between h-14 px-4
             bg-white/70 dark:bg-gray-800/70 backdrop-blur-md shadow-sm"
    >
      <div>
        <CommonHeaderLeft />
      </div>
      <div>
        <CommonHeaderRight />
      </div>
    </header>

    <div class="flex flex-1">
      <!-- 抽屉/侧边栏 -->
      <aside
          v-show="themeStore.isMenuDisplay"
          class="w-72 bg-white dark:bg-gray-800 shadow-md p-4 fixed inset-y-0 left-0 z-40 md:static md:block"
      >
        aside
      </aside>

      <!-- 主内容区 -->
      <main class="flex-1 ml-0 md:ml-72 p-4 transition-all">
        <div class="max-w-screen-xl mx-auto">
          <NuxtPage />
        </div>
      </main>
    </div>
  </div>
</template>

<style scoped>
/* 让小屏幕时自动隐藏侧边栏 */
@media (max-width: 768px) {
  aside {
    position: fixed;
    left: -100%;
    transition: left 0.3s;
  }
  aside[v-show="true"] {
    left: 0;
  }
}
</style>