<script setup lang="ts">
import { ref, onMounted } from "vue"
import PageHeader from "~/components/PageHeader.vue"
import type { Article, ArticleStatus } from "~/types/article"
import { useToast } from "primevue/usetoast"
import { useConfirm } from "primevue/useconfirm"
import { useRouter } from "vue-router"
import dayjs from "dayjs"
import MyCard from "~/components/RedesignedComponents/MyCard.vue";

const toast = useToast()
const confirm = useConfirm()
const router = useRouter()

/* ===== 数据 ===== */
// const articleList = ref<Article[]>([])

const articleList = ref<Article[]>([
  {
    id: "a1b2c3d4-001",
    title: "使用 PrimeVue + TailwindCSS 构建现代化后台管理界面",
    top: true,
    status: "published",
    content: "# PrimeVue Admin",
    tags: [
      { id: "t1", name: "PrimeVue" },
      { id: "t2", name: "TailwindCSS" },
      { id: "t3", name: "Vue3" },
    ],
    created_at: "2025-01-05T10:23:00Z",
    updated_at: "2025-01-06T09:00:00Z",
  },
  {
    id: "a1b2c3d4-002",
    title: "关于 Nuxt3 中 dark mode 设计的一些实践与踩坑记录（标题很长用于测试布局）",
    top: false,
    status: "draft",
    content: "## Dark Mode",
    tags: [
      { id: "t4", name: "Nuxt3" },
      { id: "t5", name: "DarkMode" },
    ],
    created_at: "2025-01-03T14:12:00Z",
  },
  {
    id: "a1b2c3d4-003",
    title: "文章系统中的 Tags 设计：code、slug 还是 id？",
    top: false,
    status: "published",
    content: "Tags design",
    tags: [
      { id: "t6", name: "Architecture" },
      { id: "t7", name: "Backend" },
      { id: "t8", name: "Design" },
    ],
    created_at: "2024-12-28T08:30:00Z",
  },
  {
    id: "a1b2c3d4-004",
    title: "暂时隐藏的一篇文章（后台可见）",
    top: false,
    status: "hidden",
    content: "Hidden article",
    tags: [],
    created_at: "2024-12-20T18:45:00Z",
  },
])
const total = ref<number>(0)
const page = ref<number>(1)
const size = ref<number>(10)
const loading = ref<boolean>(false)

/* ===== 拉取文章列表 ===== */
const fetchArticleList = async () => {
  loading.value = true
  try {
    const data = await $fetch<{
      page: number
      size: number
      total: number
      list: Article[]
    }>("/api/article", {
      method: "GET",
      params: {
        page: page.value,
        size: size.value,
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

/* ===== 状态 Tag 样式 ===== */
const getStatusSeverity = (status: ArticleStatus) => {
  switch (status) {
    case "published":
      return "success"
    case "draft":
      return "warn"
    case "hidden":
      return "secondary"
    default:
      return "secondary"
  }
}

// onMounted(fetchArticleList)
</script>

<template>
  <div class="mt-4">

    <!-- 页面头 -->
    <PageHeader
        title="文章管理"
        subtitle="如标题一样，这里是管理文章的地方"
    />

    <!-- 操作区 -->
    <div class="flex flex-row mb-6 space-x-3">
      <Button
          size="small"
          label="添加新文章"
          icon="pi pi-plus"
          @click="createNewArticleClick"
      />
      <Button
          size="small"
          severity="secondary"
          label="搜索 CMD+K"
          icon="pi pi-search"
      />
    </div>

    <!-- 表格 Card -->
<!--    <Card class="">-->
<!--      <template #content>-->


    <MyCard style="padding: 10px 14px">
      <DataTable
          :value="articleList"
          tableStyle="min-width: 80rem"
          class="datatable-transparent"
      >
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
                    v-for="tag in data.tags"
                    :key="tag.id"
                    class="text-sm opacity-70"
                >
                  {{ `#${tag.name}` }}
                </span>
            </div>
          </template>
        </Column>

        <!-- 创建时间 -->
        <Column header="创建时间">
          <template #body="{ data }">
              <span class="text-sm opacity-70">
                {{ dayjs(data.created_at).format('YYYY-MM-DD') }}
              </span>
          </template>
        </Column>

        <!-- 状态 -->
        <Column header="状态">
          <template #body="{ data }">
            <Tag
                :value="data.status"
                :severity="getStatusSeverity(data.status)"
            />
          </template>
        </Column>

        <!-- 操作 -->
        <Column header="操作" style="width: 120px">
          <template #body="{ data }">
            <div class="flex gap-3">
              <i
                  class="pi pi-pencil cursor-pointer opacity-70 hover:opacity-100"
                  @click="editArticle(data.id)"
              />
              <i
                  class="pi pi-trash cursor-pointer text-red-500 opacity-70 hover:opacity-100"
                  @click="deleteArticle(data.id)"
              />
            </div>
          </template>
        </Column>

      </DataTable>

      <p class="ml-[6px] mt-2 mb-2">Total: {{ articleList.length }}</p>

    </MyCard>


<!--      </template>-->
<!--    </Card>-->

    <!-- 加载状态 -->
    <div v-if="loading" class="text-center py-10 text-gray-500">
      <i class="pi pi-spin pi-spinner text-2xl"></i>
      <p class="mt-2">加载中...</p>
    </div>

    <!-- 分页 -->
    <Paginator
        v-model:first="page"
        :rows="size"
        :totalRecords="total"
        :rowsPerPageOptions="[10, 20, 50]"
        class="mt-6"
        @page="fetchArticleList"
    />

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

