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

import '~/assets/css/md-style.css'
import {useDarkMode} from "~/composables/useDarkMode";

/* ---------------- 基础逻辑 ---------------- */

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

/* ---------------- 动画 ---------------- */

useScrollFadeIn({
  selector: '.animate-card-article-id',
  direction: 'up',
  x: 200,
  stagger: 0.1,
  duration: 0.6,
  start: 'top 90%',
  useScrollTrigger: false
});

/* ---------------- 拉文章 ---------------- */

const fetchArticleById = async (id: string) => {
  try {
    const res = await $fetch<{
      id: string;
      message: string;
      article: Article | null;
    }>(`/api/admin/article/${id}`);

    if (res.article) {
      currentArticle.value = res.article;
    }
  } catch (err: any) {
    toast.add({
      severity: "error",
      summary: "加载失败，请返回重试",
      detail: String(err),
      life: 4500
    });
    router.back();
  }
};

/* ---------------- 快捷键核心逻辑 ---------------- */

/**
 * 判断当前是否在输入环境中
 */
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
  // 只处理 /
  if (e.key !== '/') return;

  if (e.shiftKey && e.key.toLowerCase() === 'q') {
    e.preventDefault()
    return router.back()
  }

  // 输入状态下放行
  if (isTypingTarget(e.target)) return;

  // 阻止浏览器滚动 / 查找
  e.preventDefault();

  // 切换目录
  visible.value = !visible.value;
};

watch(() => actionStore.triggerShowCatalog, (newVal: boolean) => {
  if (newVal) {
    actionStore.resetTriggerShowCatalog()
    // if (articleId === 'new') return saveNewArticle()
    // else return  saveArticle()
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

/* ---------------- 生命周期 ---------------- */


fetchArticleById(articleId);


onBeforeMount(() => {

})

onMounted(() => {
  themeStore.setShowCatalog(true)
  themeStore.actionCenterMsgs = [
    `按下 / 以开启文章的目录面板`,
    `单击目录区域的文章标题可以回到顶部`,
    `点击文章中的Tag以获取TagId`
  ]
  // 根据屏幕自动切换 Drawer 方向（可选）
  // position.value = window.innerWidth < 1024 ? 'bottom' : 'left';

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
      <Drawer v-model:visible="visible" header="目录" :position="position" class="!w-full md:!w-80 lg:!w-[30rem]">
        <div>
          <p class="text-xl font-bold mt-6 mb-4 hover:underline cursor-pointer" @click="toPageTop">{{ currentArticle.title }}</p>
          <MdCatalog :editorId="id" scrollElement="html" />
        </div>
      </Drawer>
    </div>

    <!-- 主体区域 -->
    <div class="flex gap-6">

      <!-- 正文 -->
      <main class="flex-1 min-w-0 pb-[120px]">
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