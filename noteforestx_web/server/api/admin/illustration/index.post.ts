import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const fetchUrl = `${config.public.apiBase}/api/v1/admin/illustration`

    // 转发原始请求 body 和 headers
    const headers = { ...event.node.req.headers }
    const method = event.node.req.method || 'POST'

    const body = event.node.req

    const res = await $fetch(fetchUrl, {
        method,
        body,
        headers,
    })

    return res
})