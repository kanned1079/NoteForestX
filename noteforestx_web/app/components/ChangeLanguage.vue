<script setup lang="ts">
import { ref, useTemplateRef, defineExpose } from "vue";
import { useI18n } from "vue-i18n";
import useUserStore from "~/store/userStore";
import { type LanguageCode, languageList } from "~/types/language";

const { setLocale } = useI18n();
const userStore = useUserStore();

const menu = ref();

const saveLanguage = (langCode: LanguageCode) => {
  userStore.languageInUsing = langCode;
  setLocale(langCode);
};

const items = languageList.map((lang) => ({
  label: lang.label,
  icon: `fi fi-${lang.flag} rounded-sm`,
  command: () => saveLanguage(lang.code),
}));

const toggleMenu = (event: Event) => {
  menu.value.toggle(event);
};

const langButton = useTemplateRef('langBtn')

const openMenu = () => {
  setTimeout(() => langButton.value?.$el.click(), 1000)

}




defineExpose({
  openMenu
})
</script>

<template>
  <!-- 🌐 当前语言按钮（卡片版） -->
  <Button
      ref="langBtn"
      class="h-20 flex-1 bg-[rgba(255,255,255,0.7)] dark:bg-gray-800 dark:text-gray-200 text-sm rounded-xl"
      variant="text"
      type="button"
      @click="toggleMenu"
      aria-haspopup="true"
      aria-controls="language_menu"
  >
    <div class="h-full w-full flex flex-col justify-between items-start">
      <i class="pi pi-language text-xl opacity-80"></i>

      <div class="text-base flex items-center">
        <span
            class="drop-shadow-md mr-2"
            :class="`fi fi-${userStore.languageInUsing?.split('_')[1] || 'us'} rounded-sm`"
        />
        {{ languageList.find(l => l.code === userStore.languageInUsing)?.label || 'Language' }}
      </div>
    </div>
  </Button>

  <!-- ⬇️ 下拉菜单 -->
  <TieredMenu
      ref="menu"
      id="language_menu"
      :model="items"
      popup
      appendTo="body"
      style="margin-top: 12px"
      class="w-48 rounded-xl"
  >
    <template #item="{ item }">
      <div class="flex flex-row justify-start items-center space-x-2 p-2" >
        <span :class="item.icon"></span>
        <span class="text-sm">{{ item.label }}</span>
      </div>
    </template>
  </TieredMenu>
</template>

<style>
.p-tieredmenu {
  margin-top: 10px;
}

.p-tieredmenu-item {
  margin-top: 1px;
  margin-bottom: 1px;
}

.p-tieredmenu-item-content {
  border-radius: 8px !important;
}

</style>