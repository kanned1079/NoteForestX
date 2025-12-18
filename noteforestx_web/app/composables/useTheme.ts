// composables/useTheme.ts
import { ref, watch, onMounted } from 'vue'

type ThemeMode = 'system' | 'light' | 'dark'

const STORAGE_KEY = 'theme-mode'

export function useTheme() {
    const mode = ref<ThemeMode>('system')

    const applyTheme = () => {
        const root = document.documentElement
        root.classList.remove('dark')

        if (mode.value === 'dark') {
            root.classList.add('dark')
        } else if (mode.value === 'system') {
            const isDark = window.matchMedia('(prefers-color-scheme: dark)').matches
            if (isDark) root.classList.add('dark')
        }
    }

    // 系统变化时同步（只在 system 模式下）
    const handleSystemChange = (e: MediaQueryListEvent) => {
        if (mode.value === 'system') {
            document.documentElement.classList.toggle('dark', e.matches)
        }
    }

    onMounted(() => {
        const saved = localStorage.getItem(STORAGE_KEY) as ThemeMode | null
        if (saved) mode.value = saved

        applyTheme()

        window
            .matchMedia('(prefers-color-scheme: dark)')
            .addEventListener('change', handleSystemChange)
    })

    watch(mode, (val) => {
        localStorage.setItem(STORAGE_KEY, val)
        applyTheme()
    })

    return {
        mode,
        setLight: () => (mode.value = 'light'),
        setDark: () => (mode.value = 'dark'),
        setSystem: () => (mode.value = 'system'),
        toggle: () => {
            mode.value = mode.value === 'dark' ? 'light' : 'dark'
        },
    }
}