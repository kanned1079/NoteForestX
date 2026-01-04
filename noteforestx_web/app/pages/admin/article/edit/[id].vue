<script setup lang="ts">
definePageMeta({
  disableShortcuts: true,
})

import {ref, watch, onMounted, onBeforeMount, onBeforeUnmount} from "vue";
import {MdEditor} from 'md-editor-v3'
import '~/assets/css/md-style.css'
import PageHeader from "~/components/PageHeader.vue";
import type {Article, ArticleTag, NewArticle, NewTag} from "~/types/article";
import {useDarkMode} from '~/composables/useDarkMode'
import {useRoute, useRouter} from 'vue-router'
import type {Illustration} from "~/types/illustration";
import {useToast} from "primevue/usetoast";
import useThemeStore from "~/store/themeStore";
import useActionStore from "~/store/actionStore";
import {CircleArrowOutUpLeft, Command} from "lucide-vue-next";
import {useScrollFadeIn} from "~/composables/useScrollFadeIn";
import {useI18n} from "vue-i18n"; // 补充 i18n 导入

// 补充缺失组件导入（模板中使用）
import Button from "primevue/button";
import Dialog from "primevue/dialog";
import IconField from "primevue/iconfield";
import InputIcon from "primevue/inputicon";
import InputText from "primevue/inputtext";
import AutoComplete from "primevue/autocomplete";
import Tag from "primevue/tag";

const {t} = useI18n(); // 初始化 t 函数
const themeStore = useThemeStore();
const actionStore = useActionStore()
const toast = useToast()
const { isDarkMode } = useDarkMode()
const router = useRouter()
const route = useRoute()
const articleId = route.params.id;
const dialogRef = ref<InstanceType<typeof Dialog>>(); // 补充 dialog 引用

const editArticle = ref<NewArticle>({
  title: '',
  status: 'draft',
  content: '',
  slug: '', // 补充模板中绑定的 slug 字段
  image_url: '', // 补充模板中绑定的 image_url 字段
  tags: [] // 每个对象是 {id?: string, name?: string} 如果存在的就直接放入id不需要name 如果是新增的放入id和name id就是'new'就行
})

useScrollFadeIn({
  selector: '.animate-card-article-edit-id',
  y: 60,
  duration: 0.6,
  stagger: 0.15,
  useScrollTrigger: false
})

// fetchTags 根据输入的tag名进行模糊搜索
const tagSearchValue = ref<string>('')
const tagSearchRes = ref<ArticleTag[]>([])
const fetchTags = async (tagName?: string) => {
  const keyword = tagName ?? tagSearchValue.value
  tagSearchRes.value = []
  try {
    const res = await $fetch<{ tags: ArticleTag[]; message: string }>('/api/admin/tag', {
      method: 'GET',
      query: { search: keyword }
    })

    if (res.tags && res.tags.length > 0) {
      tagSearchRes.value = res.tags
    } else {
      // 没有搜索到结果就显示新增
      // tagSearchRes.value = [{ id: 'new', name: keyword }]
      toast.add({
        severity: "warn",
        summary: t("article_edit.toast.warnSummary"),
        detail: t("article_edit.tags.toast.noSearchResultWarn"),
        life: 4500,
      })
    }
  } catch (err: any) {
    toast.add({
      severity: 'error',
      summary: t("article_edit.toast.loadFailed"),
      detail: `${err}`,
      life: 4500
    })
  } finally {
    tagSearchRes.value.push({ id: 'new', name: keyword })
  }
}

const selectTagItem = (option: any) => {
  const data = option.value as ArticleTag
  if (!data) return

  // 检查是否已经存在
  const exists = editArticle.value.tags.find(t => t.id === data.id && t.name === data.name)
  if (!exists) {
    if (data.id === 'new') {
      editArticle.value.tags.push({ id: 'new', name: data.name })
    } else {
      editArticle.value.tags.push({ id: data.id, name: data.name })
    }
  }

  toast.add({
    severity: "success",
    summary: t("article_edit.tags.toast.addSuccessSummary"),
    detail: `${t("article_edit.tags.option.tagPrefix")}${data.name}`,
    life: 4500,
    styleClass: 'text-mono'
  })

  // 清空输入框
  tagSearchValue.value = ''
  tagSearchRes.value = []
}

const deleteTag = (tag: { id?: string; name?: string }) => {
  editArticle.value.tags = editArticle.value.tags.filter(
      t => t.id !== tag.id || t.name !== tag.name
  )
}

const fetchArticleById = async (id: string) => {
  console.log(`load article by id ${id}`)
  try {
    const res = await $fetch<{
      id: string,
      message: string,
      article: Article | null
    }>(`/api/admin/article/${id}`, {
      method: "GET"
    })
    if (res.article) editArticle.value = res.article
  } catch (err: any) {
    toast.add({
      severity: "error",
      summary: t("article_edit.toast.loadArticleFailed"),
      detail: `${err}`,
      life: 4500,
    })
  }
}

const routerBackIn = (waitMs: number) => {
  setTimeout(() => router.back(), waitMs)
}

const saveNewArticle = async () => {
  console.log('saveNewArticle be called')
  try {
    const res = await $fetch<{
      id: string,
      message: string,
    }>(`/api/admin/article`, {
      method: "POST",
      body: {
        ...editArticle.value
      }
    })
    toast.add({
      severity: "success",
      summary: t("article_edit.toast.successSummary"),
      detail: t("article_edit.toast.saveSuccess"),
      life: 4500,
    })
  } catch (err: any) {
    toast.add({
      severity: "error",
      summary: t("article_edit.toast.loadFailed"),
      detail: `${err}`,
      life: 4500,
    })
  } finally {
    routerBackIn(1000)
  }
}

const saveArticle = async () => {
  console.log(`save article by id ${articleId}`)
  try {
    const res = await $fetch<{
      id: string,
      message: string,
    }>(`/api/admin/article/${articleId}`, {
      method: "PUT",
      body: {
        ...editArticle.value
      }
    })
    toast.add({
      severity: "success",
      summary: t("article_edit.toast.successSummary"),
      detail: t("article_edit.toast.saveSuccess"),
      life: 4500,
    })

  } catch (err: any) {
    toast.add({
      severity: "error",
      summary: t("article_edit.toast.loadFailed"),
      detail: `${err}`,
      life: 4500,
    })
  } finally {
    routerBackIn(1000)
  }
}

watch(() => actionStore.triggerArticleSave, (newVal: boolean) => {
  if (newVal) {
    actionStore.resetTriggerArticleSave()
    if (articleId === 'new') return saveNewArticle()
    else return  saveArticle()
  }
})

const handleCmdKeyDown = (e: KeyboardEvent) => {
  const isMac = navigator.platform.toUpperCase().includes('MAC')
  if (
      (isMac && e.metaKey && e.key.toLowerCase() === 'k') ||
      (!isMac && e.ctrlKey && e.key.toLowerCase() === 'k')
  ) {
    e.preventDefault()
    themeStore.setShowEditMetaDialog(true)
  } else if ((isMac && e.metaKey && e.key.toLowerCase() === 's') ||
      (!isMac && e.ctrlKey && e.key.toLowerCase() === 's')) {
    e.preventDefault()
    saveArticle()
  }

  if (e.metaKey && e.key.toLowerCase() === 'r') {
    e.preventDefault()
    // clearAndSearch()
  }
}

if (articleId !== 'new' && articleId) fetchArticleById(articleId as string)

onBeforeMount(() => {
  themeStore.setShowEditMetaBtn(true)
})

onMounted(() => {
  console.log(articleId)
  window.addEventListener('keydown', handleCmdKeyDown)
})

onBeforeUnmount(() => {
  themeStore.setShowEditMetaBtn(false)
  window.removeEventListener('keydown', handleCmdKeyDown)
})
</script>

<template>
  <div class="mt-4 pb-10 animate-card-article-edit-id" style="z-index: 1000">

    <!-- 页面头 -->
    <PageHeader
        :title="editArticle.title || t('article_edit.pageHeader.defaultTitle')"
        :subtitle="t('article_edit.pageHeader.subtitle')"
    />

    <!-- 操作区 -->
    <div class="flex flex-row mb-6 space-x-3" v-if="false">
      <Button
          size="small"
          :label="t('article_edit.operation.addNewArticle')"
          icon="pi pi-plus"
          @click=""
      />
      <Button
          size="small"
          severity="secondary"
          :label="t('article_edit.operation.searchCmdK')"
          icon="pi pi-search"
      />
    </div>

    <MdEditor class="z-50" v-model="editArticle.content" :theme="isDarkMode?'dark':undefined" :preview-theme="'github'"  />

  </div>

  <Dialog
      ref="dialogRef"
      :dismissableMask="true"
      v-model:visible="themeStore.showEditMetaDialog"
      maximizable modal
      :header="t('article_edit.metaDialog.header')"
      :style="{ width: '50rem' }"
      :breakpoints="{ '1199px': '75vw', '575px': '90vw' }"
      class="pb-4"
  >

    <div class="space-y-4">

      <div class="flex flex-col gap-2">
        <label for="name">{{ t('article_edit.metaDialog.labels.title') }}</label>
        <IconField>
          <InputIcon class="pi pi-file-edit"></InputIcon>
          <InputText autofocus variant="outlined" size="small" id="name" placeholder="Why we use Nuxt.js"
                     v-model="editArticle.title" class="w-full"/>
        </IconField>
      </div>

      <div class="flex flex-col gap-2">
        <label for="name">{{ t('article_edit.metaDialog.labels.tags') }}</label>
        <IconField class="w-full">
          <InputIcon class="pi pi-hashtag"></InputIcon>
          <AutoComplete size="small" :fluid="true" placeholder="frontend"
                        :suggestions="tagSearchRes" v-model="tagSearchValue"
                        @complete="fetchTags(undefined)"
                        @itemSelect="selectTagItem"
          >
            <template #option="slotProps">
              <div class="flex items-center space-x-2">
                <div class="text-sm opacity-90 font-mono">
                  {{ t('article_edit.tags.option.tagPrefix') }}{{ slotProps.option.name }}
                </div>
                <div class="text-xs opacity-80">
                  {{ slotProps.option.id === 'new' ? t('article_edit.tags.option.newTag') : slotProps.option.id }}
                </div>
              </div>
            </template>
          </AutoComplete>
        </IconField>
      </div>

      <div class="flex flex-row space-x-2">
        <Tag
            v-for="i in editArticle.tags"
            :key="i.id"
            icon="pi pi-hashtag"
            :severity="i.id==='new'?'warn':'primary'"
            size="small"
            class="text-xs font-normal hover:underline font-mono"
            :value="i.name"
            @click="deleteTag(i)"
        ></Tag>

        <Tag
            v-if="editArticle.tags.length === 0"
            class="text-xs font-semibold"
            icon="pi pi-exclamation-triangle"
            severity="danger"
            :value="t('article_edit.tags.noTags')"
        ></Tag>
      </div>
    </div>

    <div class="flex flex-col gap-2 mt-4">
      <label for="name">{{ t('article_edit.metaDialog.labels.slug') }}</label>
      <IconField>
        <InputIcon class="pi pi-thumbtack"></InputIcon>
        <InputText autofocus variant="outlined" size="small" id="name" placeholder="why-we-use-nuxtjs"
                   v-model="editArticle.slug" class="w-full"/>
      </IconField>
    </div>

    <div class="flex flex-col gap-2 mt-4">
      <label for="name">{{ t('article_edit.metaDialog.labels.headerImage') }}</label>
      <IconField>
        <InputIcon class="pi pi-image"></InputIcon>
        <InputText autofocus variant="outlined" size="small" id="name" placeholder="https://example.com/d/img.jpg"
                   v-model="editArticle.image_url" class="w-full"/>
      </IconField>
    </div>

  </Dialog>
</template>

<style scoped>
.p-autocomplete-option {
  padding: 0 !important;
}
</style>

<style scoped>
.p-autocomplete-option {
  padding: 0 !important;
}
</style>