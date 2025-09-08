<script setup lang="ts">
import {IllustrationItem} from "../types/illustration";
const config = useRuntimeConfig()
const props = defineProps<{
  illustration: IllustrationItem
}>()

const onClickIllItem = (id: string) => {
  navigateTo({path: `/illustration/${id}`})
}

const openNewTab = (url: string) => {
  window.open(url, '_blank', 'noopener noreferrer');
};

/*
原創 コミティア 藤の花 ボブカット 浴衣 オリジナル30000users入り 花と女の子
*
* */

</script>

<template>
  <div class="flex flex-col mb-5">
    <!-- 上半部分：正方形 + 背景 + 圆角 -->
    <div
        class="aspect-square w-full rounded-md overflow-hidden
         bg-gray-100 dark:bg-[#222222]
         flex items-center justify-center
         transition-shadow duration-200
         drop-shadow-sm
         hover:shadow-xl"
        @click="onClickIllItem(props.illustration.id)"
    >
      <img
          :src="`${config.public.apiBase}/illustration/file/${props.illustration.file_path}?size=small`"
          alt="illustration"
          class="w-full h-full object-contain"
          loading="lazy"
      />
    </div>

    <div class="mt-2 px-1">
      <div
          class="font-medium text-lg truncate hover:underline"
          @click="openNewTab(props.illustration.link)"
      >
        {{ props.illustration.name || 'Untitled Illustration' }}
      </div>
      <div
          class="text-sm text-gray-600 dark:text-gray-300 truncate hover:underline"
          @click="openNewTab(props.illustration.author.link)"
      >
        {{ props.illustration.author.name || 'Unknown Author' }}
      </div>
    </div>
  </div>
</template>

<style>
</style>