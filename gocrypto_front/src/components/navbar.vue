<script>
import { defineComponent } from "vue";
import store from "../store";
export default defineComponent({
  name: "Navbar",
  setup() {
    return {
      msg: "Navbar",
      isUserLoggedIn: store.state.userIsAuthenticated,
    };
  },
  mounted() {
    store.watch(
      (state) => state.userIsAuthenticated,
      (value) => {
        this.isUserLoggedIn = value;
      },
    );
  },
});
</script>
<template>
  <header
    class="w-full h-14 px-5 flex justify-between items-center fixed top-0 bg-transparent backdrop-blur-xl border-b-2 border-b-[#117A7830]"
    v-if="!isUserLoggedIn"
  >
    <RouterLink to="/">
      <img
        class="h-14 scale-125 duration-300 hover:scale-150"
        src="https://i.imgur.com/CAp9EhS.png"
        alt="Imagem logo"
      />
    </RouterLink>
    <nav class="flex items-center gap-6 max-[750px]:hidden">
      <RouterLink class="hover:text-white" to="/login"> Login</RouterLink>
      <RouterLink class="hover:text-white" to="/register"> Registro</RouterLink>
      <RouterLink class="hover:text-white" to="/Carteira"> Carteira</RouterLink>
      <RouterLink class="hover:text-white" to="/market"> Market</RouterLink>
      <RouterLink class="hover:text-white" to="/trade"> Trade</RouterLink>
      <RouterLink class="hover:text-white" to="/sobreNos"> Sobre Nós</RouterLink>
      <RouterLink class="hover:text-white" to="/contato"> Contato</RouterLink>
    </nav>
    <!-- trocar o conteudo do button por uma img com um icone de menu hamburguer -->
    <button class="max-[750px]:block min-[750px]:hidden">
      <img src="../assets/icons/menu-hamburguer.svg" alt="" />
    </button>
  </header>
</template>
