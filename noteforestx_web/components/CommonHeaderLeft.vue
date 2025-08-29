<script setup lang="ts">
// import useThemeStore from "~/store/themeStore";
//
// const themeStore = useThemeStore()
</script>

<template>
  <v-btn
    @click=""
    style="margin-left: 6px;"
    variant="flat"
  >
    <template v-slot:prepend>
      <v-icon size="large">{{ 'mdi-format-align-left' }}</v-icon>
    </template>
    <p class="app-name">Note Forest</p>
  </v-btn>
</template>

<style scoped>
.app-name {
  font-size: 1rem;
  opacity: 0.9;
  text-transform: uppercase !important;
}
</style>