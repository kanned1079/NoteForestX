<script setup lang="ts">
import {ref, onBeforeMount} from "vue"
import type {IllustrationAuthor} from "~/types/illustration";
import Popover from 'primevue/popover';
import dayjs from "dayjs";
import {useI18n} from "vue-i18n";
import {useToast} from "primevue/usetoast";
import {Icon} from '@vicons/utils'
import {TrashOutline, PencilOutline, AlertOutline, LinkOutline} from "@vicons/ionicons5"



const {t} = useI18n()
const toast = useToast()

const props = defineProps<{
  updateList: () => void // 父组件更新方法
}>()

const animatedAuthor = ref<boolean>(false)
const popoverRef = ref()
const currentAuthor = ref<IllustrationAuthor | null>(null)

const clickAuthor = (event: MouseEvent, author: IllustrationAuthor) => {
  currentAuthor.value = author
  popoverRef.value.toggle(event)
}

const onToast = (severity: 'info' | 'warn' | 'error' | 'success', title: string, detail?: string) => {
  toast.add({severity, summary: title, detail, life: 3000});
}

const copy = async (text: string) => {
  try {
    // Clipboard API 只能在安全上下文（https 或 localhost）下使用
    await navigator.clipboard.writeText(text)
    onToast('success', t('universal.successToast'), t('admin.illustration.tag.copySuccess', {text: `「${text}」`}))
  } catch (err: any) {
    onToast('error', t('universal.errToast'), t('admin.illustration.tag.copyFailure', {msg: `${err}`}))
  }
}

type fetchAuthorsResponseDto = {
  page: number
  size: number
  total: number
  list: IllustrationAuthor[]
}

const authorsSearchPage = ref<number>(1)
const authorsSearchSize = ref<number>(30)
const authorTotal = ref<number>(0)
const authorSearch = ref<string>("")
const lastSearch = ref<string>("")

const authorsResponse = ref<fetchAuthorsResponseDto>()

const fetchAuthorsList = async () => {
  lastSearch.value = authorSearch.value
  animatedAuthor.value = false
  try {
    const data = await $fetch<fetchAuthorsResponseDto>(`/api/admin/illustration_author`, {
      method: "GET",
      query: {
        page: authorsSearchPage.value,
        size: authorsSearchSize.value,
        search: authorSearch.value,
        related: true,
      }
    });
    if (data) {
      authorsResponse.value = data
      authorTotal.value = data.total
      animatedAuthor.value = true
      if (!data.list || data.list.length === 0) {
        toast.add({
          severity: 'warn',
          summary: t('admin.illustration.tagsSearchNotFound'),
          detail: t('admin.illustration.authorsSearchNotFoundHint', {name: authorSearch.value}),
          life: 3000
        });
      }
    }
  } catch (err: any) {
    console.error(err);
    toast.add({
      severity: 'error',
      summary: t('universal.errToast'),
      detail: `${t('universal.errToastMessage')} ${err}`,
      life: 3000
    });
  }
}

// 编辑 Author
const editAuthor = ref<IllustrationAuthor | null>(null)
const showEditDialog = ref<boolean>(false)
const newAuthorName = ref<string>("")
const newAuthorLink = ref<string>("")
const authorBeforeEditing = ref<{ name: string, link: string }>({name: "", link: ""})

const onClickEditBtn = (author: IllustrationAuthor | null) => {
  if (!author) return
  editAuthor.value = author
  authorBeforeEditing.value = {name: author.name, link: author.link}
  newAuthorName.value = author.name
  newAuthorLink.value = author.link || ""
  showEditDialog.value = true
}

const onClickSave = async () => {
  if (!editAuthor.value) return
  editAuthor.value.name = newAuthorName.value.trim()
  editAuthor.value.link = newAuthorLink.value.trim()
  if (!editAuthor.value.name || !editAuthor.value.link) {
    return onToast("warn", t('universal.warnToast'), t('admin.illustration.author.fieldsNotValid'))
  }
  if (
      editAuthor.value.name === authorBeforeEditing.value.name &&
      editAuthor.value.link === authorBeforeEditing.value.link
  ) return

  try {
    await $fetch(`/api/admin/illustration_author/${editAuthor.value.id}`, {
      method: "PUT",
      body: {name: editAuthor.value.name, link: editAuthor.value.link}
    })
    onToast("success", t('universal.successToast'), t('admin.illustration.author.editSuccess'))
    props.updateList()
    showEditDialog.value = false
  } catch (err: any) {
    console.error(err)
    onToast("error", t('universal.errToast'), `${t('universal.errToastMessage')} ${err}`)
  }
}

// 删除 Author
const showDeleteBtn = ref<boolean>(true)
const showDeleteConfirmBtn = ref<boolean>(false)

const switchToConfirm = () => {
  showDeleteBtn.value = false
  showDeleteConfirmBtn.value = true
}

const handleDeleteAuthorById = async (id?: string | null) => {
  if (!id) return
  try {
    await $fetch(`/api/admin/illustration_author/${id}`, {method: "DELETE"})
    onToast("success", t('universal.successToast'), t('admin.illustration.author.deleteTagSuccessHint'))
    props.updateList()
  } catch (err: any) {
    let code = err?.response?.status ?? err?.statusCode
    if (code === 409) {
      onToast("error", t('universal.errToast'), t('admin.illustration.author.noAuthorSatisfied'))
    } else {
      onToast("error", t('universal.errToast'), `${t('universal.errToastMessage')} ${err}`)
    }
  }
}

const onClickDeleteConfirmAuthor = async (id?: string) => {
  await handleDeleteAuthorById(id)
  await fetchAuthorsList()
}

const isAddNewAuthor = ref<boolean>(false)
const newAuthorAdd = ref<{ name: string, link: string }>({
  name: "",
  link: ""
})

const handleAddNewAuthor = async () => {
  newAuthorAdd.value.name = newAuthorAdd.value.name.trim()
  newAuthorAdd.value.link = newAuthorAdd.value.link.trim()
  if (!newAuthorAdd.value.name || !newAuthorAdd.value.link) {
    return onToast("warn", t('universal.warnToast'), t('universal.formNotInvalid'))
  }
  try {
    await $fetch<IllustrationAuthor>(`/api/admin/illustration_author`, {
      method: "POST",
      body: {
        ...newAuthorAdd.value
      }
    })
    onToast("success", t('universal.successToast'), t('admin.illustration.author.authorAdded'))
    isAddNewAuthor.value = false
    authorSearch.value = ""
    await fetchAuthorsList()
    await props.updateList()
  } catch (err: any) {
    let code = err?.response?.status ?? err?.statusCode
    if (code === 409) {
      onToast("error", t('universal.errToast'), t('admin.illustration.author.authorNameConflict'))
    } else {
      onToast("error", t('universal.errToast'), `${t('universal.errToastMessage')} ${err}`)
    }
  }
}

const onClickAdd = async () => {
  if (authorSearch.value.trim()) {
    newAuthorAdd.value.name = authorSearch.value
    isAddNewAuthor.value = true
  } else {
    onToast("info", t('universal.info'), t('admin.illustration.author.emptySearchHint'))
  }
}


fetchAuthorsList()

</script>

<template>
  <div class="card">
    <Transition name="fade-only" mode="out-in">
    <div v-if="!isAddNewAuthor">
      <IconField class="w-full border-0 mt-2 mb-4">
        <InputIcon class="pi pi-search"></InputIcon>
        <InputText
            autofocus
            size="medium" v-model="authorSearch"
            :placeholder="t('admin.illustration.authorSearchInputPlaceholder', {name: 'なつき'})"
            class="w-full"
            @keyup.enter="fetchAuthorsList"
            @keydown.enter.meta="onClickAdd"
            @keydown.enter.ctrl="onClickAdd"
        />
        <InputIcon class="pi pi-user"></InputIcon>
      </IconField>
    </div>
    </Transition>

    <Transition name="fade-only" mode="out-in">
      <div v-if="isAddNewAuthor" class="flex flex-col gap-2">

        <div class="flex flex-col gap-2">
          <label for="name">{{ t('universal.illustration.name') }}</label>
          <IconField>
            <InputIcon class="pi pi-user"></InputIcon>
            <InputText autofocus variant="outlined" size="small" id="name" placeholder="なつき"
                       v-model="newAuthorAdd.name" class="w-full"/>
          </IconField>
        </div>

        <div class="flex flex-col gap-2">
          <label for="name">{{ t('universal.illustration.name') }}</label>
          <IconField>
            <InputIcon class="pi pi-link"></InputIcon>
            <InputText autofocus variant="outlined" size="small" id="name" placeholder="https://www.pixiv.net/users/xxxxxxxx"
                       v-model="newAuthorAdd.link" class="w-full"/>
          </IconField>
        </div>

        <div class="flex justify-end gap-2 mt-4">
          <Button icon="pi pi-replay" class="p" size="small" type="button" :label="t('universal.illustration.cancel')" severity="secondary"
                  @click="() => {isAddNewAuthor = false; authorSearch=''; fetchAuthorsList()}"></Button>
          <Button icon="pi pi-save" size="small" type="button" :label="t('universal.illustration.save')" @click="handleAddNewAuthor"></Button>
        </div>
      </div>
    </Transition>

    <Transition name="fade-only" mode="out-in">
      <div v-if="authorsResponse?.list.length===0 && !isAddNewAuthor">
        <NotFoundResult
            class="mt-4"
            :title="t('admin.illustration.tagsSearchNotFound')"
            :text="t('admin.illustration.authorsSearchNotFoundHint', {name: lastSearch})"
        />
      </div>
    </Transition>

    <Transition name="fade-only">
      <div v-if="authorsResponse?.list.length!==0 && animatedAuthor && !isAddNewAuthor" class="flex flex-wrap gap-2 mt-4">
        <Tag
            v-for="i in authorsResponse?.list"
            :key="i.id"
            icon="pi pi-user"
            size="small"
            class="text-xs font-normal hover:underline cursor-pointer"
            :value="i.name"
            @click="(e) => clickAuthor(e, i)"
        />
      </div>
    </Transition>


    <Transition name="fade-only">
    <div class="mt-6" v-if="!isAddNewAuthor">
      <MyPaginationBar
          v-model:page="authorsSearchPage"
          v-model:size="authorsSearchSize"
          :total="authorTotal"
          :fetch-data="fetchAuthorsList"
          place="right"
      />
    </div>
    </Transition>

    <!-- Popover -->
    <Popover ref="popoverRef" class="opacity-95" @hide="() => {showDeleteBtn=true; showDeleteConfirmBtn=false}">
      <div v-if="currentAuthor" class="text-xs space-y-1">
        <div class="font-mono hover:underline" @click="copy(currentAuthor?.name as string)">
          <b class="opacity-80">{{ t('universal.illustration.author') }}: </b>{{ currentAuthor?.name }}
        </div>
        <div class="font-mono hover:underline">
          <b class="opacity-80">{{
              t('universal.illustration.relatedIllustrationCount')
            }}: </b>{{ currentAuthor?.related || 0 }}
        </div>
        <div class="font-mono hover:underline" @click="copy(currentAuthor?.link as string)">
          <b class="opacity-80">{{ t('universal.illustration.link') }}: </b>{{ currentAuthor?.link }}
        </div>
        <div class="font-mono"><b class="opacity-80">{{
            t('universal.illustration.createdAt')
          }}: </b>{{ dayjs(currentAuthor?.created_at).format('YYYY-MM-DD HH:mm:ss') }}
        </div>
        <div class="font-mono"><b class="opacity-80">{{
            t('universal.illustration.updatedAt')
          }}: </b>{{ dayjs(currentAuthor?.updated_at).format('YYYY-MM-DD HH:mm:ss') }}
        </div>
        <div class="pt-1">
          <span class="space-x-2">
            <Tag
                size="small"
                class="text-xs font-normal hover:underline cursor-pointer"
                :value="t('universal.illustration.edit')"
                @click="onClickEditBtn(currentAuthor)"
            >
              <template #icon>
                <Icon><PencilOutline/></Icon>
              </template>
            </Tag>

            <Transition name="fade-only" mode="out-in">
              <Tag
                  v-if="showDeleteBtn"
                  key="delete"
                  size="small"
                  severity="danger"
                  class="text-xs font-normal hover:underline cursor-pointer"
                  :value="t('universal.illustration.delete')"
                  @click="switchToConfirm"
              >
                <template #icon>
                  <Icon><TrashOutline/></Icon>
                </template>
              </Tag>

              <Tag
                  v-else
                  key="confirm"
                  size="small"
                  severity="contrast"
                  class="text-xs font-normal hover:underline cursor-pointer"
                  :value="t('universal.illustration.deleteConfirm')"
                  @click="onClickDeleteConfirmAuthor(currentAuthor?.id)"
              >
                <template #icon>
                  <Icon><AlertOutline/></Icon>
                </template>
              </Tag>
            </Transition>
          </span>
        </div>
      </div>
    </Popover>

    <!-- Edit Dialog -->
    <Dialog :show-header="true" :dismissable-mask="true" v-model:visible="showEditDialog"
            :header="`${t('universal.illustration.edit')} 「${authorBeforeEditing.name}」`" :style="{ width: '25rem' }"
            position="top" :modal="true"
            :draggable="false">
      <div class="flex flex-col gap-2">
        <label for="link">{{ t('universal.illustration.author') }}</label>
        <IconField>
          <InputIcon class="pi pi-user"></InputIcon>
          <InputText autofocus size="small" id="link" v-model="newAuthorName"
                     :placeholder="authorBeforeEditing.name"
                     class="w-full"
                     @keyup.enter="onClickSave"
          />
        </IconField>

        <label for="authorLink">{{ t('universal.illustration.link') }}</label>
        <IconField>
          <InputIcon class="pi pi-link"></InputIcon>
          <InputText size="small" id="authorLink" v-model="newAuthorLink"
                     :placeholder="authorBeforeEditing.link"
                     class="w-full"
                     @keyup.enter="onClickSave"
          />
        </IconField>
      </div>

      <span class="text-surface-500 dark:text-surface-400 block mb-4 text-sm mt-2">
        {{ t('admin.illustration.author.editTagNameHint') }}
      </span>

      <div class="flex justify-end gap-2 mt-4">
        <Button icon="pi pi-replay" class="p" size="small" type="button" :label="t('universal.illustration.cancel')" severity="secondary"
                @click="showEditDialog = false"></Button>
        <Button icon="pi pi-save" size="small" type="button" :label="t('universal.illustration.save')" @click="onClickSave"></Button>
      </div>
    </Dialog>
  </div>
</template>

<style scoped>
.card {
  position: relative;
}

/* 淡入淡出動畫 */
.fade-only-enter-active,
.fade-only-leave-active {
  transition: opacity 0.25s ease;
}

.fade-only-enter-from,
.fade-only-leave-to {
  opacity: 0;
}
</style>