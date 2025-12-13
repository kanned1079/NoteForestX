<script setup lang="ts">
import { ref, defineProps, defineEmits, onMounted, onBeforeUnmount, watch } from 'vue'

const props = defineProps<{
  modelValue: boolean
  closable?: boolean
  blur?: number
  close: () => void
}>()

// ESC 关闭
function onKeydown(event: KeyboardEvent) {
  if (event.key === 'Escape') {
    props.close()
  }
}

onMounted(() => {
  window.addEventListener('keydown', onKeydown)
})

onBeforeUnmount(() => {
  window.removeEventListener('keydown', onKeydown)
})

</script>

<template>
  <transition name="fade">
    <div
        v-if="props.modelValue"
        class="fixed inset-0 z-50"
        @click.self="close"
        :style="{ backdropFilter: `blur(${props.blur ?? 0}px)`, backgroundColor: 'rgba(227,230,237,0.9)' }"
    >
      <!-- 默认居中插槽 -->
      <div class="absolute inset-0 flex justify-center items-center">
        <slot />
      </div>

      <!-- 左下角插槽 -->
      <div class="absolute left-0 bottom-0 m-4">
        <slot name="lb" />
      </div>

      <!-- 右上角插槽 -->
      <div class="absolute right-0 top-0 m-4">
        <slot name="rt" />
      </div>
    </div>
  </transition>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>