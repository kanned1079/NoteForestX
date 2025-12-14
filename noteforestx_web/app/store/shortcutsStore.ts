import {defineStore} from "pinia";
import {ref} from "vue"
import type {KeyShortcutItem} from '~/types/shortcut'

const useShortcutsStore = defineStore('shortcutsStore', () => {

    const customShortcuts = ref<KeyShortcutItem[]>([])

    const clear = () => {
        customShortcuts.value = []
    }

    const register = (shortcuts: KeyShortcutItem[]) => {
        customShortcuts.value = shortcuts
    }

    return {
        customShortcuts,
        clear,
        register
    }

}, {
    // persist: true
})

export default useShortcutsStore