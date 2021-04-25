const path = require('path')

module.exports = {
  outputDir:path.resolve(__dirname,'../public'),
  css:{
    loaderOptions:{
      scss:{
        prependData: `@import "@/assets/colors.scss"; @import "@/assets/utils.scss";`
      }
    }
  }
}