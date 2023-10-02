<script>
import { defineComponent } from 'vue';
import store from '../store';
export default defineComponent({
    name: 'Navbar',
    setup() {
        return {
            msg: 'Navbar',
            isUserLoggedIn: store.state.userIsAuthenticated,
        }
    },
    mounted() {
        store.watch(
            (state) => state.userIsAuthenticated,
            (value) => {
                this.isUserLoggedIn = value;
            }
        );
    }
})
</script>
<template>
    <ul>
        <li>
            <RouterLink to="/">Home</RouterLink>
        </li>
        <div v-if="!isUserLoggedIn">
            <li>
                <RouterLink to="/login"> Login</RouterLink>
            </li>
            <li>
                <RouterLink to="/register"> Registro</RouterLink>
            </li>
        </div>

    </ul>
</template>