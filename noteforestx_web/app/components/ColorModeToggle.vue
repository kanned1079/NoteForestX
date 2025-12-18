<script lang="ts" setup>
import {useScrollFadeIn} from "~/composables/useScrollFadeIn";
import DaySign from '~/components/ColorSign/DaySign.vue'
import NightSign from "~/components/ColorSign/NightSign.vue";

const show = ref(false)


const toggleMode = () => {
  document.documentElement.classList.toggle('my-app-dark');
}

onMounted(() => {
  setTimeout(() => {
    show.value = true
  }, 2000)
})
</script>

<template>
  <Button
      ref="langBtn"
      class="h-20 flex-1 bg-[rgba(255,255,255,0.7)] dark:bg-gray-800 dark:text-gray-200 text-sm rounded-xl"
      variant="text"
      type="button"
      @click="toggleMode"
      aria-haspopup="true"
      aria-controls="language_menu"
  >
    <div class="h-full w-full flex flex-col justify-between items-start">

      <span>
        <DaySign v-if="show" />
        <NightSign v-if="!show"/>
      </span>

      <div class="w-full text-right">
        <span>浅色模式</span>
      </div>
    </div>
  </Button>
</template>

<style scoped>

</style>