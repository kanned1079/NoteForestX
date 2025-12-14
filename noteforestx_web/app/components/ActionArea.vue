<script setup lang="ts">
import ChangeLanguage from "~/components/ChangeLanguage.vue"
import { useScrollFadeIn } from "~/composables/useScrollFadeIn"
import {defineExpose, useTemplateRef} from "vue";

const langChangeCompRef = useTemplateRef('langBtnRef')

const showLangSelector = () => {
  console.log('action area: show')
  // setTimeout(() => langChangeCompRef.value?.openMenu(), 1000)
  langChangeCompRef.value?.openMenu()
}

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
  <div class="max-w-[320px] min-w-[300px] ">


    <div class="flex flex-row justify-start items-center space-x-2 mb-3">
      <div class="flex flex-col w-full ">
        <p class="text-xl font-semibold mb-4 animate-card">Action Center</p>

        <!-- 按钮容器：占满宽度 -->
        <div class="flex flex-row w-full space-x-4 animate-card">

          <ChangeLanguage ref="langBtnRef"/>
          <Button
              class="h-20 flex-1 bg-[rgba(255,255,255,0.7)] dark:bg-gray-800 dark:text-gray-200 text-sm rounded-xl"
              variant="text"
          >
            Focus
          </Button>
        </div>

        <p class="text-xl font-semibold mb-4 mt-8 animate-card">Message</p>
        <p class="text-sm font-light opacity-80 animate-card">Sorry, There's no message for u. lol</p>

      </div>
    </div>

  </div>
</template>

<style scoped>

</style>