<script>
import { defineComponent } from "vue";
import Navbar from "../components/navbar.vue";

export default defineComponent({
  name: "Comprar",
  components: { Navbar },
  setup() {
    return {};
  },
});
</script>

<template>
  <Navbar />
  <main
    id="buy-page"
    class="w-full min-h-screen flex flex-col justify-center items-center gap-8 pt-16 pb-6 px-2"
  >
    <div
      class="w-full max-w-[600px] flex flex-col justify-center items-center p-5 rounded-2xl bg-white text-black"
    >
      <img src="https://i.imgur.com/CAp9EhS.png" width="300" alt="Imagem logo" />
      <h1 class="text-3xl mb-5">Comprar</h1>
      <form class="w-full flex flex-col justify-center items-center gap-4">
        <div class="w-[300px] flex flex-col gap-1">
          <label>Sua Carteira</label>
          <input class="p-1 text-white" type="text" placeholder="odasjdosajkdasoj239" />
        </div>
        <div class="w-[300px] flex flex-col gap-1">
          <label>Quantidade de Criptomoedas</label>
          <input class="p-1 text-white" type="number" placeholder="0.00000000" />
        </div>
        <div class="w-[300px] flex flex-col gap-1">
          <label>Criptomoeda</label>
          <select class="p-1 text-white">
            <option value="BTC">BTC</option>
            <option value="ETH">ETH</option>
            <option value="USDT">USDT</option>
            <option value="DOGE">DOGE</option>
          </select>
        </div>
        <button
          class="w-[300px] p-1 duration-300 bg-black text-white rounded hover:bg-gray-200 hover:text-black"
        >
          Comprar
        </button>
      </form>
    </div>
  </main>
</template>

<style scoped></style>
