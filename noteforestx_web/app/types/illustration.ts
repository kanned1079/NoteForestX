export type IllustrationItem = {
    id: string
    pixiv_id: string
    file_path: string
    name: string
    author_id: string
    author: IllustrationAuthor
    link: string
    tags: IllustrationTag[]
    limited: boolean
    created_at: string | null
    updated_at: string | null
    deleted_at: string | null
}

export type IllustrationAuthor = {
    id: string
    name: string
    link: string
    created_at: string | null
    updated_at: string | null
    deleted_at: string | null
}

export type IllustrationTag = {
    id: string
    name: string
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