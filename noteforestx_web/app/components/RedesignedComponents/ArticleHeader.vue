<script setup lang="ts">
import type {ArticleTag} from "~/types/article";

const props = defineProps<{
  title: string
  tags: ArticleTag[]
  createdAt: string
}>()

import {useToast} from "primevue/usetoast";
const toast = useToast()

const clickTag = (tag: ArticleTag) => {
  toast.add({
    severity: "info",
    summary: tag.name,
    detail: tag.id,
    life: 4500
  });
}

</script>

<template>
  <header class="mb-6 space-y-1 animate-card-article-id">
    <p
        class="
      font-semibold
      text-xl
      sm:text-2xl
      lg:text-3xl
      leading-tight
      break-words
    "
    >
      {{ props.title }}
    </p>
    <div class="flex gap-2 flex-wrap">
      <span
          class="text-sm font-mono opacity-80 hover:underline cursor-pointer"
          v-for="i in props.tags"
          :key="i.id"
          @click="clickTag(i)"
      >{{ `#${i.name}` }}</span>
    </div>
    <p class="text-sm opacity-70">{{ props.createdAt }}</p>
  </header>
</template>

