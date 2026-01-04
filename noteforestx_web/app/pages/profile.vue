<script setup lang="ts">
definePageMeta({
  disableShortcuts: true,
})

import {useRouter} from "vue-router";
import useUserStore from "~/store/userStore";
import WidthTest from "~/components/RedesignedComponents/WidthTest.vue";
import dayjs from "dayjs";
import {useI18n} from "vue-i18n";
import {useScrollFadeIn} from "~/composables/useScrollFadeIn";
import {useToast} from "primevue/usetoast";
import {ref, watch} from "vue"; // 补充缺失的 ref 导入

const router = useRouter();
const {t} = useI18n();
const toast = useToast();
const userStore = useUserStore()
const showEditUsernameDialog = ref<boolean>(false)
const showUpdatePwdDialog = ref<boolean>(false)

const editUsername = ref<string>('')
const editPwd = ref<{
  previousPwd: string
  newPwd: string
}>({
  previousPwd: "",
  newPwd: ""
})

const callSaveUsername = async () => {
  if (!editUsername.value) {
    return  toast.add({
      severity: "warn",
      summary: t("profile.username.toast.errorSummary"),
      detail: t("profile.username.toast.warnEmpty"),
      life: 4000,
    })
  }
  try {
    const token = useCookie("access_token", {})
    const data = await $fetch<{
      message: string,
      new_username: string
    }>(`/api/user/${userStore.user.id}/username`, { // PATCH URL
      method: "PATCH",
      headers: {
        Authorization: `Bearer ${token.value}`  // ✅ 这里改成 Bearer
      },
      body: {
        username: editUsername.value
      }
    })
    userStore.user.username = data.new_username
    showEditUsernameDialog.value = false
    toast.add({
      severity: "success",
      summary: "Success", // 如需国际化可添加到 profile 节点
      detail: t("profile.username.toast.success", {username: data.new_username}), // 插值传参
      life: 4000,
    })
  } catch (err: any) {
    console.error(err)
    toast.add({
      severity: "error",
      summary: t("profile.toast.errorSummary"),
      detail: `${err}`,
      life: 4000,
    })
  } finally {
    // ...
  }
}

const callSavePassword = async () => {
  if (editPwd.value.newPwd.length < 6) {
    return  toast.add({
      severity: "warn",
      summary: t("profile.password.toast.errorSummary"),
      detail: t("profile.password.toast.warnLength"),
      life: 4000,
    })
  }
  try {
    const token = useCookie("access_token", {})
    const data = await $fetch<{
      message: string,
    }>(`/api/user/${userStore.user.id}/password`, { // PATCH URL
      method: "PATCH",
      headers: {
        Authorization: `Bearer ${token.value}`  // ✅ 这里改成 Bearer
      },
      body: {
        previous_password: editPwd.value.previousPwd,
        new_password: editPwd.value.newPwd,
      }
    })
    showUpdatePwdDialog.value = false
    toast.add({
      severity: "success",
      summary: "Success", // 如需国际化可添加到 profile 节点
      detail: t("profile.password.toast.success"),
      life: 4000,
    })
  } catch (err: any) {
    console.error(err)
    toast.add({
      severity: "error",
      summary: t("profile.toast.errorSummary"),
      detail: `${err}`,
      life: 4000,
    })
  } finally {
    // ...
  }
}

const onClickUpdateUsername = () => {
  editUsername.value = userStore.user.username || ''
  showEditUsernameDialog.value = true
}

const onClickUpdatePwd = () => {
  showUpdatePwdDialog.value = true
}

watch(() => userStore.user.id, (newVal: string) => {
  if (!newVal || newVal === "") {
    setTimeout(() => router.replace({path: "/"}), 2000)
  }
})

useScrollFadeIn({
  selector: '.animate-card-profile',
  y: 60,
  stagger: 0.15
})
</script>

<template>
  <div class="mt-10" v-if="userStore.user.id">
    <p class="text-lg font-bold mb-8 animate-card-profile">{{ t("profile.yourInfo") }}</p>

    <div class="space-y-8 animate-card-profile">
      <div class="space-y-2">
        <p class="text-3xl font-semibold">{{ userStore.user.username?userStore.user.username:userStore.user.email }}</p>
        <span
            @click="onClickUpdateUsername"
            class="text-sm opacity-70 mt-2 flex flex-row items-center  hover:underline cursor-pointer">
                {{ t("profile.username.editTip") }}
                <span class="pi pi-angle-right font-light"></span>
              </span>
        <p class="text-sm opacity-70">{{ t("profile.username.desc") }}</p>
      </div>

      <div class="space-y-2">
        <p class="text-3xl font-semibold">{{ userStore.user.id }}</p>
        <p class="text-sm opacity-70">{{ t("profile.userId.desc") }}</p>
      </div>

      <div class="space-y-2" v-if="userStore.user.created_at">
        <p class="text-3xl font-semibold">{{ dayjs(userStore.user.created_at).format('YYYY-MM-DD HH:mm:ss') }}</p>
        <p class="text-sm opacity-70">{{ t("profile.createTime.desc") }}</p>
      </div>
    </div>

    <div class="animate-card-profile">
      <p class="text-lg font-bold mt-16 mb-4">{{ t("profile.resetPassword") }}</p>
      <span
          @click="onClickUpdatePwd"
          class="text-sm opacity-70 flex flex-row items-center justify-start hover:underline">
          {{ t("profile.password.editTip") }}
          <span class="pi pi-refresh text-sm"></span>
      </span>
    </div>
  </div>


  <div class="mt-10" v-else>
    <div class="space-y-2 animate-card-profile">
      <p class="text-3xl font-semibold">{{ t("profile.notLogin.title") }}</p>
      <p class="text-sm opacity-70">{{ t("profile.notLogin.desc") }}</p>
    </div>
  </div>

  <Dialog :show-header="true" :dismissable-mask="true" v-model:visible="showEditUsernameDialog"
          :header="t('profile.username.dialog.title')" :style="{ width: '25rem' }"
          position="top" :modal="true"
          :draggable="false">
    <div class="flex flex-col gap-2">
      <label for="link">{{ t("profile.username.dialog.label") }}</label>
      <IconField>
        <InputIcon class="pi pi-user"></InputIcon>
        <InputText autofocus size="small" id="link" v-model="editUsername"
                   :placeholder="t('profile.username.dialog.placeholder')"
                   class="w-full"
                   @keyup.enter="callSaveUsername"
        />
      </IconField>
    </div>
    <span class="text-surface-500 dark:text-surface-400 block mb-4 text-sm mt-2">
        {{ t("profile.username.dialog.hint") }}
      </span>

    <div class="flex justify-end gap-2 mt-4">
      <Button icon="pi pi-replay" class="p" size="small" type="button" :label="t('universal.illustration.cancel')" severity="secondary"
              @click="showEditUsernameDialog = false"></Button>
      <Button icon="pi pi-save" size="small" type="button" :label="t('universal.illustration.save')" @click="callSaveUsername"></Button>
    </div>
  </Dialog>

  <!--  // -->

  <Dialog :show-header="true" :dismissable-mask="true" v-model:visible="showUpdatePwdDialog"
          :header="t('profile.password.dialog.title')" :style="{ width: '25rem' }"
          position="top" :modal="true"
          :draggable="false">
    <div class="flex flex-col space-y-4">
      <div class="space-y-2">
        <label for="link">{{ t("profile.password.dialog.oldPwdLabel") }}</label>
        <IconField>
          <InputIcon class="pi pi-key"></InputIcon>
          <InputText type="password" autofocus size="small" id="link" v-model="editPwd.previousPwd"
                     :placeholder="t('profile.password.dialog.oldPwdPlaceholder')"
                     class="w-full"
          />
        </IconField>
      </div>

      <div class="space-y-2">
        <label for="link">{{ t("profile.password.dialog.newPwdLabel") }}</label>
        <IconField>
          <InputIcon class="pi pi-key"></InputIcon>
          <InputText type="password" autofocus size="small" id="link" v-model="editPwd.newPwd"
                     :placeholder="t('profile.password.dialog.newPwdPlaceholder')"
                     class="w-full"
                     @keyup.enter="callSavePassword"
          />
        </IconField>
      </div>

    </div>
    <span class="text-surface-500 dark:text-surface-400 block mb-4 text-sm mt-2">
        {{ t("profile.password.dialog.hint") }}
      </span>

    <div class="flex justify-end gap-2 mt-4">
      <!-- 修复：此处关闭的应该是 showUpdatePwdDialog -->
      <Button icon="pi pi-replay" class="p" size="small" type="button" :label="t('universal.illustration.cancel')" severity="secondary"
              @click="showUpdatePwdDialog = false"></Button>
      <Button icon="pi pi-save" size="small" type="button" :label="t('universal.illustration.save')" @click="callSavePassword"></Button>
    </div>
  </Dialog>
</template>

<style scoped>

</style>