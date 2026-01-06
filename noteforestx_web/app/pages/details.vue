<script setup lang="ts">
import {useI18n} from "vue-i18n";
import toolslist from "~/assets/imgs/toolslistshort.png"
import {useScrollFadeIn} from "~/composables/useScrollFadeIn";

const {t} = useI18n()
const showRibbon = ref<boolean>(false)

useScrollFadeIn({
  selector: '.animate-card-details',
  y: 60,
  stagger: 0.15
})

onMounted(() => {
  setTimeout(() => {
    showRibbon.value = true
  }, 300)
})
</script>

<template>
  <ClientOnly>
    <ColoredRibbon :show="showRibbon" />
  </ClientOnly>
  <div class="mt-4 pb-60">
    <div class="container mx-auto max-w-[980px] px-4 space-y-6">
      <section class="space-y-2 animate-card-details">
        <h2 class="text-4xl font-bold opacity-90">{{ t('details.open_source') }}</h2>
        <p>
          {{ t('details.open_source_desc') }}
          &nbsp;
          <a
              class="underline font-mono opacity-80"
              :href="t('details.open_source_link')"
          >
            {{ t('details.open_source_link') }}
          </a>
        </p>
      </section>

      <!-- 背景與目標 -->
      <section class="space-y-2 animate-card-details">
        <h2 class="text-xl font-bold opacity-90">{{ t('details.background_title') }}</h2>
        <p>{{ t('details.background_desc') }}</p>
      </section>

      <section class="space-y-2 animate-card-details">
        <h2 class="text-xl font-bold opacity-90">{{ t('layout.sponsor') }}</h2>
        <p class="font-mono">0xaa53b6eb01103118a909a018a8c24d480e47bf7e <span class="font-bold">(Ethereum)</span></p>
        <p class="font-mono">0xaa53b6eb01103118a909a018a8c24d480e47bf7e <span class="font-bold">(Arbitrum)</span></p>
      </section>

      <!-- 技術棧 -->
      <section class="space-y-4 animate-card-details">
        <h2 class="text-xl font-bold opacity-90">{{ t('details.tech_stack') }}</h2>

        <!-- 前端部分 -->
        <div class="space-y-2">
          <h3 class="text-lg font-semibold opacity-90">{{ t('details.frontend_title') }}</h3>

          <ul class="list-disc list-inside space-y-1 font-mono">
            <li>{{ t('details.frontend_list_1') }}</li>
            <li>{{ t('details.frontend_list_2') }}</li>
            <li>{{ t('details.frontend_list_3') }}</li>
            <li>{{ t('details.frontend_list_4') }}</li>
            <li>{{ t('details.frontend_list_5') }}</li>
            <li>{{ t('details.frontend_list_6') }}</li>
          </ul>

          <p>{{ t('details.frontend_desc') }}</p>
        </div>

        <!-- 後端部分 -->
        <div class="space-y-2">
          <h3 class="text-lg font-semibold opacity-90">{{ t('details.backend_title') }}</h3>

          <ul class="list-disc list-inside space-y-1 font-mono">
            <li>{{ t('details.backend_list_1') }}</li>
            <li>{{ t('details.backend_list_2') }}</li>
            <li>{{ t('details.backend_list_3') }}</li>
            <li>{{ t('details.backend_list_4') }}</li>
            <li>{{ t('details.backend_list_5') }}</li>
            <li>{{ t('details.backend_list_6') }}</li>
          </ul>

          <p>{{ t('details.backend_desc') }}</p>
        </div>
      </section>
      <img :src="toolslist" :alt="t('details.tool_image_alt')" class="max-w-[800px] w-full h-auto text-left animate-card-details" />
    </div>
  </div>
</template>

<style scoped>

</style>