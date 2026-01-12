module.exports = {
    apps: [
        {
            name: 'note-forest-x',
            script: '.output/server/index.mjs',
            interpreter: 'node',
            env: {
                API_BASE_URL: "http://localhost:14001",
                PORT: 14000
            }
        }
    ]
}