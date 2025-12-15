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

import dayjs from 'dayjs';

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
  <div class="">
    <PageHeader title="文章">
      <template #subtitle>
        <p>
          文章列表按照更新順序由新到舊排序，<span class="font-mono px-1 rounded">Meta+K</span> 可調出搜索框搜索對應的文章。
        </p>
      </template>
    </PageHeader>

    <div class="flex flex-row justify-start space-x-2">
      <Button size="small" label="Profile" icon="pi pi-user" @click="blocked = true"/>
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

  </div>

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

<style scoped>

</style>