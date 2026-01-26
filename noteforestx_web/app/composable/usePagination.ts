import { ref, watch, onMounted } from 'vue'

// import { ref, watch } from 'vue'
//
// export function usePagination(defaultSize = 15) {
//     const STORAGE_KEY = 'user-pagination-size'
//
//     // 1. 初始化时直接从本地读取
//     const savedSize = localStorage.getItem(STORAGE_KEY)
//     const size = ref(savedSize ? parseInt(savedSize) : defaultSize)
//     const page = ref(1)
//
//     // 2. 监听 size 变化，一旦变化自动保存到本地
//     watch(size, (newSize) => {
//         localStorage.setItem(STORAGE_KEY, newSize.toString())
//     })
//
//     return {
//         page,
//         size
//     }
// }

// composables/usePagination.ts

export function usePagination(defaultSize = 15) {
    const STORAGE_KEY = 'user-pagination-size'
    const size = ref(defaultSize)
    const page = ref(1)

    // 初始化：仅在客户端挂载后读取本地存储
    onMounted(() => {
        if (import.meta.client) {
            const saved = localStorage.getItem(STORAGE_KEY)
            if (saved) {
                size.value = parseInt(saved)
            }
        }
    })

    // 持久化：当 size 变化时写入 localStorage
    watch(size, (newSize) => {
        if (import.meta.client) {
            localStorage.setItem(STORAGE_KEY, newSize.toString())
        }
    })

    return {
        page,
        size
    }
}
