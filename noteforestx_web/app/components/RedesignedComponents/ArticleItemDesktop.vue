<script setup lang="ts">
import type { Article } from "~/types/article";
import dayjs from "dayjs";
import { defineProps, defineEmits, ref, onMounted, onUnmounted } from "vue";
import { Pin } from 'lucide-vue-next';

const props = defineProps<{
  article: Article
}>()

const emits = defineEmits<{
  (e: 'clickTag', tag: { id: string, name: string }): void
  (e: 'clickTitle', article: Article): void
}>()

const isDesktop = ref(window.innerWidth >= 1024)

const updateWidth = () => {
  isDesktop.value = window.innerWidth >= 1024
}

onMounted(() => {
  window.addEventListener('resize', updateWidth)
})

onUnmounted(() => {
  window.removeEventListener('resize', updateWidth)
})

const toDetails = () => {
  emits('clickTitle', props.article)
}

const handleTagClick = (tag: { id: string, name: string }) => {
  emits('clickTag', tag)
}
</script>

<template>
  <div v-if="isDesktop" class="flex flex-row justify-between items-center rounded-md transition">
    <!-- 桌面布局 -->
    <div class="flex items-center space-x-4 relative">
      <span class="text-sm opacity-60">{{ dayjs(article.created_at).format('YYYY-MM-DD') }}</span>

      <span class="relative flex items-center group cursor-pointer select-none" @click="toDetails">
      <!-- 小球 -->
      <span
          class="absolute left-0 w-2 h-2 rounded-full opacity-0 group-hover:opacity-100 transition-opacity duration-200"
          :class="'bg-[#7234e9] dark:bg-[#8257f2]'"
      ></span>

        <!-- 标题 -->
      <span
          class="font-medium ml-2 relative transition-transform duration-200 group-hover:translate-x-2 flex flex-row items-center"
      >

        <span v-if="props.article.top" class="font-mono mr-1">
          <Tag severity="warn" class="font-mono text-xs font-semibold" value="Pinned" style="padding: 0 8px"></Tag>
        </span>

        {{ article.title }}
        <!-- 底部横线 -->
        <span
            class="absolute left-0 bottom-0 h-[1px] w-0 group-hover:w-full transition-all duration-200"
            :class="'bg-[#7234e9] dark:bg-[#8257f2]'"
        ></span>
      </span>
    </span>
    </div>

    <div class="space-x-2 opacity-60 flex flex-wrap">
    <span
        class="text-sm hover:underline cursor-pointer"
        v-for="tag in article.tags"
        :key="tag.id"
        @click.stop="handleTagClick(tag)"
    >
      {{ `#${tag.name}` }}
    </span>
      <span class="text-sm" v-if="article.tags.length===0">---</span>
    </div>
  </div>

  <div v-else class="flex flex-col">
    <!-- 移动端布局 -->
    <div class="flex flex-col justify-start space-y-1">
      <span @click="toDetails" class="font-bold hover:underline cursor-pointer">{{ article.title }}
       <span v-if="props.article.top" class="font-mono mr-1">
          <Tag severity="warn" class="font-mono text-xs font-semibold" value="Pinned" style="padding: 0 8px"></Tag>
        </span>

      </span>

      <div class="flex flex-wrap gap-2">
        <span
            class="text-xs opacity-70 hover:underline cursor-pointer"
            v-for="tag in article.tags"
            :key="tag.id"
            @click.stop="handleTagClick(tag)">
        {{ `#${tag.name}` }}
        </span>
        <span class="text-sm" v-if="article.tags.length===0">---</span>
      </div>
      <span class="text-xs opacity-60">{{ dayjs(article.created_at).format('YYYY-MM-DD') }}</span>
    </div>
  </div>
</template>