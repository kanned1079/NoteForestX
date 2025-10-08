<script setup lang="ts">
// import {Icon} from "@vicons/utils"
// import {useI18n} from "vue-i18n";
// import {HelpCircleSharp} from "@vicons/ionicons5"

const props = defineProps<{
  title?: string
  text?: string
}>()

</script>

<template>
  <div class="w-full">
    <Message class="w-auto" severity="warn">
      <template #container>
        <div class="text-center pt-10 pb-12 pl-4 pr-4">
<!--          <Icon size="7rem"> <HelpCircleSharp /> </Icon>-->
          <div class="text-2xl">{{ props?.title }}</div>
          <div class="opacity-70 mt-2">{{ props.text }}</div>
        </div>
      </template>
    </Message>
  </div>
</template>

<style scoped>

</style>