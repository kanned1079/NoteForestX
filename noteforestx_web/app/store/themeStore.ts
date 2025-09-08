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

    return {
        showHeaderSearchBtn,
        searchDialog,
    }

}, {
    // persist: true
})

export default useThemeStore