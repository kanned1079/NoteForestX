<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { defineProps, defineEmits, computed, onMounted } from "vue"

const { t } = useI18n()
const STORAGE_KEY = 'user-pagination-size'

const props = defineProps<{
  page: number
  size: number
  total: number
  fetchData: () => Promise<void>
  place?: "left" | "right"
}>()

const emit = defineEmits<{
  (e: "update:page", value: number): void
  (e: "update:size", value: number): void
}>()

/**
 * 挂载处理：
 * 检查本地存储的值。如果与当前 props 不一致，说明用户之前在其他页面选过不同的 size，
 * 则更新父组件的状态并重新抓取数据。
 */
onMounted(async () => {
  if (import.meta.client) {
    const savedSize = localStorage.getItem(STORAGE_KEY)
    if (savedSize) {
      const sizeNum = parseInt(savedSize)
      if (sizeNum !== props.size) {
        emit("update:size", sizeNum)
        emit("update:page", 1)
        await props.fetchData()
      }
    }
  }
})

// 页码改变
const onChangePage = async (event: { page: number }) => {
  const newPage = event.page + 1
  emit("update:page", newPage)
  await props.fetchData()
}

// 每页数量改变
const onChangeSize = async (value: number) => {
  const newSize = value > 0 ? value : 15

  if (import.meta.client) {
    localStorage.setItem(STORAGE_KEY, newSize.toString())
  }

  emit("update:size", newSize)
  emit("update:page", 1) // 切换 size 后重置回第一页
  await props.fetchData()
}

const justifyClass = computed(() => {
  return props.place === "left" ? "justify-start" : "justify-end"
})
</script>


<template>
  <div class="flex flex-row items-center gap-3" :class="justifyClass">

    <Select
        size="small"
        @update:modelValue="onChangeSize"
        :options="[15, 30, 60 ,120]"
        class="w-auto md:w-40"
        :default-value="props.size"
    >
      <template #value>
        {{ t(`universal.pagination.itemsPerPage`, {size: props.size}) }}
      </template>

      <template #option="slotProps">
        <div class="flex items-center gap-2">
          <i class="pi pi-align-right text-sm"></i>
          <div class="text-sm">{{  t(`universal.pagination.itemsPerPage`, {size: slotProps.option}) }}</div>
        </div>
      </template>

      <template #dropdownicon>
        <i class="pi pi-chevron-down" />
      </template>
    </Select>

    <Paginator
        :rows="size"
        :totalRecords="total"
        :first="(page - 1) * size"
        @page="onChangePage"
        class="w-auto text-sm"
    />
  </div>
</template>

<style>
:root {
  --p-paginator-gap: 0.25rem;
  --p-paginator-nav-button-width: 1.75rem;
  --p-paginator-nav-button-height: 1.75rem;
  --p-paginator-nav-button-border-radius: 0.375rem;
  --p-paginator-transition-duration: 150ms;
}

.p-select-option {
  height: 2rem;
}

.p-paginator {
  background-color: rgba(0,0,0,0.0);
}

:root {
  --p-select-padding-y: 0.25rem;
  --p-select-padding-x: 0.75rem;
  --p-select-sm-font-size: 0.875rem;
  --p-select-sm-padding-y: 0.25rem;
}
</style>