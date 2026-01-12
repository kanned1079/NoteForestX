<script setup lang="ts">
import { useRoute } from "#vue-router";
import { MdPreview, MdCatalog } from 'md-editor-v3';
import { useRouter } from "vue-router";
import {ref, onMounted, onBeforeUnmount, onUnmounted} from "vue";
import dayjs from 'dayjs'
import { useScrollFadeIn } from "~/composables/useScrollFadeIn";
import ArticleHeader from "~/components/RedesignedComponents/ArticleHeader.vue";
import { useToast } from "primevue/usetoast";
import type { Article } from "~/types/article";
import useThemeStore from "~/store/themeStore";
import useActionStore from "~/store/actionStore";
import {useI18n} from "vue-i18n";

import '~/assets/css/md-style.css'
import {useDarkMode} from "~/composables/useDarkMode";

const {t} = useI18n()
const themeStore = useThemeStore()
const actionStore = useActionStore()
const route = useRoute();
const router = useRouter();
const toast = useToast();

const articleId = route.params.id as string;
const id = 'article-preview';

const visible = ref<boolean>(false);
const position = ref<"bottom" | "left">("left");

const currentArticle = ref<Article>({
  id: '',
  top: false,
  status: 'published',
  title: '',
  tags: [],
  content: ''
});

const { isDarkMode } = useDarkMode();

useScrollFadeIn({
  selector: '.animate-card-article-id',
  direction: 'up',
  x: 200,
  stagger: 0.1,
  duration: 0.6,
  start: 'top 90%',
  useScrollTrigger: false
});

const fetchArticleById = async (id: string) => {
  try {
    // const res = await $fetch<{
    //   id: string;
    //   message: string;
    //   article: Article | null;
    // }>(`/api/article/${id}`);

    const res = await useHttp().get<{
        id: string;
        message: string;
        article: Article | null;
    }>(`/v1/article/${id}`)

    console.log(res)

    if (res.article) {
      currentArticle.value = res.article;
    }
  } catch (err: any) {
    toast.add({
      severity: "error",
      summary: t('article.toast.loadFailed'),
      detail: String(err),
      life: 4500
    });
    router.back();
  }
};

const isTypingTarget = (target: EventTarget | null) => {
  if (!(target instanceof HTMLElement)) return false;

  const tag = target.tagName.toLowerCase();

  return (
      tag === 'input' ||
      tag === 'textarea' ||
      target.isContentEditable
  );
};

const onKeydown = (e: KeyboardEvent) => {
  if (e.key !== '/') return;

  if (e.shiftKey && e.key.toLowerCase() === 'q') {
    e.preventDefault()
    return router.back()
  }

  if (isTypingTarget(e.target)) return;

  e.preventDefault();
  visible.value = !visible.value;
};

watch(() => actionStore.triggerShowCatalog, (newVal: boolean) => {
  if (newVal) {
    actionStore.resetTriggerShowCatalog()
    visible.value = !visible.value
  }
})

const toPageTop = () => {
  console.log('1111')
  window.scrollTo({
    top: 0,
    behavior: 'smooth'
  })
}

fetchArticleById(articleId);

onMounted(() => {
  themeStore.setShowCatalog(true)
  themeStore.actionCenterMsgs = [
    `Press / to open the article's table of contents panel`,
    `Click the article title in the table of contents area to return to the top`,
    `Click a Tag in the article to get the TagId`
  ];

  window.addEventListener('keydown', onKeydown);
});

onBeforeUnmount(() => {
  window.removeEventListener('keydown', onKeydown);

});

onUnmounted(() => {
  themeStore.setShowCatalog(false)
  themeStore.actionCenterMsgs = []
})
</script>

<template>
  <div
      class="pt-[40px] px-4 md:px-6 lg:px-8"
  >
    <div class="card flex justify-center">
      <Drawer v-model:visible="visible" :header="t('article.catalog')" :position="position" class="!w-full md:!w-80 lg:!w-[30rem]">
        <div>
          <p class="text-xl font-bold mt-6 mb-4 hover:underline cursor-pointer" @click="toPageTop">{{ currentArticle.title }}</p>
          <MdCatalog :editorId="id" scrollElement="html" />
          <p class="mt-4 text-xs font-mono opacity-50">{{ articleId }}</p>
        </div>
      </Drawer>
    </div>

    <!-- 主体区域 -->
    <div class="flex gap-6">

      <!-- 正文 -->
      <main class="flex-1 min-w-0">
        <div class="mx-auto max-w-[720px] lg:max-w-[900px] animate-card-article-id">

          <ArticleHeader
              :title="currentArticle.title"
              :tags="currentArticle.tags"
              :createdAt="dayjs(currentArticle.created_at).format('YYYY-MM-DD')"
          />

          <MdPreview
              :id="id"
              :modelValue="currentArticle.content"
              :theme="isDarkMode ? 'dark' : undefined"
              preview-theme="github"
          />

        </div>
      </main>
    </div>
  </div>
</template>

<style>

aside::-webkit-scrollbar {
  width: 6px;
}
aside::-webkit-scrollbar-thumb {
  background-color: rgba(100,100,100,0.3);
  border-radius: 3px;
}

.md-editor {
  z-index: 1 !important;
  --md-border-color: #7c7c7c !important;
  border-radius: 6px;
  --md-border-active-color: #000 !important;
  --md-bk-color: rgba(0, 0, 0, 0.0) !important;
}

.md-editor-dark {
  --md-bk-color: rgba(0, 0, 0, 0.0) !important;
  --md-border-color: #989898 !important;
}

.operate-part-root {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  margin: 20px 0;

}



</style>