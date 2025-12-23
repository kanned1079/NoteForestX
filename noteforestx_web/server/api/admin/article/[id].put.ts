// /server/api/v1/admin/illustration_tag/[id].put.ts
import { defineEventHandler, getQuery } from 'h3'

export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const query = getQuery(event)
    const id = event.context.params?.id // 从路径参数中获取 id

    if (!id) {
        throw createError({
            statusCode: 400,
            statusMessage: 'Missing id in URL'
        })
    }

    // 拼接後端請求 URL
    const url = new URL(`${config.public.apiBase}/api/v1/admin/article/${id}`)
    Object.entries(query).forEach(([key, value]) => {
        if (value !== undefined) url.searchParams.set(key, value as string)
    })

    // 获取请求体（PUT 通常要传 body）
    const body = await readBody(event)

    // 转发请求到后端
    const res = await $fetch(url.toString(), {
        method: 'PUT',
        body,
    })

    return res
})