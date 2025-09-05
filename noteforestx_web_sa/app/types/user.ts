export type User = {
    id: string,
    username?: string | null,
    email: string,
    role: 'ADMIN' | 'USER',
    token?: '',
    created_at?: Date,
    updated_at?: Date,
    deleted_at?: Date
}