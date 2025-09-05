<script setup lang="ts">
import useThemeStore from "~/store/themeStore";
import {useTheme} from "vuetify/framework";

const {locales, setLocale} = useI18n()
const theme = useTheme()
const themeStore = useThemeStore()

// type LangCode = 'en-us' | 'zh-cn' | 'ja-jp'
type LangCode = 'en' | 'cn'

type LangOption = {
  text: string,
  code: LangCode
}

const langList: LangOption[] = [
  {
    text: 'English',
    code: 'en',
  },
  // {
  //   text: '日本語',
  //   code: 'ja-jp',
  // },
  {
    text: '中文简体',
    code: 'cn'
  }
]

const colorSchemeClick = () => {
  themeStore.toggleDarkMode()
  theme.global.name.value = themeStore.isDarkModeEnabled?'dark':'light'
}

const langClick = (langCode: LangCode) => {
  themeStore.showMessage(`setlang ${langCode}`, "info")
  setLocale(langCode)
  themeStore.lang = langCode
}

</script>

<template>
  <div class="mr-4">

    <v-menu
        transition="slide-y-transition"
    >
      <template v-slot:activator="{ props }">
        <v-btn
            variant="plain"
            v-bind="props"
        >
          <v-icon style="" size="large">{{ 'mdi-translate-variant' }}</v-icon>
        </v-btn>
      </template>
      <v-list
          density="compact"
      >
        <v-list-item
            v-for="(item, index) in langList"
            :key="index"
            :value="item.text"
            @click="langClick(item.code)"
            rounded
            class="ml-2 mr-2"
        >
          <v-list-item-title>{{ item.text }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>

    <v-btn
        variant="plain"
        @click="colorSchemeClick"

    >
      <v-icon size="large">{{ 'mdi-white-balance-sunny' }}</v-icon>
    </v-btn>
  </div>
</template>

<style scoped>

</style>