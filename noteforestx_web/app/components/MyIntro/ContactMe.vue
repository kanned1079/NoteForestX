<script setup lang="ts">
import Ulli from "~/components/Ulli.vue"
import SkillRepresent from "~/components/SkillRepresent.vue"
import { useScrollFadeIn } from "~/composables/useScrollFadeIn"
import useThemeStore from "~/store/themeStore";
import { useI18n } from "vue-i18n"

const themeStore = useThemeStore()
const { t } = useI18n()

const contactInfo = {
  address: "中国上海",
  email: "kanned1079@gmail.com",
  timeZone: t("contact_me.contact_detail2")
}

const socialInfo = {
  username: "@kanned1079",
  platforms: [
    { name: "X", link: "https://x.com/kanned1079" },
    { name: "Bilibili", link: "https://space.bilibili.com/23803212" }
  ]
}

useScrollFadeIn({
  selector: '.animate-card-contact-me',
  direction: 'up',
  y: 50,
  duration: 1,
  stagger: 0.1,
  start: 'top 85%',
  toggleActions: 'play none none reverse'
})

import { openInNewTab } from "~/composables/useOpenNewTab";

</script>

<template>
  <div class="flex justify-center">
    <!-- 页面文字基准色 -->
    <div
        class="relative max-w-[1200px] container z-10 px-4 pt-20
             text-slate-800 dark:text-slate-100"
    >

      <!-- 🌟 页面标题 -->
      <div class="mb-6">
        <p class="text-2xl font-bold text-[#3261e4] dark:text-[#4f77e6]">
          {{ t("contact_me.title_tag") }}
        </p>
        <p class="text-7xl font-bold mt-5 mb-4">
          {{ t("contact_me.main_title") }}
        </p>
        <p class="text-xl opacity-85">
          {{ t("contact_me.main_desc") }}
        </p>
      </div>

      <!-- 简介 -->
      <div class="animate-card-contact-me mb-6 space-y-2 opacity-85">
        <p>{{ t("contact_me.intro_desc1") }}</p>
        <p>{{ t("contact_me.intro_desc2") }}</p>
      </div>

      <!-- 联系方式 -->
      <div class="animate-card-contact-me">
        <p class="text-3xl font-bold mt-5 mb-4">
          {{ t("contact_me.contact_info_title") }}
        </p>

        <div class="w-full border border-gray-300 dark:border-gray-600 rounded-lg overflow-hidden">
          <table class="w-full table-fixed">
            <thead>
            <tr class="border-b border-gray-300 dark:border-gray-600">
              <th class="w-[180px] px-4 py-3 text-left font-semibold">
                {{ t("contact_me.table_header1") }}
              </th>
              <th class="px-4 py-3 text-left font-semibold">
                {{ t("contact_me.table_header2") }}
              </th>
            </tr>
            </thead>
            <tbody class="opacity-85">
            <tr>
              <td class="px-4 py-4 align-top">
                <div class="space-y-2 text-sm">
                  <div>{{ t("contact_me.contact_item1") }}</div>
                  <div>{{ t("contact_me.contact_item2") }}</div>
                  <div>{{ t("contact_me.contact_item3") }}</div>
                </div>
              </td>
              <td class="px-4 py-4 align-top">
                <div class="space-y-2 text-sm">
                  <div>{{ contactInfo.address }}</div>
                  <div>{{ contactInfo.timeZone }}</div>
                  <div
                      class="underline font-semibold cursor-pointer
                             text-slate-900 dark:text-slate-50"
                      @click="openInNewTab('mailto:' + contactInfo.email)"
                  >
                    {{ contactInfo.email }}
                  </div>
                </div>
              </td>
            </tr>
            </tbody>
          </table>
        </div>

        <p class="mt-5 text-sm opacity-70">
          {{ t("contact_me.contact_remark") }}
        </p>
      </div>

      <!-- 社交媒体 -->
      <div class="animate-card-contact-me">
        <p class="text-3xl font-bold mt-8 mb-4">
          {{ t("contact_me.social_media_title") }}
        </p>

        <p class="mt-2 mb-2 text-sm opacity-85">
          {{ t("contact_me.social_desc1", { username: socialInfo.username }) }}
        </p>
        <p class="mt-2 mb-4 text-sm opacity-85">
          {{ t("contact_me.social_desc2") }}
        </p>

        <div class="w-full border border-gray-300 dark:border-gray-600 rounded-lg overflow-hidden">
          <table class="w-full table-fixed">
            <thead>
            <tr class="border-b border-gray-300 dark:border-gray-600">
              <th class="w-[180px] px-4 py-3 text-left font-semibold">
                {{ t("contact_me.social_table_header1") }}
              </th>
              <th class="px-4 py-3 text-left font-semibold">
                {{ t("contact_me.social_table_header2") }}
              </th>
            </tr>
            </thead>
            <tbody class="opacity-85">
            <tr>
              <td class="px-4 py-4 align-top">
                <div class="space-y-2 text-sm">
                  <div
                      v-for="(platform, index) in socialInfo.platforms"
                      :key="index"
                  >
                    {{ platform.name }}
                  </div>
                </div>
              </td>
              <td class="px-4 py-4 align-top">
                <div class="space-y-2 text-sm font-semibold">
                  <div
                      v-for="(platform, index) in socialInfo.platforms"
                      :key="index"
                      class="underline cursor-pointer
                             text-slate-900 dark:text-slate-50"
                      @click="openInNewTab(platform.link)"
                  >
                    {{ platform.link }}
                  </div>
                </div>
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>

    </div>
  </div>
</template>

<style scoped>
p, td, th {
  line-height: 1.6;
}
</style>