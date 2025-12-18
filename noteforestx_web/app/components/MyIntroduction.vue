<script setup lang="ts">
import MyWork from "~/components/MyIntro/MyWork.vue"
import ContactMe from "~/components/MyIntro/ContactMe.vue";
import useThemeStore from "~/store/themeStore";

const themeStore = useThemeStore()

</script>

<template>
  <section class="relative">
    <BackgroundGrid />
    <!-- 内容层 -->
    <div class="relative z-10" id="work-section">
      <MyWork v-if="themeStore.workTab==='0'" />
      <ContactMe v-if="themeStore.workTab==='1'" />
    </div>
  </section>
</template>

<style scoped>
.aaadwerf {
  padding: 0;
}
</style>