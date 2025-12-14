<script setup lang="ts">
import { onMounted, ref, nextTick } from 'vue'
import useShortcutsStore from '~/store/shortcutsStore'
import gsap from 'gsap'
import {Keyboard} from "lucide-vue-next";

const shortcutsStore = useShortcutsStore()
const containerRef = ref<HTMLElement | null>(null)

onMounted(async () => {
  await nextTick()
  const items = containerRef.value?.querySelectorAll('.shortcut-item')
  if (items && items.length) {
    gsap.fromTo(
        items,
        { x: -200, opacity: 0 },
        {
          x: 0,
          opacity: 1,
          stagger: 0.1,
          duration: 0.6,
          ease: 'power2.out',
        }
    )
  }
})
</script>

<template>
  <div ref="containerRef">
    <!-- 基础快捷键 -->

    <div class="shortcut-item">
      <div class="flex flex-row justify-start items-center space-x-2 mb-3">
        <Keyboard />
        <p class="text-xl font-semibold">TIP: Shortcuts</p>
      </div>
      <p class="opacity-80 font-light text-sm mb-6">Navigate the site with ease using keyboard shortcuts.</p>
    </div>


    <p class="font-semibold text-lg mb-3 mt-3 shortcut-item">基础</p>
    <div class="shortcut-item">
      <ShortcutItem label="Open Quick Access" :keyLabels="['Q']" />
    </div>
    <div class="shortcut-item">
      <ShortcutItem label="Close Quick Access" :keyLabels="['Q', 'Esc']" />
    </div>
    <div class="shortcut-item">
      <ShortcutItem label="Back To Home" :keyLabels="['H']" />
    </div>

    <!-- 自定义快捷键 -->
    <div v-if="shortcutsStore.customShortcuts.length > 0">
      <p class="font-semibold text-lg mb-3 mt-3 shortcut-item">该页面的快捷键</p>
      <div
          v-for="i in shortcutsStore.customShortcuts"
          :key="i.label"
          class="shortcut-item"
      >
        <ShortcutItem
            :label="i.label"
            :keyLabels="i.keyLabels"
            :pressType="i.pressType"
        />
      </div>
    </div>
  </div>
</template>

<style>
.shortcut-item {
  opacity: 0; /* GSAP 动画起始状态 */
}
</style>