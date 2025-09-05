<script lang="ts" setup>
import {NuxtLayout, NuxtPage} from "#components";
import useThemeStore from "~/store/themeStore";
const themeStore = useThemeStore()
import {useI18n} from "vue-i18n";
const {t} = useI18n()


onMounted(() => {
  themeStore.showMessage("hello", t('hello'))
})
</script>

<template>
  <v-snackbar
      class="custom-snackbar"
      v-model="themeStore.messageBody.show"
      :timeout="3000"
      location="top"
      variant="flat"
      :color="themeStore.messageBody.type"
  >
    <div style="display: flex; flex-direction: row; align-items: center; color: #f3f3f3">
      <v-icon size="small" class="mr-1">mdi-information</v-icon>
      <p>{{ themeStore.messageBody.text }}</p>
    </div>

  </v-snackbar>

  <NuxtLayout>
    <v-app>

      <NuxtPage />
    </v-app>
  </NuxtLayout>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  text-transform: none !important;
}
.page-enter-active,
.page-leave-active {
  transition: opacity 0.5s;
}
.page-enter,
.page-leave-to {
  opacity: 0;
}

.custom-snackbar {
  padding: 12px !important;
  opacity: 0.8;
}
</style>
