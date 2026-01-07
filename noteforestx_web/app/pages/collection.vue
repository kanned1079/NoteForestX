<script setup lang="ts">
import {useI18n} from "vue-i18n";
// import toolslist from "~/assets/imgs/toolslistshort.png"
import {useScrollFadeIn} from "~/composables/useScrollFadeIn";
import WidthTest from "~/components/RedesignedComponents/WidthTest.vue";
import MyCard from "~/components/RedesignedComponents/MyCard.vue";

const {t} = useI18n()
// const showRibbon = ref<boolean>(false)

type WebItem = {
  title: string,
  description: string,
  url: string,
  icon_url: string
}

const myWebAppData: WebItem[] = [
  {
    title: "Kanna's Openlist",
    description: "文件目录系统，由Openlist开源项目驱动，存放我的一些分享文件等。",
    url: `https://ikanned.com:2444`,
    icon_url: `https://ikanned.com:2444/d/upload/NXC/cloudy.svg`
  },
  {
    title: "Nxc Networks 常州",
    description: "购买边缘计算服务，提供欧美、亚太地区服务。",
    url: `https://ikanned.com:12000`,
    icon_url: `https://ikanned.com:2444/d/upload/NXC/links2.svg`
  },
  {
    title: "UptimeKuma 常州",
    description: "NxcNetworks 的状态管理页面，由UptimeKuma驱动。",
    url: `https://ikanned.com:12001/status/changzhou`,
    icon_url: `https://ikanned.com:2444/d/upload/NXC/uptimekuma.svg`
  },
]

const collectionWebAppData: WebItem[] = [
  {
    title: "Naive-UI",
    description: "一个 Vue 3 组件库，比较完整，主题可调，使用 TypeScript，快。有点意思。",
    url: `https://www.naiveui.com/`,
    icon_url: `https://ikanned.com:2444/d/upload/NXC/naivelogo-BdDVTUmz.svg`
  },
  {
    title: "Uiverse",
    description: "不可错过的参考库，有很多交互效果，包括热门的 Button、Loader 设计，内容丰富，还在不断更新。",
    url: `https://uiverse.io/`,
    icon_url: `https://ikanned.com:2444/d/upload/NXC/uiverse.png`
  },
  {
    title: "MdEditor-V3",
    description: "Markdown编辑器Vue3版本，使用jsx和typescript语法开发，支持切换主题、prettier美化文本等。",
    url: `https://imzbf.github.io/md-editor-v3`,
    icon_url: `https://ikanned.com:2444/d/upload/NXC/md-editor-v3.ico`
  },
  {
    title: "Lucide 图标库",
    description: "一个开源的图标库，提供超过1000个矢量svg图标资源，支持各类前端框架中使用。",
    url: `https://lucide.dev/`,
    icon_url: `https://ikanned.com:2444/d/upload/NXC/lucide.svg`
  },
  {
    title: "CmdNotFound",
    description: "当你有什么命令无法使用的时候，查看在各个系统上怎么安装。",
    url: `https://command-not-found.com/`,
    icon_url: `https://ikanned.com:2444/d/upload/NXC/comamndnotfound.ico`
  },
  {
    title: "IT Tools",
    description: "包含各类IT开发中需要的工具，如Token生成、Hash、Bcrypt、Diff、JSON、XML格式化等",
    url: `https://it-tools.tech/`,
    icon_url: `https://ikanned.com:2444/d/upload/NXC/ittools.ico`
  },
  // https://ikanned.com:2444/d/upload/NXC/mikutools.ico
  {
    title: "MikuTools",
    description: "有各种互联网上的小工具，图片生成，格式转换、番茄时钟等、新闻聚合、JS压缩等。",
    url: `https://tools.miku.ac/`,
    icon_url: `https://ikanned.com:2444/d/upload/NXC/mikutools.ico`
  },
  {
    title: "Anna's Archieve",
    description: "人類史上規模最大的真正開放式圖書館。截止2025年底拥有 61,654,285 本圖書，95,687,150 份文獻被永久保存。",
    url: `https://annas-archive.se/`,
    icon_url: `https://ikanned.com:2444/d/upload/NXC/annasarchieve.ico`
  },
  {
    title: "絵師100人展",
    description: "本次展覽將匯集100多位活躍於流行文化前沿的藝術家的原創作品。透過他們作品中展現的高超技藝和獨特世界觀。",
    url: `https://www.eshi100.com/`,
    icon_url: `https://ikanned.com:2444/d/upload/NXC/eshi100.svg`
  },
  {
    title: "Windows KMS Keys",
    description: "微軟官方各個Windows版本的KMS密鑰，好難搜到",
    url: `https://learn.microsoft.com/en-us/windows-server/get-started/kms-client-activation-keys`,
    icon_url: `https://ikanned.com:2444/d/upload/NXC/micrtosoft.ico`
  }

]

useScrollFadeIn({
  selector: '.animate-card-collection',
  y: 60,
  stagger: 0.15
})

</script>

<template>
  <div class="mt-4 pb-60">
    <div class="container mx-auto max-w-[1200px] px-4 space-y-6">

      <section class="space-y-2 animate-card-collection">
        <p class="text-4xl font-bold opacity-90">{{ t('collection.title') }}</p>
        <p>
          {{ t('collection.description') }}
        </p>
      </section>

      <div class="animate-card-collection space-y-2">
        <p class="text-lg font-bold">{{ t('collection.my') }}</p>
        <section class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
          <MyCard
              v-for="i in myWebAppData"
              :key="i.url"
              rounded="xl"
              :padding="20"
              hoverable
              @click="openInNewTab(i.url)"
          >
            <div class="flex items-center h-[40px] gap-3">
              <img :src="i.icon_url" class="h-[40px] rounded-md"/>
              <p class="text-xl font-bold leading-none">{{ i.title }}</p>
            </div>
            <div class="space-y-2 overflow-hidden mt-4">
              <p
                  class="underline font-mono text-xs break-all opacity-80 cursor-pointer"
                  :href="i.url"
              >
                {{ i.url }}
              </p>
              <p class="opacity-90 text-sm">{{ i.description }}</p>
            </div>

          </MyCard>
        </section>
      </div>

      <div class="animate-card-collection space-y-2">
        <p class="text-lg font-bold">{{ t('collection.collected') }}</p>
        <section class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">

          <MyCard
              v-for="i in collectionWebAppData"
              :key="i.url"
              rounded="xl"
              :padding="20"
              hoverable
              @click="openInNewTab(i.url)"
          >
            <div class="flex items-center h-[40px] gap-3">
              <img :src="i.icon_url" class="h-[40px] rounded-md"/>
              <p class="text-xl font-bold leading-none">{{ i.title }}</p>
            </div>
            <div class="space-y-2 overflow-hidden mt-4">
              <p
                  class="underline font-mono text-xs break-all opacity-80 cursor-pointer"
                  :href="i.url"
              >
                {{ i.url }}
              </p>
              <p class="opacity-90 text-sm">{{ i.description }}</p>
            </div>

          </MyCard>
        </section>
      </div>


    </div>
  </div>
</template>

<style scoped>

</style>