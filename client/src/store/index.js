import Vue from 'vue'
import Vuex from 'vuex'

// module imports
import user from './modules/UserStore'
import task from './modules/TaskStore'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    user,
    task
  }
})
