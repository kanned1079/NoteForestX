<script setup lang="ts">
import {ref, onBeforeMount} from "vue"
import type {IllustrationTag} from "~/types/illustration";
import Popover from 'primevue/popover';
import dayjs from "dayjs";
import {useI18n} from "vue-i18n";
import {useToast} from "primevue/usetoast";

const {t} = useI18n()
const toast = useToast()
const props = defineProps<{
  tag_list?: IllustrationTag[]
}>()


const tagsArr = ref<IllustrationTag[]>([])
const popoverRef = ref()
const currentTag = ref<IllustrationTag | null>(null)

// 點擊 tag 顯示 Popover
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

onBeforeMount(() => {
  if (props.tag_list?.length) {
    tagsArr.value = [...props.tag_list]
  }
})
</script>

<template>
  <div class="card">
    <div class="flex flex-wrap gap-2">
      <Tag
          v-for="i in tagsArr"
          :key="i.id"
          icon="pi pi-hashtag"
          size="small"
          class="text-xs font-normal hover:underline cursor-pointer"
          :value="i.name"
          @click="(e) => clickTag(e, i)"
      />
    </div>

    <!-- Popover -->
    <Popover ref="popoverRef" class="opacity-95" dismissable>
      <div v-if="currentTag as IllustrationTag" class="text-xs space-y-1">
        <div class="font-mono hover:underline" @click="copy(currentTag?.name as string)"><b class="opacity-80">{{ t('universal.illustration.tags') }}: </b>{{ currentTag?.name }}</div>
        <div class="font-mono hover:underline" @click="copy(currentTag?.id as string)"><b class="opacity-80">Id: </b>{{ currentTag?.id }}</div>
        <div class="font-mono"><b class="opacity-80">{{ t('universal.illustration.createdAt') }}: </b>{{ dayjs(currentTag?.created_at).format('YYYY-MM-DD HH:mm:ss') }}</div>
        <div class="font-mono"><b class="opacity-80">{{ t('universal.illustration.updatedAt') }}: </b>{{ dayjs(currentTag?.updated_at).format('YYYY-MM-DD HH:mm:ss') }}</div>
      </div>
    </Popover>
  </div>
</template>

<style scoped>
.card {
  position: relative;
}

/* 隐藏 Popover 的箭头 */
.p-popover-arrow {
  display: none !important;
}
</style>