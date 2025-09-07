<script lang="ts" setup>
</script>

<template>
  <NuxtLayout>
    <Transition
        name="fade"
        mode="out-in"
        appear
    >
      <NuxtPage />
    </Transition>
  </NuxtLayout>
</template>

<style>
* {
  padding: 0;
  margin: 0;
}

/* 定义动画 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>