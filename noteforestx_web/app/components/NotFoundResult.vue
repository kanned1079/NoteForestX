<script setup lang="ts">
import {Icon} from "@vicons/utils"
import {useI18n} from "vue-i18n";

const {t} = useI18n()

</script>

<template>
  <div class="w-full">
    <Message class="w-auto" severity="warn">
      <template #container>
        <div class="text-center p-4">
          <div class="text-2xl">Not Found</div>
        </div>
      </template>

    </Message>

  </div>
</template>

<style scoped>

</style>