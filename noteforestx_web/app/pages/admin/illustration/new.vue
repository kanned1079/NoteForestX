<script setup lang="ts">
import {ref} from "vue"
import {useI18n} from "vue-i18n"
import PageHeader from "~/components/PageHeader.vue"
import FileUpload, {type FileUploadSelectEvent} from "primevue/fileupload"
import {useToast} from 'primevue/usetoast';
import PageHeaderL2 from "~/components/PageHeaderL2.vue";
import type {Illustration, IllustrationAuthor, IllustrationTag} from "~/types/illustration";
import type {DialogProps} from 'primevue/dialog'
import {Icon} from '@vicons/utils'
import {CaretForwardOutline, GlobeOutline} from "@vicons/ionicons5"
import {Command, CircleArrowOutUpLeft, Keyboard, Option} from "lucide-vue-next"
import useThemeStore from "~/store/themeStore";
import useActionStore from "~/store/actionStore";

const {t} = useI18n()
const themeStore = useThemeStore()
const actionStore = useActionStore()
const toast = useToast();

// 表单 DTO
type createIllustrationRequestDto = {
  name: string
  author_id: string
  tags_id: string[]
  link: string
  description: string
  limited: boolean
  source: string
}

const newIllustration = ref<createIllustrationRequestDto>({
  name: "",
  author_id: "",
  tags_id: [],
  link: "",
  description: "",
  limited: false,
  source: ""
})

// 上传文件 & 预览
const uploadedFiles = ref<File[]>([])
const previewUrls = ref<string[]>([])

const onFileSelect = (event: FileUploadSelectEvent) => {
  const input = event.originalEvent?.target as HTMLInputElement | undefined
  if (!input?.files) return

  const newFiles = Array.from(input.files) // 只取新增的文件
  const newUrls = newFiles.map(f => URL.createObjectURL(f))

  uploadedFiles.value.push(...newFiles)
  previewUrls.value.push(...newUrls)
}

const removeFile = (idx: number) => {
  uploadedFiles.value.splice(idx, 1)
  previewUrls.value.splice(idx, 1)
}

const handleCommitNewIllustration = async () => {
  const formData = new FormData()
  formData.append("name", newIllustration.value.name)
  formData.append("author_id", newIllustration.value.author_id)
  newIllustration.value.tags_id.forEach(tag => formData.append("tags_id[]", tag))
  formData.append("link", newIllustration.value.link)
  formData.append("description", newIllustration.value.description)
  formData.append("limited", String(newIllustration.value.limited))
  formData.append("source", newIllustration.value.source)

  uploadedFiles.value.forEach(file => {
    formData.append("files", file) // 字段名必須是 "files"，後端才能收到多文件
  })

  const res = await fetch("/api/v1/admin/illustration", {
    method: "POST",
    body: formData
  })

  if (!res.ok) {
    console.error("上传失败", await res.text())
  } else {
    console.log("上传成功")
    uploadedFiles.value = []
    previewUrls.value = []
  }
}

type fetchTagsResponseDto = {
  page: number
  size: number
  total: number
  list: IllustrationTag[]
}

type fetchAuthorsResponseDto = {
  page: number
  size: number
  total: number
  list: IllustrationAuthor[]
}

const tagsResponse = ref<fetchTagsResponseDto>()
const authorsResponse = ref<fetchAuthorsResponseDto>()

const showTagsSelection = ref<boolean>(false)
const showAuthorSelection = ref<boolean>(false)
const tagSearch = ref<string>("")
const tagsSelected = ref<IllustrationTag[]>([])

const tagsSearchPage = ref<number>(1)
const tagsSearchSize = ref<number>(30)
const tagsSearchTotal = ref<number>(0)

const fetchIllustrationTags = async () => {
  try {
    const data = await $fetch<fetchTagsResponseDto>(`/api/admin/illustration_tag`, {
      method: "GET",
      query: {
        page: tagsSearchPage.value,
        size: tagsSearchSize.value,
        search: tagSearch.value
      }
    });
    if (data) {
      console.log(data)
      tagsResponse.value = data
      tagsSearchTotal.value = data.total
      setTimeout(() => showTagsSelection.value = true, 500)
      if (tagsResponse.value.list.length === 0) {
        toast.add({
          severity: 'warn',
          summary: t('admin.illustration.tagsSearchNotFound'),
          detail: t('admin.illustration.tagsSearchNotFoundHint', {name: tagSearch.value}),
          life: 3000
        });
      }
    }
  } catch (err: any) {
    console.error("err: ", err);
    toast.add({
      severity: 'error',
      summary: t('universal.errToast'),
      detail: `${t('universal.errToastMessage')} ${err}`,
      life: 3000
    });
  }
};

const authorSearchPage = ref<number>(1)
const authorSearchSize = ref<number>(60)
const authorSearchTotal = ref<number>(0)
const authorSearch = ref<string>("")
const authorSelected = ref<IllustrationAuthor>()

const fetchAuthors = async () => {
  try {
    const data = await $fetch<fetchAuthorsResponseDto>(`/api/admin/illustration_author`, {
      method: "GET",
      query: {
        page: authorSearchPage.value,
        size: authorSearchSize.value,
        search: authorSearch.value
      }
    });
    if (data) {
      console.log(data)
      authorsResponse.value = data
      authorSearchTotal.value = data.total
      setTimeout(() => showAuthorSelection.value = true, 800)
      if (authorsResponse.value.list.length === 0) {
        toast.add({
          severity: 'warn',
          summary: t('admin.illustration.tagsSearchNotFound'),
          detail: t('admin.illustration.tagsSearchNotFoundHint', {name: tagSearch.value}),
          life: 3000
        });
      }
    }
  } catch (err: any) {
    console.error("err: ", err);
    toast.add({
      severity: 'error',
      summary: t('universal.errToast'),
      detail: `${t('universal.errToastMessage')} ${err}`,
      life: 3000
    });
  }
};

const onPressTagSearch = async () => {
  await fetchIllustrationTags()
}

const onPressAuthorSearch = async () => {
  await fetchAuthors()
}

const delTag = (uuid: string) => {
  // 从 newIllustration.tags_id 里移除
  newIllustration.value.tags_id = newIllustration.value.tags_id.filter(id => id !== uuid)

  // 从已选 tagsSelected 里移除
  tagsSelected.value = tagsSelected.value.filter(tag => tag.id !== uuid)
}

const addTag = (tag: IllustrationTag) => {
  // 如果已经在 newIllustration.tags_id 中，就不重复添加
  if (newIllustration.value.tags_id.includes(tag.id as string)) {
    return
  }
  // 添加到已选 tag 列表
  newIllustration.value.tags_id.push(tag.id as string)
  tagsSelected.value.push(tag)
}

const addAuthor = (author: IllustrationAuthor) => {
  if (author) {
    authorSelected.value = author
    newIllustration.value.author_id = author.id
    showAuthorSelection.value = false
  }
}

// fetchIllustrationTags()

const getTotalPagesFunc = (total: number, size: number): number => Math.ceil(total / size)


const onPressSearchPaBtn = async (type: 'tag' | 'author', op: 'increase' | 'decrease') => {
  const overIndexErr = () => {
    toast.add({
      severity: 'info',
      summary: t('universal.info'),
      detail: t('universal.overIndex'),
      life: 3000,
    });
  }

  switch (type) {
    case "tag": {
      let maxPages = getTotalPagesFunc(tagsResponse.value?.total || 1, tagsResponse.value?.page || 1)
      switch (op) {
        case "increase": {
          if (tagsSearchPage.value >= maxPages) return overIndexErr()
          tagsSearchPage.value += 1;
        }
          break
        case "decrease": {
          if (tagsSearchPage.value <= 1) return overIndexErr()
          tagsSearchPage.value -= 1;
        }
          break
      }
      await fetchIllustrationTags()
    }
      break
    case "author": {
      let maxPages = getTotalPagesFunc(authorsResponse.value?.total || 1, authorsResponse.value?.page || 1)
      switch (op) {
        case "increase": {
          if (authorSearchPage.value >= maxPages) return overIndexErr()
          authorSearchPage.value += 1;
        }
          break
        case "decrease": {
          if (authorSearchPage.value <= 1) return overIndexErr()
          authorSearchPage.value -= 1;
        }
          break
      }
      await fetchAuthors()
    }
  }


}

const openLink = (url: string | undefined) => {
  if (url) {
    window.open(url, "_blank", "noopener noreferrer");
  }
};

const showModalCard = ref<{
  show: boolean
  type: 'tag' | 'author'
}>({
  show: false,
  type: 'tag'
})
const dialogRef = ref<HTMLElement | null>(null)
const dialogTitle = ref<string>("admin.illustration.tagMgr")

// POST x-www-form-urlencoded
const commitNewIllustration = async () => {
  try {
    // 1. 构建 FormData
    const formData = new FormData()
    formData.append("name", newIllustration.value.name)
    formData.append("author_id", newIllustration.value.author_id)
    newIllustration.value.tags_id.forEach(tagId => formData.append("tags_id[]", tagId))
    formData.append("link", newIllustration.value.link)
    formData.append("description", newIllustration.value.description)
    formData.append("limited", String(newIllustration.value.limited))
    formData.append("source", newIllustration.value.source)

    // 上传文件
    uploadedFiles.value.forEach(file => {
      formData.append("files", file) // 字段名必须和后端一致
    })

    // 2. 提交请求
    const res = await fetch("/api/admin/illustration", {
      method: "POST",
      body: formData
    })

    if (!res.ok) {
      const text = await res.text()
      console.error("上传失败", text)
      toast.add({
        severity: 'error',
        summary: t('universal.errToast'),
        detail: t('admin.illustration.commitFail') + ": " + text,
        life: 3000
      })
      return
    }

    // 3. 成功处理
    console.log("上传成功")
    toast.add({
      severity: 'success',
      summary: t('universal.success'),
      detail: t('admin.illustration.commitSuccess'),
      life: 3000
    })

    // 重置表单状态
    newIllustration.value = {
      name: "",
      author_id: "",
      tags_id: [],
      link: "",
      description: "",
      limited: false,
      source: ""
    }
    uploadedFiles.value = []
    previewUrls.value = []
    tagsSelected.value = []
    authorSelected.value = undefined

  } catch (err: any) {
    console.error("err: ", err)
    toast.add({
      severity: 'error',
      summary: t('universal.errToast'),
      detail: t('universal.errToastMessage') + " " + err,
      life: 3000
    })
  }
}

fetchIllustrationTags()
fetchAuthors()

watch(() => actionStore.triggerCommitNewIllustration, (newVal: boolean) => {
  if (newVal) {
    actionStore.resetTriggerCommitNewIllustration()
    console.log('commitNewIllustration', newVal)
    commitNewIllustration()
  }
})

onBeforeMount(() => {
  themeStore.setShowCommitIllustrationBtn(true)
})

onBeforeUnmount(() => {
  themeStore.setShowCommitIllustrationBtn(false)
})

onMounted(() => {
  setTimeout(() => {

  }, 500)

})

</script>

<template>
  <div class="mt-4">
    <PageHeader :title="t('admin.illustration.createIllustrationTitle')"
                :subtitle="t('admin.illustration.createIllustrationSubtitle')"/>
    <div v-if="false">
      <div class="w-full flex flex-col justify-center items-start pl-0 pr-5 pb-3 space-y-2">
        <!-- 行1：标签管理 -->
        <div class="flex flex-row justify-start items-center gap-2">
          <Tag
              size="small"
              class="text-xs font-normal cursor-pointer font-mono pt-0.5 pb-0.5"
              severity="secondary"
              value="Secondary"
          >
            <Option class="w-3 h-3" />
            + t
          </Tag>
          <span class="text-sm opacity-80">进入以管理所有的标签</span>
        </div>

        <!-- 行2：作者管理 -->
        <div class="flex flex-row justify-start items-center gap-2">
          <Tag
              size="small"
              class="text-xs font-normal cursor-pointer font-mono pt-0.5 pb-0.5"
              severity="secondary"
              value="Secondary"
          >
            <Option class="w-3 h-3" />
            + a
          </Tag>
          <span class="text-sm opacity-80">进入以管理所有的作者</span>
        </div>

        <!-- 行3：关闭Dialog -->
        <div class="flex flex-row justify-start items-center gap-2">
          <Tag
              size="small"
              class="text-xs font-normal cursor-pointer font-mono pt-0.5 pb-0.5"
              severity="secondary"
              value="Secondary"
          >
            <CircleArrowOutUpLeft class="w-3 h-3" />
            Esc
          </Tag>
          <span class="text-sm opacity-80">关闭当前的Dialog或者提示框</span>
        </div>
      </div>
    </div>

    <div class="mb-20 mt-10">
      <div class="grid grid-cols-1 md:grid-cols-12 gap-8 px-4 lg:px-20">

        <!-- 左侧：上传 + 预览 -->
        <div class="md:col-span-6">
          <div class="mb-4">
            <Message severity="warn" size="small" closable>{{ t('admin.illustration.noImageUploaded') }}</Message>
          </div>

          <div v-if="previewUrls.length > 0" class="flex flex-col gap-4 mb-4">
            <div
                v-for="(url, idx) in previewUrls"
                :key="idx"
                class="relative group w-full drop-shadow-lg rounded-lg"
            >
              <img
                  :src="url"
                  class="w-full h-auto object-contain rounded-lg"/>
              <div
                  class="absolute bottom-0 left-0 right-0 bg-black/30 text-white text-sm px-3 py-2
             opacity-0 group-hover:opacity-100 transition-opacity duration-300
             flex justify-between items-center rounded-bl-lg rounded-br-lg"
              >
                <span class="truncate">{{ uploadedFiles[idx]?.name || "NO NAME" }}</span>
                <Button
                    label="刪除"
                    icon="pi pi-times"
                    variant="link"
                    severity="danger"
                    size="small"
                    @click="removeFile(idx)"
                />
              </div>
            </div>
          </div>


          <FileUpload
              name="files[]"
              :customUpload="true"
              auto
              multiple
              accept="image/*"
              @select="onFileSelect"
          >
            <!-- 自定義上傳區 -->
            <template #header="{ chooseCallback }">
              <Card @click="chooseCallback" class="shadow-none">
                <template #content>
                  <div
                      class="w-full aspect-[3/1] flex flex-col justify-center items-center cursor-pointer"
                  >
                    <p class="opacity-80 text-center pointer-events-none">
                      {{ t('admin.illustration.imageUploadedTitle') }}
                    </p>
                    <p class="pl-2 pr-2 opacity-50 text-sm mt-2">{{
                        t('admin.illustration.uploadImageDescription')
                      }}</p>
                  </div>
                </template>
              </Card>
            </template>
          </FileUpload>

          <Card v-if="uploadedFiles.length===0" class="rounded-lg mt-4 border-1 shadow-none">

            <template #content>
              <div
                  class="w-full h-64 rounded-lg flex flex-col justify-center items-center cursor-pointer"
              >
                还没有上传图片
              </div>
            </template>
          </Card>
        </div>

        <!-- 右侧：表单 -->

        <div class="md:col-span-6 flex flex-col gap-5">
          <PageHeaderL2 :title="t('universal.illustration.illustrationName')"
                        :subtitle="t('admin.illustration.editIllustrationMeta')"/>


          <div class="flex flex-col gap-2">
            <label for="name">{{ t('universal.illustration.name') }}</label>
            <IconField>
              <InputIcon class="pi pi-images"></InputIcon>
              <InputText autofocus variant="outlined" size="small" id="name" placeholder="夏夜に咲く星"
                         v-model="newIllustration.name" class="w-full"/>
            </IconField>
          </div>

          <div class="flex flex-col gap-2">
            <label for="link">{{ t('universal.illustration.link') }}</label>
            <IconField>
              <InputIcon class="pi pi-link"></InputIcon>
              <InputText size="small" id="link" v-model="newIllustration.link"
                         placeholder="https://www.pixiv.net/artworks/000000000"
                         class="w-full"/>
            </IconField>
          </div>

          <div class="flex flex-col gap-2">
            <label for="description">{{ t('universal.illustration.description') }}</label>
            <Textarea size="small" id="description"
                      placeholder="ひさかたの あめゆくつきを あみにさし わがおほきみは きぬがさにせり。"
                      v-model="newIllustration.description" rows="4" class="w-full"/>
          </div>

          <div class="flex flex-col gap-2">
            <span class="space-x-2">
              <label for="source">{{ t('universal.illustration.original') }}</label>
            <Button link size="small" class="p-0 underline" @click="newIllustration.source='Pixiv'">Pixiv</Button>
            <Button link size="small" class="p-0 underline" @click="newIllustration.source='X'">X</Button>
            <Button link size="small" class="p-0 underline"
                    @click="newIllustration.source='Instagram'">Instagram</Button>
            </span>
            <InputText size="small" id="source" placeholder="Pixiv" v-model="newIllustration.source" class="w-full"/>
          </div>

          <div class="flex flex-col gap-2">
            <label class="text-sm">{{ t('universal.illustration.isLimited') }}</label>
            <div class="flex items-center gap-2">
              <Checkbox size="small" v-model="newIllustration.limited" :binary="true" inputId="limited"/>
              <label for="limited" class="text-sm">{{ t('universal.illustration.isLimitedHint') }}</label>
            </div>
          </div>

          <div class="flex flex-col gap-2">
            <span class="space-x-2">
              <label for="tag">{{ t('universal.illustration.tags') }}</label>
            <Button link size="small" class="p-0 underline"
                    @click="showTagsSelection=!showTagsSelection">選擇標籤</Button>
            </span>
            <div class="flex flex-wrap gap-2">
              <Tag
                  v-for="i in tagsSelected"
                  :key="i.id"
                  icon="pi pi-hashtag"
                  severity="success"
                  size="small"
                  class="text-xs font-normal hover:underline"
                  :value="i.name"
                  @click="delTag(i.id)"
              ></Tag>
              <Tag size="small" class="text-xs font-normal" v-if="tagsSelected.length===0" icon="pi pi-info-circle"
                   severity="warn" :value="t('admin.illustration.tag.noTagsSatisfied')"></Tag>
            </div>

            <Panel :collapsed="!showTagsSelection">

              <template #header>
                <IconField class="w-full border-0">
                  <InputIcon class="pi pi-search"></InputIcon>
                  <InputText size="small" id="link" v-model="tagSearch"
                             :placeholder="t('admin.illustration.tagSearchInputPlaceholder', {name: '女の子'})"
                             class="w-full"
                             @keyup.enter="onPressTagSearch"
                  />
                  <InputIcon class="pi pi-hashtag"></InputIcon>
                </IconField>
              </template>
              <div class="pt-2">
                <Message v-if="tagsResponse?.list.length===0" class="mt-1 mb-1" severity="warn" size="small">
                  {{ t('admin.illustration.tagsSearchNotFound') }}
                </Message>
                <div class="flex flex-row justify-between items-baseline mb-1">
                  <span class="space-x-2" v-if="tagsResponse?.list.length!==0">

                    <Button link size="small" class="p-0 underline" @click="onPressSearchPaBtn('tag', 'decrease')"
                    >{{ t('universal.pageF') }}
                    </Button>
                    <Button link size="small" class="p-0 underline" @click="onPressSearchPaBtn('tag', 'increase')"
                    >{{ t('universal.pageA') }}
                  </Button>
                    <label class="font-normal text-xs font-mono pb-1 opacity-70">
                      {{ t('universal.pages') }}
                    [{{
                        tagsResponse?.page || 1
                      }}/{{ getTotalPagesFunc(tagsResponse?.total || 1, tagsResponse?.page || 1) }}]
                  </label>
                </span>
                  <span class="text-sm">
                  <label>{{ t('admin.illustration.tag.noTagsSatisfied') }}</label>
                   <Button link size="small" class="p-0 underline"
                           @click="() => {dialogTitle='admin.illustration.tagMgr' ;showModalCard.type='tag'; showModalCard.show=true}"
                   >
                    {{ t('admin.illustration.tagMgr') }}
                  </Button>
                </span>
                </div>
                <div class="flex flex-wrap gap-2 justify-start">
                  <Tag
                      v-for="i in tagsResponse?.list"
                      :key="i.id"
                      icon="pi pi-hashtag"
                      size="small"
                      class="text-xs font-normal hover:underline"
                      :value="i.name"
                      @click="addTag(i)"
                  />
                </div>
              </div>
            </Panel>
          </div>

          <div class="flex flex-col gap-2">
            <span class="space-x-2">
              <label for="tag">{{ `${t('universal.illustration.author')}/${t('universal.illustration.eshi')}` }}</label>
            <Button link size="small" class="p-0 underline"
                    @click="showAuthorSelection=!showAuthorSelection">選擇作者</Button>
            </span>
            <div class="gap-2 mb-2">
              <div v-if="authorSelected" class="flex flex-col justify-start items-start">
                <span class="flex flex-row items-center space-x-1">
<!--                  <i class="pi pi-circle"></i>-->
                  <Icon class="opacity-80"><CaretForwardOutline/></Icon>
                  <span class="text-xl font-semibold opacity-75">{{ authorSelected?.name }}</span>
                </span>
                <span class="flex flex-row items-center space-x-1">
                  <Icon class="opacity-80"><CaretForwardOutline/></Icon>
                  <Button link class="p-0 text-sm font-extralight underline" :label="authorSelected.link"
                          @click="openLink(authorSelected.link)">
                </Button>
                </span>

              </div>
              <div v-else>
                <Skeleton class="mb-2" width="16rem" height="1.75rem"></Skeleton>
                <Skeleton width="7rem" height="1rem"></Skeleton>
              </div>

            </div>

            <Panel :collapsed="!showAuthorSelection">

              <template #header>
                <IconField class="w-full border-0">
                  <InputIcon class="pi pi-search"></InputIcon>
                  <InputText size="small" id="link" v-model="authorSearch"
                             :placeholder="t('admin.illustration.authorSearchInputPlaceholder', {name: 'なつき'})"
                             class="w-full"
                             @keyup.enter="onPressAuthorSearch"
                  />
                  <InputIcon class="pi pi-user"></InputIcon>
                </IconField>
              </template>
              <div class="pt-2">
                <Message v-if="authorsResponse?.list.length===0" class="mt-1 mb-1" severity="warn" size="small">
                  {{ t('admin.illustration.tagsSearchNotFound') }}
                </Message>
                <div class="flex flex-row justify-between items-baseline mb-1">
                  <span class="space-x-2" v-if="authorsResponse?.list.length!==0">

                    <Button link size="small" class="p-0 underline" @click="onPressSearchPaBtn('author', 'decrease')"
                    >{{ t('universal.pageF') }}
                    </Button>
                    <Button link size="small" class="p-0 underline" @click="onPressSearchPaBtn('author', 'increase')"
                    >{{ t('universal.pageA') }}
                  </Button>
                    <label class="font-normal text-xs font-mono pb-1 opacity-70">
                      {{ t('universal.pages') }}
                    [{{
                        tagsResponse?.page || 1
                      }}/{{ getTotalPagesFunc(authorsResponse?.total || 1, authorsResponse?.page || 1) }}]
                  </label>
                </span>
                  <span class="text-sm">
                  <label>{{ t('admin.illustration.tag.noAuthorSatisfied') }}</label>
                   <Button link size="small" class="p-0 underline"
                           @click="() => {dialogTitle='admin.illustration.authorMgr' ;showModalCard.type='author'; showModalCard.show=true}"
                   >
                    {{ t('admin.illustration.authorMgr') }}
                  </Button>
                </span>
                </div>
                <div class="flex flex-wrap gap-2 justify-start">
                  <Tag
                      v-for="i in authorsResponse?.list"
                      :key="i.id"
                      icon="pi pi-user"
                      size="small"
                      class="text-xs font-normal hover:underline"
                      :value="i.name"
                      @click="addAuthor(i)"
                  />
                </div>

              </div>
            </Panel>
          </div>

          <Message severity="info">
            <template #icon>
              <Keyboard class="w-4 h-4"/>
            </template>

            <template #default>
              <div class="flex flex-col justify-start">
                <div>该页面的快捷键</div>

              </div>
            </template>
          </Message>

          <div class="space-x-2 space-y-2">

          </div>

          <Button

              size="small"
              label="提交插画"
              icon="pi pi-check"
              class="mt-4 w-fit"
              @click="handleCommitNewIllustration"
          />
        </div>

      </div>
    </div>
  </div>

  <Dialog
      ref="dialogRef"
      :dismissableMask="true"
      v-model:visible="showModalCard.show"
      maximizable modal
      :header="t(dialogTitle as string)"
      :style="{ width: '50rem' }"
      :breakpoints="{ '1199px': '75vw', '575px': '90vw' }"
  >
    <IllustrationTagsMgrPanel v-if="showModalCard.type==='tag'" :update-list="fetchIllustrationTags"/>
    <IllustrationAuthorsMgrPanel v-if="showModalCard.type==='author'" :update-list="fetchAuthors"/>

    <template #footer>
      <Divider class="mt-0 mb-2"/>
      <div class="w-full flex flex-col justify-center items-start pl-5 pr-5 pb-3">
        <div class="gap-2 flex flex-row justify-start items-center">
          <Tag
              size="small"
              class="text-xs font-normal cursor-pointer font-mono pt-0.5 pb-0.5"
              severity="secondary" value="Secondary">
            <Command class="w-3 h-3"/>
            + Enter
          </Tag>
          <span class="text-sm opacity-80">{{ t('admin.illustration.toggleCreateMode') }}</span>

          <Tag
              size="small"
              class="text-xs font-normal cursor-pointer font-mono pt-0.5 pb-0.5 ml-2"
              severity="secondary" value="Secondary">
            <CircleArrowOutUpLeft class="w-3 h-3"/>
            Esc
          </Tag>
          <span class="text-sm opacity-80">{{ t('admin.illustration.esc') }}</span>
        </div>

      </div>
    </template>

  </Dialog>

</template>

<style>

.p-fileupload-header {
  padding: 0;

}

.p-fileupload-content {
  display: none;
}

.p-panel-header {
  height: 35px;
  padding: 0 !important;
}

.p-panel-content {
  padding: 0 !important;
}

.p-panel {
  border: 0;
  background-color: rgba(0, 0, 0, 0) !important;
}

.p-dialog-footer {
  padding: 0;
  display: flex;
  flex-direction: column;
}

</style>
