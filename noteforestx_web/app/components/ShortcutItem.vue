<script setup lang="ts">
import Kbd from "~/components/MyIntro/Kbd.vue"

const props = defineProps<{
  label: string
  keyLabels: string[]
  pressType?: 'or' | 'and'
}>()

const separator = computed(() =>
    props.pressType === 'and' ? '+' : '/'
)
</script>

<template>
  <div class="flex flex-col w-full">
    <div class="flex flex-row justify-between items-center ">
      <!-- 左侧描述 -->
      <p class="text-sm text-gray-600 dark:text-gray-300">
        {{ label }}
      </p>

      <!-- 右侧快捷键 -->
      <span class="flex items-center gap-1 font-mono text-sm">
        <template v-for="(key, index) in keyLabels" :key="key">
          <Kbd :key-label="key" size="sm" />

          <!-- 分隔符 -->
          <span
              v-if="index < keyLabels.length - 1"
              class="px-1 text-gray-400"
          >
            {{ separator }}
          </span>
        </template>
      </span>
    </div>

    <Divider class="my-2" />
  </div>
</template>

<style scoped>
</style>