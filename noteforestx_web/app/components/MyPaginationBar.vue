<script setup lang="ts">
import { defineProps, defineEmits } from "vue"

const props = defineProps<{
  page: number
  size: number
  total: number
  fetchData: () => Promise<void>
}>()

const emit = defineEmits<{
  (e: "update:page", value: number): void
  (e: "update:size", value: number): void
}>()

// 页码改变
const onChangePage = async (event: { page: number }) => {
  emit("update:page", event.page) // ✅ 双向绑定 page
  await props.fetchData()
}

// 每页数量改变
const onChangeSize = async (value: number) => {
  emit("update:size", value) // ✅ 双向绑定 size
  await props.fetchData()
}



</script>

<template>
  <div class="flex flex-row items-center gap-3 justify-end">
    <Paginator
        :rows="size"
        :totalRecords="total"
        :first="(page - 1) * size"
        @page="onChangePage"
        class="w-auto"
    />
    <Select
        :modelValue="size"
        @update:modelValue="onChangeSize"
        :options="[15, 30, 60, 120]"
        class="w-auto md:w-36"
    >
      <template #option="slotProps">
        <div class="flex items-center gap-2">
          <i class="pi pi-align-right"></i>
          <div>{{ slotProps.option }}</div>
        </div>
      </template>
      <template #dropdownicon>
        <i class="pi pi-chevron-down"/>
      </template>
      <template #header>
        <div class="font-medium p-3">每页显示数量</div>
      </template>
    </Select>
  </div>
</template>

<style>
:root {
  --p-paginator-gap: 0.25rem; /* 按钮之间间距 */
  --p-paginator-nav-button-width: 2rem; /* 按钮宽度 */
  --p-paginator-nav-button-height: 2rem; /* 按钮高度 */
  --p-paginator-nav-button-border-radius: 0.375rem; /* 圆角 */
  --p-paginator-transition-duration: 150ms; /* 动画时间 */
  //--p-paginator-nav-button-background: #f9fafb; /* 默认背景 */
  //--p-paginator-nav-button-hover-background: #e5e7eb; /* hover 背景 */

}

.p-select-option {
  height: 2rem;
}

:root {
  --p-select-padding-y: 0.25rem; /* 减小上下 padding */
  --p-select-padding-x: 0.75rem;
  --p-select-sm-font-size: 0.875rem;
  --p-select-sm-padding-y: 0.25rem; /* 小尺寸时的 padding */
}
</style>