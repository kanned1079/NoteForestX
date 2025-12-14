import {defineStore} from "pinia";
import {ref} from "vue"
import type {KeyShortcutItem} from '~/types/shortcut'

const useShortcutsStore = defineStore('shortcutsStore', () => {

    const customShortcuts = ref<KeyShortcutItem[]>([
        // 通用基础操作
        {
            label: '复制',
            keyLabels: ['Ctrl + C', 'Cmd + C'],
            pressType: 'or' // 不同系统二选一
        },
        {
            label: '粘贴',
            keyLabels: ['Ctrl + V', 'Cmd + V'],
            pressType: 'or'
        },
        {
            label: '剪切',
            keyLabels: ['Ctrl + X', 'Cmd + X'],
            pressType: 'or'
        },
        {
            label: '撤销',
            keyLabels: ['Ctrl + Z', 'Cmd + Z'],
            pressType: 'or'
        },
        {
            label: '重做',
            keyLabels: ['Ctrl + Y', 'Cmd + Shift + Z'],
            pressType: 'or'
        },


        {
            label: '关闭当前标签页',
            keyLabels: ['Ctrl + W', 'Cmd + W'],
            pressType: 'or'
        },
        {
            label: '刷新页面',
            keyLabels: ['F5', 'Ctrl + R', 'Cmd + R'],
            pressType: 'or' // 多种方式可选
        },
    ])

    const clear = () => {
        if (customShortcuts.value.length > 0) {
            customShortcuts.value = []
        }
    }

    const register = (shortcuts: KeyShortcutItem[]) => {
        if (shortcuts.length === 0) {
            customShortcuts.value = shortcuts
        }
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