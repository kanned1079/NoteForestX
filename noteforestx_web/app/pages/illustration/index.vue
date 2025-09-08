<script setup lang="ts">
import {ref} from "vue"
import {useI18n} from "vue-i18n";

import PageHeader from "~/components/PageHeader.vue";
import IllustrationItemPreview from "../../components/IllustrationItemPreview.vue";
import useThemeStore from "../../store/themeStore";
import {IllustrationItem} from "../../types/illustration";
import { Icon } from '@vicons/utils'
import {
  ReturnDownBackOutline,
  ArrowUpOutline,
  ArrowDownOutline
} from "@vicons/ionicons5"

const {t} = useI18n()
const themeStore = useThemeStore()
const illustrationList = ref<IllustrationItem[]>([])

const page = ref<number>(1)
const size = ref<number>(15)
const total = ref<number>(0)
const showLimited = ref<boolean>(true)

const fetchIllustrationList = async () => {
  illustrationList.value = []
  const data = await $fetch<{
    page: number
    size: number
    total: number
    list: IllustrationItem[]
  }>(`http://localhost:8081/api/v1/illustration`, {
    method: "GET",
    params: {
      page: page.value,
      size: size.value,
      show_limited: showLimited.value
    }

  })
  console.log(data.total)
  total.value = data.total
  data.list.forEach((item: IllustrationItem) => illustrationList.value.push(item))

}

fetchIllustrationList()

const onChangePage = async (v: { page: number, first: number, pageCount: number }) => {
  console.log(v.page + 1)
  page.value = v.page
  await fetchIllustrationList()
}

// const showSearchDialog = ref<boolean>(false)
// const toggleSearchDialog = () => {
//   showSearchDialog.value = !showSearchDialog.value
// }

const searchIll = ref<{
  content: string
  search_as: string
}>({
  content: '',
  search_as: ''
})

const searchHistory = ref([
  {
    content: "キミの考えは、すべてまるっとお見通しだ！",
    type: "illustration"
  },
  {
    content: "女の子",
    type: "tag"
  },
  {
    content: "あらめ＠お仕事募集中",
    type: "author"
  },
  {
    content: "七原しえ",
    type: "author"
  }
])

onUnmounted(() => {
  themeStore.showHeaderSearchBtn = false
})

onMounted(() => {
  themeStore.showHeaderSearchBtn = true
})

</script>

<template>
  <div class="p-4">
    <PageHeader :title="t('layout.illustrationLib')"
                subtitle="在这里所有收藏的插画项目将会以它们的创建时间进行排序，此处显示的为预览图，如您需要获取原图请点击目标插画进入二级菜单。"/>
    <Message class="mb-5" severity="warn">部分图片带有Limit属性，需要管理员来赋予您权限才可以进行查看。</Message>
    <!--    <Divider />-->
    <div class="flex flex-row items-center space-x-3" v-if="false">
      <Button
          class="h-8 px-3"
          icon="pi pi-search"
          severity="primary"
          outlined
          size="small"
          :label="`在 ${ illustrationList.length } 中搜索`"
      >
      </Button>

      <Button
          class="h-8 px-3"
          label="重置搜索"
          icon="pi pi-refresh"
          severity="primary"
          size="small"
          variant="text"
      />
    </div>
    <div class="mx-auto w-full max-w-[1920px] mt-6">
      <div class="grid gap-5 grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6">
        <IllustrationItemPreview
            :illustration="i"
            v-for="i in illustrationList" :key="i.id"></IllustrationItemPreview>
      </div>
    </div>

    <MyPaginationBar
        v-model:page="page"
        v-model:size="size"
        :total="total"
        :fetchData="fetchIllustrationList"
    />
  </div>

  <Dialog
      v-model:visible="themeStore.searchDialog.show"
      :show-header="false"
      modal
      :content-class="'p-0'"
  :mask-class="'backdrop-blur-sm bg-black/50'"
  class="w-full sm:w-3/4 md:w-2/3 lg:w-1/2 xl:w-1/3 mx-2 sm:mx-auto"
  >
  <div class="flex flex-col p-4">
    <!-- 你的内容 -->
    <IconField>
      <InputIcon class="pi pi-search" />
      <InputText class="w-full" size="large" v-model="searchIll.content" placeholder="Search for something..." />
    </IconField>

    <div class="mt-2 mb-2 flex flex-row justify-between items-center">
      <span class="font-medium">最近搜索</span>
      <Button class="h-8 text-xs font-light" size="small" link label="清除搜索历史"></Button>
    </div>

    <div v-for="i in searchHistory" class="flex items-center justify-between gap-3 p-2 rounded-md hover:bg-gray-100 dark:hover:bg-gray-800 mb-2">
      <!-- 左侧图标 -->
      <div class="flex-shrink-0">
        <i class="pi pi-history text-lg text-gray-500"></i>
      </div>

      <!-- 中间文字，上下排列 -->
      <div class="flex flex-col flex-1">
        <span class="font-medium text-gray-900 dark:text-gray-100">{{ i.content }}</span>
        <span class="text-sm text-gray-500 dark:text-gray-400">{{ i.type }}</span>
      </div>

      <!-- 右侧关闭按钮 -->
      <div class="flex-shrink-0 cursor-pointer">
        <i class="pi pi-times text-gray-400 hover:text-gray-600 dark:hover:text-gray-200"></i>
      </div>
    </div>




  </div>
    <Divider class="mt-2" />
    <div class="ml-2 mr-2 mb-3 text-xs flex flex-row items-center gap-2 opacity-80">
      <Tag severity="secondary">
        <template #icon>
          <Icon><ReturnDownBackOutline /></Icon>
        </template>
      </Tag>
      <span>to select</span>
      <Tag class="ml-2" severity="secondary">
        <template #icon>
          <Icon><ArrowUpOutline /></Icon>
        </template>
      </Tag>
      <Tag severity="secondary">
        <template #icon>
          <Icon><ArrowDownOutline /></Icon>
        </template>
      </Tag>
      <span>to navigate</span>
    </div>
  </Dialog>

</template>

<style>


</style>
