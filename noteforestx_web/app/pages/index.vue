<script setup lang="ts">
definePageMeta({
  layout: 'empty',
  layoutTransition: true
})
import {useI18n} from "vue-i18n";
const {t} = useI18n()
</script>

<template>
  <div>
    index {{ t('hello') }}
  </div>
</template>

<style scoped>

</style>