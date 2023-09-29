<script>
import { defineComponent } from 'vue';
import axios from '../api'
import store from '../store';
export default defineComponent({
    name: 'Login',
    setup() {
        return {
            msg: 'Login',
            LoginDetails: {
                email: '',
                password: ''
            },
            onSu(e) {
                e.preventDefault()
                axios.post('/login', this.LoginDetails)
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
    <form @submit="onSu">
        <input type="text" v-model="LoginDetails.email" placeholder="Email">
        <input type="password" v-model="LoginDetails.password" placeholder="Password">
        <button type="submit">Login</button>
    </form>
    <h1> {{ LoginDetails }}</h1>
</template>


<style scoped></style>