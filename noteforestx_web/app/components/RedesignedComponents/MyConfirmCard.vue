<script setup lang="ts">
import { useConfirm } from 'primevue/useconfirm'

const props = defineProps<{
  header: string
  title: string
  subtitle: string
  cancelled: () => void
  confirmed: () => void
  confirmBtnSeverity: '' | 'danger'
}>()

const confirm = useConfirm()

const showConfirm = () => {
  confirm.require({
    group: 'common-confirm',
    header: props.header,
    icon: 'pi pi-exclamation-triangle',

    rejectProps: {
      label: '取消',
      icon: 'pi pi-times',
      outlined: true,
      size: 'small'
    },

    acceptProps: {
      label: '确认',
      icon: 'pi pi-check',
      size: 'small',
      severity: props.confirmBtnSeverity || undefined
    },

    accept: () => {
      props.confirmed()
    },

    reject: () => {
      props.cancelled()
    }
  })
}

// 👇 暴露方法，让父组件调用
defineExpose({
  showConfirm
})
</script>

<template>
  <ConfirmDialog group="common-confirm">
    <template #message="slotProps">
      <div
          class="w-full"
      >
        <div>
          <p class="font-semibold text-base sm:text-lg">
            {{ props.title }}
          </p>
          <p
              class="mt-1 text-sm sm:text-base text-surface-600 dark:text-surface-400">
            {{ props.subtitle }}
          </p>
        </div>
      </div>
    </template>
  </ConfirmDialog>
</template>

<style scoped>
/* 可选：微调对话框视觉 */
</style>