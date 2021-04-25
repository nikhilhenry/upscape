import Vue from 'vue'
import Vuex from 'vuex'

// module imports
import user from './modules/UserStore'
import task from './modules/TaskStore'
import tag from './modules/TagStore'
import objective from './modules/ObjectiveStore'

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
    task,
    tag,
    objective
  }
})
