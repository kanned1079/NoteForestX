// https://nuxt.com/docs/api/configuration/nuxt-config
import vuetify, {transformAssetUrls} from 'vite-plugin-vuetify'

import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
    build: {transpile: ['vuetify'],},
    compatibilityDate: '2025-07-15',
    devtools: {enabled: true},
    modules: [
        '@nuxtjs/i18n',
        (_options, nuxt) => {
            nuxt.hooks.hook('vite:extendConfig', (config) => {
                // @ts-expect-error
                config.plugins.push(vuetify({autoImport: true}))
            })
        }, ``
    ],
    css: ['~/assets/css/main.css'],
    i18n: {
        strategy: 'prefix',
        defaultLocale: 'zh-cn',
        locales: [
            {code: 'zh-cn', name: 'Chinese', file: 'cn.json'},
            {code: 'en-us', name: 'English', file: 'en.json'},
            {code: 'ja-jp', name: 'Japanese', file: 'ja.json'},
        ]
    },
    vite: {
        vue: {template: {transformAssetUrls,},},
        plugins: [tailwindcss(),]
    },
    app: {
        head: {
            title: 'Note Forest',
            link: [
                {
                    rel: 'icon',
                    type: 'image/png',
                    href: '/note_forest_icon.png'
                }
            ]
        },
        layoutTransition: {name: 'fade', mode: 'out-in'}
    },
})