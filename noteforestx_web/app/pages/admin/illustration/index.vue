<script setup lang="ts">
import { ref, onMounted } from "vue"
import PageHeader from "../../../components/PageHeader.vue"
import type { Illustration } from "../../../types/illustration"
import { useToast } from "primevue/usetoast"
import { useConfirm } from "primevue/useconfirm"

const config = useRuntimeConfig()
const toast = useToast()
const confirm = useConfirm()

// 数据
const illustrationList = ref<Illustration[]>([])
const total = ref<number>(0)
const page = ref<number>(1) // 从 1 开始
const size = ref<number>(10) // 每页 10 条
const loading = ref<boolean>(false)

const imagePrefix: string = `${config.public.apiBase}/api/v1`

// 拉取数据
const fetchIllustList = async () => {
  loading.value = true
  try {
    const data = await $fetch<{
      page: number
      size: number
      total: number
      list: Illustration[]
    }>("/api/illustration", {
      method: "GET",
      params: {
        page: page.value,
        size: size.value,
        show_limited: true,
      },
    })
    illustrationList.value = data.list
    total.value = data.total
  } catch (err: any) {
    console.error(err)
    toast.add({
      severity: "error",
      summary: "加载失败",
      detail: `${err}`,
      life: 4500,
    })
  } finally {
    loading.value = false
  }
}

// 删除
const deleteIllust = async (id: string) => {
  confirm.require({
    message: "确定要删除这张插画吗？",
    header: "删除确认",
    icon: "pi pi-exclamation-triangle",
    acceptClass: "p-button-danger",
    accept: async () => {
      try {
        await $fetch(`/api/illustration/${id}`, { method: "DELETE" })
        toast.add({
          severity: "success",
          summary: "删除成功",
          detail: "插画已删除",
          life: 3000,
        })
        fetchIllustList()
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

// 初始化加载
onMounted(() => {
  fetchIllustList()
})
</script>

<template>
  <div class="mt-4">
    <PageHeader title="插画管理" subtitle="如标题一样这里是管理插画的地方" />

    <!-- 卡片网格 -->
    <div v-if="!loading" class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 mt-4">
      <Card
          v-for="illust in illustrationList"
          :key="illust.id"
          class="shadow-md rounded-xl overflow-hidden"
      >
        <!-- 图片 -->
        <template #header>
          <img
              v-if="illust.images && illust.images.length > 0"
              :src="`${imagePrefix}/illustration/file/${illust.images[0].file_path}?size=medium`"
              class="w-full h-48 object-cover"
          />
        </template>

        <!-- 内容 -->
        <template #title>
          <div class="flex justify-between items-center">
            <span class="font-semibold">{{ illust.name }}</span>
            <Tag
                :value="illust.limited ? '限制型' : '非限制型'"
                :severity="illust.limited ? 'warn' : 'success'"
            />
          </div>
        </template>

        <template #content>
          <p class="text-sm text-gray-600 mb-2">
            作者：<span class="font-medium">{{ illust.author?.name }}</span>
          </p>
          <div class="flex flex-wrap gap-1">
            <Tag
                v-for="tag in illust.tags"
                :key="tag.id"
                :value="tag.name"
                severity="secondary"
                class="text-xs"
            />
          </div>
        </template>

        <!-- 操作按钮 -->
        <template #footer>
          <div class="flex justify-end gap-2">
            <Button
                icon="pi pi-pencil"
                text
                rounded
                severity="info"
                @click="$router.push(`/admin/illustration/edit/${illust.id}`)"
            />
            <Button
                icon="pi pi-trash"
                text
                rounded
                severity="danger"
                @click="deleteIllust(illust.id)"
            />
          </div>
        </template>
      </Card>
    </div>

    <!-- 加载中 -->
    <div v-else class="text-center py-10 text-gray-500">
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
        @page="fetchIllustList"
    />
  </div>
</template>

<style scoped>
.p-card {
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}
.p-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.12);
}
</style>