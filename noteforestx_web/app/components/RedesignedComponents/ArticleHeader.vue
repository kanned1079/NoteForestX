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
      {{ title }}
    </p>
    <div class="flex gap-2 flex-wrap">
      <span
          class="text-sm font-mono opacity-80"
          v-for="i in tags" :key="i.id"
      >{{ `#${i.name}` }}</span>
    </div>
    <p class="text-sm opacity-70">{{ createdAt }}</p>
  </header>
</template>

<script setup lang="ts">
defineProps<{
  title: string
  tags: Array<{ id: number | string; name: string }>
  createdAt: string
}>()
</script>