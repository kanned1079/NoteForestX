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

    return {
        showHeaderSearchBtn,
        searchDialog,
        workTab,
        setWorkTab

    }

}, {
    // persist: true
})

export default useThemeStore