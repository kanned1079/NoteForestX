import {defineStore} from "pinia";
import {ref} from "vue"
import type {User} from "~/types/user";
import {useI18n} from "vue-i18n";
import router from "#app/plugins/router";

const useUserStore = defineStore('userStore', () => {
    const isAuthed = ref<boolean>(false)
    const user = ref<User>({
        id: '',
        email: '',
        username: null,
        role: 'USER'
    })

    const clearUserData = () => {
        console.log('清除用户数据')
        isAuthed.value = false
        Object.assign(user.value, {
            id: '',
            email: '',
            username: '',
            role: "USER"
        } as User)
        const token = useCookie('token')
        token.value = null // 或 ''
    }

    const logout = (): boolean => {
        console.log('登出操作')
        const token = useCookie('access_token', {})
        token.value = null
        clearUserData()
        // useRouter().replace({path: '/'})
        return true
    }

    const languageInUsing = ref<string>('')

    return {
        user,
        isAuthed,
        clearUserData,
        logout,
        languageInUsing
    }
}, {
    persist: true
})

export default useUserStore