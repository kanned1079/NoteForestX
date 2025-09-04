<script setup lang="ts">
import useThemeStore from "~/store/themeStore";

const themeStore = useThemeStore()
</script>

<template>
  <v-btn
      @click="themeStore.toggleMenuDisplay"
      variant="text"
      class="ml-1.5"
      border="0"
  >
    <template v-slot:prepend>
      <v-icon size="large">{{ 'mdi-format-align-left' }}</v-icon>
    </template>
    <p class="text-base opacity-90 !uppercase">Note Forest</p>
  </v-btn>
</template>

<style scoped>

</style>