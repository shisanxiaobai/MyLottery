module.exports = {
  transpileDependencies: ['vuetify'],
  assetsDir: 'static',
  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = '欢迎来到十三小白的购物系统'
      return args
    })
  },
  productionSourceMap: false
}
