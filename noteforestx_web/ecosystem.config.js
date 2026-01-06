module.exports = {
    apps: [
        {
            name: 'note-forest-x',
            script: '/home/kanna/NoteForestX/noteforestx_web/.output/server/index.mjs',
            interpreter: 'node',
            env: {
                API_BASE_URL: "https://ikanned.com:14000",
                PORT: 14000
            }
        }
    ]
}