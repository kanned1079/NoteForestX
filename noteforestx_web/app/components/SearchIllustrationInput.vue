<template>
  <IconField class="w-auto">
    <InputIcon :class="currentIcon" />
    <InputText
        class="w-full"
        size="large"
        v-model="searchIll.content"
        placeholder="Search for something..."
        @input="handleInput"
    />
  </IconField>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'

const searchIll = ref({ content: '' })

// 根据前缀动态返回图标
const currentIcon = computed(() => {
  const val = searchIll.value.content.trim()
  if (val.startsWith('/user')) return 'pi pi-user'
  if (val.startsWith('/tag')) return 'pi pi-hashtag'
  if (val.startsWith('/name')) return 'pi pi-palette'
  if (val.startsWith('/id')) return 'pi pi-file'
  return 'pi pi-search'
})

// 可选：如果想在输入改变时立即更新图标
function handleInput() {
  // 当前图标由 computed 自动更新，无需额外逻辑
}
</script>