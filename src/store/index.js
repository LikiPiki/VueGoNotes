import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    token: '',
    userId: ''
  },
  actions: {
    setToken ({commit}, token) {
      commit('set', {type: 'token', item: token})
    },
    setUserId ({commit}, id) {
      commit('set', {type: 'userId', item: id})
    }
  },
  mutations: {
    set (state, {type, item}) {
      console.log('changing')
      state[type] = item
    }
  },
  getters: {
    getToken (state) {
      return state.token
    },
    getUserId (state) {
      return state.userId
    }
  }
})

export default store
