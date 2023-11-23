<script>
import api from "../api";
import { defineComponent } from "vue";
import Navbar from "../components/navbar.vue";
import Trasacao from "../components/transacao.vue"
import { RouterLink } from "vue-router";
import { ref } from 'vue';


export default defineComponent({
  name: "Carteira",

  components: { Navbar, Trasacao },
  setup() {
    return {
      msg: "Carteira",
      transactions: [
      ],
      saldo: 0,
      renderTx: ref([]),
      txReady: true,
    }
  },
  mounted() {
    api.get('/api/wallet/my', {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`
      }
    }).then((res) => {
      this.transactions = res.data;
      this.transactions.forEach((tx) => {
        if (tx.operation == 'buy') {
          this.saldo += tx.amount;
        } else {
          this.saldo -= tx.amount;
        }
      });
      this.renderTx = res.data;
      
      console.log(renderTx)
    }).catch((err) => {
      console.log(err);
    });
  
  }

});
</script>

<template>
  <Navbar />
  <main id="wallet-page" class="w-full min-h-screen flex items-end pt-14 p-14 text-black">
    <div
      class="w-full min-h-[70vh] flex justify-between max-[900px]:gap-0 gap-5 px-12 rounded-t-3xl rounded-b-3xl bg-white max-[900px]:flex-col max-[900px]:px-0">
      <div
        class="w-1/2 mt-[-80px] flex flex-col gap-2 pt-5 px-5 overflow-auto max-[900px]:rounded-none rounded-t-3xl max-[900px]:shadow shadow-2xl bg-white max-[900px]:w-full max-[900px]:mt-0">
        <div class="flex flex-col gap-4">
          <div class="flex flex-col">
            <h1 class="text-2xl font-medium">Carteira</h1>
            <h2 class="text-2xl font-medium text-green-600">R$ {{ saldo }}</h2>
          </div>
          <div class="flex flex-col">
            <h1 class="flex items-center gap-2 text-2xl font-medium">
              <img src="../assets/images/Bitcoin.svg.png" width="28" alt="" />
              Bitcoin
            </h1>
            <h2 class="text-2xl font-medium text-green-600">0.0004324029</h2>
          </div>
          <div class="flex flex-col">
            <h1 class="flex items-center gap-2 text-2xl font-medium">
              <img src="https://th.bing.com/th/id/OIP.5rXpT4zjW6BBgP_z-97bqwHaHa?pid=ImgDet&rs=1" width="28" alt="" />
              Monero
            </h1>
            <h2 class="text-2xl font-medium text-green-600">0.0001832029</h2>
            <div class="flex items-center gap-2 mt-5">
              <RouterLink class="p-2 px-4 rounded hover:opacity-80 text-white bg-green-600" to="/comprar">Comprar </RouterLink>
               <RouterLink class="p-2 px-4 rounded hover:opacity-80 text-white bg-red-600" to="/vender">Vender
              </RouterLink>
            </div>
          </div>
          <hr />
        </div>
        <div class="w-full h-full">
          <img src="../assets/images/graph.svg" alt="" />
        </div>
      </div>
      <div
        class="w-1/2 mt-[-80px] flex flex-col gap-2 pt-5 pl-5 overflow-auto max-[900px]:rounded-none rounded-t-3xl max-[900px]:shadow shadow-2xl bg-white max-[900px]:w-full max-[900px]:mt-0 max-[900px]:pl-2 max-[900px]:pr-2">
        <h1 class="text-2xl font-medium">Transações</h1>
        <table class="" v-if="txReady">
          <span v-for="taa in renderTx">
            <Trasacao :tx="taa" />
          </span>
        </table>
      </div>
    </div>
  </main>
</template>
