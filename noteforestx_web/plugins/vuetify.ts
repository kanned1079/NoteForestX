// import this after install `@mdi/font` package
import '@mdi/font/css/materialdesignicons.css'

import { type ThemeDefinition } from "vuetify"

import 'vuetify/styles'
import { createVuetify } from 'vuetify'
// import {grayNight, grayLight} from "~/plugins/themes/gray";
// import {glacierBlueNight, glacierBlueLight} from "~/plugins/themes/glacierBlue";

export default defineNuxtPlugin((app) => {
    const vuetify = createVuetify({
        // ... your configuration
        // theme: {
        //     defaultTheme: 'glacierBlueLight',
        //     themes: {
        //         grayNight,
        //         grayLight,
        //         glacierBlueNight,
        //         glacierBlueLight
        //     }
        // }
        theme: {
            defaultTheme: "glacierBlueLight",
            themes: {
                glacierBlueLight,
                glacierBlueNight
            }
        }
    })
    app.vueApp.use(vuetify)
})


export const glacierBlueLight: ThemeDefinition = {
    dark: false,
    colors: {
        background: '#f5f5f5',
        surface: '#ffffff',
        primary: '#6390b9',
        'primary-darken-1': '#50749a',
        secondary: '#90a4ae',
        'secondary-darken-1': '#708690',
        error: '#d32f2f',
        info: '#1976d2',
        success: '#388e3c',
        warning: '#fbc02d',
    },
}

export const glacierBlueNight: ThemeDefinition = {
    dark: true,
    colors: {
        background: '#000000',
        surface: '#2c2c2c',
        primary: '#6390b9',
        'primary-darken-1': '#4f728f',
        secondary: '#90a4ae',
        'secondary-darken-1': '#708690',
        error: '#cf6679',
        info: '#90caf9',
        success: '#81c784',
        warning: '#ffb74d',
        a: '#393939'
    },
}