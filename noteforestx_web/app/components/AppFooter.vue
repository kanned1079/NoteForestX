<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {openInNewTab} from "~/composables/useOpenNewTab";

const {t} = useI18n()
const links = [
  {link: 'https://github.com/kanned1079', icon: 'github'},
  {link: 'https://x.com/kanned1079', icon: 'twitter'},
  {link: 'https://www.instagram.com/kanned1079/', icon: 'instagram'},
]

const siteLinks = [
  {link: 'https://github.com/kanned1079/NoteForestX', label: 'app_footer.source_code'},
  {link: 'https://uiverse.io/', label: 'app_footer.uiverse'},
  {link: 'https://primevue.org/', label: 'app_footer.primevue'}
]

</script>

<template>
  <div class="relative w-full overflow-hidden pt-[160px] pb-10">
    <!-- 背景 -->
    <BackgroundGrid :flipVertical="true" />

    <!-- ================= 主内容区 ================= -->
    <div class="relative z-10 w-full flex justify-center">
      <div
          class="
          w-full
          max-w-[1200px]
          grid
          grid-cols-1
          md:grid-cols-2
          gap-8
          md:gap-12
          px-4
        "
      >
        <!-- 左侧 -->
        <div class="flex flex-col justify-start items-start opacity-80">
          <div class="text-base font-bold mb-2">
            {{ t('app_footer.about_me') }}
          </div>
<!--          <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">-->
<!--            I'm kanned1079, a front-end developer who loves intuitive,-->
<!--            clean and modern UI design.-->
<!--          </p>-->
          <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">
            <i18n-t keypath="home.intro_desc" tag="span">
              <template #name>
                <span class="text-[#1e88a8] underline">kanned1079</span>
              </template>
              <template #role>
<!--                <span class="font-bold">Front-End Developer</span>-->
                <span class="font-bold">{{ t('home.fullstack_developer') }}</span>
              </template>
            </i18n-t>
          </p>
          <div class="mt-2 space-x-4 flex flex-row items-center justify-start">
            <span
                v-for="i in links"
                :key="i.link"
                :class="['pi', `pi-${i.icon}`, 'text-2xl', 'cursor-pointer']"
                @click="openInNewTab(i.link)"
            ></span>
          </div>
        </div>

        <!-- 右侧 -->
        <div class="flex flex-col justify-start opacity-80">
          <div class="text-base font-bold mb-2">
            {{ t('app_footer.this_site') }}
          </div>
          <div class="mt-2 space-y-2 text-sm text-gray-800 dark:text-gray-300">
            <div
                v-for="i in siteLinks"
                :key="i.link"
                class="flex flex-row items-center justify-start space-x-1 hover:underline"
                @click="openInNewTab(i.link)"
            >
              <p>{{ t(i.label) }}</p>
              <span class="pi pi-external-link text-xs font-semibold"></span>
            </div>
          </div>

        </div>
      </div>
    </div>

    <!-- ================= Footer ================= -->
    <footer class="relative z-10 mt-10 w-full flex justify-center px-4">
      <div class="w-full max-w-[1200px] text-center text-sm text-gray-500">
        © 2026 NoteForestX · By kanned1079
      </div>
    </footer>
  </div>
</template>


<style scoped>
/* 如需微调可在这里加 */
</style>