// middleware/auth.global.ts
import useUserStore from '~/store/userStore'
// import useThemeStore from "~/store/themeStore";

export default defineNuxtRouteMiddleware((to) => {
    const userStore = useUserStore()
    // const themeStore = useThemeStore()

    const isAuthed = userStore.isAuthed && userStore.user.id
    const isAdmin = userStore.user?.role === 'ADMIN'

    const path = to.path

    if (path.endsWith('/profile') && !isAuthed) {
        console.log('you are not authed.')
        // themeStore.showMessage('PLEASE LOGIN FIRST', 'warning')

        // return navigateTo(`/${to.path.split('/')[1]}/`)
        return navigateTo(`/`)
    }

    if (path.includes('/admin') && (!isAuthed || !isAdmin)) {
        console.log('sorry, you have no admin privilege.')
        return navigateTo(`/`)
    }
})