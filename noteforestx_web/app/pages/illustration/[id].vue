<script setup lang="ts">
import {ref} from "vue"
import {useRoute} from "vue-router";
import type {Illustration} from "../../types/illustration";
const config = useRuntimeConfig()
const route = useRoute()

const illustrationUuid = route.params.id as string
var illustration = ref<Illustration | null>(null)

const fetchIllustrationList = async () => {
  try{
    // illustrationList.value = []
    const data = await $fetch<Illustration>(`/api/v1/illustration/${illustrationUuid}`, {
      method: "GET",
      // params: {
      //   page: page.value,
      //   size: size.value,
      //   show_limited: showLimited.value
      // }
    })
    if (data) {

      illustration.value = data
      console.log(data)
    }

  } catch (err: any) {
    console.log("err: ",err)
  }
}

fetchIllustrationList()

onMounted(() => {
  console.log(illustrationUuid)
})

</script>

<template>
<div class="mb-20">
  <img
      v-for="i in illustration?.images"
      :src="`${config.public.apiBase}/illustration/file/${i.file_path}?size=medium`"
      alt="illustration"
      class="rounded-xl mb-4"
      loading="lazy"
  />

  <div class="text-lg font-bold mt-10">{{ illustration.name }}</div>
  <div class="font-light opacity-70 mt-2">{{ illustration.description }}</div>

  <Tag
      v-for="i in illustration?.images"
      icon="pi pi-info-circle" severity="info" value="Info"></Tag>

</div>


</template>

<style scoped>

</style>