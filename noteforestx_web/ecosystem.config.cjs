module.exports = {
    apps: [
        {
            name: 'note-forest-x',
            script: '.output/server/index.mjs',
            interpreter: 'node',
            env: {
                API_BASE_URL: "http://127.0.0.1:8081",
                PORT: 14000
            }
        }
    ]
}