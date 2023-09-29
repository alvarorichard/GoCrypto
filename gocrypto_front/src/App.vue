<script>
import { RouterLink } from 'vue-router';
import { defineComponent } from 'vue';
import axios from './api'
import store from './store';
export default defineComponent({
  name: 'App',
  setup() {
    return {
      isUserAuthenticated: store.state.userIsAuthenticated,
      user: store.state.user,
      getUserData: async () => {
        const token = localStorage.getItem('token');
        if (token) {
          const info = await axios.get('/users/me', {
            headers: {
              Authorization: `${token}`
            }
          })
         store.commit('setUser', info.data.claims)
        }
      },

    };
  },
  computed: {
  },
  mounted() {
    this.getUserData()
  },
});
</script>

<template>
  <ul>
    <li>
      <RouterLink to="/">Home</RouterLink>
    </li>
    <div v-if="isUserAuthenticated == false">
      <li>
        <RouterLink to="/login"> Login</RouterLink>
      </li>
      <li>
        <RouterLink to="/register"> Registro</RouterLink>
      </li>
    </div>

  </ul>
  <RouterView />
</template>

<style scoped></style>
