<script setup lang="ts">
// definePageMeta({
//   layout: "preview"
// })
import {useI18n} from "vue-i18n";
import {ref, watch, onMounted, onBeforeUnmount, onBeforeMount} from "vue"
import useThemeStore from "~/store/themeStore";
import useActionStore from "~/store/actionStore";
import { ArrowUp, ArrowDown, CornerDownLeft } from 'lucide-vue-next';
// import {Icon} from "@vicons/utils";
// import {ArrowDownOutline, ArrowUpOutline, ReturnDownBackOutline} from "@vicons/ionicons5";
import InputIcon from "primevue/inputicon";
import IconField from "primevue/iconfield";
import {useToast} from "primevue/usetoast";
import dayjs from 'dayjs';
import type {Article} from "~/types/article";
import {useScrollFadeIn} from "~/composables/useScrollFadeIn";
import type {SearchQuery} from "~/types/article_search";
import ArticleItemDesktop from "~/components/RedesignedComponents/ArticleItemDesktop.vue";
import {navigateTo} from "#app";
import {useHttp} from "~/composables/useCommonFetch";
import {usePagination} from "~/composable/usePagination"; // 补充 navigateTo 导入（Nuxt 环境）

useScrollFadeIn({
  selector: '.animate-card-article-index',
  direction: 'up',
  x: 200,
  stagger: 0.1,
  duration: 0.4,
  start: 'top 90%',
  useScrollTrigger: false
})

const toast = useToast()
const actionStore = useActionStore()
const themeStore = useThemeStore();
const {t} = useI18n();

// todo
// const blocked = ref<boolean>(false)

const activeTagName = ref<string>("")
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
  const trimSearchTitle = searchTitle.value.trim()

  if (!trimSearchTitle || trimSearchTitle === lastSearchTitle.value) {
    isValid.value = false
    return
  }

  isValid.value = true

  if (id) {
    searchHistory.value = searchHistory.value.filter(item => item.id !== id)
  }

  if (searchHistory.value.length >= 5) {
    searchHistory.value.pop()
  }

  lastSearchTitle.value = trimSearchTitle

  if (searchHistory.value[0]) {
    searchHistory.value[0].active = false
  }

  searchHistory.value.unshift({
    id: dayjs().valueOf(),
    active: false,
    title: trimSearchTitle,
    searched_at: dayjs().format("YYYY/MM/DD HH:mm:ss"),
  })

  localStorage.setItem(
      'articleSearchHistory',
      JSON.stringify(searchHistory.value)
  )

  await doSearch(trimSearchTitle)
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
  // ✅ 如果焦点在输入框里，直接跳过
  const target = e.target as HTMLElement
  if (
      target?.tagName === 'INPUT' ||
      target?.tagName === 'TEXTAREA' ||
      target?.isContentEditable
  ) {
    return
  }

  if (searchHistory.value.length === 0) return

  if (e.key === "ArrowDown") {
    e.preventDefault()
    activeIndex.value =
        (activeIndex.value + 1) % searchHistory.value.length
    updateActive()
  }
  else if (e.key === "ArrowUp") {
    e.preventDefault()
    activeIndex.value =
        (activeIndex.value - 1 + searchHistory.value.length) % searchHistory.value.length
    updateActive()
  }
  else if (e.key === "Enter") {
    e.preventDefault()
    const item = searchHistory.value[activeIndex.value]
    if (item) {
      searchTitle.value = item.title
      handleSearch(item.id)
    }
  }
}

const handleCmdKeyDown = (e: KeyboardEvent) => {
  if (!import.meta.client) return
  const isMac = navigator.platform.toUpperCase().includes('MAC')
  if (
      (isMac && e.metaKey && e.key.toLowerCase() === 'k') ||
      (!isMac && e.ctrlKey && e.key.toLowerCase() === 'k')
  ) {
    e.preventDefault() // 阻止浏览器默认行为
    themeStore.searchDialog.show = true
  }

  if (e.metaKey && e.key.toLowerCase() === 'r') {
    e.preventDefault()
    clearAndSearch()
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
  if (import.meta.client) window.addEventListener('keydown', handleKeyDown)
}

const onSearchDialogClose = () => {
  refreshSearch()
  if (import.meta.client) {
    window.removeEventListener('keydown', handleKeyDown)
  }
}

// -------------------------------------------------------------------

const { page, size } = usePagination(15)

// const page = ref<number>(1)
// const size = ref<number>(30)
const total = ref<number>(1)

const articleList = ref<Article[]>([])

const searchQuery = ref<SearchQuery>({})

const loading = ref<boolean>(false)
const errMsg = ref<any>()
const doSearch = async (keyword: string) => {
  searchQuery.value = {
    search: keyword,
    tag: '',
    tag_id: '',
    status: searchQuery.value.status // 保留状态过滤（如果你之后加）
  }
  activeTagName.value = ""

  page.value = 1
  await fetchArticleList()
  themeStore.searchDialog.show = false
}

const searchByTag = async (tag: { id: string; name: string }) => {
  searchQuery.value = {
    tag_id: tag.id, // ✅ 最高优先级
    tag: '',
    search: '',
    status: searchQuery.value.status
  }
  activeTagName.value = tag.name

  page.value = 1
  await fetchArticleList()
}

const clearSearch = async () => {
  searchQuery.value = {}
  activeTagName.value = ""
  page.value = 1
  await fetchArticleList()
}

const clearAndSearch = async () => {
  await clearSearch()
  await handleSearch()
}

const fetchArticleList = async () => {
  loading.value = true
  try {
    const query: Record<string, any> = {
      page: page.value,
      size: size.value,
    }

    // tag_id 优先级最高
    if (searchQuery.value.tag_id) {
      query.tag_id = searchQuery.value.tag_id
    } else {
      if (searchQuery.value.tag) {
        query.tag = searchQuery.value.tag
      }
      if (searchQuery.value.search) {
        query.search = searchQuery.value.search
      }
    }

    if (searchQuery.value.status) {
      query.status = searchQuery.value.status
    }

    const data = await useHttp().get<{
      page: number
      size: number
      total: number
      list: Article[]
    }>("/v1/article", {
      query
    })
    articleList.value = data.list
    total.value = data.total
  } catch (err: any) {
    errMsg.value = err
    toast.add({
      severity: "error",
      summary: t('articleIndex.loadFailed'),
      detail: `${err}`,
      life: 4000,
    })
  } finally {
    loading.value = false
  }
}

const toDetails = (article: Article) => navigateTo({path: `/article/${article.id}/${article.slug || 'empty-slug'}`})


watch(() => actionStore.triggerSearchArticle, (newVal: boolean) => {
  if (newVal) {
    actionStore.resetTriggerSearchArticle()
    clearAndSearch()
  }
})

// ✅ 監聽 page 與 size 變化並持久化到本地儲存，注意安全判斷客戶端避免 SSR 報錯
watch(page, (newPage) => {
  if (import.meta.client) {
    localStorage.setItem('article-list-page', newPage.toString())
  }
})

watch(size, (newSize) => {
  if (import.meta.client) {
    localStorage.setItem('article-list-size', newSize.toString())
  }
})


onMounted(() => {
  // ✅ 僅在客戶端掛載後讀取上次保存的頁碼和每頁數量
  if (import.meta.client) {
    const savedPage = localStorage.getItem('article-list-page')
    const savedSize = localStorage.getItem('article-list-size')
    if (savedPage) page.value = parseInt(savedPage)
    if (savedSize) size.value = parseInt(savedSize)
  }

  fetchArticleList()

  themeStore.showHeaderSearchBtn = true
  themeStore.actionCenterMsgs = [
    t('articleIndex.openSearchBox'),
    t('articleIndex.resetSearch'),
    t('articleIndex.clickTagToSearch'),
  ]
  const savedHistory = localStorage.getItem('articleSearchHistory')
  if (savedHistory) {
    searchHistory.value = JSON.parse(savedHistory)
  }
  if (import.meta.client) {
    window.addEventListener("keydown", handleCmdKeyDown)
  }
})

onBeforeUnmount(() => {
  themeStore.showHeaderSearchBtn = false
  if (import.meta.client) {
    window.removeEventListener("keydown", handleCmdKeyDown)
  }
  themeStore.actionCenterMsgs = []
})
</script>

<template>

<!--  <ClientOnly>-->
    <div class="w-full flex flex-row justify-center">
      <div class="max-w-[900px] container">
        <PageHeader
            :title="t('articleIndex.pageTitle')"
            class="mb-8 animate-card-article-index"
        >
          <template #subtitle>
            <p>{{ t('articleIndex.pageSubtitle') }}</p>
          </template>
        </PageHeader>

        <!-- Active Filter Banner -->
        <div v-if="searchQuery.search || searchQuery.tag_id" class="mb-6 flex flex-row items-center gap-3 py-2 px-4 bg-slate-100 dark:bg-slate-800/40 border border-slate-200/50 dark:border-slate-700/50 rounded-xl animate-card-article-index">
          <span class="text-sm opacity-90 flex items-center gap-2">
            <i class="pi pi-filter text-[#7234e9] dark:text-[#8257f2]"></i>
            <span v-if="searchQuery.search" class="flex items-center gap-1">
              {{ t('articleIndex.filterSearch', { query: searchQuery.search }) }}
            </span>
            <span v-else-if="searchQuery.tag_id" class="flex items-center gap-1">
              {{ t('articleIndex.filterTag', { query: activeTagName }) }}
            </span>
          </span>
          <Button
              icon="pi pi-times"
              severity="secondary"
              size="small"
              variant="text"
              rounded
              class="h-7 w-7 ml-auto hover:bg-slate-200 dark:hover:bg-slate-700/60"
              @click="clearSearch"
          />
        </div>

        <div v-if="loading" class="space-y-2 mb-10 animate-card-article-index">
          <div v-for="n in 5" :key="n" class="flex flex-row justify-between items-center py-1.5 px-0 border-b border-dashed border-slate-200 dark:border-slate-800/60">
            <div class="flex items-center gap-4 w-2/3">
              <Skeleton width="96px" height="1.25rem" class="opacity-60" />
              <Skeleton width="50%" height="1.25rem" />
            </div>
            <Skeleton width="100px" height="1.25rem" class="opacity-60" />
          </div>
        </div>

        <transition name="slide-fade" mode="out-in" v-else>
          <div class="mb-10" v-if="articleList.length > 0">
            <div class="space-y-2">
              <ArticleItemDesktop
                  v-for="item in articleList"
                  :key="item.id"
                  :article="item"
                  @clickTitle="toDetails"
                  @clickTag="searchByTag"
              />
            </div>

            <MyPaginationBar
                class="mt-10 animate-card-article-index"
                v-model:page="page"
                v-model:size="size"
                :total="total"
                :fetchData="fetchArticleList"
            />
          </div>
          <!-- Empty State -->
          <div v-else class="flex flex-col items-center justify-center py-16 px-4 bg-white dark:bg-[#141414] border border-slate-200/50 dark:border-slate-700/50 rounded-2xl shadow-sm text-center mb-10 animate-card-article-index">
            <div class="h-16 w-16 bg-[#7234e9]/10 dark:bg-[#8257f2]/10 text-[#7234e9] dark:text-[#8257f2] rounded-2xl flex items-center justify-center mb-4">
              <i class="pi pi-inbox text-3xl" />
            </div>
            <h3 class="text-xl font-bold mb-2">{{ t('articleIndex.emptyArticleList') }}</h3>
            <p class="text-sm opacity-70 max-w-sm mb-6">{{ t('articleIndex.emptyArticleListDesc') }}</p>
            <Button
                v-if="searchQuery.search || searchQuery.tag_id"
                :label="t('article.operation.resetSearch')"
                severity="primary"
                outlined
                size="small"
                class="px-4 h-9"
                @click="clearSearch"
            />
          </div>
        </transition>

        <transition name="slide-fade">
          <div v-if="errMsg" class="mt-4">
            <Message severity="error">{{ t('articleIndex.searchError') }} {{ errMsg }}</Message>
          </div>
        </transition>
      </div>
      <!--    主體部分結束-->
    </div>
<!--  </ClientOnly>-->



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
            :placeholder="t('articleIndex.searchPlaceholder')"
            @input="() => {isValid = !!searchTitle.trim()}"
            autofocus
            :invalid="!isValid"
            @keyup.enter.stop="handleSearch()"
        />
      </IconField>
      <div v-if="searchHistory.length > 0" class="mt-2 mb-2 flex flex-row justify-between items-center">
        <span class="font-medium">{{ t('articleIndex.recentSearches') }}</span>
        <Button
            class="h-8 text-xs font-light"
            size="small"
            link
            :label="t('articleIndex.clearSearchHistory')"
            @click="removeAllSearchHistory"
        ></Button>
      </div>
      <div v-else class="text-center w-full pt-8 pb-6">
        <span>{{ t('articleIndex.noRecentSearches') }}</span>
      </div>

      <div
          v-for="i in searchHistory"
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
<!--            <Icon>-->
<!--              <ReturnDownBackOutline/>-->
<!--            </Icon>-->
            <CornerDownLeft :size="14" :strokeWidth="2.25" />
          </template>
        </Tag>
        <span>{{ t('articleIndex.searchToSearch') }}</span>
        <Tag class="ml-2" severity="secondary">
          <template #icon>
<!--            <Icon>-->
<!--              <ArrowUpOutline/>-->
<!--            </Icon>-->
            <ArrowUp :size="14" :strokeWidth="2.25" />
          </template>
        </Tag>
        <Tag severity="secondary">
          <template #icon>
<!--            <Icon>-->
<!--              <ArrowDownOutline/>-->
<!--            </Icon>-->
            <ArrowDown :size="14" :strokeWidth="2.25" />
          </template>
        </Tag>
        <span>{{ t('articleIndex.arrowToNavigate') }}</span>
      </div>
    </div>
  </Dialog>
</template>

<style>
</style>