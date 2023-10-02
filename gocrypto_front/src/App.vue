<script>
import { defineComponent } from 'vue';
import Navbar from './components/navbar.vue'
import axios from './api'
import store from './store';
export default defineComponent({
  name: 'App',
   setup() {
        return {
            user: store.state.user,
            store : store,
            async getUser() {
                const token = localStorage.getItem('token')
                const user = await axios.get("/users/me", {
                    headers: {
                        Authorization: `${token}`
                    }
                })
                console.log(user.data)
                store.commit('setUser', user.data)
            }
        }
    },
    mounted() {
        this.getUser()
    }
  
});
</script>

<template>
   
  <h1 class="text-3xl font-bold underline text-blue-500">
    Hello world!
  </h1>
  <button class=" btn btn-primary">
    teste botao DaisyUI
    </button>

    <RouterView />
</template>

<style scoped></style>
