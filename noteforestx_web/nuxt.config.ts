// https://nuxt.com/docs/api/configuration/nuxt-config
import Aura from '@primeuix/themes/aura';

export default defineNuxtConfig({
    compatibilityDate: '2025-07-15',
    devtools: {enabled: true},
    modules: [
        '@primevue/nuxt-module',
        '@nuxt/icon',
        '@pinia/nuxt',
        'pinia-plugin-persistedstate/nuxt',
        '@nuxtjs/i18n',
        '@nuxtjs/tailwindcss'

    ],
    primevue: {
        options: {
            theme: {
                preset: Aura,
                darkModeSelector: 'system',
                cssLayer: false
            }
        }
    },
    css: [
        '~/assets/css/main.css',
        '~/assets/css/theme.css',
        'primeicons/primeicons.css'
    ],
    // postcss: {
    //     plugins: {
    //         tailwindcss: {},
    //         autoprefixer: {},
    //     }
    // },
    i18n: {
        locales: [
            {code: 'en', language: 'en-US', file: 'en.json'},
            {code: 'cn', language: 'zh-CN', file: 'cn.json'}
        ],
        strategy: 'prefix',
        defaultLocale: "en"
    },
})