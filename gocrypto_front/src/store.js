import { createStore } from 'vuex'

const store = createStore({
  state() {
    return {
      userIsAuthenticated: false,
      user: {
        admin: false,
        name: '',
        id: '',
        exp: 0,
      }
    }
  },
  mutations: {
    setUserIsAuthenticated(state, payload) {
      if (typeof payload === 'boolean') {
        state.userIsAuthenticated = payload;
      } else {
        console.error('Expected payload to be a boolean for setUserIsAuthenticated mutation');
      }
    },
    setUser(state, payload) {
      if (payload && payload.claims) {
        state.user = { ...payload.claims };
      } else {
        console.error('Expected payload to have a claims property for setUser mutation');
      }
    }
  }
});

export default store;
