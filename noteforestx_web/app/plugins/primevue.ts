import { defineNuxtPlugin } from '#app'
import PrimeVue from 'primevue/config'
import ConfirmationService from 'primevue/confirmationservice'
import ToastService from 'primevue/toastservice'

export default defineNuxtPlugin((nuxtApp) => {
    // 安裝 PrimeVue 核心與服務
    nuxtApp.vueApp.use(PrimeVue)
    nuxtApp.vueApp.use(ConfirmationService)
    nuxtApp.vueApp.use(ToastService)
})