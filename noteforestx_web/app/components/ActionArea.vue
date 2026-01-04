<script setup lang="ts">
import ChangeLanguage from "~/components/ChangeLanguage.vue"
import { useScrollFadeIn } from "~/composables/useScrollFadeIn"
import {defineExpose, useTemplateRef} from "vue";
import ColorModeToggle from "~/components/ColorModeToggle.vue";
import useThemeStore from "~/store/themeStore";
import {useI18n} from "vue-i18n";

const {t} = useI18n()
const themeStore = useThemeStore()
const langChangeCompRef = useTemplateRef('langBtnRef')

const showLangSelector = () => {
  console.log('action area: show')
  // setTimeout(() => langChangeCompRef.value?.openMenu(), 1000)
  langChangeCompRef.value?.openMenu()
}


// const { mode, setLight, setDark, setSystem } = useTheme()

useScrollFadeIn({
  selector: '.animate-card',
  direction: 'left',
  x: 80
})


defineExpose({
  showLangSelector
})
</script>


<template>
  <div class="max-w-[320px] min-w-[300px]">
    <div class="flex flex-row justify-start items-center space-x-2 mb-3">
      <div class="flex flex-col w-full ">
        <p class="text-xl font-semibold mb-4 animate-card">{{ t('actionCenter.title') }}</p>

        <!-- 按钮容器：占满宽度 -->
        <div class="flex flex-row w-full space-x-4 animate-card">

          <ChangeLanguage ref="langBtnRef"/>

        </div>

        <p class="text-xl font-semibold mb-4 mt-8 animate-card">{{ t('actionCenter.msg') }}</p>
        <p
            v-if="themeStore.actionCenterMsgs.length===0"
            class="text-sm font-light opacity-80 animate-card">{{ t('actionCenter.noMsg') }}</p>
        <p
            v-for="i in themeStore.actionCenterMsgs"
            :key="i"
            class="text-sm font-light opacity-80 animate-card mb-2">
          {{ i }}
        </p>

      </div>
    </div>

  </div>
</template>

<style scoped>

</style>