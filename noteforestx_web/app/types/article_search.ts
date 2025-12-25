export type SearchQuery = {
    search?: string // 搜索文章标题
    tag?: string // 搜索tag名称
    tag_id?: string // 传入tag的uuid 优先级最高 如果提供了忽略上面两个
    status?: string // 不传就查询全部
}