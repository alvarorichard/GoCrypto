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


       
            <RouterLink to="/">Home</RouterLink>
      
        <div class="flex gap-6" v-if="!isUserLoggedIn">
           
                <RouterLink to="/login"> Login</RouterLink>
                <RouterLink to="/register"> Registro</RouterLink>
                <RouterLink to="/market"> Market</RouterLink>
                <RouterLink to="/trade"> Trade</RouterLink>
                <RouterLink to="/sobreNos"> Sobre Nós</RouterLink>
                <RouterLink to="/contato"> Contato</RouterLink>
         
        </div>

   
   
</template>