<script setup lang="ts">
import { ref } from "vue";
import { useI18n } from "vue-i18n";
import useUserStore from "~/store/userStore";
import type { LanguageItem, LanguageCode } from "~/types/language";

const { setLocale } = useI18n();
const userStore = useUserStore();

const menu = ref();

const languageList: LanguageItem[] = [
  { label: "English", flag: "us", code: "en_us" },
  { label: "简体中文", flag: "cn", code: "zh_cn" },
  { label: "繁體中文", flag: "hk", code: "zh_hk" },
  { label: "日本語", flag: "jp", code: "ja_jp" },
  { label: "Español", flag: "es", code: "es_es" },
  { label: "Français", flag: "fr", code: "fr_fr" },
  { label: "Suomi", flag: "fi", code: "fi_fi" },
  { label: "Deutsch", flag: "de", code: "de_de" },
  { label: "Nederlands", flag: "nl", code: "nl_nl" },
];

const items = languageList.map((lang) => ({
  label: lang.label,
  icon: `fi fi-${lang.flag} rounded-sm`,
  command: () => saveLanguage(lang.code),
}));

const saveLanguage = (langCode: LanguageCode) => {
  userStore.languageInUsing = langCode;
  setLocale(langCode);
};

const toggleMenu = (event: Event) => {
  menu.value.toggle(event);
};
</script>

<template>
  <div class="p-6 max-w-4xl mx-auto">
    <!-- 页面标题 -->
    <h1 class="text-2xl font-bold mb-6">
      🌐 {{ $t("layout.changLang") }}
    </h1>

    <!-- 语言选择卡片 -->
    <Card class="shadow-md">
      <template #title>
        {{ $t("layout.languageSettings") }}
      </template>
      <template #content>
        <div class="flex flex-col space-y-4">
          <p class="text-gray-500 dark:text-gray-400">
            {{ $t("layout.chooseLanguageDescription") }}
          </p>

          <!-- 当前语言按钮 -->
          <div class="flex items-center space-x-3">
            <Button
                size="small"
                variant="outlined"
                type="button"
                @click="toggleMenu"
                aria-haspopup="true"
                aria-controls="language_menu"
            >
              <span
                  class="drop-shadow-sm mr-2"
                  :class="`fi fi-${userStore.languageInUsing?.split('_')[1] || 'us'} rounded-sm`"
              />
              {{ languageList.find(l => l.code === userStore.languageInUsing)?.label || 'Language' }}
            </Button>

            <span class="text-sm text-gray-400">
              {{ $t("layout.currentLang") }}
            </span>
          </div>

          <!-- 下拉菜单 -->
          <TieredMenu
              ref="menu"
              id="language_menu"
              :model="items"
              popup
              appendTo="body"
              class="w-48"
          >
            <template #item="{ item }">
              <div class="flex items-center space-x-2">
                <span :class="item.icon"></span>
                <span class="text-sm">{{ item.label }}</span>
                <i
                    v-if="item.label === languageList.find(l => l.code === userStore.languageInUsing)?.label"
                    class="pi pi-check ml-auto text-primary-500"
                />
              </div>
            </template>
          </TieredMenu>
        </div>
      </template>
    </Card>
  </div>
</template>