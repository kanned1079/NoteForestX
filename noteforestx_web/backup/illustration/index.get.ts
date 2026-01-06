import { defineEventHandler, getQuery } from 'h3'

export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const query = getQuery(event)

    // 构建后端请求 URL
    const url = new URL(`${config.public.apiBase}/api/v1/illustration`)
    Object.entries(query).forEach(([key, value]) => {
        if (value !== undefined) url.searchParams.set(key, value as string)
    })

    // 使用 fetch 发送请求
    const res = await $fetch(url.toString(), {
        method: 'GET'
    })

    return res
})