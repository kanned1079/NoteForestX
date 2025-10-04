<script setup lang="ts">
import {ref} from "vue"
import {useI18n} from "vue-i18n"
import PageHeader from "~/components/PageHeader.vue"
import FileUpload, {type FileUploadSelectEvent} from "primevue/fileupload"
import InputText from "primevue/inputtext"
import Textarea from "primevue/textarea"
import Checkbox from "primevue/checkbox"
import Button from "primevue/button"
import {useToast} from 'primevue/usetoast';
import PageHeaderL2 from "~/components/PageHeaderL2.vue";
import type {Illustration, IllustrationTag} from "~/types/illustration";

const {t} = useI18n()
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

const tagsResponse = ref<fetchTagsResponseDto>()

const showTagsSelection = ref<boolean>(false)
const tagSearch = ref<string>("")
const tagsSelected = ref<IllustrationTag[]>([])

const tagsSearchPage = ref<number>(1)
const tagsSearchSize = ref<number>(40)
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
      tagsResponse.value = data
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

const onPressTagSearch = async () => {
  console.log("tag search")
  console.log(tagSearch.value)
  await fetchIllustrationTags()
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

// fetchIllustrationTags()

const getTotalPagesFunc = (total: number, size: number): number => {
  if (!tagsResponse.value) return 1
  return Math.ceil(tagsResponse.value.total / tagsResponse.value.size)
}

const onPressTagSearchPaBtn = async (op: 'increase' | 'decrease') => {
  let overIndexErr = () => {
    toast.add({
      severity: 'info',
      summary: t('universal.info'),
      detail: t('universal.overIndex'),
      life: 3000,
    });
  }

  let maxPages = getTotalPagesFunc(tagsResponse.value?.total || 1, tagsResponse.value?.page || 1)
  switch (op) {
    case "increase": {// 向後一頁
      if (tagsSearchPage.value >= maxPages) return overIndexErr()
      tagsSearchPage.value += 1;
    }
      break
    case "decrease": { //
      if (tagsSearchPage.value <= 1) return overIndexErr()
      tagsSearchPage.value -= 1;
    }
      break
  }
  await fetchIllustrationTags()
}


onMounted(() => {
  setTimeout(() => {
    fetchIllustrationTags()
  }, 500)
})

</script>

<template>
  <div class="mt-4">
    <PageHeader :title="t('admin.illustration.createIllustrationTitle')"
                :subtitle="t('admin.illustration.createIllustrationSubtitle')"/>

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
              <!-- 圖片填滿 -->
              <img
                  :src="url"
                  class="w-full h-auto object-contain rounded-lg"
              />

              <!-- 底部懸浮信息 -->
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


          <!--          <Skeleton width="100%" height="150px" class="rounded-lg">-->


          <!--          </Skeleton>-->

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
          <PageHeaderL2 title="插畫集信息" subtitle="編輯該插畫集信息"/>


          <div class="flex flex-col gap-2">
            <label for="name">Illustration name</label>
            <IconField>
              <InputIcon class="pi pi-images"></InputIcon>
              <InputText autofocus variant="outlined" size="small" id="name" placeholder="夏夜に咲く星"
                         v-model="newIllustration.name" class="w-full"/>
            </IconField>
          </div>

          <!--          <div class="flex flex-col gap-2">-->
          <!--            <label for="author">作者ID</label>-->
          <!--            <InputText id="author" v-model="newIllustration.author_id" class="w-full"/>-->
          <!--          </div>-->

          <div class="flex flex-col gap-2">
            <label for="link">Original link</label>
            <IconField>
              <InputIcon class="pi pi-link"></InputIcon>
              <InputText size="small" id="link" v-model="newIllustration.link"
                         placeholder="https://www.pixiv.net/artworks/000000000"
                         class="w-full"/>
            </IconField>
          </div>

          <div class="flex flex-col gap-2">
            <label for="description">Description</label>
            <Textarea size="small" id="description"
                      placeholder="ひさかたの あめゆくつきを あみにさし わがおほきみは きぬがさにせり。"
                      v-model="newIllustration.description" rows="4" class="w-full"/>
          </div>

          <div class="flex flex-col gap-2">
            <span class="space-x-2">
              <label for="source">Original</label>
            <Button link size="small" class="p-0 underline" @click="newIllustration.source='Pixiv'">Pixiv</Button>
            <Button link size="small" class="p-0 underline" @click="newIllustration.source='X'">X</Button>
            <Button link size="small" class="p-0 underline"
                    @click="newIllustration.source='Instagram'">Instagram</Button>
            </span>
            <InputText size="small" id="source" placeholder="Pixiv" v-model="newIllustration.source" class="w-full"/>
          </div>

          <div class="flex flex-col gap-2">
            <label class="text-sm">Is Limited (NSFW)</label>
            <div class="flex items-center gap-2">
              <Checkbox size="small" v-model="newIllustration.limited" :binary="true" inputId="limited"/>
              <label for="limited" class="text-sm">如果被选择，该插画集将对部分用户不可见。</label>
            </div>
          </div>

          <div class="flex flex-col gap-2">
            <span class="space-x-2">
              <label for="tag">Tags</label>
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
              <Tag size="small" class="text-sm font-normal" v-if="tagsSelected.length===0" icon="pi pi-info-circle"
                   severity="warn" value="你还没有选择标签"></Tag>
            </div>

            <Panel :collapsed="!showTagsSelection">

              <template #header>
                <IconField class="w-full border-0">
                  <InputIcon class="pi pi-search"></InputIcon>
                  <InputText size="small" id="link" v-model="tagSearch" placeholder="女の子"
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
                <div class="flex flex-row justify-between items-center mb-1">
                  <span class="space-x-2" v-if="tagsResponse?.list.length!==0">
<!--                  <label for="tag">Tags</label>-->
                  <label class="text-sm font-semibold pb-1">選取Tag</label>
                  <Button link size="small" class="p-0 underline"
                          @click="onPressTagSearchPaBtn('decrease')"
                  >
                    {{ t('universal.pageF') }}
                  </Button>
                    <Button link size="small" class="p-0 underline"
                            @click="onPressTagSearchPaBtn('increase')"
                    >
                    {{ t('universal.pageA') }}
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
                <!--                <div class="flex flex-wrap gap-2">-->
                <!--                  <Tag-->
                <!--                      v-for="i in tagsResponse?.list"-->
                <!--                      :key="i.id"-->
                <!--                      icon="pi pi-hashtag"-->
                <!--                      size="small"-->
                <!--                      class="text-xs font-normal hover:underline"-->
                <!--                      :value="i.name"-->
                <!--                      @click="addTag(i)"-->
                <!--                  ></Tag>-->
                <!--                </div>-->

                <div class="space-x-2 mt-2 opacity-70">
                  <label class="text-sm font-semibold pb-1">
                    分頁參數
                  </label>
                  <label class="text-sm font-normal font-mono pb-1">
                    [{{ t('universal.current') }} {{ tagsResponse?.page || 1 }}/{{ t('universal.total') }}
                    {{ getTotalPagesFunc(tagsResponse?.total || 1, tagsResponse?.page || 1) }}]
                  </label>
                  <!--                  <label class="text-sm font-normal font-mono pb-1 ">{{ tagsResponse?.total }}</label>-->
                </div>
              </div>
            </Panel>
          </div>

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
</style>
