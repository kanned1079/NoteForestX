// ecosystem.config.js
module.exports = {
    apps: [
        {
            name: 'note-forest-x_server',
            script: './bin/noteforestx-server-darwin-arm64',
            args: 'server',
            exec_mode: 'fork',
            instances: 1,
            autorestart: true,
            watch: false,
            env: {
                GIN_MODE: 'release'
            }
        }
    ]
}