export type Article = {
    id: string              // uuid
    slug?: string // 现在不用，以后不破坏结构
    title: string
    top: boolean            // 是否置顶（可以暂时保留 boolean）
    status: ArticleStatus   // 文章状态
    content: string         // Markdown
    tags: ArticleTag[]      // ✅ 关键
    created_at?: string
    updated_at?: string
    deleted_at?: string
}

export type ArticleTag = {
    id: string
    name: string
    code?: string           // ✅ 稳定标识（即使你不用 slug）
    created_at?: string
}

export type ArticleStatus =
    | 'draft'
    | 'published'
    | 'hidden'

export const ARTICLE_STATUS_FLOW: ArticleStatus[] = [
    'draft',
    'published',
    'hidden',
]

export const NEXT_ARTICLE_STATUS_MAP: Record<ArticleStatus, ArticleStatus> = {
    draft: 'published',
    published: 'hidden',
    hidden: 'draft',
}

export type NewTag = {
    id?: string
    name?: string
}

export type NewArticle = {
    title: string
    slug?: string
    status: ArticleStatus
    content: string
    tags: NewTag[]
    image_url?: string
}