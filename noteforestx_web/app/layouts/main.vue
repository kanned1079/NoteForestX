<script setup lang="ts">
import { ref, onBeforeMount, onMounted } from "vue"
import AppHeader from "../components/AppHeader.vue"
import { useRoute } from 'vue-router';
import { useI18n } from "vue-i18n"
const { locale } = useI18n()
const route = useRoute();

</script>

<template>
  <div>
    <AppHeader />
    <div class="">
      <Toast />
      <ConfirmDialog />
      <slot />
    </div>
  </div>
</template>

<style scoped>
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.5s ease;
}

.fade-slide-enter-from {
  opacity: 0;
  transform: translateY(-20px);
}
.fade-slide-enter-to {
  opacity: 1;
  transform: translateY(0);
}

.fade-slide-leave-from {
  opacity: 1;
  transform: translateY(0);
}
.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(-20px);
}
</style>