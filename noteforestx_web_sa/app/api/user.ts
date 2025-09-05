import {useCommonFetch} from "~/composables/useCommonFetch";
import type {User} from "~/types/user";
import {da} from "vuetify/locale";
import useUserStore from "~/store/userStore";
import type {UniversalApiResponse} from "~/types/responseDto";
import userStore from "~/store/userStore";

// type UniversalApiResponse<T> = {
//     code: number,
//     message: string[],
//     user: T
// }


export const processUserAuth = async (email: string, password: string, type: 'login' | 'register'): Promise<{
    code: number
}> => {
    // const userStore = useUserStore()
    // const {code, data, error} = await useCommonFetch<User>(`/api/v1/user/${type}`, {
    //     method: 'POST',
    //     auth: false,
    //     body: {
    //         email, password
    //     }
    // })
    // if (!error && code === 200 && data) {
    //     Object.assign(userStore.user, data)
    //     userStore.isAuthed = true
    //     const token = useCookie('token')
    //     token.value = data.token
    //     return {code: 200}
    // } else {
    //     return {code: code}
    // }

    try {
        const userStore = useUserStore()
        const data = await $fetch<{
            user: User,
            token: string
        }>(`/api/v2/user/${type}`, {
            method: "POST",
            body: {
                email, password
            }
        })
        Object.assign(userStore.user, data.user)
        userStore.isAuthed = true
        const token = useCookie('token')
        token.value = data.token
        return {code: 200}
    } catch (err: any) {
        if (err?.statusCode) {
            return {code: err?.statusCode}
        }
        return {code: 500}
    }
}

type UpdateUsernameRespData = {
    username: string
}

export const updateUsername = async (newUsername: string): Promise<{ code: number }> => {
    const userStore = useUserStore()
    // const {
    //     code,
    //     data,
    //     error
    // } = await useCommonFetch<UpdateUsernameRespData>(`/api/user/username/update/${userStore.user.id}`, {
    //     method: 'PATCH',
    //     auth: true,
    //     body: {
    //         username: newUsername
    //     }
    // })
    // if (!error && code === 200 && data) {
    //     userStore.user.username = data.username
    //     return {code: 200}
    // } else {
    //     return {code: code}
    // }

    try {
        const token = useCookie('token')
        const data = await $fetch<{
            id: string,
            username: string
        }>(`/api/v2/user/username/update/${userStore.user.id}`, {
            method: 'PATCH',
            headers: {
                Authorization: `Bearer ${token.value}`
            },
            body: {
                id: userStore.user.id,
                username: newUsername
            }
        })

        userStore.user.username = data.username
        return {code: 200}
    } catch (err: any) {
        if (err?.statusCode) {
            return {code: err?.statusCode}
        }
        return {code: 500}
    }
}

export const updateUserPasswordById = async (previousPwd: string, newPwd: string): Promise<{ code: number }> => {
    // console.log('11111')
    try {
        const userStore = useUserStore()
        const token = useCookie('token')
        const data = await $fetch<{
            message: string
        }>(`/api/v2/user/password/update/${userStore.user.id}`, {
            method: 'PATCH',
            headers: {
                Authorization: `Bearer ${token.value}`
            },
            body: {
                previousPassword: previousPwd,
                newPassword: newPwd
            }
        })
        return data ? {code: 200} : {code: 500}
    } catch (err: any) {
        if (err?.statusCode) {
            return {code: err?.statusCode}
        }
        return {code: 500}
    }
    // const {
    //     code,
    //     data,
    //     error
    // } = await useCommonFetch<UpdateUsernameRespData>(`/api/v1/user/password/update/${userStore.user.id}`, {
    //     method: 'PATCH',
    //     auth: true,
    //     body: {
    //         previousPassword: previousPwd,
    //         newPassword: newPwd
    //     }
    // })
    // if (!error && code === 200 && data) {
    //     return {code: 200}
    // } else {
    //     return {code: code}
    // }
}

// export const newUserRegister = async (email: string, password: string): {code: number} => {
//
// }