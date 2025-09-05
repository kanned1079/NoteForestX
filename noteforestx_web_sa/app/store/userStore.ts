import {defineStore} from "pinia";
import type {User} from "~/types/user";
import {ref} from 'vue'

const useUserStore = defineStore('userStore', () => {
    const isAuthed = ref<boolean>(true)
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
            role: 'USER'
        } as User)
        const token = useCookie('token')
        token.value = null // 或 ''
    }

    const logout = async () => {
        console.log('see u')
        clearUserData()
        await useRouter().replace({path: '/'})
    }

    return {
        user,
        isAuthed,
        clearUserData,
        logout,
    }
}, {
    persist: true,
})

export default useUserStore