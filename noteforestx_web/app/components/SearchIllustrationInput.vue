<template>
  <IconField class="w-auto">
    <InputIcon :class="currentIcon" />
    <InputText
        class="w-full font-mono justify-center"
        size="large"
        v-model="searchIll.content"
        :placeholder="command"
        @input="handleInput"
        autofocus
        :invalid="!isValid"
        @keyup.enter="handleEnter"
    />

    <InputIcon v-if="searchLimited" class="pi pi-sparkles" />
  </IconField>
</template>

<script setup lang="ts">
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import { ref, computed } from 'vue'

const searchIll = ref({ content: '' })

// 占位符提示
const command : string = `/{tag|name|user|id} [-limited]`

// 根据前缀动态返回图标
const currentIcon = computed(() => {
  const val = searchIll.value.content.trim()
  if (val.startsWith('/user')) return 'pi pi-user'
  if (val.startsWith('/tag')) return 'pi pi-hashtag'
  if (val.startsWith('/name')) return 'pi pi-palette'
  if (val.startsWith('/id')) return 'pi pi-file'
  return 'pi pi-search'
})

// 解析结果
const searchAs = ref<string>("")        // 前缀命令
const searchLimited = ref<boolean>(false) // 是否加了 -limited
const searchVal = ref<string[]>([])     // 搜索值
const isValid = ref<boolean>(true)     // 输入是否合法

function handleInput() {
  const raw = searchIll.value.content.trim()

  if (raw.length === 0) {
    // 空输入不算错误
    searchAs.value = ""
    searchLimited.value = false
    searchVal.value = []
    isValid.value = true
    return
  }

  if (!raw.startsWith('/')) {
    searchAs.value = 'tag'
    searchLimited.value = false
    searchVal.value = raw.split(/\s+/)
    isValid.value = searchVal.value.length > 0
    return
  }

  const parts = raw.split(/\s+/).filter(Boolean)
  const prefix = parts[0]?.slice(1)
  const allowed = ['tag','tags', 'user', 'name', 'id']

  if (!allowed.includes(prefix)) {
    isValid.value = false
    return
  }

  searchAs.value = prefix
  searchLimited.value = parts.includes('-limited')
  const args = parts.slice(1).filter(p => p !== '-limited')

  if (prefix === 'tag' || prefix === 'tags') {
    searchVal.value = args
    isValid.value = args.length > 0
  } else {
    if (args.length === 1) {
      searchVal.value = args
      isValid.value = true
    } else {
      searchVal.value = []
      isValid.value = false
    }
  }
}

const handleEnter = () => {
  console.log("--------------------------------")
  if (isValid && searchIll.value) {
    console.log(searchAs.value)
    console.log(searchVal.value)
    console.log(searchLimited.value)
  }
}

onMounted(() => {
  console.log('MOUNT')
  console.log(isValid.value)
})

</script>