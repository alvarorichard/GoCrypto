import { createStore } from 'vuex'


const store = createStore({
  state () {
    return {
    userIsAuthenticated: false,
    user : {
        admin: false,
        name: '',
        id: '',
        exp: 0,
    }
    }
  },
  mutations: {
    setUserIsAuthenticated(state, payload) {
      state.userIsAuthenticated = payload}
    },
    setUser(state, payload) {
        state.user = payload
    },
})


export default store;