import {defineStore} from "pinia";
import {useTheme} from "vuetify";

const useThemeStore = defineStore("themeStore", () => {
    const isMenuDisplay = ref<boolean>(false)
    const isDarkModeEnabled = ref<boolean>(false)
    const lang = ref<string>('cn')

    const toggleMenuDisplay = () => isMenuDisplay.value = !isMenuDisplay.value
    const toggleDarkMode = () => {
        isDarkModeEnabled.value = !isDarkModeEnabled.value
    }

    const messageBody = ref({
        show: false,
        text: '',
        type: 'success',
        timeout: 16000
    })
    const showMessage = (text: string, type: 'success' | 'warning' | 'info' | 'error' | string = 'info', timeout?: number) => {
        messageBody.value.text = text
        messageBody.value.type = type
        messageBody.value.show = true
        messageBody.value.timeout = timeout | 1600
        setTimeout(() => messageBody.value.show = false, messageBody.value.timeout)
    }

    return {
        isMenuDisplay,
        isDarkModeEnabled,
        toggleMenuDisplay,
        toggleDarkMode,
        lang,
        messageBody,
        showMessage

    }
})

export default useThemeStore