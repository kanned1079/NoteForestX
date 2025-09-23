<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import type { Illustration } from "../../types/illustration";
import dayjs from "dayjs";
import {useAsyncData} from "nuxt/app";

const route = useRoute();
const illustId = route.params.id as string;

const config = useRuntimeConfig();

const illustration = ref<Illustration | null>(null);

const fetchIllustrationById = async () => {
  try {
    const data = await $fetch<Illustration>(`/api/illustration/${illustId}`, {
    //   const data = await useFetch<Illustration>(`/api/illustration/${illustId}`, {
      method: "GET",
    });
    if (data) {
      illustration.value = data;
    }
  } catch (err: any) {
    console.error("err: ", err);
  }
};

const openLink = (url: string | undefined) => {
  if (url) {
    window.open(url, "_blank", "noopener noreferrer");
  }
};

const downloadImageClick = (path: string) => {
  const url = `${config.public.apiBase}/api/v1/illustration/file/${path}?size=original`;

  const link = document.createElement("a");
  link.href = url;
  link.download = path; // 文件名，可以自定义，例如 `${path}.jpg`
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
};

fetchIllustrationById()

onMounted(() => {

});
</script>

<template>
  <div class="mb-20 mt-10">
    <!-- 大布局 -->
    <div class="grid grid-cols-1 md:grid-cols-12 gap-8 px-4 lg:px-20">
      <!-- 左侧插画 -->
<!--      <div class="md:col-span-6">-->
<!--        <div class="flex flex-col">-->
<!--          <img-->
<!--              v-for="i in illustration?.images"-->
<!--              :key="i.id"-->
<!--              :src="`${config.public.apiBase}/api/v1/illustration/file/${i.file_path}?size=medium`"-->
<!--              alt="illustration"-->
<!--              class="rounded-md mb-4 drop-shadow-md object-contain"-->
<!--              loading="lazy"-->
<!--          />-->
<!--        </div>-->
<!--      </div>-->
      <!-- 左侧插画 -->
      <!-- 左侧插画 -->
      <div class="md:col-span-6">
        <div class="flex flex-col">
          <div
              v-for="(i, index) in illustration?.images"
              :key="i.id"
              class="relative group rounded-md mb-4 drop-shadow-md overflow-hidden"
          >
            <!-- 图片 -->
            <img
                :src="`${config.public.apiBase}/api/v1/illustration/file/${i.file_path}?size=medium`"
                alt="illustration"
                class="w-full h-full object-contain"
                loading="lazy"
            />

            <!-- 悬浮遮罩 -->
            <div
                class="absolute inset-0 bg-black/60 opacity-0 group-hover:opacity-100 transition-opacity duration-300 flex items-end"
            >
              <!-- 左下角信息 -->
              <div class="p-3 text-white text-sm space-y-1">
                <div class="text-sm font-thin">
                  {{ i.width }} x {{ i.height }}px
                </div>
<!--                <div class="text-sm font-thin">-->
<!--                  {{ dayjs(i?.created_at).format("YYYY/MM/DD HH:mm") }}-->
<!--                </div>-->
                <div class="text-sm font-light">
                  {{ i.id }}
                </div>

                <Button
                  variant="link"
                  class="p-0 underline"
                  @click.prevent="downloadImageClick(i.file_path)"
                >Get Original</Button>

              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 右侧信息 -->
      <div class="md:col-span-6">
        <!-- 插画信息 -->
        <div>
          <Tag :severity="illustration?.limited?'warn':'primary'" :value="illustration?.limited?'Limited (NSFW)':'Unlimited (SFW)'"></Tag>
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
              @click="openLink(illustration?.link as string)"
          >
            转到原链接地址
          </Button>
        </div>

        <!-- 作者信息 -->
        <div class="mt-6 border-t pt-4">
          <div class="text-2xl font-bold hover:underline">{{ illustration?.author.name }}</div>
          <div class="text-sm font-light opacity-60 mt-2">
            创建于 {{ dayjs(illustration?.author.created_at).format("YYYY/MM/DD HH:mm") }}
          </div>
          <Button
              link
              class="w-auto p-0 mt-2"
              size="small"
              @click="openLink(illustration?.author.link as string)"
          >
            转到作者主页
          </Button>

          <!-- 其他作品 -->
          <div class="text-md font-bold mt-4 hover:underline">该作者的其他作品</div>
          <div class="grid grid-cols-2 gap-3 mt-4">
            <IllustrationItemPreview
                v-for="i in 4"
                :key="i"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>