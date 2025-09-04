import {defineStore} from "pinia";
import type {User} from "~/types/user";
import {useI18n} from "vue-i18n";
import router from "#app/plugins/router";

const useUserStore = defineStore('userStore', () => {
    const isAuthed = ref<boolean>(true)
    const user = ref<User>({
        // id: 1,
        // email: 'kanned1079@icloud.com',
        // username: 'kanned1079',
        // role: 'admin'
        id: '',
        email: '',
        username: null,
        role: 'USER'
    })

    const clearUserData = () => {
        console.log('清除用户数据')
        isAuthed.value = false
        // Object.assign(user.value, {
        //     id: 0,
        //     email: '',
        //     username: '',
        // } as User)
        const token = useCookie('token')
        token.value = null // 或 ''
    }

    const logout = () => {
        console.log('登出操作')
        clearUserData()
        useRouter().replace({path: '/'})
    }

    return {
        user,
        isAuthed,
        clearUserData,
        logout,
        // messageBody,
        // showMessage
    }
}, {
    persist: true
})

export default useUserStore