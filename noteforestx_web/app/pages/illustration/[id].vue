<script setup lang="ts">
import {useRoute} from "vue-router";

const route = useRoute()

const illustrationUuid = route.params.id as string

onMounted(() => {
  console.log(illustrationUuid)
})

</script>

<template>
{{ illustrationUuid }}
</template>

<style scoped>

</style>