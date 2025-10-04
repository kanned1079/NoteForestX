<script lang="ts" setup>

</script>

<template>

  <NuxtLayout>
<!--    <Transition-->
<!--        name="fade"-->
<!--        mode="out-in"-->
<!--        appear-->
<!--    >-->
      <NuxtPage />
<!--    </Transition>-->
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


 :root{
   --body-bg: #f8fafc;
   --text: #111;
 }

@media (prefers-color-scheme: dark){
  :root{
    --body-bg: #111113; /* 你提到想用的更深色 */
    --text: #eee;
  }
}

html,body{
  height:100%;
  margin:0;
}

body{
  background: var(--body-bg);
  color: var(--text);
  transition: background .25s ease, color .25s ease;
}

</style>