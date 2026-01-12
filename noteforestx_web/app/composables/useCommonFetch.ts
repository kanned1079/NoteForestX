import type { FetchOptions } from 'ofetch'

// 定义基础配置接口
interface CustomOptions extends Omit<FetchOptions, 'method'> {
    includeToken?: boolean
}

type HttpMethods = "GET" | "POST" | "PUT" | "PATCH" | "DELETE" | "HEAD" | "CONNECT" | "OPTIONS" | "TRACE" | "get" | "post" | "put" | "patch" | "delete" | "head" | "connect" | "options" | "trace" | undefined

// 定义不带 Body 的请求配置（GET, DELETE）
type RequestConfigNoBody = CustomOptions

// 定义带 Body 的请求配置（POST, PUT, PATCH）
type RequestConfigWithBody = CustomOptions

// 1. 直接从 FetchOptions 提取 method 类型，确保与 $fetch 完全一致
type HttpMethod = FetchOptions['method']

// 2. 扩展配置接口：继承自 FetchOptions 但排除 method
// 显式添加 includeToken，并约束 body 满足 ofetch 要求
interface RequestConfig<B extends FetchOptions['body'] = any> extends Omit<FetchOptions, 'method' | 'body'> {
    includeToken?: boolean
    body?: B
}

// export const useHttp = () => {
//     const config = useRuntimeConfig()
//     const token = useCookie<string | null>('access_token')
//
//     //内部核心请求方法
//     const request = async <T = any>(
//         url: string,
//         method: FetchOptions['method'],
//         options: FetchOptions = {},
//         includeToken = true
//     ): Promise<T> => {
//
//         // 处理 Headers
//         const headers: Record<string, string> = {
//             ...(options.headers as Record<string, string>)
//         }
//
//         if (includeToken && token.value) {
//             headers['Authorization'] = `Bearer ${token.value}`
//         }
//
//         return $fetch<T>(url, {
//             baseURL: config.public.apiBase,
//             ...options,
//             method: method as HttpMethods,
//             headers,
//         })
//     }
//
//     // 导出 Axios 风格的 API
//     return {
//         get: <T = any>(url: string, config?: RequestConfigNoBody) =>
//             request<T>(url, 'GET', config),
//
//         post: <T = any, B extends FetchOptions['body'] = any>(url: string, body?: B, config?: RequestConfigWithBody) =>
//             request<T>(url, 'POST', { ...config, body }),
//
//         put: <T = any, B extends FetchOptions['body'] = any>(url: string, body?: B, config?: RequestConfigWithBody) =>
//             request<T>(url, 'PUT', { ...config, body }),
//
//         patch: <T = any, B extends FetchOptions['body'] = any>(url: string, body?: B, config?: RequestConfigWithBody) =>
//             request<T>(url, 'PATCH', { ...config, body }),
//
//         delete: <T = any>(url: string, config?: RequestConfigNoBody) =>
//             request<T>(url, 'DELETE', config),
//     }
// }

export const useHttp = () => {
    const config = useRuntimeConfig()
    const token = useCookie<string | null>('access_token')

    /**
     * 内部核心请求方法
     */
    const request = async <T = any, B extends FetchOptions['body'] = any>(
        url: string,
        method: HttpMethod,
        options: RequestConfig<B> = {}
    ): Promise<T> => {
        // 解构出我们的自定义配置 includeToken，默认为 true
        // 剩余的 fetchOptions 直接透传给 $fetch
        const { includeToken = true, ...fetchOptions } = options

        // 处理 Headers：使用 Record 避免索引类型报错
        const headers: Record<string, string> = {
            ...(fetchOptions.headers as Record<string, string>)
        }

        // 注入 Token 逻辑
        if (includeToken && token.value) {
            headers['Authorization'] = `Bearer ${token.value}`
        }

        return $fetch<T>(url, {
            // 注意顺序：baseURL 和 fetchOptions 先展开，method 和 headers 后展开以确保优先级
            baseURL: config.public.apiBase,
            ...fetchOptions,
            method: method as HttpMethods,
            headers,
        })
    }

    // 3. 导出 Axios 风格的 API
    return {
        get: <T = any>(url: string, config?: RequestConfig) =>
            request<T>(url, 'GET', config),

        post: <T = any, B extends FetchOptions['body'] = any>(url: string, body?: B, config?: RequestConfig<B>) =>
            request<T, B>(url, 'POST', { ...config, body }),

        put: <T = any, B extends FetchOptions['body'] = any>(url: string, body?: B, config?: RequestConfig<B>) =>
            request<T, B>(url, 'PUT', { ...config, body }),

        patch: <T = any, B extends FetchOptions['body'] = any>(url: string, body?: B, config?: RequestConfig<B>) =>
            request<T, B>(url, 'PATCH', { ...config, body }),

        delete: <T = any>(url: string, config?: RequestConfig) =>
            request<T>(url, 'DELETE', config),
    }
}