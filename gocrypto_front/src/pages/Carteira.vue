<script>
import { defineComponent } from "vue";
import Navbar from "../components/navbar.vue";

export default defineComponent({
  name: "Carteira",
  components: { Navbar },
  setup() {
    return {};
  },
});
</script>

<template>
  <Navbar />
  <main id="wallet-page" class="w-full min-h-screen flex items-end pt-14 p-14 text-black">
    <div
      class="w-full min-h-[70vh] flex justify-between max-[900px]:gap-0 gap-5 px-12 rounded-t-3xl rounded-b-3xl bg-white max-[900px]:flex-col max-[900px]:px-0"
    >
      <div
        class="w-1/2 mt-[-80px] flex flex-col gap-2 pt-5 px-5 overflow-auto max-[900px]:rounded-none rounded-t-3xl max-[900px]:shadow shadow-2xl bg-white max-[900px]:w-full max-[900px]:mt-0"
      >
        <div class="flex flex-col gap-4">
          <div class="flex flex-col">
            <h1 class="text-2xl font-medium">Carteira</h1>
            <h2 class="text-2xl font-medium text-green-600">R$ 192,592.39</h2>
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
              <img
                src="https://th.bing.com/th/id/OIP.5rXpT4zjW6BBgP_z-97bqwHaHa?pid=ImgDet&rs=1"
                width="28"
                alt=""
              />
              Monero
            </h1>
            <h2 class="text-2xl font-medium text-green-600">0.0001832029</h2>
            <div class="flex items-center gap-2 mt-5">
              <button class="p-2 px-4 rounded hover:opacity-80 text-white bg-green-600">
                Enviar
              </button>
              <button class="p-2 px-4 rounded hover:opacity-80 text-white bg-green-600">
                Receber
              </button>
              <button class="p-2 px-4 rounded hover:opacity-80 text-white bg-green-600">
                Comprar
              </button>
            </div>
          </div>
          <hr />
        </div>
        <div class="w-full h-full">
          <img src="../assets/images/graph.svg" alt="" />
        </div>
      </div>
      <div
        class="w-1/2 mt-[-80px] flex flex-col gap-2 pt-5 pl-5 overflow-auto max-[900px]:rounded-none rounded-t-3xl max-[900px]:shadow shadow-2xl bg-white max-[900px]:w-full max-[900px]:mt-0 max-[900px]:pl-2 max-[900px]:pr-2"
      >
        <h1 class="text-2xl font-medium">Transações</h1>
        <table class="">
          <tr>
            <th>Transação</th>
            <th>Data</th>
            <th>Horário</th>
            <th>Tipo</th>
            <th>Valor</th>
            <th>Status</th>
          </tr>
          <tr>
            <td>Transação 1</td>
            <td>01/11/2023</td>
            <td>14:30</td>
            <td>Bitcoin</td>
            <td>R$50,00</td>
            <td class="text-red-600 flex items-center gap-2">
              <span class="w-[10px] h-[10px] rounded-full bg-red-600"></span> Failed
            </td>
          </tr>
          <tr>
            <td>Transação 2</td>
            <td>01/11/2023</td>
            <td>14:30</td>
            <td>Bitcoin</td>
            <td>R$50,00</td>
            <td class="text-green-600 flex items-center gap-2">
              <span class="w-[10px] h-[10px] rounded-full bg-green-600"></span>
              Success
            </td>
          </tr>
          <tr>
            <td>Transação 3</td>
            <td>01/11/2023</td>
            <td>14:30</td>
            <td>Bitcoin</td>
            <td>R$50,00</td>
            <td class="text-green-600 flex items-center gap-2">
              <span class="w-[10px] h-[10px] rounded-full bg-green-600"></span>
              Success
            </td>
          </tr>
          <tr>
            <td>Transação 4</td>
            <td>01/11/2023</td>
            <td>14:30</td>
            <td>Bitcoin</td>
            <td>R$50,00</td>
            <td class="text-red-600 flex items-center gap-2">
              <span class="w-[10px] h-[10px] rounded-full bg-red-600"></span>
              Failed
            </td>
          </tr>
          <tr>
            <td>Transação 5</td>
            <td>01/11/2023</td>
            <td>14:30</td>
            <td>Bitcoin</td>
            <td>R$50,00</td>
            <td class="text-green-600 flex items-center gap-2">
              <span class="w-[10px] h-[10px] rounded-full bg-green-600"></span>
              Success
            </td>
          </tr>
        </table>
      </div>
    </div>
  </main>
</template>
