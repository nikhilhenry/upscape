const path = require('path')

module.exports = {
  css:{
    loaderOptions:{
      scss:{
        prependData: `@import "@/assets/colors.scss"; @import "@/assets/utils.scss";`
      }
    }
  }
}