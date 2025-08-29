<script setup lang="ts">
import {useTheme} from "vuetify/framework";
import useUserStore from "~/store/userStore";
import useThemeStore from "~/store/themeStore";
import {useI18n} from "vue-i18n";
import LoginForm from "~/components/LoginForm.vue";

const {t, locale} = useI18n()

const router = useRouter()
const route = useRoute()
const theme = useTheme()
const themeStore = useThemeStore()
const userStore = useUserStore()

type MenuItem = {
  text: string,
  icon: string,
  subtitle?: string
  path: string
}

const items = computed<MenuItem[]>(() => {
  return !(userStore.isAuthed && userStore.user.role==='ADMIN')?[
    {
      text: 'menu.home',
      icon: 'mdi-home',
      path: '/'
    },
    {
      text: 'menu.user',
      icon: 'mdi-account',
      path: '/profile'
    },
    {
      text: 'menu.knowledge',
      icon: 'mdi-book',
      path: '/knowledge'
    },
  ]:[
    {
      text: 'menu.home',
      icon: 'mdi-home',
      path: '/'
    },
    {
      text: 'menu.overview',
      icon: 'mdi-view-dashboard-variant',
      path: '/admin'
    },
    {
      text: 'menu.user',
      icon: 'mdi-account',
      path: '/profile'
    },
    {
      text: 'menu.knowledge',
      icon: 'mdi-book',
      path: '/knowledge'
    },
    {
      text: 'menu.write',
      icon: 'mdi-book-edit',
      path: '/admin/knowledge'
    },
  ]
})

const bottomItems = computed<MenuItem[]>(() => {
  let list: MenuItem[] = []
  if (userStore.isAuthed) list.push({text: 'menu.logout', icon: 'mdi-exit-to-app', path: 'exit'})
  if (userStore.user.role === 'ADMIN' && userStore.isAuthed) {
    list.unshift({text: 'menu.admin', icon: 'mdi-folder', path: 'admin'})
  }
  return list
})

const menuClick = (path: string) => {
  themeStore.toggleMenuDisplay()
  navigateTo({
    path: `/${locale.value}${path}`,
  })
}

const showLoginCard = ref<boolean>(false)
const closeLoginCard = () => setTimeout(() => showLoginCard.value = false, 1000)

const loginOrRegClick = () => {
  if (!userStore.user.id) {
    showLoginCard.value = true
  }
}

const bottomClick = (path: string) => {
  if (path === 'admin') {
    // 登入
    if (userStore.isAuthed && userStore.user.role==='ADMIN') {
      navigateTo({path: `/${locale.value}/admin`})
    }
  } else if (path === 'exit' && userStore.isAuthed) {
    // 登出
    userStore.clearUserData()
    themeStore.toggleMenuDisplay()
    if (route.path !== `/${locale.value}`) {
      router.replace(`/${locale.value}`)
    }
  }
}

onMounted(() => {
  // userStore.user.email = 'swdw'
  // console.log(locale)
})


</script>

<template>

  <div class="root">
    <div class="up-part">
      <div class="primary-info-root">
        <div class="primary-line">
          <div :style="{backgroundColor: theme.global.current.value.colors.primary}" class="color-desc"/>
        </div>
        <div class="profile-root">
          <p class="username">{{ userStore.user.username ? userStore.user.username : userStore.isAuthed?t('menu.noUsername'):t('menu.visitor') }}</p>
          <p @click="loginOrRegClick" class="email">{{
              userStore.user.email ? userStore.user.email : t('menu.clickLogin')
            }}</p>
        </div>

      </div>

      <v-list
          class="ma-3 mt-1"
      >
        <v-list-subheader>{{ t('menu.title') }}</v-list-subheader>

        <v-list-item
            v-for="(item, i) in items"
            :key="i"
            :value="item"
            color="primary"
            rounded="lg"
            class="mt-2"
            @click="menuClick(item.path)"
            density="comfortable"
        >
          <template v-slot:prepend>
            <v-icon :icon="item.icon"></v-icon>
          </template>
          <v-list-item-title v-text="t(item.text)"></v-list-item-title>
        </v-list-item>
      </v-list>
    </div>
    <div class="down-part">

      <v-btn
          variant="plain"
          v-for="(item, i) in bottomItems"
          :key="item.text"
          class="bottom-btn"
          @click="bottomClick(item.path)"
      >{{ t(item.text) }}
        <template v-slot:prepend>
          <v-icon>{{ item.icon }}</v-icon>
        </template>
      </v-btn>
    </div>

  </div>

  <div class="text-center pa-4">
    <v-dialog
        v-model="showLoginCard"
        width="auto"
    >
      <LoginForm
        :close-login-card="closeLoginCard"
      ></LoginForm>
    </v-dialog>
  </div>


</template>

<style lang="less" scoped>
.root {
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.primary-info-root {
  height: 90px;
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-content: center;

  .primary-line {
    height: 100%;
    width: 4px;
    display: flex;
    flex-direction: column;
    justify-content: center;

    .color-desc {
      height: 60%;
    }
  }
}

.profile-root {
  padding: 16px;

  .username {
    font-size: 1.4rem;
    font-weight: bold;
  }

  .email {
    font-size: 1rem;
    opacity: 0.8;
    text-decoration: underline;
    transition: ease-in-out 200ms;
  }

  .email:hover {
    opacity: 0.5;
  }
}

.down-part {
  padding-bottom: 20%;
  margin: 0 16px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;

  .bottom-btn {
    margin-top: 6px;
    width: 100%;
    text-transform: none !important;
  }
}
</style>