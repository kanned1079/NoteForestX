<script setup lang="ts">
// 如果 IllustrationItem 是个组件，import 它
import {ref} from "vue"
import type {IllustrationItem} from '~/types/illustration'
import PageHeader from "~/components/PageHeader.vue";
import {useCommonFetch} from "~/composable/useCommonFetch";

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

const illustrationList = ref<IllustrationItem[]>([])

// const { data, pending, error } = useCommonFetch("/illustration", {
//   requireToken: false,
//   method: "GET"
// })
//
// console.log(data, pending, error)


// const fetchIllustrationList = () => {
//   const { data, pending, error } = useCommonFetch("/illustration", {
//     requireToken: false,
//     method: "GET"
//   })
//
//   console.log(data, pending, error)
// }

const fetchIllustrationList = async () => {
  const data = await $fetch<{
    page: number
    size: number
    total: number
    list: IllustrationItem[]
  }>("http://localhost:8081/api/v1/illustration", {
    method: "GET",

  })
  //
  // console.log(data)

  data.list.forEach((item: IllustrationItem) => illustrationList.value.push(item))
  // console.log(illustrationList.value)
}

fetchIllustrationList()

// console.log(illustrationList.value)

//
// const data = useCommonFetch<{
//   page: number
//   size: number
//   total: number
//   list: IllustrationItem[]
// }>("/illustration")

// illustrationList.value = data.list

onMounted(() => {
  // fetchIllustrationList()

  // console.log(illustrationList.value)
})

</script>

<template>


  <div class="p-4">
    <PageHeader title="插画收藏" subtitle="在这里所有收藏的插画项目将会以它们的创建时间进行排序，此处显示的为预览图，如您需要获取原图请点击目标插画进入二级菜单。部分图片带有Limit属性，需要管理员来赋予您权限才可以进行查看。" />
    <span class="flex-row space-x-3">
      <Button :label="`在 ${illustrationList.length} 中搜索`" icon="pi pi-search" severity="primary" size="small" />
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
            :name="i.name"
            :author="i.author.name"
            :img_url="`http://localhost:8081/api/v1/illustration/file/${i.file_path}?size=original`"
            v-for="i in illustrationList" :key="i.id" ></IllustrationItem>
      </div>
    </div>

  </div>
</template>