import {defineStore} from "pinia";
import {ref} from "vue"
import type {User} from "~/types/user";

export type SearchDialog = {
    show: boolean;
    type: ""
}

const useThemeStore = defineStore('themeStore', () => {
    const showHeaderSearchBtn = ref<boolean>(false)
    const searchDialog = ref<SearchDialog>({
        show: false,
        type: ""
    })

    const workTab = ref<'0' | '1'>('0')
    const setWorkTab = (value: '0' | '1') => workTab.value = value

    const showLeftIntroBtn = ref<boolean>(false)
    const setShowLeftIntroBtn = (val: boolean) => showLeftIntroBtn.value = val

    const showEditMetaBtn = ref<boolean>(false)
    const setShowEditMetaBtn = (val: boolean) => showEditMetaBtn.value = val
    const showEditMetaDialog = ref<boolean>(false)
    const setShowEditMetaDialog = (val: boolean) => showEditMetaDialog.value = val

    const showCatalog = ref<boolean>(false)
    const setShowCatalog = (val: boolean) => showCatalog.value = val

    const showCommitIllustrationBtn = ref<boolean>(false)
    const setShowCommitIllustrationBtn = (val: boolean) => showCommitIllustrationBtn.value = val

    const actionCenterMsgs = ref<string[]>([])

    return {
        showHeaderSearchBtn,
        searchDialog,
        workTab,
        setWorkTab,
        showLeftIntroBtn,
        setShowLeftIntroBtn,
        showEditMetaBtn,
        setShowEditMetaBtn,
        showEditMetaDialog,
        setShowEditMetaDialog,
        showCatalog,
        setShowCatalog,
        actionCenterMsgs,
        showCommitIllustrationBtn,
        setShowCommitIllustrationBtn
    }

}, {
    // persist: true
})

export default useThemeStore