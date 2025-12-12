<script setup lang="ts">
import {ref, onBeforeMount} from "vue"
import type {IllustrationAuthor, IllustrationTag} from "~/types/illustration";
import Popover from 'primevue/popover';
import dayjs from "dayjs";
import {useI18n} from "vue-i18n";
import {useToast} from "primevue/usetoast";
import {Icon} from '@vicons/utils'
import {TrashOutline, PencilOutline, AlertOutline} from "@vicons/ionicons5"
import NotFoundResult from "~/components/NotFoundResult.vue";

const {t} = useI18n()
const toast = useToast()
const props = defineProps<{
  updateList: () => void //　副組件執行更新
}>()

const animatedTag = ref<boolean>(false)

const popoverRef = ref()
const currentTag = ref<IllustrationTag | null>(null)
const curr = ref<IllustrationAuthor>

const clickTag = (event: MouseEvent, tag: IllustrationTag) => {
  currentTag.value = tag
  popoverRef.value.toggle(event)
}

const onToast = (severity: 'info' | 'warn' | 'error' | 'success', title: string, detail?: string) => {
  toast.add({
    severity: severity,
    summary: title,
    detail: detail,
    life: 3000,
  });
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

type fetchTagsResponseDto = {
  page: number
  size: number
  total: number
  list: IllustrationTag[]
}

const tagsSearchPage = ref<number>(1)
const tagsSearchSize = ref<number>(30)
const tagTotal = ref<number>(0)
const tagSearch = ref<string>("")
const lastSearch = ref<string>("")

const tagsResponse = ref<fetchTagsResponseDto>()

const fetchTagsList = async () => {
  lastSearch.value = tagSearch.value
  animatedTag.value = false
  try {
    const data = await $fetch<fetchTagsResponseDto>(`/api/admin/illustration_tag`, {
      method: "GET",
      query: {
        page: tagsSearchPage.value,
        size: tagsSearchSize.value,
        search: tagSearch.value,
        related: true
      }
    });
    if (data) {
      tagsResponse.value = data
      tagTotal.value = data.total
      animatedTag.value = true
      if (!data.list || tagsResponse.value?.list.length == 0) {
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

const editTag = ref<IllustrationTag | null>(null)

const showEditDialog = ref<boolean>(false)
const newTagName = ref<string>("")
const tagNameBeforeEditing = ref<string>("")

const onClickEditBtn = (tag: IllustrationTag | null) => {
  if (tag) {
    editTag.value = tag
    newTagName.value = tag.name
    tagNameBeforeEditing.value = tag.name
    showEditDialog.value = true
  }
}

const onClickSave = async () => {
  if (editTag.value) {
    editTag.value.name = newTagName.value.trim()
    if (editTag.value.name && editTag.value.name !== tagNameBeforeEditing.value) {
      // 允許保存修改
      try {
        const data = await $fetch<{
          tag?: IllustrationTag,
          message: string
        }>(`/api/admin/illustration_tag/${editTag.value.id}`, {
          method: "PUT",
          body: {
            name: editTag.value.name
          }
        });
        if (data) {
          onToast("success", t('universal.successToast'), t('admin.illustration.tag.editTagNameSuccessHint'))
          props.updateList()
          showEditDialog.value = false
        }
      } catch (err: any) {
        console.error(err)
        onToast(
            "error",
            t('universal.errToast'),
            `${t('universal.errToastMessage')} ${err}`
        )
      }
      // end...
    } else {
      return onToast("warn", t('universal.warnToast'), t('admin.illustration.tag.tagNameNotValid'))
    }
  }
}

const handleDeleteTagById = async (id: string | null | undefined) => {
  if (id) {
    try {
      const data = await $fetch<{
        message: string
      }>(`/api/admin/illustration_tag/${id}`, {
        method: "DELETE",
      });
      if (data) {
        onToast("success", t('universal.successToast'), t('admin.illustration.tag.deleteTagSuccessHint'))
        props.updateList()
      }
    } catch (err: any) {
      let code = err?.response?.status ?? err?.statusCode
      if (code === 409) {
        onToast(
            "error",
            t('universal.errToast'),
            `${t('admin.illustration.tag.deleteTagConflict')}`
        )
      } else {
        onToast(
            "error",
            t('universal.errToast'),
            `${t('universal.errToastMessage')} ${err}`
        )
      }
    }
  }
}

const showDeleteBtn = ref<boolean>(true)
const showDeleteConfirmBtn = ref<boolean>(false)

const switchToConfirm = () => {
  showDeleteBtn.value = false
  showDeleteConfirmBtn.value = true
}

const onClickDeleteConfirmTag = async (id: string | undefined) => {
  console.log(id)
  await handleDeleteTagById(id)
  await fetchTagsList()
}


const isAddNewTag = ref<boolean>(false)
const newTagAddName = ref<string>('')

const handleAddNewTag = async () => {
  // newAuthorAdd.value.name = newAuthorAdd.value.name.trim()
  // newAuthorAdd.value.link = newAuthorAdd.value.link.trim()
  newTagAddName.value = newTagAddName.value.trim()
  if (!newTagAddName.value) {
    return onToast("warn", t('universal.warnToast'), t('universal.formNotInvalid'))
  }
  try {
    await $fetch<IllustrationTag>(`/api/admin/illustration_tag`, {
      method: "POST",
      body: {
        name: newTagAddName.value
      }
    })
    onToast("success", t('universal.successToast'), t('admin.illustration.tag.authorAdded'))
    isAddNewTag.value = false
    tagSearch.value = ""
    // await fetchAuthorsList()
    await fetchTagsList()
    await props.updateList()
    // await props.updateList()
  } catch (err: any) {
    let code = err?.response?.status ?? err?.statusCode
    if (code === 409) {
      onToast("error", t('universal.errToast'), `${t('admin.illustration.tag.tagNameConflict')} ${err}`)
    } else {
      onToast("error", t('universal.errToast'), `${t('universal.errToastMessage')} ${err}`)
    }
  }
}

const onClickAdd = async () => {
  if (tagSearch.value.trim()) {
    newTagAddName.value = tagSearch.value
    isAddNewTag.value = true
  } else {
    // todo: defefwwf hint change
    onToast("info", t('universal.info'), t('admin.illustration.author.emptySearchHint'))
  }
}


fetchTagsList()

</script>

<template>
  <div class="card">

    <Transition name="fade-only" mode="out-in">

      <div v-if="!isAddNewTag">
        <IconField class="w-full border-0 mt-2 mb-4">
          <InputIcon class="pi pi-search"></InputIcon>
          <InputText size="medium" id="link" v-model="tagSearch"
                     :placeholder="t('admin.illustration.tagSearchInputPlaceholder', {name: '女の子'})"
                     class="w-full"
                     @keyup.enter="fetchTagsList"
                     @keydown.enter.meta="onClickAdd"
                     @keydown.enter.ctrl="onClickAdd"
          />
          <InputIcon class="pi pi-hashtag"></InputIcon>
        </IconField>
      </div>
    </Transition>

    <Transition name="fade-only" mode="out-in">
      <div v-if="isAddNewTag" class="flex flex-col gap-2">

        <div class="flex flex-col gap-2">
          <label for="name">{{ t('admin.illustration.tag.newTagName') }}</label>
          <IconField>
            <InputIcon class="pi pi-hashtag"></InputIcon>
            <InputText autofocus variant="outlined" size="small" id="name" placeholder="なつき"
                       v-model="newTagAddName" class="w-full"/>
          </IconField>
        </div>

        <div class="flex justify-end gap-2 mt-4">
          <Button icon="pi pi-replay" class="p" size="small" type="button" :label="t('universal.illustration.cancel')" severity="secondary"
                  @click="() => {isAddNewTag = false; tagSearch=''; fetchTagsList()}"></Button>
          <Button icon="pi pi-save" size="small" type="button" :label="t('universal.illustration.save')" @click="handleAddNewTag"></Button>
        </div>
      </div>
    </Transition>


    <Transition name="fade-only" mode="out-in">
      <div v-if="tagsResponse?.list.length===0 && !isAddNewTag">
      <NotFoundResult
          v-if="tagsResponse?.list.length===0"
          class="mt-4"
          :title="t('admin.illustration.tagsSearchNotFound')"
          :text="t('admin.illustration.tagsSearchNotFoundHint', {name: lastSearch})"/>
      </div>
    </Transition>


    <Transition name="fade-only">
      <div v-if="tagsResponse?.list.length!==0 && animatedTag && !isAddNewTag" class="flex flex-wrap gap-2 mt-4">
        <Tag
            v-for="i in tagsResponse?.list"
            :key="i.id"
            icon="pi pi-hashtag"
            size="small"
            class="text-xs font-normal hover:underline cursor-pointer"
            :value="i.name"
            @click="(e) => clickTag(e, i)"
        />
      </div>
    </Transition>


    <Transition name="fade-only">
      <div class="mt-6" v-if="!isAddNewTag">
      <MyPaginationBar
          v-model:page="tagsSearchPage"
          v-model:size="tagsSearchSize"
          :total="tagTotal"
          :fetch-data="fetchTagsList"
          place="right"
      />
    </div>
    </Transition>

    <!-- Popover -->
    <Popover ref="popoverRef" class="opacity-95" @hide="() => {showDeleteBtn=true; showDeleteConfirmBtn=false}">
      <div v-if="currentTag as IllustrationTag" class="text-xs space-y-1">
        <div class="font-mono hover:underline" @click="copy(currentTag?.name as string)"><b
            class="opacity-80">{{ t('universal.illustration.tags') }}: </b>{{ currentTag?.name }}
        </div>
        <div class="font-mono"><b class="opacity-80">{{
            t('universal.illustration.relatedIllustrationCount')
          }}: </b>{{ currentTag?.related || "NO RESULT" }}
        </div>
        <div class="font-mono hover:underline" @click="copy(currentTag?.id as string)"><b
            class="opacity-80">Id: </b>{{ currentTag?.id }}
        </div>
        <div class="font-mono"><b class="opacity-80">{{
            t('universal.illustration.createdAt')
          }}: </b>{{ dayjs(currentTag?.created_at).format('YYYY-MM-DD HH:mm:ss') }}
        </div>
        <div class="font-mono"><b class="opacity-80">{{
            t('universal.illustration.updatedAt')
          }}: </b>{{ dayjs(currentTag?.updated_at).format('YYYY-MM-DD HH:mm:ss') }}
        </div>
        <div class="pt-1">
          <span class="space-x-2">
           <Tag
               size="small"
               class="text-xs font-normal hover:underline cursor-pointer"
               :value="t('universal.illustration.edit')"
               @click="onClickEditBtn(currentTag)"
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
          @click="onClickDeleteConfirmTag(currentTag?.id)"
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
  </div>

  <Dialog :show-header="true" :dismissable-mask="true" v-model:visible="showEditDialog"
          :header="`${t('universal.illustration.edit')} 「${tagNameBeforeEditing}」`" :style="{ width: '25rem' }"
          position="top" :modal="true"
          :draggable="false">

    <div class="flex flex-col gap-2 ">
      <label for="link">{{ t('universal.illustration.tags') }}</label>
      <IconField>
        <InputIcon class="pi pi-hashtag"></InputIcon>
        <InputText autofocus size="small" id="link" v-model="newTagName"
                   :placeholder="tagNameBeforeEditing"
                   class="w-full"
                   @keyup.enter="onClickSave"
        />
      </IconField>
    </div>

    <span
        class="text-surface-500 dark:text-surface-400 block mb-4 text-sm mt-2">{{
        t('admin.illustration.tag.editTagNameHint')
      }}</span>

    <div class="flex justify-end gap-2 mt-4">
      <Button class="p" size="small" type="button" :label="t('universal.illustration.cancel')" severity="secondary"
              @click="showEditDialog = false"></Button>
      <Button size="small" type="button" :label="t('universal.illustration.save')" @click="onClickSave"></Button>
    </div>
  </Dialog>

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