import { defineEventHandler, getQuery, getRouterParams } from 'h3'

export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const query = getQuery(event)
    const params = getRouterParams(event) // 获取路径参数
    const illustId = params.id as string

    // 构建 Golang API URL
    const url = new URL(`${config.public.apiBase}/api/v1/illustration/${illustId}?size=medium`)
    Object.entries(query).forEach(([key, value]) => {
        if (value !== undefined) url.searchParams.set(key, value as string)
    })

    // 使用 fetch 请求 Golang 后端
    return await $fetch(url.toString(), { method: 'GET' })
})