// https://nuxt.com/docs/api/configuration/nuxt-config
import Aura from '@primeuix/themes/aura';
// import {MyPreset} from "./app/plugins/MyPresetTheme";
import { definePreset } from '@primeuix/themes'

const MyPreset = definePreset(Aura, {
    semantic: {
        focusRing: {
            width: '2px',
            style: 'dashed',
            color: '{primary.color}',
            offset: '1px'
        },
        primary: {
            50: '#e5eff7',
            100: '#c6d9eb',
            200: '#a7c2df',
            300: '#88add3',
            400: '#6998c7',
            500: '#4f7da6',
            600: '#3f657f',
            700: '#2f4e5f',
            800: '#1f3640',
            900: '#0f1b20',
            950: '#080f10'
        },
        colorScheme: {
            light: {
                primary: {
                    color: '#4f7da6',
                    inverseColor: '#ffffff',
                    hoverColor: '#3f657f',
                    activeColor: '#2f4e5f'
                },
                highlight: {
                    background: '#4f7da6',
                    focusBackground: '#3f657f',
                    color: '#ffffff',
                    focusColor: '#ffffff'
                },
                formField: {
                    hoverBorderColor: '#4f7da6'
                }
            },
            dark: {
                primary: {
                    color: '#4f7da6',
                    inverseColor: '#ffffff',
                    hoverColor: '#6998c7',
                    activeColor: '#88add3'
                },
                highlight: {
                    background: 'rgba(79,125,166,0.16)',
                    focusBackground: 'rgba(79,125,166,0.24)',
                    color: '#ffffff',
                    focusColor: '#ffffff'
                },
                formField: {
                    hoverBorderColor: '#4f7da6'
                }
            }
        }
    }
})

export default defineNuxtConfig({
    compatibilityDate: '2025-07-15',
    devtools: {enabled: true},
    modules: [
        '@pinia/nuxt',
        'pinia-plugin-persistedstate/nuxt',
        '@primevue/nuxt-module',
        '@nuxt/icon',
        '@nuxtjs/i18n',
        '@nuxtjs/tailwindcss'

    ],
    primevue: {
        options: {
            theme: {
                preset: MyPreset,
                // cssLayer: false
            }
        }
    },
    css: [
        '~/assets/css/main.css',
        '~/assets/css/theme.css',
        'primeicons/primeicons.css',
        'flag-icons/css/flag-icons.min.css'
    ],
    i18n: {
        locales: [
            { code: 'en_us', language: 'en-US', file: 'en_US.json' },
            { code: 'zh_cn', language: 'zh-CN', file: 'zh_CN.json' },
            { code: 'zh_hk', language: 'zh-HK', file: 'zh_HK.json' },
            { code: 'ja_jp', language: 'ja-JP', file: 'ja_JP.json' },
            { code: 'es_es', language: 'es-ES', file: 'es_ES.json' },
            { code: 'fr_fr', language: 'fr-FR', file: 'fr_FR.json' },
            { code: 'fi_fi', language: 'fi-FI', file: 'fi_FI.json' },
            { code: 'de_de', language: 'de-DE', file: 'de_DE.json' },
            { code: 'nl_nl', language: 'nl-NL', file: 'nl_NL.json' }
        ],
        strategy: 'prefix',
        defaultLocale: "zh_cn"
    },
    runtimeConfig: {
        public: {
            apiBase: process.env.API_BASE_URL || "http://localhost:8081/api/v1",
        }
    }
})