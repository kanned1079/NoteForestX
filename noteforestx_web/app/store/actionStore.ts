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



    return {
        triggerArticleSave,
        fireTriggerArticleSave,
        resetTriggerArticleSave
    }

}, {
    // persist: true
})

export default useActionStore