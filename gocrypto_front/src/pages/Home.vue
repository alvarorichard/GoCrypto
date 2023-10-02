<script>
import { defineComponent } from 'vue';
import store from '../store';
import axios from '../api'
import Navbar from '../components/navbar.vue';
export default defineComponent({
    name: 'Home',
    components: { Navbar },
    setup() {
        return {
            msg: 'Home',
            user: store.state.user,
            store: store,
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
})
</script>
<template>
    <Navbar />
    <h1>Home</h1>
    <h1 v-if="user">
        Olá {{ user.name }}
    </h1>
</template>