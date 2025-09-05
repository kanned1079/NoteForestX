<script setup lang="ts">
import useUserStore from "~/store/userStore";
import useThemeStore from "~/store/themeStore";
import { useI18n } from "vue-i18n";
import LoginForm from "~/components/LoginForm.vue";

const { t, locale } = useI18n()
const router = useRouter()
const route = useRoute()
const themeStore = useThemeStore()
const userStore = useUserStore()

type MenuItem = {
  text: string,
  icon: string,
  subtitle?: string
  path: string
}

const items = computed<MenuItem[]>(() => {
  return !(userStore.isAuthed && userStore.user.role === 'ADMIN') ? [
    { text: 'menu.home', icon: 'mdi-home', path: '/' },
    { text: 'menu.user', icon: 'mdi-account', path: '/profile' },
    { text: 'menu.knowledge', icon: 'mdi-book', path: '/knowledge' },
  ] : [
    { text: 'menu.home', icon: 'mdi-home', path: '/' },
    { text: 'menu.overview', icon: 'mdi-view-dashboard-variant', path: '/admin' },
    { text: 'menu.user', icon: 'mdi-account', path: '/profile' },
    { text: 'menu.knowledge', icon: 'mdi-book', path: '/knowledge' },
    { text: 'menu.write', icon: 'mdi-book-edit', path: '/admin/knowledge' },
  ]
})

const bottomItems = computed<MenuItem[]>(() => {
  let list: MenuItem[] = []
  if (userStore.isAuthed) list.push({ text: 'menu.logout', icon: 'mdi-exit-to-app', path: 'exit' })
  if (userStore.user.role === 'ADMIN' && userStore.isAuthed) {
    list.unshift({ text: 'menu.admin', icon: 'mdi-folder', path: 'admin' })
  }
  return list
})

const menuClick = (path: string) => {
  themeStore.toggleMenuDisplay()
  navigateTo({ path: `/${locale.value}${path}` })
}

const showLoginCard = ref<boolean>(false)
const closeLoginCard = () => setTimeout(() => showLoginCard.value = false, 1000)

const loginOrRegClick = () => {
  if (!userStore.user.id) showLoginCard.value = true
}

const bottomClick = (path: string) => {
  if (path === 'admin') {
    if (userStore.isAuthed && userStore.user.role === 'ADMIN') {
      navigateTo({ path: `/${locale.value}/admin` })
    }
  } else if (path === 'exit' && userStore.isAuthed) {
    userStore.clearUserData()
    themeStore.toggleMenuDisplay()
    if (route.path !== `/${locale.value}`) {
      router.replace(`/${locale.value}`)
    }
  }
}
</script>

<template>
  <div class="h-screen flex flex-col justify-between">
    <!-- 顶部用户信息 + 菜单 -->
    <div>
      <div class="h-[90px] flex items-center">
        <!-- 左侧彩色竖线 -->
        <div class="h-full w-1 flex items-center">
          <div class="h-[60%] w-full" :style="{ backgroundColor: themeStore.primaryColor }"></div>
        </div>
        <!-- 用户信息 -->
        <div class="px-4">
          <p class="text-xl font-bold">
            {{ userStore.user.username ? userStore.user.username : userStore.isAuthed ? t('menu.noUsername') : t('menu.visitor') }}
          </p>
          <p
              @click="loginOrRegClick"
              class="text-base opacity-80 underline transition duration-200 cursor-pointer hover:opacity-50"
          >
            {{ userStore.user.email ? userStore.user.email : t('menu.clickLogin') }}
          </p>
        </div>
      </div>

      <!-- 菜单 -->
      <div class="mt-3 px-3">
        <p class="text-gray-500 text-sm mb-2">{{ t('menu.title') }}</p>
        <div v-for="(item, i) in items" :key="i" class="mt-2">
          <button
              @click="menuClick(item.path)"
              class="flex items-center w-full px-3 py-2 rounded-lg hover:bg-gray-100 transition"
          >
            <span class="material-icons mr-3"><v-icon>{{ item.icon }}</v-icon></span>
            <span>{{ t(item.text) }}</span>
          </button>
        </div>
      </div>
    </div>

    <!-- 底部按钮 -->
    <div class="pb-[20%] mx-4 flex flex-col items-start">
      <button
          v-for="(item, i) in bottomItems"
          :key="item.text"
          @click="bottomClick(item.path)"
          class="mt-1 w-full flex items-center px-3 py-2 rounded-lg hover:bg-gray-100 transition"
      >
        <span class="material-icons mr-2">{{ item.icon }}</span>
        {{ t(item.text) }}
      </button>
    </div>
  </div>

  <!-- 登录弹窗 -->
  <div class="text-center p-4">
    <dialog v-if="showLoginCard" class="rounded-lg shadow-xl bg-white p-6">
      <LoginFormNew />
    </dialog>
  </div>
</template>