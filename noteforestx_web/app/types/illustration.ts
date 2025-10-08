export type IllustrationImage = {
    id: string
    illustration_id: string
    file_path: string
    order: number
    width: number
    height: number
    created_at: string | null
    updated_at: string | null
    deleted_at: string | null
}

export type Illustration = {
    id: string
    pixiv_id: string
    name: string
    author_id: string
    author: IllustrationAuthor
    link: string
    source: string
    description: string
    tags: IllustrationTag[]
    images: IllustrationImage[]
    limited: boolean
    created_at: string | null
    updated_at: string | null
    deleted_at: string | null
}

export type IllustrationAuthor = {
    id: string
    name: string
    link: string
    related?: number
    created_at: string | null
    updated_at: string | null
    deleted_at: string | null
}

export type IllustrationTag = {
    id: string
    name: string
    related?: number
    created_at: string | null
    updated_at: string | null
    deleted_at: string | null
}

export type IllustrationTagMapping = {
    illustration_id: string
    tag_id: string
    created_at: string | null
    updated_at: string | null
    deleted_at: string | null
}