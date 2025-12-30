<script setup lang="ts">
import { ref, onMounted } from "vue"
import PageHeader from "../../../components/PageHeader.vue"
import type { Illustration } from "~/types/illustration"
import { useToast } from "primevue/usetoast"
import { useConfirm } from "primevue/useconfirm"
import IllustrationItemPreview from "~/components/IllustrationItemPreview.vue";
import {useRouter} from "vue-router";
const config = useRuntimeConfig()
const toast = useToast()
const confirm = useConfirm()
const router = useRouter()

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
    // for(let i=0; i< 9; i++) {
    //   illustrationList.value.push(illustrationList.value[0])
    //   total.value += 1
    // }
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
        await fetchIllustList()
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

const createNewIllustrationClick = async () => router.push({
  path: "/admin/illustration/new",
})

fetchIllustList()

// 初始化加载
onMounted(() => {

})
</script>

<template>
  <div class="mt-4">
    <PageHeader title="插画管理" subtitle="如标题一样这里是管理插画的地方" />

    <div class="flex flex-row w-auto mb-6 space-x-3">
      <Button size="small"  label="添加新插画" icon="pi pi-plus" @click="createNewIllustrationClick" />
      <Button size="small" variant="outlined" severity="warn" label="显示Limited类型" icon="pi pi-sparkles" />
      <Button size="small"  severity="secondary" label="搜索 CMD+K" icon="pi pi-search" />

    </div>

    <div v-if="!loading" class="mx-auto w-full max-w-[1920px] mt-6">
      <div class="grid gap-5 grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6">
        <IllustrationItemPreview
            admin
            :illustration="i"
            v-for="i in illustrationList" :key="i.id"
          @click=""
         />
      </div>
    </div>
    <div v-else class="text-center py-10 text-gray-500">
      <i class="pi pi-spin pi-spinner text-2xl"></i>
      <p class="mt-2">加载中...</p>
    </div>


    <div class="m-6"></div>

    <!-- 卡片网格 -->
<!--    <div v-if="!loading" class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-2  mt-">-->
<!--      <IllustrationItemPreviewMgr-->
<!--          v-for="illust in illustrationList"-->
<!--          :key="illust.id"-->
<!--        :illustration="illust"-->
<!--      ></IllustrationItemPreviewMgr>-->
<!--    </div>-->
    <!-- 加载中 -->


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