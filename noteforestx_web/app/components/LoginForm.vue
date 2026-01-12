<template>
  <!--  <div class="bg-surface-50 dark:bg-surface-950 px-6 py-20 md:px-20 lg:px-80">-->
  <div class="bg-surface-0 dark:bg-surface-900 p-8 md:p-2 rounded-2xl w-full max-w-sm mx-auto flex flex-col gap-8 mt-4 min-w-[320px]">
    <div class="flex flex-col items-center gap-4">
      <div class="flex items-center gap-4">
      </div>
      <div class="flex flex-col items-center gap-2 w-full">
        <div class="text-surface-900 dark:text-surface-0 text-2xl font-semibold leading-tight text-center w-full">
          {{ t("login.page.title") }}
        </div>
        <div class="text-center w-full">
            <span class="text-surface-700 dark:text-surface-200 leading-normal text-sm">
              {{ t("login.page.subtitle") }}
            </span>
        </div>
      </div>
    </div>
    <div class="flex flex-col gap-6 w-full">
      <div class="flex flex-col gap-2 w-full">
        <label for="email1" class="text-surface-900 text-sm dark:text-surface-0 font-medium leading-normal">
          {{ t("login.form.labels.email") }}
        </label>
        <InputText size="small" v-model="formData.email" id="email1" type="text" placeholder="Email address" class="w-full px-3 py-2 shadow-sm rounded-lg" />
      </div>
      <div class="flex flex-col gap-2 w-full">
        <label for="password1" class="text-surface-900 text-sm dark:text-surface-0 font-medium leading-normal">
          {{ t("login.form.labels.password") }}
        </label>
        <InputText @keyup.enter.stop="callLoginReq" size="small" v-model="formData.password"  id="password1" type="password" placeholder="Password" class="w-full px-3 py-2 shadow-sm rounded-lg" />
        <a class="text-xs cursor-pointer hover:text-primary-emphasis opacity-80 hover:underline">
          {{ t("login.form.forgotPwd") }}
        </a>
      </div>
    </div>
    <Button size="small" @click="callLoginReq" :label="t('login.button.signIn')" icon="pi pi-user" class="w-full py-2 rounded-lg flex justify-center items-center gap-2">
      <template #icon>
        <i class="pi pi-user text-base! leading-normal!" />
      </template>
    </Button>
  </div>
  <!--  </div>-->
</template>

<script lang="ts" setup>
const props = defineProps<{
  closeDialog: () => void
}>()

import {useI18n} from "vue-i18n";
import type {User} from '~/types/user'
import { ref } from 'vue';
import {useToast} from "primevue/usetoast";
import useUserStore from "~/store/userStore";

const {t} = useI18n(); // 初始化 i18n
const token = useCookie('access_token', {})
const toast = useToast()
const userStore = useUserStore()
const formData = ref<{
  email: string,
  password: string
}>({
  email: '',
  password: ''
})

const callLoginReq = async () => {
  try {
    // const data = await $fetch<{
    //   user: User,
    //   token: string
    // }>(`/api/public/login`, {
    //   method: "POST",
    //   body: {
    //     ...formData.value
    //   }
    // })

    const data = await useHttp().post<{
      user: User,
      token: string
    }>(`/v1/public/user/login`, {
      ...formData.value
    })
    userStore.user = data.user
    token.value = data.token
    userStore.isAuthed = true
    props.closeDialog()
    toast.add({
      severity: 'success',
      summary: t("login.toast.loginSuccessSummary"),
      detail: t("login.toast.loginSuccessDetail", { email: data.user.email }), // 插值传参
      life: 3000
    });
    setTimeout(() => navigateTo("/profile"), 1000)
  } catch (err :any) {
    toast.add({
      severity: 'warn',
      summary: t("login.toast.loginErrorSummary"),
      detail: `${err}`,
      life: 3000
    });
  }
}
</script>