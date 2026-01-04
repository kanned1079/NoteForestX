<script setup lang="ts">
import { ref } from 'vue'
import useShortcutsStore from '~/store/shortcutsStore'
import { Keyboard } from 'lucide-vue-next'
import { useScrollFadeIn } from '~/composables/useScrollFadeIn'
import {useI18n} from "vue-i18n";

const {t} = useI18n()
const shortcutsStore = useShortcutsStore()
const containerRef = ref<HTMLElement | null>(null)


  useScrollFadeIn({
    selector: '.shortcut-item',
    direction: 'right',
    x: 200,
    stagger: 0.1,
    duration: 0.6,
    start: 'top 90%',
    useScrollTrigger: false
  })
</script>

<template>
  <div ref="containerRef" class="pb-10 w-[280px]">
    <!-- 基础快捷键 -->

    <div class="shortcut-item">
      <div class="flex flex-row justify-start items-center space-x-2 mb-3">
        <Keyboard />
        <p class="text-xl font-semibold">TIP: {{ t('shortcuts.header') }}</p>
      </div>
      <p class="opacity-80 font-light text-sm mb-6">{{ t('shortcuts.description') }}</p>
    </div>

<!--    <p class="font-semibold text-lg mb-3 mt-3 shortcut-item">基础</p>-->
    <div class="shortcut-item">
      <ShortcutItem :label="t('shortcuts.quickAccess.open')" :keyLabels="['Q']" />
    </div>
    <div class="shortcut-item">
      <ShortcutItem :label="t('shortcuts.quickAccess.close')" :keyLabels="['Q', 'Esc']" />
    </div>
    <div class="shortcut-item">
      <ShortcutItem :label="t('shortcuts.backHome')" :keyLabels="['H']" />
    </div>
    <div class="shortcut-item">
      <ShortcutItem :label="t('shortcuts.article')" :keyLabels="['A']" />
    </div>
    <div class="shortcut-item">
      <ShortcutItem label="It's My Duty" :keyLabels="['D']" />
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