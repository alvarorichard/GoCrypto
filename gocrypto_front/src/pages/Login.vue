<script>
import { defineComponent } from 'vue';
import axios from '../api'
import store from '../store';
import Navbar from '../components/navbar.vue';
export default defineComponent({
    name: 'Login',
    components: { Navbar },
    setup() {
        return {
            msg: 'Login',
            LoginDetails: {
                email: '',
                password: ''
            },
            onSu(e) {
                e.preventDefault();
                axios.post('/login', this.LoginDetails)
                    .then(res => {
                    if (res.status == 200) {
                        localStorage.setItem('token', res.data.token);
                        localStorage.setItem('user', JSON.stringify(res.data.user));
                        store.commit('setUserIsAuthenticated', true);
                        this.$router.push('/');
                    }
                })
                    .catch(err => {
                    alert(err);
                });
            }
        };
    },
    components: { Navbar }
})
</script>
<template>
    

    <!---Colocando os elementos e alinhando a Header-->
    <div class="bg-black min-h-screen  justify-between">
        <header class="flex justify-between items-center bg-black p-5">
          <Navbar />
        </header>
      
        <div className="absolute mt-20">
        <h1 className=" text-white font-bold text-4xl relative text-left ">WELCOME</h1>
        <h2 className=" text-white font-bold text-4xl relative  text-left " >BUY AND SELL CRYPTOCURRENCY</h2>
        </div>
       
        <div class="flex justify-center items-center ml-40">
           
            <!-- Lado esquerdo com inputs -->
            <div class="text-left mr-auto ">
              <form @submit="onSu">
                <div class="mb-4">
                  <label for="email" class="block text-white">Login</label>
                  <input
                    type="text"
                    id="email"
                    v-model="LoginDetails.email"
                    placeholder="Email"
                    class="border border-white rounded p-2 w-64"
                  />
                </div>
                <div class="mb-4 ">
                  <label for="password" class="block text-white">Senha</label>
                  <input
                    type="password"
                    id="password"
                    v-model="LoginDetails.password"
                    placeholder="Password"
                    class="border border-white rounded p-2 w-64"
                  />
                </div>
                <button type="submit" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                  Login
                </button>
              </form>
              <router-link to="/registro">
                <button class=" py-2 px-4 rounded bg-blue-500 text-white rounded hover:bg-blue-600 mt-10">Não tenho Conta</button>
              </router-link>
            </div>


          
            <!-- Canto direito com imagem e strong -->
            <div class="relative text-white  text-4xl fixed">
              <img src="../components/imagem/macaco.jpg" width="680" height="1000" alt="macaco" />
            </div>
          </div>
          
      
    
  
   
    </div>
</template>


<style scoped></style>