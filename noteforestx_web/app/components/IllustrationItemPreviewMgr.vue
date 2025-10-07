<script setup lang="ts">
import { computed } from "vue";
import { useI18n } from "vue-i18n";
import type { Illustration } from "~/types/illustration";
const { locale } = useI18n();
const config = useRuntimeConfig();

const props = defineProps<{
  illustration: Illustration;
}>();

// mask：0-20% 完全不透明，20%-100% 渐变到透明
const maskGradient = "linear-gradient(to right, rgba(0,0,0,1) 0%, rgba(0,0,0,1) 20%, rgba(0,0,0,0) 100%)";

const maskStyle = computed(() => ({
  WebkitMaskImage: maskGradient, // Safari
  maskImage: maskGradient,       // 其他现代浏览器
  WebkitMaskSize: "100% 100%",
  maskSize: "100% 100%",
  WebkitMaskRepeat: "no-repeat",
  maskRepeat: "no-repeat",
}));

const onClickIllItem = () => {
  navigateTo({
    path: `/admin/illustration/${props.illustration.id}`,
  });
};

const onClickDelete = () => {
  // TODO: 删除逻辑
};
</script>

<template>

  <Card style="width: 18rem; overflow: hidden">
    <template #header>
<!--      <img alt="user header" :src="`${config.public.apiBase}/api/v1/illustration/file/${props.illustration?.images[0].file_path}?size=small`" />-->
      <div
          class="relative aspect-square w-full rounded-tl-md overflow-hidden
             flex items-center justify-center
             transition-shadow duration-200
             drop-shadow-sm hover:shadow-xl cursor-pointer"
          @click="onClickIllItem"
      >
        <!-- 背景模糊图 -->
        <img
            v-if="props.illustration?.images?.length"
            :src="`${config.public.apiBase}/api/v1/illustration/file/${props.illustration?.images[0].file_path}?size=small`"
            alt="blur-background"
            class="absolute inset-0 w-full h-full object-cover scale-110 filter blur-sm brightness-100 opacity-30"
        />

        <!-- 前景正常插画 -->
        <img
            v-if="props.illustration?.images?.length"
            :src="`${config.public.apiBase}/api/v1/illustration/file/${props.illustration?.images[0].file_path}?size=small`"
            alt="illustration"
            class="relative w-full h-full object-contain"
            loading="lazy"
        />
      </div>
    </template>
    <template #title>{{ props.illustration?.name }}</template>
    <template #subtitle>{{ props.illustration?.author?.name }}</template>
    <template #content>
      <p class="m-0">
        Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore sed cupiditate neque
        quas!
      </p>
    </template>
    <template #footer>
      <div class="flex gap-4 mt-1">
        <Button label="Cancel" severity="secondary" variant="outlined" class="w-full" />
        <Button label="Save" class="w-full" />
      </div>
    </template>
  </Card>
</template>

<style scoped>
/* 如果需要更精细的浏览器兼容处理，可以在这里加额外 CSS */
/* 目前 mask-image 的行内样式已包含 -webkit- 前缀，覆盖大多数情况 */
</style>