<script setup lang="ts">
import Ulli from "~/components/Ulli.vue"
import SkillRepresent from "~/components/SkillRepresent.vue"
import { useScrollFadeIn } from "~/composables/useScrollFadeIn"
import useThemeStore from "~/store/themeStore";

const themeStore = useThemeStore()

useScrollFadeIn({
  selector: '.animate-card-contact-me',
  direction: 'up',
  y: 50,
  duration: 1,
  stagger: 0.1,
  start: 'top 85%',
  toggleActions: 'play none none reverse'
})

import {openInNewTab} from "~/composables/useOpenNewTab";

</script>

<template>
  <div class="flex justify-center">
    <div class="relative max-w-[1000px] container z-10 px-4 pt-20 text-slate-800 dark:text-slate-100">

      <!-- 🌟 页面标题 -->
      <div class="mb-6">
        <p class="text-2xl font-bold text-[#3261e4] dark:text-[#4f77e6]">作品展示</p>
        <p class="text-7xl font-bold mt-5 mb-4">联系我</p>
        <p class="text-xl">您可以随时通过社交媒体、电子邮件或电话与我联系。</p>
      </div>

      <!-- 简介 -->
      <div class="animate-card-contact-me mb-6 text-gray-700 space-y-2">
        <p>友情提醒：此處提供的資訊僅供商業用途。如有任何疑問，歡迎直接透過我的社群媒體與我聯繫。</p>
        <p>感謝您瞭解並負責任地使用本產品。</p>
      </div>

      <!-- 联系方式 -->
      <div class="animate-card-contact-me">
        <p class="text-3xl font-bold mt-5 mb-4">联系方式</p>
        <div class="w-full border border-gray-300 rounded-lg overflow-hidden">
          <table class="w-full table-fixed">
            <thead>
            <tr class="border-b border-gray-300">
              <th class="w-[180px] px-4 py-3 text-left font-semibold">联系项</th>
              <th class="px-4 py-3 text-left font-semibold">详情</th>
            </tr>
            </thead>
            <tbody>
            <tr>
              <td class="px-4 py-4 align-top">
                <div class="space-y-2 text-sm">
                  <div>地址</div>
                  <div>时区</div>
                  <div>邮箱</div>
                </div>
              </td>
              <td class="px-4 py-4 align-top">
                <div class="space-y-2 text-sm">
                  <div>中国上海</div>
                  <div>UTC+8:00（中国标准时间）</div>
                  <div class="underline font-semibold cursor-pointer" @click="openInNewTab('mailto:kanned1079@gmail.com')">kanned1079@gmail.com</div>
                </div>
              </td>
            </tr>
            </tbody>
          </table>
        </div>
        <p class="mt-5 text-sm">如需进一步联系，例如获取我的电话，请先发送电子邮件给我。</p>
      </div>


      <div class="animate-card-contact-me">
        <p class="text-3xl font-bold mt-8 mb-4">社交媒体</p>
        <p class="mt-2 mb-2 text-sm">您可以通过搜索 <span class="font-bold underline">@kanned1079</span> 找到我的社交媒体账号，这个用户名几乎在所有平台都一致。</p>
        <p class="mt-2 mb-4 text-sm">如果时间紧张，我在下面列出了最常使用的社交平台链接：</p>

        <div class="w-full border border-gray-300 rounded-lg overflow-hidden animate-card-contact-me">
          <table class="w-full table-fixed">
            <thead>
            <tr class="border-b border-gray-300">
              <th class="w-[180px] px-4 py-3 text-left font-semibold">社交平台</th>
              <th class="px-4 py-3 text-left font-semibold">链接</th>
            </tr>
            </thead>
            <tbody>
            <tr>
              <td class="px-4 py-4 align-top">
                <div class="space-y-2 text-sm">
                  <div>X</div>
                  <div>Bilibili</div>
                </div>
              </td>
              <td class="px-4 py-4 align-top">
                <div class="space-y-2 text-sm font-semibold">
                  <div class="underline cursor-pointer" @click="openInNewTab('https://x.com/kanned1079')">https://x.com/kanned1079</div>
                  <div class="underline cursor-pointer" @click="openInNewTab('https://space.bilibili.com/23803212')">https://space.bilibili.com/23803212</div>
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
/* 可以增加一些整体行高优化 */
p, td, th {
  line-height: 1.6;
}
</style>