import type { Illustration } from "../../types/illustration"

export default defineEventHandler(async (event) => {
    const query = getQuery(event)
    const page = query.page || 1
    const size = query.size || 15
    const show_limited = query.show_limited || false

    // 转发请求到后端 Gin
    const data = await $fetch<{
        page: number
        size: number
        total: number
        list: Illustration[]
    }>(`http://localhost:8081/api/v1/illustration`, {
        method: "GET",
        params: { page, size, show_limited }
    })

    return data
})