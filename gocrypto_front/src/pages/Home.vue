<script>
import { defineComponent } from "vue";
import store from "../store";
import axios from "../api";
import Navbar from "../components/navbar.vue";
import Footer from "../components/footer.vue";

export default defineComponent({
  name: "Home",
  components: { Navbar, Footer },
  setup() {
    return {
      msg: "Home",
      user: store.state.user,
      store: store,
      async getUser() {
        const token = localStorage.getItem("token");
        const user = await axios.get("/users/me", {
          headers: {
            Authorization: `${token}`,
          },
        });
        console.log(user.data);
        store.commit("setUser", user.data);
      },
    };
  },
});
</script>

<template>
  <Navbar />
  <main
    id="home-page"
    class="w-full min-h-screen flex flex-col justify-center items-center pt-16 px-2"
  >
    <div class="flex justify-center items-center flex-wrap">
      <h1 class="max-w-sm text-4xl text-white font-medium max-[783px]:text-center">
        A maneira mais popular do mundo de comprar vender e trocar criptomoeda
      </h1>
      <img class="h-[400px]" src="../assets/images/hero-home.svg" alt="" />
    </div>
    <h1 class="mb-8 text-2xl text-center text-white">BUY AND SELL CRIPTOCURRENCY</h1>
    <div class="w-full flex justify-center items-center flex-wrap gap-5">
      <div
        class="w-48 h-44 flex flex-col justify-center items-center gap-3 p-4 rounded-xl duration-300 hover:opacity-80 bg-white text-black cursor-default"
      >
        <img src="../assets/images/Bitcoin.svg.png" width="52" alt="" />
        <div class="flex flex-col justify-center items-center gap-1">
          <h1 class="text-xl">Bitcoin <span class="font-medium">BTC</span></h1>
          <p class="text-lg font-semibold">$34,687.42</p>
          <p class="text-green-700">+R$1,520,20</p>
        </div>
      </div>
      <div
        class="w-48 h-44 flex flex-col justify-center items-center gap-3 p-4 rounded-xl duration-300 hover:opacity-80 bg-white text-black cursor-default"
      >
        <img
          src="https://th.bing.com/th/id/R.86b356aa4636232f3e200c65d2a8b6b4?rik=MjmYLeoNiBw6Bg&pid=ImgRaw&r=0"
          width="52"
          alt=""
        />
        <div class="flex flex-col justify-center items-center gap-1">
          <h1 class="text-xl">Ethereum <span class="font-medium">ETH</span></h1>
          <p class="text-lg font-semibold">$1,796.64</p>
          <p class="text-red-700">-R$1,520,20</p>
        </div>
      </div>
      <div
        class="w-48 h-44 flex flex-col justify-center items-center gap-3 p-4 rounded-xl duration-300 hover:opacity-80 bg-white text-black cursor-default"
      >
        <img
          src="https://th.bing.com/th/id/OIP.5rXpT4zjW6BBgP_z-97bqwHaHa?pid=ImgDet&rs=1"
          width="52"
          alt=""
        />
        <div class="flex flex-col justify-center items-center gap-1">
          <h1 class="text-xl">Monero <span class="font-medium">XMR</span></h1>
          <p class="text-lg font-semibold">R$828.65</p>
          <p class="text-green-700">+R$1,520,20</p>
        </div>
      </div>
    </div>
  </main>
  <Footer />
</template>

<style scoped></style>
