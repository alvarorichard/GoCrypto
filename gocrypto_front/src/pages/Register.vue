<script>
import { defineComponent } from 'vue';

import Navbar from '../components/navbar.vue';
import axios from '../api'
import store from '../store';
export default defineComponent({
    name: 'Register',
    components: { Navbar },
    setup() {
        return {
            msg: 'Register',
            LoginDetails: {
                email: '',
                password: '',
                name: ''
            },
            onSu(e) {
                e.preventDefault()
                axios.post('/register', this.LoginDetails)
                    .then(res => {
                        if (res.status == 200) {
                            localStorage.setItem('token', res.data.token)
                            localStorage.setItem('user', JSON.stringify(res.data.user))
                            store.commit('setUserIsAuthenticated', true)
                            this.$router.push('/')
                        }
                    })
                    .catch(err => {
                        alert(err)
                    })
            }

        }
    },
})
</script>

<template>
    <Navbar />
    <form @submit="onSu">
        <input type="text" v-model="LoginDetails.name" placeholder="Name">
        <input type="text" v-model="LoginDetails.email" placeholder="Email">
        <input type="password" v-model="LoginDetails.password" placeholder="Password">
        <button type="submit">Register</button>
    </form>

</template>