<script setup lang="ts">
import {useI18n} from "vue-i18n";
import { defineProps, defineEmits, computed, ref } from "vue"

const {t} = useI18n()
const sizeLabel = ref<number>()

const props = defineProps<{
  page: number
  size: number
  total: number
  fetchData: () => Promise<void>
  place?: "left" | "right" // ✅ 可选，默认为右侧
}>()

const emit = defineEmits<{
  (e: "update:page", value: number): void
  (e: "update:size", value: number): void
}>()

// 页码改变（PrimeVue 的 event.page 从 0 开始）
const onChangePage = async (event: { page: number }) => {
  const newPage = event.page + 1 // ✅ 转换为 1 起始
  emit("update:page", newPage)
  await props.fetchData()
}


// 每页数量改变 -> 同时重置页码为1
const onChangeSize = async (value: number) => {
  const newSize = value > 0 ? value : 15
  emit("update:size", newSize)
  emit("update:page", 1) // ✅ 重置页码
  await props.fetchData()
}

// ✅ 动态对齐样式
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

      <template #header>
        <div v-if="false" class="font-light text-xs pl-3 pr-3 pt-2">每页显示数量</div>
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

:root {
  --p-select-padding-y: 0.25rem;
  --p-select-padding-x: 0.75rem;
  --p-select-sm-font-size: 0.875rem;
  --p-select-sm-padding-y: 0.25rem;
}
</style>