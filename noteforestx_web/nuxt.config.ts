// https://nuxt.com/docs/api/configuration/nuxt-config
// import {fileURLToPath, URL} from 'node:url'
import Aura from '@primeuix/themes/aura';
// import ToastService from 'primevue/toastservice';
// import {MyPreset} from "./app/plugins/MyPresetTheme";
import {definePreset} from '@primeuix/themes'


const Noir = definePreset(Aura, {
    semantic: {
        primary: {
            50: '#6998c7',
            100: '#6998c7',
            200: '#6998c7',
            300: '#6998c7',
            400: '#6998c7',
            500: '#6998c7',
            600: '#6998c7',
            700: '{zinc.700}',
            800: '{zinc.800}',
            900: '{zinc.900}',
            950: '{zinc.950}'
        },
        colorScheme: {
            light: {
                surface: {
                    0: '#c3c3c3',
                    50: '{zinc.50}',
                    100: '{zinc.100}',
                    200: '{zinc.200}',
                    300: '{zinc.300}',
                    400: '{zinc.400}',
                    500: '{zinc.500}',
                    600: '{zinc.600}',
                    700: '{zinc.700}',
                    800: '{zinc.800}',
                    900: '{zinc.900}',
                    950: '{zinc.950}'
                },
                overlay: {
                    0: '#c3c3c3'
                },
                highlight: {
                    background: '#343434',
                    color: "#453454"
                }

            },
            dark: {
                semantic: {
                    highlight: {
                        background: '{primary.200}',
                        color: '{primary.900}',
                    }
                }
            }
        }
    }
});


const MyPreset = definePreset(Aura, {
    semantic: {
        formField: {
            borderRadius: '8px'
        },
        card: {
            borderRadius: '8px'
        },
        borderRadius: {
            small: '4px',   // 默认小圆角
            medium: '8px',  // 中等圆角
            large: '12px',  // 大圆角
            pill: '9999px'  // pill 类型
        },
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
                surface: {
                    0: '#f9fafc',
                    50: '{zinc.50}',
                    100: '{zinc.100}',
                    200: '{zinc.200}',
                    300: '{zinc.300}',
                    400: '{zinc.400}',
                    500: '{zinc.500}',
                    600: '{zinc.600}',
                    700: '{zinc.700}',
                    800: '{zinc.800}',
                    900: '{zinc.900}',
                    950: '{zinc.950}'
                },
                root: {
                    background: '#f9fafc',
                    color: '#111827'
                },
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
                    hoverBorderColor: '#4f7da6',
                    background: '#ffffff',
                }
            },
            dark: {
                0: '#ffffff',
                50: '{zinc.50}',
                100: '{zinc.100}',
                200: '{zinc.200}',
                300: '{zinc.300}',
                400: '{zinc.400}',
                500: '{zinc.500}',
                600: '{zinc.600}',
                700: '{zinc.700}',
                800: '{zinc.800}',
                900: '{zinc.900}',
                950: '{zinc.950}',
                root: {
                    background: '#1f1f1f',
                    color: '#f9fafc'
                },
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
                    hoverBorderColor: '#4f7da6',
                    background: '#141414',
                },
                card: {
                    background: '#141414'

                }
            }
        },
        // inputVariant: "filled"

    },
    components: {
        card: {
            colorScheme: {
                light: {
                    root: {
                        background: '#ffffff',
                        color: '{surface.700}'
                    },
                },
                dark: {
                    root: {
                        background: '#141414',
                        color: '{surface.0}'
                    },
                    subtitle: {
                        color: '{surface.400}'
                    }
                }
            }
        },

    }
})

// const MyPreset = definePreset(Aura, {
//     semantic: {
//         focusRing: {
//             width: '2px',
//             style: 'dashed',
//             color: '{primary.color}',
//             offset: '1px'
//         },
//         primary: {
//             50: '#e5eff7',
//             100: '#c6d9eb',
//             200: '#a7c2df',
//             300: '#88add3',
//             400: '#6998c7',
//             500: '#4f7da6',
//             600: '#3f657f',
//             700: '#2f4e5f',
//             800: '#1f3640',
//             900: '#0f1b20',
//             950: '#080f10'
//         },
//         colorScheme: {
//             light: {
//                 primary: {
//                     color: '#4f7da6',
//                     inverseColor: '#ffffff',
//                     hoverColor: '#3f657f',
//                     activeColor: '#2f4e5f'
//                 },
//                 highlight: {
//                     background: '#4f7da6',
//                     focusBackground: '#3f657f',
//                     color: '#ffffff',
//                     focusColor: '#ffffff'
//                 },
//                 formField: {
//                     hoverBorderColor: '#4f7da6'
//                 },
//                 surface: {
//                     // 👇这里设置浅色模式背景色
//                     0: '#f9fafc',
//                     50: '#6f7685',   // 页面主要背景色
//                     100: '#f1f3f5',  // 比如卡片背景
//                 }
//             },
//             dark: {
//                 primary: {
//                     color: '#4f7da6',
//                     inverseColor: '#ffffff',
//                     hoverColor: '#6998c7',
//                     activeColor: '#88add3'
//                 },
//                 highlight: {
//                     background: 'rgba(79,125,166,0.16)',
//                     focusBackground: 'rgba(79,125,166,0.24)',
//                     color: '#ffffff',
//                     focusColor: '#ffffff'
//                 },
//                 formField: {
//                     hoverBorderColor: '#4f7da6'
//                 },
//                 surface: {
//                     0: '#181818',
//                     50: '#2f302f',   // 你之前深色模式想要的背景
//                     100: '#3a3b3a'
//                 }
//             }
//         }
//     },
//     inputVariant: "filled"
// })

export default defineNuxtConfig({
    compatibilityDate: '2025-07-15',
    devtools: {enabled: true},
    modules: [
        '@pinia/nuxt',
        'pinia-plugin-persistedstate/nuxt',
        '@primevue/nuxt-module',
        '@nuxt/icon',
        '@nuxtjs/i18n',
        '@nuxtjs/tailwindcss',
        '@primevue/nuxt-module',
        '@vueuse/motion/nuxt'
    ],
    primevue: {
        options: {
            theme: {
                options: {
                    prefix: 'p',
                    // darkModeSelector: '.my-app-dark',
                    darkModeSelector: 'system',
                    cssLayer: false
                },
                preset: MyPreset,
                // cssLayer: false
            }
        }
    },
    css: [
        'primeicons/primeicons.css',
        '~/assets/css/main.css',
        '~/assets/css/theme.css',
        'flag-icons/css/flag-icons.min.css'
    ],
    i18n: {
        locales: [
            {code: 'en_us', language: 'en-US', file: 'en_US.json'},
            {code: 'zh_cn', language: 'zh-CN', file: 'zh_CN.json'},
            // {code: 'zh_hk', language: 'zh-HK', file: 'zh_HK.json'},
            // {code: 'ja_jp', language: 'ja-JP', file: 'ja_JP.json'},
            // {code: 'es_es', language: 'es-ES', file: 'es_ES.json'},
            // {code: 'fr_fr', language: 'fr-FR', file: 'fr_FR.json'},
            // {code: 'fi_fi', language: 'fi-FI', file: 'fi_FI.json'},
            // {code: 'de_de', language: 'de-DE', file: 'de_DE.json'},
            // {code: 'nl_nl', language: 'nl-NL', file: 'nl_NL.json'}
        ],
        strategy: 'prefix',
        defaultLocale: "zh_cn"
    },
    // resolve: {
    //     alias: {
    //         '~': fileURLToPath(new URL('./app', import.meta.url)), // ~ 指向 src
    //         // '@': fileURLToPath(new URL('./src', import.meta.url)), // @ 也可以一起
    //     }
    // },
    runtimeConfig: {
        public: {
            apiBase: process.env.API_BASE_URL || "http://127.0.0.1:8081",
        }
    },
    vite: {
        optimizeDeps: {
            exclude: ['primevue']
        }
    }
    // nitro: {
    //     devProxy: {
    //         "/api/": {
    //             target: "http://localhost:8081/api/",
    //             changeOrigin: true,
    //         }
    //     }
    // },
    // devServer: {
    //     proxy: {
    //         '/api': 'http://localhost:8081'
    //     }
    // }
})