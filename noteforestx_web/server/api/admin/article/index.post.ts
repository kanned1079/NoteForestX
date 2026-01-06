import {defineEventHandler, getCookie, getQuery} from 'h3'

export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const query = getQuery(event)

    // 拼接後端請求 URL
    const url = new URL(`${config.public.apiBase}/api/v1/admin/article`)
    Object.entries(query).forEach(([key, value]) => {
        if (value !== undefined) url.searchParams.set(key, value as string)
    })

    // 获取请求体（PUT 通常要传 body）
    const body = await readBody(event)

    const token = getCookie(event, 'access_token')
    if (!token) {
        throw createError({
            statusCode: 401,
            statusMessage: 'Missing access token'
        })
    }

    // 转发请求到后端
    const res = await $fetch(url.toString(), {
        method: 'POST',
        body,
        headers: {
            Authorization: `Bearer ${token}` // ✅ Bearer 形式
        }
    })

    return res
})