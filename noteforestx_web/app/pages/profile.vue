<script setup lang="ts">
import useUserStore from "~/store/userStore";
import WidthTest from "~/components/RedesignedComponents/WidthTest.vue";
import dayjs from "dayjs";
import {useScrollFadeIn} from "~/composables/useScrollFadeIn";

const userStore = useUserStore()

useScrollFadeIn({
  selector: '.animate-card-profile',
  y: 60,
  stagger: 0.15
})

</script>

<template>
  <div class="mt-10">
    <p class="text-lg font-bold mb-8 animate-card-profile">您的信息</p>

    <div class="space-y-8 animate-card-profile">
      <div class="space-y-2">
        <p class="text-3xl font-semibold">{{ userStore.user.username?userStore.user.username:userStore.user.email }}</p>
        <p class="text-sm opacity-70">这是您在发表评论时使用的用户名，如果您没有设置，将显示为邮箱地址。</p>
      </div>

      <div class="space-y-2">
        <p class="text-3xl font-semibold">{{ userStore.user.id }}</p>
        <p class="text-sm opacity-70">您在本站的用户唯一 ID，不能修改。</p>
      </div>

      <div class="space-y-2" v-if="userStore.user.created_at">
        <p class="text-3xl font-semibold">{{ dayjs(userStore.user.created_at).format('YYYY-MM-DD HH:mm:ss') }}</p>
        <p class="text-sm opacity-70">您创建账户时候的时间。</p>
      </div>
    </div>

    <div class="animate-card-profile">
      <p class="text-lg font-bold mt-16 mb-4">重置密码</p>
      <span class="text-sm opacity-70 flex flex-row items-center justify-start hover:underline">如果需要，轻点这里以重置密码。<span class="pi pi-refresh text-sm"></span> </span>
    </div>


  </div>
</template>

<style scoped>

</style>