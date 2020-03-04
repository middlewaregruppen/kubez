module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  devServer: {
    proxy: {
      '^/kubez': {
        target: 'http://localhost:3000/',
        //target: 'http://192.168.64.11:31503/',
        ws: true,
        changeOrigin: true
      }
    }
  }
}
