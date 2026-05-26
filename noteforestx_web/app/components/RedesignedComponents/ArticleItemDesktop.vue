<script setup lang="ts">
import type { Article } from "~/types/article";
import dayjs from "dayjs";
// Nuxt 3 自动导入常用 Composition API，不需要手动 import defineProps 等

const props = defineProps<{
  article: Article
}>()

const emits = defineEmits<{
  (e: 'clickTag', tag: { id: string, name: string }): void
  (e: 'clickTitle', article: Article): void
}>()

const toDetails = () => {
  emits('clickTitle', props.article)
}

const handleTagClick = (tag: { id: string, name: string }) => {
  emits('clickTag', tag)
}
</script>

<template>
  <div class="w-full">
    <!-- 桌面布局：lg (1024px) 以上顯示，以下隱藏 -->
    <div class="hidden lg:flex flex-row justify-between items-center rounded-lg py-1.5 px-0 transition-all duration-200 hover:bg-slate-100/60 dark:hover:bg-slate-800/30 hover:translate-x-1 cursor-pointer" @click="toDetails">
      <div class="flex items-center space-x-4 relative">
        <span class="text-sm opacity-60 w-24 flex-shrink-0 select-none">{{ dayjs(article.created_at).format('YYYY-MM-DD') }}</span>
        <span class="relative flex items-center group select-none">
          <span class="absolute left-0 w-2 h-2 rounded-full opacity-0 group-hover:opacity-100 transition-opacity duration-200 bg-[#7234e9] dark:bg-[#8257f2]"></span>
          <span class="font-medium ml-2 relative transition-transform duration-200 group-hover:translate-x-2 flex flex-row items-center">
            <Tag v-if="article.top" severity="warn" class="mr-1 text-xs" value="Pinned" style="padding: 0 8px" />
            {{ article.title }}
            <span class="absolute left-0 bottom-0 h-[1px] w-0 group-hover:w-full transition-all duration-200 bg-[#7234e9] dark:bg-[#8257f2]"></span>
          </span>
        </span>
      </div>

      <div class="space-x-2 opacity-60 flex flex-wrap">
        <span v-for="tag in article.tags" :key="tag.id" class="text-sm hover:underline cursor-pointer" @click.stop="handleTagClick(tag)">
          {{ `#${tag.name}` }}
        </span>
        <span v-if="article.tags.length === 0" class="text-sm">---</span>
      </div>
    </div>

    <!-- 移动端布局：lg (1024px) 以上隐藏 -->
    <div class="flex lg:hidden flex-col justify-start space-y-1">
      <span @click="toDetails" class="font-bold hover:underline cursor-pointer">
        {{ article.title }}
        <Tag v-if="article.top" severity="warn" class="ml-1 text-xs" value="Pinned" style="padding: 0 8px" />
      </span>

      <div class="flex flex-wrap gap-2">
        <span v-for="tag in article.tags" :key="tag.id" class="text-xs opacity-70 hover:underline cursor-pointer" @click.stop="handleTagClick(tag)">
          {{ `#${tag.name}` }}
        </span>
        <span v-if="article.tags.length === 0" class="text-sm">---</span>
      </div>
      <span class="text-xs opacity-60">{{ dayjs(article.created_at).format('YYYY-MM-DD') }}</span>
    </div>
  </div>
</template>
