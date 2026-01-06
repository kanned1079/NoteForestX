import {defineEventHandler, getCookie, getQuery} from 'h3'

export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const query = getQuery(event)

    // const id = event.context.params?.id // 从路径参数中获取 id

    // if (!id) {
    //     throw createError({
    //         statusCode: 400,
    //         statusMessage: 'Missing article id in URL'
    //     })
    // }


    const token = getCookie(event, 'access_token')
    if (!token) {
        throw createError({
            statusCode: 401,
            statusMessage: 'Missing access token'
        })
    }

    // 构建后端请求 URL
    const url = new URL(`${config.public.apiBase}/api/v1/admin/tag`)
    Object.entries(query).forEach(([key, value]) => {
        if (value !== undefined) url.searchParams.set(key, value as string)
    })

    // 使用 fetch 发送请求
    const res = await $fetch(url.toString(), {
        method: 'GET',
        headers: {
            Authorization: `Bearer ${token}` // ✅ Bearer 形式
        }
    })

    return res
})