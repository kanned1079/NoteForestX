<script setup lang="ts">
import { ref } from 'vue'
import useShortcutsStore from '~/store/shortcutsStore'
import { Keyboard } from 'lucide-vue-next'
import { useScrollFadeIn } from '~/composables/useScrollFadeIn'

const shortcutsStore = useShortcutsStore()
const containerRef = ref<HTMLElement | null>(null)


  useScrollFadeIn({
    selector: '.shortcut-item',
    direction: 'right',
    x: 200,
    stagger: 0.1,
    duration: 0.6,
    start: 'top 90%'
  })
</script>

<template>
  <div ref="containerRef" class="pb-10">
    <!-- 基础快捷键 -->

    <div class="shortcut-item">
      <div class="flex flex-row justify-start items-center space-x-2 mb-3">
        <Keyboard />
        <p class="text-xl font-semibold">TIP: Shortcuts</p>
      </div>
      <p class="opacity-80 font-light text-sm mb-6">Navigate the site with ease using keyboard shortcuts.</p>
    </div>

<!--    <p class="font-semibold text-lg mb-3 mt-3 shortcut-item">基础</p>-->
    <div class="shortcut-item">
      <ShortcutItem label="Open Quick Access" :keyLabels="['Q']" />
    </div>
    <div class="shortcut-item">
      <ShortcutItem label="Close Quick Access" :keyLabels="['Q', 'Esc']" />
    </div>
    <div class="shortcut-item">
      <ShortcutItem label="Back To Home" :keyLabels="['H']" />
    </div>
    <!--    <div>-->
<!--      <p-->
<!--          class="font-semibold text-lg mb-3 mt-3 shortcut-item"-->
<!--          v-scroll-fade="{ direction: 'right', x: 200, duration: 0.6 }"-->
<!--      >-->
<!--        该页面的快捷键-->
<!--      </p>-->

<!--      <div-->
<!--          v-for="(i, index) in shortcutsStore.customShortcuts"-->
<!--          :key="i.label"-->
<!--          class="shortcut-item"-->
<!--          v-scroll-fade="{ direction: 'right', x: 200, duration: 0.6, delay: 0.1 * index }"-->
<!--      >-->
<!--        <ShortcutItem-->
<!--            :label="i.label"-->
<!--            :keyLabels="i.keyLabels"-->
<!--            :pressType="i.pressType"-->
<!--        />-->
<!--      </div>-->
<!--    </div>-->



  </div>
</template>

<style scoped>
</style>