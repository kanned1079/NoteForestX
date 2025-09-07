<script lang="ts" setup>
import {ref, computed} from "vue";
import {updateSurfacePalette} from '@primeuix/themes';
import {useI18n} from "vue-i18n";
import {Icon} from '@vicons/utils'
import {useRouter} from 'vue-router';
import {TicketSharp} from "@vicons/ionicons5"
import useUserStore from "@/store/userStore";

const {t, setLocale} = useI18n()
const router = useRouter();
const userStore = useUserStore()

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
          icon: 'pi pi-book'
        },
        {
          label: 'layout.illustrationLib',
          icon: 'pi pi-images',
          route: '/illustration'
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
          shortcut: '⌥+Q'
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

onMounted(() => {
  console.log(userStore)
})

</script>

<template>
  <div class="p-3 flex-row flex justify-between items-center">
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
    <div class="mr-3">
<!--      <Button size="small" variant="link" type="button" @click="" aria-haspopup="true">-->
<!--        {{ t('layout.knowledgeLib') }}-->
<!--      </Button>-->
<!--      <Button size="small" variant="link" type="button" @click="" aria-haspopup="true">-->
<!--        {{ t('layout.illustrationLib') }}-->
<!--      </Button>-->
<!--      <Button size="small" variant="link" type="button" @click="" aria-haspopup="true">-->
<!--        {{ t('layout.aboutMe') }}-->
<!--      </Button>-->
<!--      <Button size="small" variant="link" type="button" @click="onClickSetLocale" aria-haspopup="true"-->
<!--              aria-controls="lang_menu">-->
<!--        <i class="pi pi-language"></i>-->
<!--        语言切换-->
<!--      </Button>-->
<!--      <TieredMenu size="small" appendTo="body" ref="langMenuRef" :model="items" popup class="mt-2 w-auto">-->
<!--        <template #item="{ item }">-->
<!--          <div class="flex flex-row items-center justify-between">-->
<!--            <div class="flex justify-start items-center flex-row space-x-1.5 pt-1.5 pb-1.5 pl-1.5 mb-1.25">-->
<!--              <span class="drop-shadow-sm" :class="`fi fi-${item.code} rounded-sm`"></span>-->
<!--              <span class="text-sm">{{ item.label }}</span>-->
<!--            </div>--
<!--            <div>-->
<!--              <i v-if="item.code==='hk'" class="pi pi-check font-bold text-xs text-cyan-800 mr-3"></i>-->
<!--            </div>-->
<!--          </div>-->
<!--        </template>-->
<!--      </TieredMenu>-->
    </div>
  </div>

  <Dialog v-model:visible="showLoginModal" :show-header="false" modals :style="null">
    <LoginForm/>
  </Dialog>


</template>

<style>

</style>