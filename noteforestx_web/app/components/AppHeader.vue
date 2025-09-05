<script lang="ts" setup>
import {ref} from "vue";
import {updateSurfacePalette} from '@primeuix/themes';
import {useI18n} from "vue-i18n";

const {t, setLocale} = useI18n()
import flagPlaceholder from "~/assets/imgs/flag_placeholder.png"

const changeSurfaceColor = () => {
  updateSurfacePalette({
    50: '{zinc.50}',
    // ...
    950: '{zinc.950}'
  });
}

const onClickLogoButton = () => {
  console.log("pressed")
}

const menu = ref();

type LangItem = { label: string; code: string };
const items: LangItem[] = [
  {
    label: "English",
    code: 'us',
  },
  {
    label: "简体中文",
    code: 'cn',
  },
  {
    label: "繁體中文",
    code: 'hk',
  },
  {
    label: "日本語",
    code: 'jp',
  },
  {
    label: "Spain",
    code: 'es',
  },
  {
    label: "Français",
    code: 'fr',
  },
  {
    label: "Suomi",
    code: 'fi',
  },
  {
    label: "Deutsch",
    code: 'de',
  },
  {
    label: "Nederlands",
    code: 'nl',
  },

];


const toggle = (event) => {
  menu.value.toggle(event);
};

function toggleDarkMode() {
  console.log("toggleDarkMode");
  document.documentElement.classList.toggle('my-app-dark');
}

</script>

<template>
  <div class="p-3 flex-row flex justify-between items-center">
    <div class="ml-3">
      <Button
          class="ml-3"
          size="small"
          icon="pi pi-home"
          label="NoteForest X"
          severity="secondary"
          variant="text"
      >


      </Button>

    </div>
    <div class="mr-3">


      <Button size="small" variant="text" type="button" @click="toggle" aria-haspopup="true"
              aria-controls="overlay_tmenu">
        <i class="pi pi-language"></i>
        语言切换
      </Button>
      <TieredMenu size="small" appendTo="body" ref="menu" :model="items" popup class="mt-2 w-auto">
        <template #item="{ item }">
          <div class="flex flex-row items-center justify-between">
            <div class="flex justify-start items-center flex-row space-x-1.5 pt-1.5 pb-1.5 pl-1.5 mb-1.25">
              <span class="drop-shadow-sm" :class="`fi fi-${item.code} rounded-sm`"></span>
              <span class="text-sm">{{ item.label }}</span>
            </div>
            <div>
              <i v-if="item.code==='hk'" class="pi pi-check font-bold text-xs text-cyan-800 mr-3"></i>
<!--              <Tag rounded severity="success" class="pi pi-check-circle text-sm" />-->
            </div>
          </div>

        </template>
      </TieredMenu>


      <!--      <Button-->
      <!--          class="ml-3"-->
      <!--          size="small"-->
      <!--          icon="pi pi-sun"-->
      <!--          label="Sun"-->
      <!--          severity="secondary" variant="primary"-->
      <!--          @click="changeSurfaceColor"-->
      <!--      >-->
      <!--      </Button>-->

    </div>
  </div>

</template>

<style>
/* 调整 TieredMenu 弹出距离按钮 */
.p-tieredmenu.p-overlaypanel {
  margin-top: 0.5rem; /* 控制下方间距 */
  margin-left: 0.25rem; /* 控制横向偏移 */
}
</style>