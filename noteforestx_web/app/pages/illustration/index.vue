<script setup lang="ts">
// 如果 IllustrationItem 是个组件，import 它
import IllustrationItem from '~/components/IllustrationItem.vue'
import PageHeader from "~/components/PageHeader.vue";

const illustrationTopMenu = [
  {
    label: 'File',
    icon: 'pi pi-file',
    items: [
      {
        label: 'New',
        icon: 'pi pi-plus',
        command: () => {
          toast.add({ severity: 'success', summary: 'Success', detail: 'File created', life: 3000 });
        }
      },
      {
        label: 'Print',
        icon: 'pi pi-print',
        command: () => {
          toast.add({ severity: 'error', summary: 'Error', detail: 'No printer connected', life: 3000 });
        }
      }
    ]
  },
  {
    label: 'Search',
    icon: 'pi pi-search',
    command: () => {
      toast.add({ severity: 'warn', summary: 'Search Results', detail: 'No results found', life: 3000 });
    }
  },
  {
    separator: true
  },
  {
    label: 'Sync',
    icon: 'pi pi-cloud',
    items: [
      {
        label: 'Import',
        icon: 'pi pi-cloud-download',
        command: () => {
          toast.add({ severity: 'info', summary: 'Downloads', detail: 'Downloaded from cloud', life: 3000 });
        }
      },
      {
        label: 'Export',
        icon: 'pi pi-cloud-upload',
        command: () => {
          toast.add({ severity: 'info', summary: 'Shared', detail: 'Exported to cloud', life: 3000 });
        }
      }
    ]
  }
]
</script>

<template>


  <div class="p-4">
    <PageHeader title="插画收藏" subtitle="在这里所有收藏的插画项目将会以它们的创建时间进行排序，此处显示的为预览图，如您需要获取原图请点击目标插画进入二级菜单。" />
    <span class="flex-row space-x-3">
      <Button label="在 1532 中搜索" icon="pi pi-search" severity="primary" size="small" />
      <Button label="重置搜索" icon="pi pi-refresh" severity="primary" size="small" variant="outlined" />
    </span>
<!--    <div class="card">-->
<!--      <Menubar size="small" :model="illustrationTopMenu" />-->
<!--      <Toast />-->
<!--    </div>-->

    <!-- 容器：间距 + 最大宽度（可选） -->
    <div class="mx-auto w-full max-w-[1920px] mt-5">
      <!-- grid：手机 2 列，sm:3 md:4 lg:5 xl:6 -->
      <div class="grid gap-5 grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6">
        <!-- 假设有 12 张示例 -->
        <IllustrationItem
            name="Sample title"
            :author="`author 1`"
            :img_url="`https://ikanned.com:2444/d/upload/NXC/134529242_p0.png`"
            v-for="n in 12" :key="n" ></IllustrationItem>
      </div>
    </div>

  </div>
</template>