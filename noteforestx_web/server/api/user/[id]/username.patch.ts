import { defineEventHandler, getQuery, readBody, getCookie } from 'h3'

export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const query = getQuery(event)
    const id = event.context.params?.id

    if (!id) {
        throw createError({
            statusCode: 400,
            statusMessage: 'Missing id in URL'
        })
    }

    // 拼接後端請求 URL
    const url = new URL(`${config.public.apiBase}/api/v1/user/${id}/username`)
    Object.entries(query).forEach(([key, value]) => {
        if (value !== undefined) url.searchParams.set(key, value as string)
    })

    // 获取请求体
    const body = await readBody(event)

    // ✅ 使用 getCookie 获取 token
    const token = getCookie(event, 'access_token')

    // const token = getCookie(event, 'access_token')
    console.log('[username.patch.ts] token:', token)

    if (!token) {
        throw createError({
            statusCode: 401,
            statusMessage: 'Missing access token'
        })
    }

    // 转发请求到后端
    const res = await $fetch(url.toString(), {
        method: 'PATCH',
        body,
        headers: {
            Authorization: `Bearer ${token}` // ✅ Bearer 形式
        }
    })

    return res
})