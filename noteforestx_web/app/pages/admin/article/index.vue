<script setup lang="ts">
import {ref, onMounted, computed, onBeforeUnmount} from "vue"
import {useI18n} from "vue-i18n";
import PageHeader from "~/components/PageHeader.vue"
import {type Article, type ArticleStatus} from "~/types/article"
import {type SearchQuery} from "~/types/article_search";
import {NEXT_ARTICLE_STATUS_MAP} from "~/types/article"
import { useToast } from "primevue/usetoast"
import { useConfirm } from "primevue/useconfirm"
import { useRouter } from "vue-router"
import dayjs from "dayjs"
import MyCard from "~/components/RedesignedComponents/MyCard.vue";
import {useScrollFadeIn} from "~/composables/useScrollFadeIn";
import IconField from "primevue/iconfield";
import InputIcon from "primevue/inputicon";
import type InputText from "primevue/inputtext";
import MyConfirmCard from "~/components/RedesignedComponents/MyConfirmCard.vue";
import Button from "primevue/button"; // 补充缺失的组件导入
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import Tag from "primevue/tag";
import MyPaginationBar from "~/components/MyPaginationBar.vue"; // 补充缺失的组件导入
import {useHttp} from '~/composables/useCommonFetch'
import {usePagination} from "~/composable/usePagination";

const {t} = useI18n()
const toast = useToast()
const confirm = useConfirm()
const router = useRouter()
const confirmRef = ref<InstanceType<typeof MyConfirmCard>>()

useScrollFadeIn({
  selector: '.animate-card-article-index',
  y: 60,
  duration: 0.6,
  stagger: 0.15,
  useScrollTrigger: false
})

/* ===== 数据 ===== */
// const articleList = ref<Article[]>([])

// const onEnterPress = () => {
//   // 1. 解析输入
//   parseSearchInput(searchInput.value)
//   // 2. 搜索时必须回到第一页
//   page.value = 1
//   // 3. 执行请求
//   fetchArticleList()
// }

const inputRef = ref<InstanceType<typeof InputText> | null>(null)

const articleList = ref<Article[]>([])
// const appMenu = ref()
const total = ref<number>(0)
// const page = ref<number>(1)
// const size = ref<number>(10)
const { page, size } = usePagination(15)
const loading = ref<boolean>(false)

const getNextArticleStatus = (
    status: ArticleStatus
): ArticleStatus => {
  return NEXT_ARTICLE_STATUS_MAP[status]
}

const searchQuery = ref<SearchQuery>({})
const searchInput = ref<string>('')

const parseSearchInput = (input: string) => {
  const trimmed = input.trim()

  // 1. 初始化一个干净的查询对象
  const newQuery: SearchQuery = {
    search: undefined,
    tag: undefined,
    tag_id: undefined
  }

  if (!trimmed) {
    searchQuery.value = newQuery
    return
  }

  // 2. 执行正则匹配
  // (t|tag) 匹配 t 或 tag；\s+ 匹配至少一个空格；(.+) 匹配后面的内容
  const tagMatch = trimmed.match(/^\/(t|tag)\s+(.+)/i)
  const titleMatch = trimmed.match(/^\/(i|title)\s+(.+)/i)

  // 3. 安全判断与赋值
  if (tagMatch && tagMatch[2]) {
    // 只有当 tagMatch 不为 null 且存在第二个捕获组时进入
    newQuery.tag = tagMatch[2].trim()
  } else if (titleMatch && titleMatch[2]) {
    // 同理处理标题搜索
    newQuery.search = titleMatch[2].trim()
  } else {
    // 4. 如果没有匹配到任何命令（如直接输入 "vue" 或输入了 "/t" 但没写内容）
    // 自动过滤掉可能存在的开头斜杠，防止带斜杠发送给后端
    newQuery.search = trimmed.replace(/^\//, '')
  }

  searchQuery.value = newQuery
}

// const parseSearchInput = (input: string) => {
//   const trimmed = input.trim()
//   if (!trimmed) {
//     searchQuery.value = {}
//     return
//   }
//
//   if (trimmed.startsWith('/i ')) {
//     // /i <title>
//     searchQuery.value = {
//       search: trimmed.slice(3).trim(),
//       tag: undefined,
//       tag_id: undefined
//     }
//   } else if (trimmed.startsWith('/t ')) {
//     // /t <tag>
//     searchQuery.value = {
//       search: undefined,
//       tag: trimmed.slice(3).trim(),
//       tag_id: undefined
//     }
//   } else {
//     // 默认按标题搜索
//     searchQuery.value = {
//       search: trimmed,
//       tag: undefined,
//       tag_id: undefined
//     }
//   }
// }

const currentIcon = computed(() => {
  const val = searchInput.value.trim()
  if (val.startsWith('/i') || val.startsWith('/title')) return 'pi pi-file-word'
  if (val.startsWith('/t') || val.startsWith('/tag')) return 'pi pi-hashtag'
  return 'pi pi-file-word'
})

const onEnterPress = () => {
  // 1. 解析输入
  parseSearchInput(searchInput.value)
  // 2. 搜索时必须回到第一页
  page.value = 1
  // 3. 执行请求
  fetchArticleList()
}

/* ===== 拉取文章列表 ===== */
const fetchArticleList = async () => {
  loading.value = true
  try {

    const data = await useHttp().get<{
        page: number
        size: number
        total: number
        list: Article[]
    }>(`/v1/admin/article`, {
      query: {
        page: page.value,
        size: size.value,
        ...searchQuery.value,
      }
    })

    articleList.value = data.list
    total.value = data.total
  } catch (err: any) {
    console.error(err)
    toast.add({
      severity: "error",
      summary: t("article.toast.loadFailed"),
      detail: `${err}`,
      life: 4000,
    })
  } finally {
    loading.value = false
  }
}

/* ===== 新建 ===== */
const createNewArticleClick = () => {
  navigateTo({ path: `/admin/article/edit/new` })
}

/* ===== 编辑 ===== */
const editArticle = (id: string) => {
  navigateTo({ path: `/admin/article/edit/${id}` })
}

const delArticle = ref<{
  id: string
  title: string
}>({
  id: "",
  title: ""
})

const deleteArticleClick = (article: Article) => {
  delArticle.value.id = article.id
  delArticle.value.title = article.title
  confirmRef.value?.showConfirm()
}

/* ===== 删除 ===== */
const deleteArticle = async () => {
  if (delArticle.value.id) {
    try {
      await useHttp().delete(`/v1/admin/article/${delArticle.value.id}`)

      toast.add({
        severity: "success",
        summary: t("article.toast.operateSuccess"),
        detail: t("article.toast.deleteSuccess"),
        life: 3000,
      })
      await fetchArticleList()
    } catch (err: any) {
      toast.add({
        severity: "error",
        summary: t("article.toast.deleteFailed"),
        detail: err?.data?.error || err?.message || String(err),
        life: 4000,
      })
    }
  }
}

const updateArticleStatus = async (
    id: string,
    field: 'top' | 'status',
    value: boolean | string
) => {
  try {
    await useHttp().patch(`/v1/admin/article/${id}`, {
      [field]: value,
    }, {
      includeToken: true
    })

    toast.add({
      severity: 'success',
      summary: t("article.toast.operateSuccess"),
      detail: t("article.toast.operateSuccess"),
      life: 3000,
    })

    await fetchArticleList()
  } catch (err: any) {
    toast.add({
      severity: 'error',
      summary: t("article.toast.operateFailed"),
      detail: err?.data?.error || err?.message || String(err),
      life: 4000,
    })
  }
}

/* ===== 状态 Tag 样式 ===== */
const getStatusSeverity = (status: ArticleStatus): {color: string, icon: string} => {
  switch (status) {
    case "published":
      return {color: 'success', icon: 'pi pi-check-square'}
    case "draft":
      return {color: 'warn', icon: 'pi pi-pencil'}
    case "hidden":
      return {color: 'secondary', icon: 'pi pi-filter-slash'}
    default:
      return {color: 'secondary', icon: 'pi pi-filter-slash'}
  }
}

const handleKeyDown = (e: KeyboardEvent) => {
  const target = e.target as HTMLElement
  // 如果焦点已经在 input 内，忽略
  if (target.tagName === 'INPUT' || target.tagName === 'TEXTAREA') {
    return
  }

  if (e.key === '/') {
    e.preventDefault() // 阻止默认行为
    inputRef.value?.$el.focus()
  }
}

onMounted(() => {
  fetchArticleList()
  window.addEventListener('keydown', handleKeyDown)
})

onBeforeUnmount(() => {
  window.removeEventListener('keydown', handleKeyDown)
})
</script>

<template>
  <div class="mt-4 animate-card-article-index">

    <!-- 页面头 -->
    <PageHeader
        :title="t('article.pageHeader.title')"
        :subtitle="t('article.pageHeader.subtitle')"
    />

    <!-- 操作区 -->
    <div class="flex flex-row mb-6 space-x-3">
      <Button
          size="small"
          :label="t('article.operation.addNewArticle')"
          icon="pi pi-plus"
          class="text-sm"
          @click="createNewArticleClick"
      />
      <Button
          size="small"
          severity="secondary"
          :label="t('article.operation.resetSearch')"
          icon="pi pi-refresh"
          class="text-sm"
          variant="outlined"
          @click="() => { searchQuery={}; fetchArticleList() }"
      />
      <IconField class="w-auto">
        <InputIcon :class="currentIcon" />
        <InputText
            ref="inputRef"
            class="w-full font-mono justify-center"
            size="small"
            v-model="searchInput"
            :placeholder="`/cmd {title|tag}`"
            autofocus
            @keyup.enter="onEnterPress"
        />
      </IconField>
    </div>

    <!-- 表格 Card -->
    <!--    <Card class="">-->
    <!--      <template #content>-->

    <transition name="slide-fade">
      <div v-if="!loading">
        <MyCard style="padding: 14px 14px" class="">
          <DataTable
              :value="articleList"
              tableStyle="min-width: 80rem"
              class="datatable-transparent"
              size="small"
          >

            <template #empty>
              <div class="text-center mt-10 mb-10 space-y-3">
                <div class="text-3xl font-bold">{{ t('article.table.empty.title') }}</div>
                <div>{{ t('article.table.empty.desc') }}</div>
              </div>
            </template>

            <!-- 标题 -->
            <Column field="title" :header="t('article.table.headers.title')">
              <template #body="{ data }">
            <span
                class="cursor-pointer hover:underline block truncate max-w-[60ch]"
                @click="editArticle(data.id)"
            >
  {{ data.title }}
</span>
              </template>
            </Column>

            <!-- Tags -->
            <Column :header="t('article.table.headers.tags')">
              <template #body="{ data }">
                <div class="flex gap-2 flex-wrap">
                <span
                    v-if="data.tags.length > 0"
                    v-for="tag in data.tags"
                    :key="tag.id"
                    class="text-sm opacity-70 hover:underline cursor-pointer"
                    @click="() => { searchQuery.tag_id=tag.id; fetchArticleList()}"
                >
                  {{ t('article.table.tags.tagPrefix') }}{{ tag.name }}
                </span>
                  <span v-else class="text-sm opacity-70 hover:underline cursor-pointer">
                    {{ t('article.table.tags.noTags') }}
                  </span>
                </div>
              </template>
            </Column>

            <!-- 创建时间 -->
            <Column :header="t('article.table.headers.createTime')">
              <template #body="{ data }">
              <span class="text-sm opacity-70 font-mono">
                {{ dayjs(data.created_at).format('YYYY-MM-DD') }}
              </span>
              </template>
            </Column>

            <!-- 状态 -->
            <Column :header="t('article.table.headers.status')">
              <template #body="{ data }">
                <Tag
                    size="small"
                    :value="data.status"
                    :severity="getStatusSeverity(data.status).color"
                    :icon="getStatusSeverity(data.status).icon"
                    class="hover:underline text-sm font-mono cursor-pointer"
                    style="padding: 2px 6px"
                    @click="updateArticleStatus(data.id, 'status', getNextArticleStatus(data.status))"
                />
              </template>
            </Column>

            <!-- 操作 -->
            <Column :header="t('article.table.headers.operation')" style="width: 120px">
              <template #body="{ data }">
                <div class="flex gap-3">
                  <i
                      class="pi pi-pencil cursor-pointer opacity-80 hover:opacity-100"
                      @click="editArticle(data.id)"
                  />
                  <i
                      :class="data.top?['text-blue-600']:[]"
                      class="pi pi-paperclip cursor-pointer opacity-80 hover:opacity-100"
                      @click="updateArticleStatus(data.id, 'top', !data.top)"
                  />
                  <i
                      class="pi pi-trash cursor-pointer text-red-500 opacity-80 hover:opacity-100"
                      @click="deleteArticleClick(data)"
                  />
                </div>
              </template>
            </Column>

          </DataTable>

          <p class="ml-[6px] mt-4 mb-2 opacity-80 text-sm font-bold">
            {{ t('article.table.total', { count: articleList.length }) }}
          </p>

        </MyCard>

        <MyPaginationBar
            class="mt-6"
            v-model:page="page"
            v-model:size="size"
            :total="total"
            :fetchData="fetchArticleList"
        />
      </div>
    </transition>

    <transition name="slide-fade">
      <div v-if="loading" class="text-center py-10 text-gray-500">
        <i class="pi pi-spin pi-spinner text-2xl"></i>
        <p class="mt-2">{{ t('article.loading.text') }}</p>
      </div>
    </transition>

  </div>

  <MyConfirmCard
      ref="confirmRef"
      class="w-full max-w-[90vw] sm:max-w-[420px] lg:max-w-[500px]"
      :header="t('article.confirm.deleteHeader')"
      :title="delArticle.title"
      :subtitle="delArticle.id"
      :cancelled="() => {delArticle = {id: '', title: ''}; console.log('cleared')}"
      :confirmed="deleteArticle"
      confirm-btn-severity="danger"
  />

</template>

<style>
/* 表格主体透明 */
.p-datatable .p-datatable-tbody > tr,
.p-datatable .p-datatable-tbody > tr > td,
.p-datatable .p-datatable-tbody > tr > td > div {
  background-color: transparent !important;
  padding: 10px 6px;
}

/* 如果有 hover 效果也要透明 */
.p-datatable .p-datatable-tbody > tr.p-highlight,
.p-datatable .p-datatable-tbody > tr:hover {
  background-color: rgba(255,255,255,0.1) !important; /* 或 transparent */
}

/* 表头也可以透明 */
.p-datatable .p-datatable-thead > tr > th {
  background-color: transparent !important;
  padding: 10px 6px;
}
</style>

<style>
/* 表格主体透明 */
.p-datatable .p-datatable-tbody > tr,
.p-datatable .p-datatable-tbody > tr > td,
.p-datatable .p-datatable-tbody > tr > td > div {
  background-color: transparent !important;
  padding: 10px 6px;
}

/* 如果有 hover 效果也要透明 */
.p-datatable .p-datatable-tbody > tr.p-highlight,
.p-datatable .p-datatable-tbody > tr:hover {
  background-color: rgba(255,255,255,0.1) !important; /* 或 transparent */
}

/* 表头也可以透明 */
.p-datatable .p-datatable-thead > tr > th {
  background-color: transparent !important;
  padding: 10px 6px;
}


</style>

