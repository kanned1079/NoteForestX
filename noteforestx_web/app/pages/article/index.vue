<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {ref} from "vue"
import useThemeStore from "~/store/themeStore";
import {Icon} from "@vicons/utils";
import {ArrowDownOutline, ArrowUpOutline, ReturnDownBackOutline} from "@vicons/ionicons5";
import SearchIllustrationInput from "~/components/SearchIllustrationInput.vue";
import InputIcon from "primevue/inputicon";
import IconField from "primevue/iconfield";
import { Keyboard } from 'lucide-vue-next';
import {useToast} from "primevue/usetoast";
import dayjs from 'dayjs';
import WidthTest from "~/components/RedesignedComponents/WidthTest.vue";
import type {Article} from "~/types/article";
import {useScrollFadeIn} from "~/composables/useScrollFadeIn";

useScrollFadeIn({
  selector: '.animate-card-article-index',
  // y: 60,
  // duration: 0.6,
  // stagger: 0.15
  direction: 'up',
  x: 200,
  stagger: 0.1,
  duration: 0.4,
  start: 'top 90%',
  useScrollTrigger: false
})

// useScrollFadeIn({
//   selector: '.animate-card-article-index-i',
//   direction: 'right',
//   x: 100,
//   stagger: 0.1,
//   duration: 0.4,
//   start: 'right 90%',
//   useScrollTrigger: false
// })

const toast = useToast()
const themeStore = useThemeStore();
const {t} = useI18n();

// todo
const blocked = ref<boolean>(false)

const searchTitle = ref<string>("")
const isValid = ref<boolean>(true)
const searchHistory = ref<{
  id: number
  active: boolean
  title: string,
  searched_at: string
}[]>([])

const lastSearchTitle = ref<string>("")
const handleSearch = async (id?: number) => {
  let trimSearchTitle = searchTitle.value.trim()
  if (trimSearchTitle == lastSearchTitle.value || trimSearchTitle == "") {
    // TODO
    return isValid.value = false
  }
  if (id) searchHistory.value = searchHistory.value.filter(item => item.id !== id)
  if (searchHistory.value.length >= 5) searchHistory.value.pop()  // 删除数组最后一个元素
  lastSearchTitle.value = trimSearchTitle
  let nowTime = Date.now()
  if (searchHistory.value[0]) searchHistory.value[0].active = false
  searchHistory.value.unshift({
    id: dayjs().valueOf(),
    active: false,
    title: trimSearchTitle,
    searched_at: dayjs(nowTime).format("YYYY/MM/DD HH:mm:ss")
  })

  localStorage.setItem('articleSearchHistory', JSON.stringify(searchHistory.value))
  // ...


}

const removeAllSearchHistory = () => {
  searchHistory.value = []
  localStorage.removeItem('articleSearchHistory')
}

const removeSearchItem = (id: number) => {
  if (searchHistory.value[0]) { // 检查要删除的是否是第一项（最新记录）
    const isRemovingFirstItem = searchHistory.value.length > 0 && searchHistory.value[0].id === id
    if (isRemovingFirstItem) lastSearchTitle.value = ""  // 如果删的是最新项，那么清除 lastSearchTitle
  }
  searchHistory.value = searchHistory.value.filter(item => item.id !== id)
  localStorage.setItem('articleSearchHistory', JSON.stringify(searchHistory.value))
}

const activeIndex = ref<number>(0)
const updateActive = () => {
  searchHistory.value.forEach((item, idx) => {
    item.active = (idx === activeIndex.value)
  })
}

const handleKeyDown = (e: KeyboardEvent) => {
  if (searchHistory.value.length === 0) return

  if (e.key === "ArrowDown") {
    activeIndex.value =
        (activeIndex.value + 1) % searchHistory.value.length
    updateActive()
  }
  else if (e.key === "ArrowUp") {
    activeIndex.value =
        (activeIndex.value - 1 + searchHistory.value.length) % searchHistory.value.length
    updateActive()
  }
}

const handleCmdKeyDown = (e: KeyboardEvent) => {
  const isMac = navigator.platform.toUpperCase().includes('MAC')

  if (
      (isMac && e.metaKey && e.key.toLowerCase() === 'k') ||
      (!isMac && e.ctrlKey && e.key.toLowerCase() === 'k')
  ) {
    e.preventDefault() // 阻止浏览器默认行为
    themeStore.searchDialog.show = true
  }
}

const refreshSearch = () => {
  if (searchHistory.value.length === 0) {
    searchTitle.value = ""
    return
  }
  activeIndex.value = 0
  searchHistory.value.forEach((item, idx) => {
    item.active = (idx === 0)
  })
  searchTitle.value = ""
  localStorage.setItem('articleSearchHistory', JSON.stringify(searchHistory.value))
}

const onSearchDialogOpen = () => {
  refreshSearch()
  if (searchHistory.value.length > 0) window.addEventListener('keydown', handleKeyDown)
}

const onSearchDialogClose = () => {
  refreshSearch()
  window.removeEventListener('keydown', handleKeyDown)
}

// -------------------------------------------------------------------

const page = ref<number>(1)
const size = ref<number>(30)
const total = ref<number>(1)

const articleList = ref<Article[]>([])

// for (let i = 0; i < 3; i++) {
//   let copy = articleList.value
//   articleList.value = [...articleList.value, ...copy]
// }

const loading = ref<boolean>(false)
const errMsg = ref<any>()
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
      query: {
        page: page.value,
        size: size.value,
        // ...searchQuery.value,
      },
    })
    articleList.value = data.list
    total.value = data.total
  } catch (err: any) {
    errMsg.value = err
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

const toDetails = (article: Article) => navigateTo({path: `/article/${article.id}/${article.slug || 'empty-slug'}`})

// -------------------------------------------------------------------

fetchArticleList()

onMounted(() => {
  themeStore.showHeaderSearchBtn = true
  const savedHistory = localStorage.getItem('articleSearchHistory')
  if (savedHistory) {
    searchHistory.value = JSON.parse(savedHistory)
  }
  window.addEventListener("keydown", handleCmdKeyDown)
})

onBeforeUnmount(() => {
  themeStore.showHeaderSearchBtn = false
  window.removeEventListener("keydown", handleCmdKeyDown)
})

</script>

<template>
  <div class="w-full flex flex-row justify-center">

    <div class="max-w-[900px] container">

      <PageHeader title="文章" class="mb-8 animate-card-article-index">
        <template #subtitle>
          <p>
            文章列表按照更新順序由新到舊排序，<span class="font-mono px-1 rounded">Meta+K</span> 可調出搜索框搜索對應的文章。
          </p>
        </template>
      </PageHeader>

      <transition name="slide-fade">
        <div class="mb-10" v-if="!loading && articleList.length > 0">
          <div class="space-y-4">
            <div
                v-for="i in articleList"
                :key="i.id"
                class="flex flex-row justify-between items-center"
            >
              <div class="space-x-4 flex items-center">
            <span class="text-sm opacity-60">
             {{ dayjs(i.created_at).format("YYYY-MM-DD") }}
            </span>
                <span class="relative flex items-center group cursor-pointer select-none">
               <span class="absolute left-0 w-2 h-2 bg-red-600 dark:bg-blue-400 rounded-full opacity-0 group-hover:opacity-100 transition-opacity duration-200"
               ></span>
                <span
                    @click="toDetails(i)"
                    class="ml-2 relative after:block after:absolute after:h-[1px] after:bg-red-600 after:dark:dark:bg-blue-400 after:bottom-0 after:left-0 after:w-0 group-hover:after:w-full after:transition-all after:duration-200 after:ease-in-out transform transition-transform duration-200 ease-in-out group-hover:translate-x-2.5">{{ i.title }}</span>
            </span>
              </div>
              <div class="space-x-2 opacity-60">
          <span
              class="text-sm hover:underline"
              v-for="tag in i.tags"
              :key="tag.id"
          >{{ `#${tag.name}` }}</span>
              </div>
            </div>
          </div>

          <MyPaginationBar
              class="mt-10 animate-card-article-index"
              v-model:page="page"
              v-model:size="size"
              :total="total"
              :fetchData="fetchArticleList"
          />
        </div>
        <div v-else>
          <Message severity="warn">文章列表為空或查詢不到對應的文章</Message>
        </div>
      </transition>

      <transition name="slide-fade">
        <div v-if="errMsg" class="mt-4">
          <Message severity="error">查询时遇到错误 {{ errMsg }}</Message>
        </div>
      </transition>




    </div>
<!--    主體部分結束-->
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
      @after-hide="onSearchDialogClose"
      @show="onSearchDialogOpen"
  >

    <div class="flex flex-col p-4">
      <!--      <SearchIllustrationInput/>-->
      <IconField class="w-auto">
        <InputIcon class="pi pi-search"/>
        <InputText
            class="w-full font-mono justify-center"
            size="medium"
            v-model="searchTitle"
            placeholder="Ciallo"
            @input="() => {isValid = !!searchTitle.trim()}"
            autofocus
            :invalid="!isValid"
            @keyup.enter="handleSearch()"
        />
      </IconField>
      <div v-if="searchHistory.length > 0" class="mt-2 mb-2 flex flex-row justify-between items-center">
        <span class="font-medium">最近搜索</span>
        <Button class="h-8 text-xs font-light" size="small" link label="清除搜索历史"
                @click="removeAllSearchHistory"></Button>
      </div>
      <div v-else class="text-center w-full pt-8 pb-6">
        <span>No recent searches</span>
      </div>

      <div v-for="i in searchHistory"
           class="transition ease-in-out duration-150 flex items-center justify-between gap-3 p-2 rounded-md hover:bg-gray-200 dark:hover:bg-gray-800 mb-2"
           :class="{'bg-gray-200 dark:bg-gray-800': i.active}"
           @click="searchTitle = i.title; handleSearch(i.id)"
      >
        <!-- 左侧图标 -->
        <div class="flex-shrink-0">
          <i class="pi pi-history text-lg text-gray-500"></i>
        </div>

        <!-- 中间文字，上下排列 -->
        <div class="flex flex-col flex-1">
          <span class="font-medium text-gray-900 dark:text-gray-100">{{ i.title }}</span>
          <span class="text-xs text-gray-500 dark:text-gray-400">{{ i.searched_at }}</span>
        </div>

        <!-- 右侧刪除按钮 -->
        <div class="flex-shrink-0 cursor-pointer">
          <i
              @click="removeSearchItem(i.id)"
              class="pi pi-times text-gray-400 hover:text-gray-600 dark:hover:text-gray-200"></i>
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
        <span>to search</span>
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
        <Button variant="link" label="查看搜索規則" iconPos="right" class="text-xs p-0">
          <template #icon>
            <i class="pi pi-info-circle text-sm"></i>
          </template>
        </Button>
      </div>

    </div>


  </Dialog>

  <Mask :modelValue="blocked" :close="() => {blocked = false}" closeable >
    <template #rt>
      <div>
        <Button size="small" severity="secondary" icon="pi pi-times" aria-label="Cancel" class="opacity-80" />

      </div>
    </template>

    <template #lb>
      <div class="max-w-[320px] m-4">
        <div class="flex flex-row justify-start items-center space-x-2 mb-3 ">
          <Keyboard />
          <p class="text-xl font-semibold">TIP: Shortcuts</p>
        </div>
        <p class="opacity-80 font-light text-sm mb-6">Navigate the site with ease using keyboard shortcuts.</p>

        <div class="flex flex-row justify-between items-center">
            <p class="text-sm font-extralight">Open Quick Access</p>
            <span>
              <Kbd keyLabel="Q" size="sm" />

            </span>
        </div>

        <Divider class="mt-2 mb-1" />
        <div class="flex flex-row justify-between items-center">
          <p class="text-sm">Close Quick Access</p>
          <span class="font-mono">
              <Kbd keyLabel="Q" size="sm" />
              /
              <Kbd keyLabel="Esc" size="sm" />

            </span>
        </div>
        <Divider class="mt-2 mb-1" />

        <div class="flex flex-row justify-between items-center">
          <p class="text-sm">Open Quick Access</p>
          <span>
              <Kbd keyLabel="Cmd" size="sm" />
              +
              <Kbd keyLabel="Q" size="sm" />

            </span>
        </div>
        <Divider class="mt-2 mb-1" />

        <div class="flex flex-row justify-between items-center">
          <p class="text-sm">Open Quick Access</p>
          <span>
              <Kbd keyLabel="Cmd" size="sm" />
              +
              <Kbd keyLabel="Q" size="sm" />

            </span>
        </div>
        <Divider class="mt-2 mb-1" />



      </div>
    </template>
  </Mask>

</template>

<style>
</style>