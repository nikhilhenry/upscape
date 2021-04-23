import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

// external libraries
import axios from 'axios'

// axios baseURL
axios.defaults.baseURL = process.env.VUE_APP_API_URL

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
