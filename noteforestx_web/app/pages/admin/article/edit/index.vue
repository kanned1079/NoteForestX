<script setup lang="ts">
definePageMeta({
  layout: 'default'
})

import {useRouter} from "vue-router";

const router = useRouter();

onMounted(() => {
  setTimeout(() => router.back(), 2000)
})

</script>

<template>
  <div class="text-base opacity-60 mt-20 p-4">
    Please access the editing page from the article list.
  </div>
</template>