<script setup lang="ts">
// definePageMeta({
//   layout: 'default',
//   layoutTransition: true,
// })
import {ref} from "vue"
import {useI18n} from "vue-i18n";

import PageHeader from "../../components/PageHeader.vue";
import IllustrationItemPreview from "../../components/IllustrationItemPreview.vue";
import useThemeStore from "../../store/themeStore";
import useShortcutsStore from "~/store/shortcutsStore";
import type {Illustration} from "../../types/illustration";
import {Icon} from '@vicons/utils'
import {
  ReturnDownBackOutline,
  ArrowUpOutline,
  ArrowDownOutline
} from "@vicons/ionicons5"
import SearchIllustrationInput from "~/components/SearchIllustrationInput.vue";
const config = useRuntimeConfig()

const {t} = useI18n()
const themeStore = useThemeStore()
const shortcutsStore = useShortcutsStore()
const illustrationList = ref<Illustration[]>([])

const page = ref<number>(1)
const size = ref<number>(15)
const total = ref<number>(0)
const showLimited = ref<boolean>(true)

const fetchIllustrationList = async () => {
  try{
    illustrationList.value = []
    const data = await $fetch<{
        page: number
        size: number
        total: number
        list: Illustration[]
    }>('/api/illustration', {
      method: 'GET',
      params: {
        page: page.value,
        size: size.value,
        show_limited: showLimited.value
      }
    })
    if (data) {
      console.log(data)
      total.value = data.total
      illustrationList.value = []
      data.list.forEach((item: Illustration) => illustrationList.value.push(item))
      // if (illustrationList.value) {
      //   for (let i = 0; i < illustrationList.value[0].tags.length; i++ ){
      //     console.log(`${illustrationList.value[0].tags[i].id} - ${illustrationList.value[0].tags[i].name}`)
      //   }
      // }
    }
  } catch (err: any) {
    console.log("err: ",err)
  }
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

const searchHistory = ref<{ content: string, type: string }[]>([
  {
    content: "キミの考えは、すべてまるっとお見通しだ！",
    type: "illustration"
  },
  {
    content: "着物 母の日",
    type: "tag"
  },
  {
    content: "和かない",
    type: "author"
  }

])

const searchHint = [
  {
    command: '/tag',
    hint: '按照tag进行搜索 (可以空格分开多个)',
  },
  {
    command: '/user',
    hint: '按照画师名进行搜索',
  },
  {
    command: '/name',
    hint: '直接按照插画名进行搜索',
  },
  {
    command: '/id',
    hint: '直接搜索文件或插画Id',
  },
  {
    command: '/limited',
    hint: '显示Limited标记类型插画 (可接在前面的命令后)',
  },
]

const setTagType = (command: string): string => {
  return command === '/limited' ? "info" : "secondary"
}

const keyDownHandler = (e: KeyboardEvent) => {
  const openSearchDialogByShortCut = () => themeStore.searchDialog.show = true
  // macOS: Meta + K
  if (e.metaKey && e.key.toLowerCase() === "k") {
    e.preventDefault()
    openSearchDialogByShortCut()
  }

  // Windows/Linux: Ctrl + K
  if (e.ctrlKey && e.key.toLowerCase() === "k") {
    e.preventDefault()
    openSearchDialogByShortCut()
  }
}

onBeforeUnmount(() => {
  window.removeEventListener("keydown", keyDownHandler)
})

onUnmounted(() => {
  themeStore.showHeaderSearchBtn = false
  shortcutsStore.clear()
})

onBeforeMount(() => {
  shortcutsStore.register([
    {
      label: "打开搜寻窗口",
      keyLabels: ['Cmd', 'K'],
      pressType: 'and'
    }
  ])
  console.log(shortcutsStore.customShortcuts.length)
})

onMounted(() => {
  themeStore.showHeaderSearchBtn = true
  window.addEventListener("keydown", keyDownHandler)

  // themeStore.searchDialog.show = true



})

</script>

<template>
  <div class="">
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
      :closable="true"
      :dismissable-mask="true"
      :content-class="'p-0'"
      :mask-class="'backdrop-blur-sm bg-black/50'"
      class="w-full sm:w-3/4 md:w-2/3 lg:w-1/2 xl:w-1/3 mx-2 sm:mx-auto"
  >

    <div class="flex flex-col p-4">
      <SearchIllustrationInput/>
<!--      <div class="mt-2 font-light text-sm">-->
<!--        使用Enter键以搜索-->
<!--      </div>-->
      <div class="mt-2 mb-2 flex flex-row justify-between items-center">
        <span class="font-medium">最近搜索</span>
        <Button class="h-8 text-xs font-light" size="small" link label="清除搜索历史"></Button>
      </div>

      <div v-for="i in searchHistory"
           class="flex items-center justify-between gap-3 p-2 rounded-md hover:bg-gray-100 dark:hover:bg-gray-800 mb-2">
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

    <Divider v-if="false" class="mt-0 mb-0"/>

    <div class="m-3 text-xs flex flex-row items-center gap-2" v-if="false">
      <div class="flex flex-col justify-start">
        <!--        <span class="font-bold text-sm">支持的查询命令</span>-->
        <span
            v-for="i in searchHint"
            :key="i.command"
            class="mt-2 mb-1 flex items-center gap-2 text-xs"
            v-if="false"
        >
          <Tag
              class="h-5 font-mono font-normal shrink-0 w-20 justify-start text-xs"
              :value="i.command"
              :severity="setTagType(i.command)"
          />
          <span class="flex-1">{{ i.hint }}</span>
        </span>
        <span class="text-xs mt-2 opacity-70">* 默认使用tag进行搜索。键入的命令作为前缀，将启用搜索模式，搜索框前的图标将会变为搜索的类型图标，命令后续键入关键词后即可搜索，仅在使用tag搜索时允许多个关键词。</span>

        <div class="w-full">
          <!--          <Button variant="link" size="small" class="p-0 underline text-sm">-->
          <!--            查看搜索規則-->
          <!--            <template #icon>-->
          <!--              <i class="pi pi-chevron-right"></i>-->
          <!--            </template>-->
          <!--          </Button>-->

        </div>
      </div>
    </div>

    <Divider class="mt-2"/>

    <div class="ml-2 mr-2 mb-3 flex flex-row justify-between">
      <div class="text-xs flex flex-row items-center gap-2 opacity-80">
        <Tag severity="secondary">
          <template #icon>
            <Icon>
              <ReturnDownBackOutline/>
            </Icon>
          </template>
        </Tag>
        <span>to select</span>
        <Tag class="ml-2" severity="secondary">
          <template #icon>
            <Icon>
              <ArrowUpOutline/>
            </Icon>
          </template>
        </Tag>
        <Tag severity="secondary">
          <template #icon>
            <Icon>
              <ArrowDownOutline/>
            </Icon>
          </template>
        </Tag>
        <span>to navigate</span>
      </div>

      <div>
        <Button variant="link" label="查看搜索規則" iconPos="right" class="text-sm p-0">
          <template #icon>
            <i class="pi pi-info-circle text-sm"></i>
          </template>
        </Button>
      </div>

    </div>


  </Dialog>

</template>

<style>


</style>
