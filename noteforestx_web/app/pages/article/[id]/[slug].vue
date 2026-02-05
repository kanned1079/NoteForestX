<script setup lang="ts">
definePageMeta({
  layout: "preview"
})

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

const currentArticle = ref<Article>();

// const currentArticle = ref<Article>({
//   id: '',
//   top: false,
//   status: 'published',
//   title: '',
//   tags: [],
//   content: ''
// });

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

const readingTime = computed(() => {
  if (!currentArticle.value?.content) return 0;

  // 1. 去除 HTML 标签和 Markdown 特殊字符
  const text = currentArticle.value.content.replace(/[#*`>]/g, '').trim();

  // 2. 统计字数（兼容中英文）
  // 中文匹配：[\u4e00-\u9fa5]，英文单词匹配：\w+
  const wordCount = (text.match(/[\u4e00-\u9fa5]/g) || []).length +
      (text.match(/\w+/g) || []).length;

  // 3. 计算时长（假设每分钟阅读 250 字）
  const wpm = 250;
  const minutes = Math.ceil(wordCount / wpm);

  return minutes > 0 ? minutes : 1; // 至少显示 1 分钟
});

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

setTimeout(() => fetchArticleById(articleId), 500)

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
  <!-- 这里移除可能冲突的 pt-[40px]，统一交给 Layout 控制或作为内边距 -->
  <div class="w-full py-8 flex-1">

    <!-- Drawer 保持原样 -->
    <div class="card flex justify-center" v-if="currentArticle">
      <Drawer v-model:visible="visible" :header="t('article.catalog')" :position="position" class="!w-full md:!w-80 lg:!w-[30rem]">
        <div>
          <p class="text-xl font-bold mt-6 mb-4 hover:underline cursor-pointer" @click="toPageTop">{{ currentArticle.title }}</p>
          <MdCatalog :scrollElementOffsetTop="80" :editorId="id" scrollElement="html" />
          <p class="mt-4 text-xs font-mono opacity-50">{{ articleId }}</p>
        </div>
      </Drawer>
    </div>

    <!-- 主体区域 -->
    <div class="flex flex-col">
      <main class="w-full">
        <div class="mx-auto max-w-[720px] lg:max-w-[900px] animate-card-article-id">

          <div v-if="currentArticle">
            <ArticleHeader
                :title="currentArticle.title"
                :tags="currentArticle.tags"
                :createdAt="dayjs(currentArticle.created_at).format('YYYY-MM-DD')"
            />

            <!-- 关键点：用一个 div 包裹预览，并确保它没有被溢出隐藏 -->
            <div class="md-content-wrapper mt-8 clear-both overflow-visible">
              <MdPreview
                  :id="id"
                  :modelValue="currentArticle.content"
                  :theme="isDarkMode ? 'dark' : undefined"
                  preview-theme="github"
              />
            </div>

            <Divider />


            <div class="text-sm opacity-70">
              <p>About {{ readingTime }} min read.</p>
            </div>


            <Message severity="warn" size="small" class="mt-10">
              因为一些原因，评论区暂时不可见。 / 因為一些原因，評論區暫時不可見。 / 何らかの理由により、コメント欄は一時的に利用できません。 / 어떤 이유에서인지, 댓글 섹션을 일시적으로 이용할 수 없습니다. / For some reason, the comments section is temporarily unavailable. / Pour une raison quelconque, la section commentaires est temporairement indisponible. / Jostain syystä kommenttiosio on tilapäisesti poissa käytöstä. / Бо баъзе сабабҳо, бахши шарҳҳо муваққатан дастрас нест.
            </Message>

          </div>



          <!-- Loading 占位 -->
          <div v-else class="">
            <div class="text-xl font-semibold opacity-80 animate-pulse mb-2">Loading article content...</div>
            <div class="text-sm font-mono opacity-70">{{ articleId }} (UUID)</div>
          </div>

        </div>
      </main>
    </div>


    <!-- 在文章末尾加一个强制撑开的占位符，防止渲染延迟导致的 Footer 抽风 -->
    <div class="h-10 w-full clear-both"></div>
  </div>
</template>


<!--<template>-->
<!--  <div-->
<!--      class="pt-[40px] md:px-6 lg:px-8"-->
<!--  >-->
<!--    <div class="card flex justify-center" v-if="currentArticle">-->
<!--      <Drawer v-model:visible="visible" :header="t('article.catalog')" :position="position" class="!w-full md:!w-80 lg:!w-[30rem]">-->
<!--        <div>-->
<!--          <p class="text-xl font-bold mt-6 mb-4 hover:underline cursor-pointer" @click="toPageTop">{{ currentArticle.title }}</p>-->
<!--          <MdCatalog :scrollElementOffsetTop="80" :editorId="id" scrollElement="html" />-->
<!--          <p class="mt-4 text-xs font-mono opacity-50">{{ articleId }}</p>-->
<!--        </div>-->
<!--      </Drawer>-->
<!--    </div>-->

<!--    &lt;!&ndash; 主体区域 &ndash;&gt;-->
<!--    <div class="flex gap-6">-->

<!--      &lt;!&ndash; 正文 &ndash;&gt;-->
<!--      <main class="flex-1 min-w-0">-->

<!--        <div class="mx-auto max-w-[720px] lg:max-w-[900px] animate-card-article-id">-->

<!--          <div v-if="currentArticle">-->
<!--            <ArticleHeader-->
<!--                :title="currentArticle.title"-->
<!--                :tags="currentArticle.tags"-->
<!--                :createdAt="dayjs(currentArticle.created_at).format('YYYY-MM-DD')"-->
<!--            />-->

<!--            <MdPreview-->
<!--                :id="id"-->
<!--                :modelValue="currentArticle.content"-->
<!--                :theme="isDarkMode ? 'dark' : undefined"-->
<!--                preview-theme="github"-->
<!--            />-->
<!--          </div>-->
<!--          <div v-else class="space-y-4">-->
<!--            <div class="text-xl font-semibold opacity-80">Loading article ...</div>-->
<!--            <div class="font-mono opacity-50">{{ articleId }}</div>-->
<!--          </div>-->


<!--        </div>-->


<!--      </main>-->
<!--    </div>-->
<!--  </div>-->
<!--</template>-->

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

/* 确保 md-editor 不会因为内部绝对定位导致高度塌陷 */
.md-editor {
  height: auto !important; /* 强制自适应高度 */
  z-index: 1 !important;
  --md-bk-color: transparent !important;
}

/* 解决目录或编辑器可能带来的浮动问题 */
.md-content-wrapper {
  display: block;
  width: 100%;
  position: relative;
  min-height: 200px; /* 给一个最小高度，防止加载时瞬间塌陷 */
}
</style>