<script setup lang="ts">
import {ref, onMounted, computed} from "vue"
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

const {t} = useI18n()
const toast = useToast()
const confirm = useConfirm()
const router = useRouter()

useScrollFadeIn({
  selector: '.animate-card-article-index',
  y: 60,
  duration: 0.6,
  stagger: 0.15,
  useScrollTrigger: false
})

/* ===== 数据 ===== */
// const articleList = ref<Article[]>([])

const inputRef = ref<InstanceType<typeof InputText> | null>(null)

const articleList = ref<Article[]>([])
// const appMenu = ref()
const total = ref<number>(190)
const page = ref<number>(1)
const size = ref<number>(10)
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
  if (!trimmed) {
    searchQuery.value = {}
    return
  }

  if (trimmed.startsWith('/i ')) {
    // /i <title>
    searchQuery.value = {
      search: trimmed.slice(3).trim(),
      tag: undefined,
      tag_id: undefined
    }
  } else if (trimmed.startsWith('/t ')) {
    // /t <tag>
    searchQuery.value = {
      search: undefined,
      tag: trimmed.slice(3).trim(),
      tag_id: undefined
    }
  } else {
    // 默认按标题搜索
    searchQuery.value = {
      search: trimmed,
      tag: undefined,
      tag_id: undefined
    }
  }
}

const currentIcon = computed(() => {
  const val = searchInput.value.trim()
  if (val.startsWith('/i')) return 'pi pi-file-word'
  if (val.startsWith('/t')) return 'pi pi-hashtag'
  return 'pi pi-file-word'
})

const onEnterPress = () => {
  parseSearchInput(searchInput.value)
  fetchArticleList()
}


/* ===== 拉取文章列表 ===== */
const fetchArticleList = async () => {
  loading.value = true
  try {
    const data = await $fetch<{
      page: number
      size: number
      total: number
      list: Article[]
    }>("/api/admin/article", {
      method: "GET",
      query: {
        page: page.value,
        size: size.value,
        ...searchQuery.value,
      },
    })
    articleList.value = data.list
    total.value = data.total
  } catch (err: any) {
    toast.add({
      severity: "error",
      summary: "加载失败",
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

/* ===== 删除 ===== */
const deleteArticle = (id: string) => {
  confirm.require({
    message: "确定要删除这篇文章吗？",
    header: "删除确认",
    icon: "pi pi-exclamation-triangle",
    acceptClass: "p-button-danger",
    accept: async () => {
      try {
        await $fetch(`/api/article/${id}`, { method: "DELETE" })
        toast.add({
          severity: "success",
          summary: "删除成功",
          detail: "文章已删除",
          life: 3000,
        })
        fetchArticleList()
      } catch (err: any) {
        toast.add({
          severity: "error",
          summary: "删除失败",
          detail: `${err}`,
          life: 4000,
        })
      }
    },
  })
}

const updateArticleStatus = async (
    id: string,
    field: 'top' | 'status',
    value: boolean | string
) => {
  try {
    await $fetch(`/api/admin/article/${id}`, {
      method: 'PATCH',
      body: {
        [field]: value, // ⭐ 关键
      },
    })

    toast.add({
      severity: 'success',
      summary: '成功',
      detail: '操作成功',
      life: 3000,
    })

    await fetchArticleList()
  } catch (err: any) {
    toast.add({
      severity: 'error',
      summary: '失败',
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

fetchArticleList()

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
        title="文章管理"
        subtitle="如标题一样，这里是管理文章的地方，新建的文章默认为不可见，可以对每一篇文章设置不同的状态。"
    />

    <!-- 操作区 -->
    <div class="flex flex-row mb-6 space-x-3">
      <Button
          size="small"
          label="添加新文章"
          icon="pi pi-plus"
          class="text-sm"
          @click="createNewArticleClick"
      />
      <Button
          size="small"
          severity="secondary"
          label="重制搜索"
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
            placeholder="/cmd {title|tag}"
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
        <MyCard style="padding: 10px 14px" class="">
          <DataTable
              :value="articleList"
              tableStyle="min-width: 80rem"
              class="datatable-transparent"
              size="small"
          >

            <template #empty>
              <div class="text-center mt-10 mb-10 space-y-3">
                <div class="text-3xl font-bold">没有结果</div>
                <div>没有找到对应的所有结果，尝试换一个关键词</div>
              </div>
            </template>


            <!-- 标题 -->
            <Column field="title" header="标题">
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
            <Column header="Tags">
              <template #body="{ data }">
                <div class="flex gap-2 flex-wrap">
                <span
                    v-if="data.tags.length > 0"
                    v-for="tag in data.tags"
                    :key="tag.id"
                    class="text-sm opacity-70 hover:underline cursor-pointer font-mono"
                    @click="() => {searchQuery.tag_id=tag.id; fetchArticleList()}"
                >
                  {{ `#${tag.name}` }}
                </span>
                  <span v-else class="text-sm opacity-70 hover:underline cursor-pointer">---</span>
                </div>
              </template>
            </Column>

            <!-- 创建时间 -->
            <Column header="创建时间">
              <template #body="{ data }">
              <span class="text-sm opacity-70 font-mono">
                {{ dayjs(data.created_at).format('YYYY-MM-DD') }}
              </span>
              </template>
            </Column>

            <!-- 状态 -->
            <Column header="状态">
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
            <Column header="操作" style="width: 120px">
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
                      @click="deleteArticle(data.id)"
                  />
                </div>
              </template>
            </Column>

          </DataTable>

          <p class="ml-[6px] mt-4 mb-2 opacity-80 text-sm font-bold">Total: {{ articleList.length }}</p>

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
        <p class="mt-2">加载中...</p>
      </div>
    </transition>

  </div>
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

