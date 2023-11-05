<script>
import { defineComponent } from "vue";
import Navbar from "../components/navbar.vue";

export default defineComponent({
  name: "Contato",
  components: { Navbar },
  setup() {
    return {};
  },
});
</script>

<template>
  <Navbar />
  <main class="w-full min-h-screen flex flex-col justify-center items-center gap-8 pt-16 pb-6 px-2">
    <img src="../assets/images/hero-support.svg" width="400" alt="" />
    <h1 class="text-4xl font-bold text-center">Central de Ajuda - Contato</h1>
    <div class="flex flex-col gap-4">
      <h1 class="text-2xl">Principais assuntos</h1>
      <div class="w-full max-w-[500px] flex items-center flex-wrap gap-2">
        <a class="p-2 bg-gray-900 duration-300 hover:text-white" href="/">Redefinir Senha</a>
        <a class="p-2 bg-gray-900 duration-300 hover:text-white" href="/">Desbloquear Conta</a>
        <a class="p-2 bg-gray-900 duration-300 hover:text-white" href="/"
          >Alterar Endereço de email</a
        >
        <a class="p-2 bg-gray-900 duration-300 hover:text-white" href="/">Scam ou Phising</a>
        <a class="p-2 bg-gray-900 duration-300 hover:text-white" href="/"
          >Problemas ao comprar criptomoedas</a
        >
      </div>
    </div>
    <button class="fixed bottom-0 right-0 m-5 bg-gray-900 rounded-full p-3 hover:bg-gray-800">
      <img src="../assets/icons/chat-icon.png" width="32" alt="" />
    </button>
  </main>
</template>

<style scoped></style>
