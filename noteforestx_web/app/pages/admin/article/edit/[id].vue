<script setup lang="ts">
definePageMeta({
  disableShortcuts: true,
})

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

const themeStore = useThemeStore();
const actionStore = useActionStore()
const toast = useToast()
const { isDarkMode } = useDarkMode()
const router = useRouter()
const route = useRoute()
const articleId = route.params.id;

const editArticle = ref<NewArticle>({
  title: '',
  status: 'draft',
  content: '',
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
        summary: "警告",
        detail: `没有搜索结果 可以尝试新增`,
        life: 4500,
      })
    }
  } catch (err: any) {
    toast.add({ severity: 'error', summary: '加载失败', detail: `${err}`, life: 4500 })
  } finally {
    tagSearchRes.value.push({ id: 'new', name: keyword })
  }
}

const selectTagItem = (option: any) => {
  const data = option.value as ArticleTag
  if (!data) return

  // 检查是否已经存在
  const exists = editArticle.value.tags.find(t => t.id === data.id && t.name === data.name)
  if (exists) return

  // 新增 tag
  if (data.id === 'new') {
    editArticle.value.tags.push({ id: 'new', name: data.name })
  } else {
    editArticle.value.tags.push({ id: data.id, name: data.name })
  }

  toast.add({
    severity: "success",
    summary: `添加Tag`,
    detail: `#${data.name}`,
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
      summary: "加载失败 请返回重试",
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
      summary: "成功",
      detail: `文章保存成功`,
      life: 4500,
    })
  } catch (err: any) {
    toast.add({
      severity: "error",
      summary: "加载失败",
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
      summary: "成功",
      detail: `文章保存成功`,
      life: 4500,
    })

  } catch (err: any) {
    toast.add({
      severity: "error",
      summary: "加载失败",
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

if (articleId !== 'new' && articleId) fetchArticleById(articleId as string)

onBeforeMount(() => {
  themeStore.setShowEditMetaBtn(true)
})

onMounted(() => {
  console.log(articleId)
})

onBeforeUnmount(() => {
  themeStore.setShowEditMetaBtn(false)
})

</script>

<template>
  <div class="mt-4 pb-10 animate-card-article-edit-id" style="z-index: 1000">

    <!-- 页面头 -->
    <PageHeader
        :title="editArticle.title || '文章標題'"
        subtitle="點擊右上角以設置文章Meta信息"
    />

    <!-- 操作区 -->
    <div class="flex flex-row mb-6 space-x-3" v-if="false">
      <Button
          size="small"
          label="添加新文章"
          icon="pi pi-plus"
          @click=""
      />
      <Button
          size="small"
          severity="secondary"
          label="搜索 CMD+K"
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
      :header="'編輯Meta'"
      :style="{ width: '50rem' }"
      :breakpoints="{ '1199px': '75vw', '575px': '90vw' }"
      class="pb-4"
  >

    <div class="space-y-4">

      <div class="flex flex-col gap-2">
        <label for="name">{{ '标题' }}</label>
        <IconField>
          <InputIcon class="pi pi-file-edit"></InputIcon>
          <InputText autofocus variant="outlined" size="small" id="name" placeholder="Why we use Nuxt.js"
                     v-model="editArticle.title" class="w-full"/>
        </IconField>
      </div>

      <div class="flex flex-col gap-2">
        <label for="name">{{ 'Tags' }}</label>
        <IconField class="w-full">
          <InputIcon class="pi pi-hashtag"></InputIcon>
          <AutoComplete size="small" :fluid="true" placeholder="frontend"
                        :suggestions="tagSearchRes" v-model="tagSearchValue"
                        @complete="fetchTags(undefined)"
                        @itemSelect="selectTagItem"
          >
            <template #option="slotProps">
              <div class="flex items-center space-x-2">
                <div class="text-sm opacity-90 font-mono">{{ `#${slotProps.option.name}` }}</div>
                <div class="text-xs opacity-80">{{ slotProps.option.id === 'new'? '新增': slotProps.option.id }}</div>
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
            icon="pi pi-exclamation-triangle" severity="danger" value="还没有选择标签"></Tag>
        </div>
    </div>

    <div class="flex flex-col gap-2 mt-4">
      <label for="name">{{ 'Slug (可选)' }}</label>
      <IconField>
        <InputIcon class="pi pi-thumbtack"></InputIcon>
        <InputText autofocus variant="outlined" size="small" id="name" placeholder="why-we-use-nuxtjs"
                   v-model="editArticle.slug" class="w-full"/>
      </IconField>
    </div>

    <div class="flex flex-col gap-2 mt-4">
      <label for="name">{{ '头部图片 (可选)' }}</label>
      <IconField>
        <InputIcon class="pi pi-image"></InputIcon>
        <InputText autofocus variant="outlined" size="small" id="name" placeholder="https://example.com/d/img.jpg"
                   v-model="editArticle.image_url" class="w-full"/>
      </IconField>
    </div>



<!--    <IllustrationTagsMgrPanel v-if="showModalCard.type==='tag'" :update-list="fetchIllustrationTags"/>-->
<!--    <IllustrationAuthorsMgrPanel v-if="showModalCard.type==='author'" :update-list="fetchAuthors"/>-->




<!--    <template #footer>-->
<!--      <Divider class="mt-0 mb-2"/>-->
<!--      <div class="w-full flex flex-col justify-center items-start pl-5 pr-5 pb-3">-->
<!--        <div class="gap-2 flex flex-row justify-start items-center">-->

<!--        </div>-->

<!--      </div>-->
<!--    </template>-->

  </Dialog>


</template>

<style scoped>
.p-autocomplete-option {
  padding: 0 !important;
}
</style>