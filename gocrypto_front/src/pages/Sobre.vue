<script>
import { defineComponent } from "vue";
import Navbar from "../components/navbar.vue";

export default defineComponent({
  name: "Sobre",
  components: { Navbar },
  setup() {
    return {};
  },
});
</script>

<template>
  <Navbar />
  <main class="w-full min-h-screen flex flex-col justify-center items-center gap-8 pt-16 pb-6 px-2">
    <img src="../assets/images/hero-about.svg" width="400" alt="" />
    <h1 class="text-4xl font-bold text-center">Welcome to GoCrypto</h1>
    <p class="max-w-[800px] text-lg text-center">
      Bem-vindo à GoCrypto, a sua plataforma confiável para a compra e venda de
      criptomoedas. Fundada em 1941, a nossa missão é tornar o mundo das criptomoedas
      acessível a todos, independentemente do nível de experiência.
    </p>
    <div class="w-full max-w-[800px] flex justify-between items-center gap-5">
      <div class="flex flex-col items-center flex-wrap gap-2">
        <h1 class="text-xl font-bold text-center text-[#eee]">+1M Users</h1>
        <img src="../assets/icons/user-icon.png" alt="" />
      </div>
      <div class="flex flex-col items-center gap-2">
        <h1 class="text-xl font-bold text-center text-[#eee]">+150M Transactions</h1>
        <img src="../assets/icons/money-icon.png" alt="" />
      </div>
      <div class="flex flex-col items-center gap-2">
        <h1 class="text-xl font-bold text-center text-[#eee]">24/7 Support</h1>
        <img src="../assets/icons/support-icon.png" alt="" />
      </div>
    </div>
  </main>
</template>

<style scoped></style>
