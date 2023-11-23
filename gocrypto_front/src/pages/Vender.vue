<script>
import { defineComponent } from "vue";
import Navbar from "../components/navbar.vue";
import {ref } from "vue";
import api from "../api";
export default defineComponent({
  name: "Vender",
  components: { Navbar },
  setup() {
    const Operation = ref({
      amount: "",
      coin: "",
      type: "sell"
    });
    const msg = ref("Vender");

    const onSubmit = (e) => {
      e.preventDefault();
      api.post("/api/wallet/sell", {
        amount: Operation.value.amount,
        coin_id: Number(Operation.value.coin),
        operation: Operation.value.type,
      }, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
      }).then((res) => {
        if (res.status == 200) {
          alert("Venda realizada com sucesso");
        }
      }).catch((err) => {
        alert(err);
      });
    };

    return {
      Operation,
      msg,
      onSubmit,
    };
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
      <h1 class="text-3xl mb-5">Vender</h1>
      <form class="w-full flex flex-col justify-center items-center gap-4">
        <div class="w-[300px] flex flex-col gap-1">
          <label>Quantidade de Criptomoedas</label>
          <input class="p-1 text-white" type="number" v-model="Operation.amount" placeholder="0.00000000" />
        </div>
        <div class="w-[300px] flex flex-col gap-1">
          <label>Criptomoeda</label>
          <select v-model="Operation.coin" class="p-1 text-white">
            <option value="1">BTC</option>
            <option value="2">ETH</option>
            <option value="3">USDT</option>
            <option value="4">DOGE</option>
          </select>
        </div>
        <button
        @click="onSubmit"
          class="w-[300px] p-1 duration-300 bg-black text-white rounded hover:bg-gray-200 hover:text-black"
        >
          Comprar
        </button>
      </form>
    </div>
  </main>
</template>

<style scoped></style>
