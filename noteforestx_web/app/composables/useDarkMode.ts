import { ref, onMounted, onBeforeUnmount } from 'vue'

export function useDarkMode() {
    const isDarkMode = ref<boolean>(false)
    let mediaQuery: MediaQueryList | null = null

    const update = (e?: MediaQueryList | MediaQueryListEvent) => {
        isDarkMode.value = !!e?.matches
    }

    onMounted(() => {
        if (typeof window === 'undefined') return

        mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
        update(mediaQuery)

        mediaQuery.addEventListener('change', update)
    })

    onBeforeUnmount(() => {
        mediaQuery?.removeEventListener('change', update)
    })

    return {
        isDarkMode,
    }
}