import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

// external libraries
import axios from 'axios'
import TimeAgo from 'javascript-time-ago'
import en from 'javascript-time-ago/locale/en'

// requirement to set locale
TimeAgo.addDefaultLocale(en)

// axios baseURL
axios.defaults.baseURL = process.env.VUE_APP_API_URL


Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
