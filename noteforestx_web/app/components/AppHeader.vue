<script lang="ts" setup>
import {ref} from "vue";
import {updateSurfacePalette} from '@primeuix/themes';
import {useI18n} from "vue-i18n";
import {Icon} from '@vicons/utils'
import { useRouter } from 'vue-router';
import {TicketSharp} from "@vicons/ionicons5"
import useUserStore from "../store/userStore";

const {t, setLocale} = useI18n()
const router = useRouter();
const userStore = useUserStore()

const showLoginModal = ref<boolean>(false)
const onClickProfileButton = () => {
  if (userStore.isAuthed) {
    showLoginModal.value = true
  }
}

const showLanguageModal = ref<boolean>(false)

const langMenuRef = ref()
type LangItem = { label: string; code: string };
const items: LangItem[] = [
  {
    label: "English",
    code: 'us',
    command: () => {
      console.log("en")
      setLocale('us')
    }
  },
  {
    label: "简体中文",
    code: 'cn',
    command: () => setLocale('cn')
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

const onClickSetLocale = (event) => {
  langMenuRef.value.toggle(event)
}

const itemsMenu = ref([
  {
    label: 'appMenu.general',
    items: [
      {
        label: 'layout.homePage',
        icon: 'pi pi-home',
        route: '/'
      },
      {
        label: 'layout.aboutMe',
        icon: 'pi pi-github',
        route: '/about'
      },
      {
        label: 'layout.knowledgeLib',
        icon: 'pi pi-book',
      },
      {
        label: 'layout.illustrationLib',
        icon: 'pi pi-images',
      },
      {
        label: 'layout.sponsor',
        icon: 'pi pi-gift'
      }
    ]
  },
  {
    label: 'appMenu.lang',
    items: [
      {
        label: 'layout.changLang',
        icon: 'pi pi-language',
        route: '/config/language'
      },
    ]
  },
  {
    label: 'appMenu.admin',
    items: [
      {
        label: 'layout.statistic',
        icon: 'pi pi-chart-bar'
      },
      {
        label: 'layout.knowledgeMgr',
        icon: 'pi pi-file-edit',
      },
      {
        label: 'layout.illustrationMgr',
        icon: 'pi pi-palette',
      },
    ]
  },
  {
    label: 'appMenu.profile',
    items: [
      {
        label: 'layout.myAccount',
        icon: 'pi pi-user',
        shortcut: '⌥+I'
      },
      {
        label: 'layout.logout',
        icon: 'pi pi-sign-out',
        shortcut: '⌥+Q'
      }
    ]
  },
  {
    separator: true
  }
]);

const appMenu = ref()
const toggleAppMenu = (event) => {
  appMenu.value.toggle(event);
};

</script>

<template>
  <div class="p-3 flex-row flex justify-between items-center">
    <div class="ml-0">
      <!--     <LogoButton :click="toggleDrawerOpen" />-->
      <Button
          size="medium"
          label="NoteForest X"
          severity="secondary"
          variant="text"
          aria-haspopup="true"
          aria-controls="app_menu"
          @click="toggleAppMenu"
      >
        <template #icon>
          <Icon>
            <TicketSharp/>
          </Icon>
        </template>
      </Button>
      <Menu ref="appMenu" id="app_menu" :model="itemsMenu" class="w-full md:min-w-60 md:max-w-max mt-2" :popup="true">
        <template #start>

        </template>
        <template #submenulabel="{ item }">
          <span class="text-primary font-bold">{{ t(item.label) }}</span>
        </template>
        <template #item="{ item, props }">
          <router-link v-slot="{ href, navigate }" :to="item.route" custom>
            <a v-ripple :href="href" v-bind="props.action" @click="navigate">
              <span :class="item.icon" />
              <span class="ml-2">{{ t(item.label) }}</span>
            </a>
          </router-link>
        </template>
        <template #end>
          <button
              @click="onClickProfileButton"
              v-ripple
              class="relative overflow-hidden w-full border-0 bg-transparent flex items-start p-2 pl-4 hover:bg-surface-100 dark:hover:bg-surface-800 rounded-none cursor-pointer transition-colors duration-200"
          >
            <div class="inline-flex flex-col items-start text-left w-full">
      <span class="font-bold whitespace-normal break-words w-full">
        kanned1079@gmail.com
      </span>
              <span class="text-sm">管理员</span>
            </div>
          </button>
          <button
              @click="onClickProfileButton"
              v-ripple
              class="relative overflow-hidden w-full border-0 bg-transparent flex items-start p-2 pl-4 hover:bg-surface-100 dark:hover:bg-surface-800 rounded-none cursor-pointer transition-colors duration-200"
          >
            777
          </button>
        </template>
      </Menu>

    </div>
    <div class="mr-3" v-if="false">
      <Button size="small" variant="link" type="button" @click="toggle" aria-haspopup="true">
        {{ t('layout.knowledgeLib') }}
      </Button>
      <Button size="small" variant="link" type="button" @click="toggle" aria-haspopup="true">
        {{ t('layout.illustrationLib') }}
      </Button>
      <Button size="small" variant="link" type="button" @click="toggle" aria-haspopup="true">
        {{ t('layout.aboutMe') }}
      </Button>
      <Button size="small" variant="link" type="button" @click="onClickSetLocale" aria-haspopup="true"
              aria-controls="lang_menu">
        <i class="pi pi-language"></i>
        语言切换
      </Button>
      <TieredMenu size="small" appendTo="body" ref="langMenuRef" :model="items" popup class="mt-2 w-auto">
        <template #item="{ item }">
          <div class="flex flex-row items-center justify-between">
            <div class="flex justify-start items-center flex-row space-x-1.5 pt-1.5 pb-1.5 pl-1.5 mb-1.25">
              <span class="drop-shadow-sm" :class="`fi fi-${item.code} rounded-sm`"></span>
              <span class="text-sm">{{ item.label }}</span>
            </div>
            <div>
              <i v-if="item.code==='hk'" class="pi pi-check font-bold text-xs text-cyan-800 mr-3"></i>
            </div>
          </div>

        </template>
      </TieredMenu>

    </div>
  </div>

  <Dialog v-model:visible="showLoginModal" :show-header="false" modals :style="null">
   <LoginForm />

  </Dialog>

  <Dialog v-model:visible="showLanguageModal" :show-header="true" modals>
    <TieredMenu size="small" appendTo="body" ref="langMenuRef" :model="items" class="mt-2 w-auto">
      <template #item="{ item }">
        <div class="flex flex-row items-center justify-between">
          <div class="flex justify-start items-center flex-row space-x-1.5 pt-1.5 pb-1.5 pl-1.5 mb-1.25">
            <span class="drop-shadow-sm" :class="`fi fi-${item.code} rounded-sm`"></span>
            <span class="text-sm">{{ item.label }}</span>
          </div>
          <div>
            <i v-if="item.code==='hk'" class="pi pi-check font-bold text-xs text-cyan-800 mr-3"></i>
          </div>
        </div>

      </template>
    </TieredMenu>
  </Dialog>


</template>

<style>

</style>