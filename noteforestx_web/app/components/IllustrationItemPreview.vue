<script setup lang="ts">
import { useI18n } from "vue-i18n";
import type { Illustration } from "../types/illustration";
const { locale } = useI18n();
const config = useRuntimeConfig();
const props = defineProps<{
  illustration: Illustration;
  admin?: boolean
}>();

const onClickIllItem = (id: string) => {
  if (!props.admin) {
    navigateTo({
      path: `/${locale.value}/illustration/${id}`,
    });
  }
};

const openNewTab = (url: string) => {
  window.open(url, "_blank", "noopener noreferrer");
};
</script>

<template>
  <div class="flex flex-col mb-5">
    <!-- 正方形容器 -->
    <div
        class="relative aspect-square w-full rounded-md overflow-hidden
             flex items-center justify-center
             transition-shadow duration-200
             drop-shadow-sm hover:shadow-xl cursor-pointer"
        @click="onClickIllItem(props.illustration?.id)"
    >
      <!-- 背景模糊图 -->
      <img
          v-if="props.illustration?.images?.length"
          :src="`${config.public.apiBase}/api/v1/illustration/file/${props.illustration?.images[0].file_path}?size=small`"
          alt="blur-background"
          class="absolute inset-0 w-full h-full object-cover scale-110 filter blur-sm brightness-100 opacity-80"
      />

      <!-- 前景正常插画 -->
      <img
          v-if="props.illustration?.images?.length"
          :src="`${config.public.apiBase}/api/v1/illustration/file/${props.illustration?.images[0].file_path}?size=small`"
          alt="illustration"
          class="relative w-full h-full object-contain"
          loading="lazy"
      />
    </div>

    <!-- 标题 + 作者 -->
    <div class="mt-2 px-1">
      <div
          class="font-medium text-lg truncate hover:underline"
          @click="openNewTab(props.illustration?.link)"
      >
        {{ props.illustration?.name || "Untitled Illustration" }}
      </div>
      <div
          class="text-sm text-gray-600 dark:text-gray-300 truncate hover:underline"
          @click="openNewTab(props.illustration?.author.link)"
      >
        {{ props.illustration?.author.name || "Unknown Author" }}
      </div>
    </div>
  </div>
</template>

<style>
/* 无需额外 CSS，Tailwind 完全能实现 */
</style>