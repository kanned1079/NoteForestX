<script lang="ts" setup>
import {ref, computed} from "vue";
import {updateSurfacePalette} from '@primeuix/themes';
import {useI18n} from "vue-i18n";
import {Icon} from '@vicons/utils'
import {useRouter} from 'vue-router';
import {TicketSharp} from "@vicons/ionicons5"
import useUserStore from "../store/userStore";
import useThemeStore from "../store/themeStore";


const {t, setLocale} = useI18n()
const router = useRouter();
const userStore = useUserStore()
const themeStore = useThemeStore()

const showLoginModal = ref<boolean>(false)
const onClickProfileButton = () => {
  console.log(userStore.isAuthed)
  if (!userStore.isAuthed) {
    showLoginModal.value = true
  }
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
        // {
        //   label: 'layout.aboutMe',
        //   icon: 'pi pi-github',
        //   route: '/about'
        // },
        {
          label: 'layout.knowledgeLib',
          icon: 'pi pi-book',
          route: '/article'
        },
        {
          label: 'layout.illustrationLib',
          icon: 'pi pi-images',
          route: '/illustration'
        },
        {
          label: 'layout.sponsor',
          icon: 'pi pi-gift',
          route: '/sponsor'
        },
        {
          label: 'layout.details',
          icon: 'pi pi-wrench',
          route: '/details'
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
        }
      ]
    }
  ]

  // ✅ 管理员菜单
  if (userStore.isAuthed && userStore.user.role === 'ADMIN') {
    items.push({
      label: 'appMenu.admin',
      items: [
        {
          label: 'layout.statistic',
          icon: 'pi pi-chart-bar'
        },
        {
          label: 'layout.knowledgeMgr',
          icon: 'pi pi-file-edit'
        },
        {
          label: 'layout.illustrationMgr',
          icon: 'pi pi-palette'
        }
      ]
    })
  }

  // ✅ 用户菜单
  if (userStore.isAuthed && userStore.user.role === 'USER') {
    items.push({
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
          shortcut: '⌥+Q',
          command: () => userStore.logout()
        }
      ]
    })
  }

  items.push({separator: true})

  return items
}

const itemsMenu = buildMenu()

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

</script>

<template>
  <div class="p-1 flex-row flex justify-between items-center" style="backdrop-filter: blur(20px)">
    <div class="ml-0">
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
        <template #submenuheader="{ item }">
          <span class="text-primary font-bold">{{ t(item.label as string) }}</span>
        </template>
        <template #item="{ item, props }">
          <router-link v-slot="{ href, navigate }" :to="item.route" custom>
            <a v-ripple :href="href" v-bind="props.action" @click="navigate">
              <span :class="item.icon"/>
              <span class="ml-2">{{ t(item.label as string) }}</span>
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
        {{ getUserName }}
      </span>
              <span class="text-sm">{{ t(getUserRole) }}</span>
            </div>
          </button>
        </template>
      </Menu>

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
      <span>搜寻</span>
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
      />
    </div>
  </div>

  <Dialog v-model:visible="showLoginModal" :show-header="false" modals :style="null">
    <LoginForm/>
  </Dialog>


</template>

<style>

</style>