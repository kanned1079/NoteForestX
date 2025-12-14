import { defineNuxtPlugin } from '#app'
import ScrollFade from '~/directives/v-scroll-fade'

export default defineNuxtPlugin((nuxtApp) => {
    nuxtApp.vueApp.directive('scroll-fade', ScrollFade)
})