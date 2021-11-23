module.exports = {
    devServer: {
        proxy: {
            "/get": {
                target: "http://127.0.0.1:8888",
                ws: true
            },
            "/update": {
                target: "http://127.0.0.1:8888",
            }
        }
    },

    publicPath: '',
    outputDir: '../back/dist'
}