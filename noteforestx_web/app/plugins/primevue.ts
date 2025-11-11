import { defineNuxtPlugin } from '#app'
import PrimeVue from 'primevue/config'
import ConfirmationService from 'primevue/confirmationservice'
import ToastService from 'primevue/toastservice'
import AnimateOnScroll from 'primevue/animateonscroll'

export default defineNuxtPlugin((nuxtApp) => {
    nuxtApp.vueApp.use(PrimeVue)
    nuxtApp.vueApp.use(ConfirmationService)
    nuxtApp.vueApp.use(ToastService)


    nuxtApp.vueApp.directive('animateonscroll', AnimateOnScroll)

    // // 注册组件用法
    // nuxtApp.vueApp.component('AnimateOnScroll', AnimateOnScroll)
    //
    // // 注册指令用法
    // nuxtApp.vueApp.directive('v-animateonscroll', AnimateOnScroll)
})