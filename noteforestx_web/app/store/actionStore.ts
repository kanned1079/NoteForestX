import {defineStore} from "pinia";
import {ref} from "vue"
import type {User} from "~/types/user";


const useActionStore = defineStore('actionStore', () => {
    const triggerArticleSave = ref<boolean>(false)
    const triggerArticleSavePayload = ref<any>(null)

    // 触发函数
    const fireTriggerArticleSave = (data?: any) => {
        triggerArticleSavePayload.value = data
        triggerArticleSave.value = true
    }

    // 重置状态
    const resetTriggerArticleSave = () => {
        triggerArticleSave.value = false
        triggerArticleSavePayload.value = null
    }

    const triggerShowCatalog = ref<boolean>(false)
    const triggerShowCatalogPayload = ref<any>(null)

    const fireTriggerShowCatalog = (data?: any) => {
        triggerShowCatalogPayload.value = data
        triggerShowCatalog.value = true
    }

    const resetTriggerShowCatalog = () => {
        triggerShowCatalog.value = false
        triggerShowCatalogPayload.value = null
    }

    const triggerSearchArticle = ref<boolean>(false)

    const fireTriggerSearchArticle = (data?: any) => {
        triggerSearchArticle.value = true
    }

    const resetTriggerSearchArticle = () => {
        triggerSearchArticle.value = false
    }


    return {
        triggerArticleSave,
        fireTriggerArticleSave,
        resetTriggerArticleSave,
        triggerShowCatalog,
        fireTriggerShowCatalog,
        triggerShowCatalogPayload,
        resetTriggerShowCatalog,
        triggerSearchArticle,
        fireTriggerSearchArticle,
        resetTriggerSearchArticle
    }

}, {
    // persist: true
})

export default useActionStore