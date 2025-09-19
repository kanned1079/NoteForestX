<script setup lang="ts">
import {ref, onMounted} from "vue"
import {useRoute} from "vue-router";
import type {Illustration} from "../../types/illustration";
import dayjs from "dayjs"

const route = useRoute()
const config = useRuntimeConfig()
const illustrationUuid = route.params.id as string
const illustration = ref<Illustration | null>(null)

const fetchIllustrationList = async () => {
  try{
    const data = await $fetch<Illustration>(`/api/v1/illustration/${illustrationUuid}`, {
      method: "GET",
    })
    if (data) {
      illustration.value = data
    }
  } catch (err: any) {
    console.error("err: ",err)
  }
}

const openLink = (url: string) => {
  if (url) {
    window.open(url, '_blank', 'noopener noreferrer');
  }
}

onBeforeMount(() => {

})

onMounted(() => {
  fetchIllustrationList()
})
</script>

<template>
  <div class="mb-20 mt-10">
    <!-- 大布局 -->
    <div class="grid grid-cols-1 lg:grid-cols-8 gap-8" v-if="illustration">
      <!-- 左侧插画 (2/3) -->


      <div class="lg:col-span-1">
      </div>
      <div class="lg:col-span-3">
        <div class="flex flex-col">
          <img
              v-for="i in illustration?.images"
              :key="i.id"
              :src="`${config.public.apiBase}/illustration/file/${i.file_path}?size=medium`"
              alt="illustration"
              class="rounded-md mb-4 drop-shadow-md  object-contain"
              loading="lazy"
          />
        </div>
      </div>

      <!-- 右侧信息 (1/3) -->
      <div class="lg:col-span-3">
        <!-- 插画信息 -->
        <div>
          <Tag :severity="illustration?.limited?'warn':'primary'" :value="illustration?.limited?'限制型':'非限制型'"></Tag>
          <div class="text-2xl font-bold mt-2">{{ illustration?.name || "NO NAME" }}</div>
          <div class="font-light opacity-70 mt-2">{{ illustration?.description || "NO DESCRIPTION" }}</div>
          <div class="flex flex-row flex-wrap gap-2 mt-2">
            <Button
                v-for="i in illustration?.tags"
                :key="i.id"
                variant="link"
                severity="secondary"
                class="w-auto hover:underline p-0 text-sm font-light"
                @click="() => {console.log(i.id)}"
            >
              {{ `#${i.name}` }}
            </Button>
          </div>
          <div class="text-sm font-light opacity-60 mt-2">
            创建于 {{ dayjs(illustration?.created_at).format("YYYY/MM/DD HH:mm") }}
          </div>
          <Button
              link
              class="w-auto p-0 mt-2"
              size="small"
              @click="openLink(illustration?.link)"
          >
            转到原链接地址
          </Button>
        </div>

        <!-- 作者信息 -->
        <div class="mt-6 border-t pt-4">
          <div class="text-2xl font-bold hover:underline">{{ illustration?.author.name }}</div>
<!--          <div class="text-sm font-light mt-2 hover:underline">{{ illustration?.author.id }}</div>-->
          <div class="text-sm font-light opacity-60 mt-2">
            创建于 {{ dayjs(illustration?.author.created_at).format("YYYY/MM/DD HH:mm") }}
          </div>
          <Button
              link
              class="w-auto p-0 mt-2"
              size="small"
              @click="openLink(illustration?.author.link)"
          >
            转到作者主页
          </Button>

          <!-- 其他作品 -->
          <div class="text-md font-bold mt-4 hover:underline">该作者的其他作品</div>
          <div class="grid grid-cols-2 gap-3 mt-4">
<!--            <IllustrationItemPreview-->
<!--                v-for="i in 4"-->
<!--                :key="i"-->
<!--            />-->
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style >
</style>