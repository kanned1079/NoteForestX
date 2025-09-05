export type UniversalApiResponse<T, Extra = {}> = {
    code: number
    message: string[]
    data: T
} & Extra