
  
  <script>
  import { defineComponent } from 'vue';
  import store from '../store';
  import axios from '../api';
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
          const token = localStorage.getItem('token');
          const user = await axios.get("/users/me", {
            headers: {
              Authorization: `${token}`
            }
          });
          console.log(user.data);
          store.commit('setUser', user.data);
        }
      }
    },
  });
  </script>

<template>


 <!---Colocando os elementos e alinhando a Header-->
    <div class="bg-black min-h-screen flex flex-col justify-between">
      <header class="flex justify-between items-center bg-black p-5">
        <Navbar />
      </header>
  

       <!---Alinhando a imagem e o texto no centro da Pagina-->
      <div class="flex-grow flex justify-center items-center text-white font-bold text-4xl relative">
        <div class="text-center">
          <img src="../components/imagem/macaco.jpg" width="600" height="800" alt="macaco" />
          <strong class="absolute inset-x-0 bottom-28">BUY AND SELL CRYPTOCURRENCY</strong>
        </div>
      </div>
  
          <!---Rodapé (a parte branca da pagina)-->
        
    </div>

    <footer class="absolute inset-x-0 bottom-10 bg-white px-2 py-2 mt-5 mx-96 rounded-lg flex justify-between items-center max-md:px-44 max-sm:inset- bottom-1 ">
      
      <div class="flex items-center space-x-7 mx-2">
        <img src="../components/imagem/Bitcoin.svg.png" alt="bitcoin" class="h-10 w-10 rounded-full">
        <p class="text-gray-700">BTC</p> <p class="text-gray-700">|</p>
        <p >Bitcoin</p> 
      </div>
      <!---Função para assim que apertar o botão ir para a pagina do Market-->
      <router-link to="/market">
          <button class="px-10 py-2.7-8 bg-blue-500 text-white rounded hover:bg-blue-600">Buy</button>
        </router-link>
    </footer>
  </template>
  
  <style scoped>
  
  </style>
  