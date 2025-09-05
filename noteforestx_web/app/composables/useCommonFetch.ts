import useUserStore from '~/store/userStore'

export async function useCommonFetch<T>(
    url: string,
    options: {
        auth?: boolean
        headers?: Record<string, string>
        [key: string]: any
    } = {}
): Promise<{ data: T | null; error: string | null; code: number }> {
    const userStore = useUserStore()
    const config = useRuntimeConfig()
    const token = useCookie('token')

    const headers: Record<string, string> = {
        ...(options.headers || {})
    }

    if (options.auth && token.value) {
        headers.Authorization = `Bearer ${token.value}`
    }

    try {
        const raw = await $fetch<{
            code: number
            message: string | string[]
            data: T
        }>(url, {
            baseURL: config.public.apiBase as string,
            credentials: 'include',
            headers,
            // ❌ 删除 throw: false（因为不是合法字段）
            ...options
        })

        console.log('raw: ', raw)

        if (raw.code === 401) {
            userStore.logout()
            return { code: 401, data: null, error: '未授权，请重新登录' }
        }

        if (raw.code !== 200) {
            const msg = Array.isArray(raw.message) ? raw.message.join(', ') : raw.message
            return { code: raw.code, data: null, error: msg || '请求失败' }
        }

        return {
            code: 200,
            data: raw.data,
            error: null
        }

    } catch (err: any) {
        console.error('Fetch error:', err)

        const status = err?.response?.status || err?.status

        if (status === 401) {
            userStore.logout()
            return { code: 401, data: null, error: '未授权，请重新登录' }
        }

        if (status === 400) {
            return { code: 400, data: null, error: '数据格式不正确' }
        }

        if (status === 404) {
            return { code: 404, data: null, error: '数据不存在' }
        }

        return { code: 500, data: null, error: '网络错误或服务器异常' }
    }
}