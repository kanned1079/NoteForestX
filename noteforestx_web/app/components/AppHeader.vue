<script lang="ts" setup>
import {ref, computed} from "vue";
import {updateSurfacePalette} from '@primeuix/themes';
import {useI18n} from "vue-i18n";
import {Icon} from '@vicons/utils'
import {useRouter, useRoute} from 'vue-router';
import {TicketSharp} from "@vicons/ionicons5"
import useUserStore from "../store/userStore";
import useThemeStore from "../store/themeStore";
import useActionStore from "~/store/actionStore";
import {Keyboard} from "lucide-vue-next";
import ShortcutsDefault from "~/components/ShortcutsDefault.vue";
import ChangeLanguage from "~/components/ChangeLanguage.vue";
import {languageList} from "~/types/language"
import ActionArea from '~/components/ActionArea.vue'
import { useScrollFadeIn } from '~/composables/useScrollFadeIn'
import {useToast} from "primevue/usetoast";


import { useTemplateRef } from 'vue'
import Kbd from "~/components/MyIntro/Kbd.vue";

const toast = useToast()
const {t, setLocale, locale} = useI18n()
const route = useRoute()
const router = useRouter();
const userStore = useUserStore()
const themeStore = useThemeStore()
const actionStore = useActionStore()


const blocked = ref<boolean>(false)
// const showAff = computed<boolean>(() => )

const showLoginModal = ref<boolean>(false)
const onClickProfileButton = () => {
  console.log(userStore.isAuthed)
  if (!userStore.isAuthed) {
    showLoginModal.value = true
  }
}

const actionAreaRefItem = useTemplateRef('actionAreaRef')

const showLangSelector = () => {
  blocked.value = true
  // langChangeCompRef.value?.openMenu()
  // setTimeout(() => langChangeCompRef.value?.openMenu(), 500)
  actionAreaRefItem.value?.showLangSelector()
}

const buildMenu = () => {
  const items: any[] = [
    {
      label: 'appMenu.general',
      items: [
        {
          label: 'layout.homePage',
          icon: 'pi pi-home',
          route: '/'
        },
        {
          label: 'layout.knowledgeLib',
          icon: 'pi pi-book',
          // route: '/article'
          route: `/${locale.value}/article`
        },
        {
          label: 'collection.title',
          icon: 'pi pi-star',
          // route: '/collection',
          route: `/${locale.value}/collection`
        },
        {
          label: 'layout.details',
          icon: 'pi pi-wrench',
          // route: '/details'
          route: `/${locale.value}/details`
        },
      ]
    },
    {
      label: 'appMenu.lang',
      items: [
        {
          label: 'layout.changLang',
          icon: 'pi pi-language',
          command: () => showLangSelector()
        }
      ]
    }
  ]

  // ✅ 管理员菜单
  if (userStore.isAuthed && userStore.user.role === 'ADMIN') {
  //   if (userStore.isAuthed) {
    items.push({
      label: 'appMenu.admin',
      items: [
        {
          label: 'layout.knowledgeMgr',
          icon: 'pi pi-file-edit',
          // route: '/admin/article'
          route: `/${locale.value}/admin/article`
        },
      ]
    })
  }

    if (userStore.isAuthed) {
    items.push({
      label: 'appMenu.profile',
      items: [
        {
          label: 'layout.myAccount',
          icon: 'pi pi-user',
          shortcut: '⌥+I',
          route: `/${locale.value}/profile`
        },
        {
          label: 'layout.logout',
          icon: 'pi pi-sign-out',
          shortcut: '⌥+Q',
          command: () => logoutClick()
        }
      ]
    })
  }

  items.push({separator: true})

  return items
}

// const itemsMenu = computed(() => buildMenu())

const itemsMenu = computed(() => {
  if (!import.meta.client) return [] // SSR 下返回空数组
  return buildMenu()
})

const appMenu = ref()
const toggleAppMenu = (event: MouseEvent) => {
  appMenu.value?.toggle(event);
};

const getUserRole = computed<string>((): string => {
  if (userStore.isAuthed) {
    switch (userStore.user.role) {
      case "ADMIN":
        return "appMenu.adminRole"
      case "USER":
        return "appMenu.userRole"
      default:
        return "appMenu.noLoginRole"
    }
  } else {
    return 'appMenu.noLoginRole'
  }
})

const getUserName = computed(() => {
  return userStore.isAuthed
      ? (userStore.user.username || userStore.user.email)
      : "sample@domain.com"
})


const onSearchBtnClicked = () => themeStore.searchDialog.show = true

function onKeydown(e: KeyboardEvent) {
  if (route.meta.disableShortcuts) return
  // 忽略输入框中的按键，避免干扰输入
  const target = e.target as HTMLElement
  if (target?.tagName === 'INPUT' || target?.tagName === 'TEXTAREA') {
    return
  }

  if (e.key === 'q' || e.key === 'Q') {
    blocked.value = !blocked.value
  }

  if (e.key.toLowerCase() === 'l') showLangSelector()

  if (e.key.toLowerCase() === 'h') {
    navigateTo("/")
    blocked.value = false
  }

  if (e.key.toLowerCase() === 'a') {
    navigateTo("/article")
    blocked.value = false
  }

  if (e.key.toLowerCase() === 'i') {
    navigateTo("/illustration")
    blocked.value = false
  }

  if (e.key.toLowerCase() === 'd') {
    navigateTo("/its-my-duty")
    blocked.value = false
  }

  if (e.key === 'Escape' && blocked.value) {
    blocked.value = false
  }
}

// const onClickWorkBtn = (code: '0' | '1' | '2') => {
//   if (code === '2')  {
//     const designEl = document.querySelector('#work-design')
//     if (designEl) designEl.scrollIntoView({ behavior: 'smooth' })
//   } else {
//     themeStore.setWorkTab(code)
//     const el = document.querySelector('#work-section')
//     if (el) el.scrollIntoView({ behavior: 'smooth' })
//   }
// }

const onClickWorkBtn = (code: '0' | '1' | '2') => {
  if (process.client) { // Nuxt 3 提供 process.client 判断
    const designEl = document.querySelector('#work-design')
    const el = document.querySelector('#work-section')
    if (designEl) designEl.scrollIntoView({ behavior: 'smooth' })
    if (el) el.scrollIntoView({ behavior: 'smooth' })
  }
}

const logoutClick = () => {
  if (userStore.logout()) {
    toast.add({
      severity: 'success',
      summary: t('layout.logoutSuccess'),
      life: 3000
    });
  }
}

onMounted(() => {
  window.addEventListener('keydown', onKeydown)
})

onBeforeUnmount(() => {
  window.removeEventListener('keydown', onKeydown)
})

</script>

<template>
  <div class="p-3 flex-row flex justify-between items-center" style="backdrop-filter: blur(1px)">
    <div class="ml-1 flex flex-row items-center">
      <Button
          size="medium"
          severity="secondary"
          variant="text"
          aria-haspopup="true"
          aria-controls="app_menu"
          @click="toggleAppMenu"
          class="p-1 mr-4"
      >
        <template #default>
          <div class="text-xl font-semibold flex flex-row items-center space-x-2">
            <Icon class="text-[#7234e9] dark:text-[#8257f2]">
              <TicketSharp/>
            </Icon>
            <p class="text-[#7234e9] dark:text-[#8257f2]">NoteForest <span class="font-mono">X</span></p>
          </div>
        </template>
      </Button>
      <Menu ref="appMenu" id="app_menu" :model="itemsMenu" style="width: auto" class="w-full md:min-w-60 md:max-w-max mt-2" :popup="true">
        <template #start>

        </template>
        <template #submenuheader="{ item }">
          <span class="text-primary font-bold">{{ t(item.label as string) }}</span>
        </template>
        <template #item="{ item, props }">
          <!-- ① 有路由的菜单项 -->
          <router-link
              v-if="item.route"
              :to="item.route"
              custom
              v-slot="{ href, navigate }"
          >
            <a
                v-ripple
                :href="href"
                v-bind="props.action"
                @click="navigate"
            >
              <span :class="item.icon" />
              <span class="ml-2">{{ t(item.label as string) }}</span>
            </a>
          </router-link>

          <!-- ② 纯行为菜单项 -->
          <button
              v-else
              v-ripple
              type="button"
              v-bind="props.action"
              @click="item.command?.()"
              class="w-full text-left flex items-center"
          >
            <span :class="item.icon" />
            <span class="ml-2">{{ t(item.label as string) }}</span>
          </button>
        </template>
        <template #end>
          <button
              @click="onClickProfileButton"
              v-ripple
              class="relative overflow-hidden w-full border-0 bg-transparent flex items-start pt-2 pb-2 pl-4 pr-4 hover:bg-surface-100 dark:hover:bg-surface-800 rounded-none cursor-pointer transition-colors duration-200"
          >
            <div class="inline-flex flex-col items-start text-left w-full">
      <span class="font-bold whitespace-normal break-words w-full">
        {{ getUserName }}
      </span>
              <span class="text-sm opacity-50">{{ t(getUserRole) }}</span>
            </div>
          </button>
        </template>
      </Menu>

      <Button
          v-if="themeStore.showLeftIntroBtn"
          size="medium"
          severity="secondary"
          variant="text"
          aria-haspopup="true"
          aria-controls="app_menu"
          class="p-1 mr-2"
          @click="onClickWorkBtn('0')"
      >
        <span class="pi pi-sparkles"></span>
        <p class="text-base font-base">{{ t('my_work.header_title') }}</p>
      </Button>

      <Button
          v-if="themeStore.showLeftIntroBtn"
          size="medium"
          severity="secondary"
          variant="text"
          aria-haspopup="true"
          aria-controls="app_menu"
          class="p-1 mr-2"
          @click="navigateTo({path: '/article'})"
      >
        <span class="pi pi-book" />
      </Button>

      <Button
          v-if="themeStore.showLeftIntroBtn"
          size="medium"
          severity="secondary"
          variant="text"
          aria-haspopup="true"
          aria-controls="app_menu"
          class="p-1 mr-2"
          @click="navigateTo({path: '/collection'})"
      >
        <span class="pi pi-star hover:pi-star-fill"></span>
      </Button>


    </div>

    <div class="mr-3 flex flex-row gap-3" v-if="themeStore.showHeaderSearchBtn">
      <Button
          class="h-8 px-3 flex items-center gap-2"
          icon="pi pi-search"
          severity="primary"
          outlined
          size="small"
          :label="`Search`"
          @click="onSearchBtnClicked"
      >
        <template #default>
    <span class="flex items-center gap-1">
      <i class="pi pi-search text-xs"></i>
<!--      <span>搜寻</span>-->
      <span
          class="px-2 py-0.5 text-xs font-light "
      >
        ⌘+K
      </span>

    </span>
        </template>
      </Button>
      <Button
          class="h-8 px-3"
          icon="pi pi-refresh"
          severity="secondary"
          size="small"
          outlined
          @click="actionStore.fireTriggerSearchArticle()"
      />
    </div>

    <div class="mr-3 flex flex-row gap-3" v-if="themeStore.showEditMetaBtn">
      <Button
          class="h-8 px-3 flex items-center gap-2"
          icon="pi pi-search"
          severity="primary"
          outlined
          size="small"
          :label="`Search`"
          @click="themeStore.setShowEditMetaDialog(true)"
      >
        <template #default>
          <span class="flex items-center gap-1">
            <i class="pi pi-search text-xs"></i>
          <span>{{ t('layout.editMeta') }}</span>
          <span class="px-2 py-0.5 text-xs font-light ">⌘+K</span>
          </span>
        </template>
      </Button>

      <Button
          class="h-8 px-3 flex items-center gap-2"
          icon="pi pi-save"
          severity="primary"
          outlined
          size="small"
          :label="`Search`"
          @click="actionStore.fireTriggerArticleSave(null)"
      >
        <template #default>
          <span class="flex items-center gap-1">
            <i class="pi pi-file text-xs"></i>
          <span>{{ t('layout.save') }}</span>
          <span class="px-2 py-0.5 text-xs font-light ">⌘+S</span>
          </span>
        </template>
      </Button>
    </div>

    <div class="mr-3 flex flex-row gap-3" v-if="themeStore.showCatalog">
      <Button
          class="h-8 px-3 flex items-center gap-2"
          icon="pi pi-search"
          severity="primary"
          outlined
          size="small"
          :label="`Search`"
          @click="actionStore.fireTriggerShowCatalog(null)"
      >
        <template #default>
          <span class="flex items-center gap-1">
            <i class="pi pi-align-left text-xs"></i>
          <span>{{ t('article.catalog') }}</span>
          <span class="px-2 py-0.5 text-xs font-light font-mono"> / </span>
          </span>
        </template>
      </Button>
    </div>

    <div class="mr-3 flex flex-row gap-3" v-if="themeStore.showCommitIllustrationBtn">
      <Button
          class="h-8 px-3 flex items-center gap-2"
          icon="pi pi-search"
          severity="primary"
          outlined
          size="small"
          :label="`Search`"
          @click="actionStore.fireTriggerCommitNewIllustration()"
      >
        <template #default>
          <span class="flex items-center gap-1">
            <i class="pi pi-search text-xs"></i>
          <span>Save</span>
          <span class="px-2 py-0.5 text-xs font-light ">⌘+S</span>
          </span>
        </template>
      </Button>
    </div>
  </div>

  <Dialog
      :dismissableMask="true"
      maximizable modal
      v-model:visible="showLoginModal" :show-header="false" modals class="p-0">
    <LoginForm :close-dialog="() => {showLoginModal=false}"/>
  </Dialog>

  <Mask :modelValue="blocked" :close="() => {blocked = false}" closeable >
    <template #rt>
      <div class="text-right mb-6">
        <Button size="small" severity="secondary" icon="pi pi-times" aria-label="Cancel" class="opacity-80" @click="blocked = false" />
      </div>
      <ActionArea ref="actionAreaRef" />
    </template>

    <template #lb>
      <div class="max-w-[320px]">

        <ShortcutsDefault />

      </div>
    </template>

  </Mask>

</template>

<style>

</style>