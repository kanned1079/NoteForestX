// https://nuxt.com/docs/api/configuration/nuxt-config
import vuetify, {transformAssetUrls} from 'vite-plugin-vuetify'

export default defineNuxtConfig({
    build: {
        transpile: ['vuetify'],
    },
    compatibilityDate: '2025-07-15',
    devtools: {enabled: true},
    modules: [
        (_options, nuxt) => {
            nuxt.hooks.hook('vite:extendConfig', (config) => {
                // @ts-expect-error
                config.plugins.push(vuetify({autoImport: true}))
            })
        },
        '@nuxt/icon',
        '@pinia/nuxt',
        'pinia-plugin-persistedstate/nuxt',
        '@nuxtjs/tailwindcss',
        '@nuxtjs/i18n',
    ],
    vite: {
        vue: {
            template: {
                transformAssetUrls,
            },
        },
    },
    i18n: {
        locales: [
            {code: 'us', language: 'en-US', file: 'us.json'},
            {code: 'cn', language: 'zh-CN', file: 'cn.json'}
        ],
        strategy: 'prefix',
        defaultLocale: "en"
    },
    css: ['~/assets/css/tailwind.css']
})