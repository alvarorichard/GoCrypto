<script>
import { defineComponent } from "vue";

import Navbar from "../components/navbar.vue";
import axios from "../api";
import store from "../store";
import { RouterLink } from "vue-router";

export default defineComponent({
  name: "Register",
  components: { Navbar },
  setup() {
    return {
      msg: "Register",
      RegisterDetails: {
        email: "",
        password: "",
      },
      onSu(e) {
        e.preventDefault();
        //alert("cadastro com sucesso");
        console.log(this.RegisterDetails);
        axios
          .post("/api/user/register", this.RegisterDetails)
          .then((res) => {
            if (res.status == 200) {
              localStorage.setItem("token", res.data.token);
              localStorage.setItem("user", JSON.stringify(res.data.user));
              store.commit("setUserIsAuthenticated", true);
              this.$router.push("/");
            }
          })
          .catch((err) => {
            alert(err);
          });
      },
    };
  },
});
</script>

<template>
  <Navbar />
  <main class="w-full min-h-screen flex items-center pt-16 pb-6 px-2 bg-black">
    <section class="w-6/12 flex flex-col justify-center items-center gap-5 px-2 max-md:w-full">
      <img class="md:hidden" src="https://i.imgur.com/CAp9EhS.png" alt="Imagem logo" />
      <h1 class="text-3xl text-center text-white font-bold">Are you ok to go?</h1>
      <form class="max-w-[300px] w-full flex flex-col gap-3" @submit.prevent="onSu">
        <div class="flex flex-col gap-1">
        
        </div>
        <div class="flex flex-col gap-1">
          <label>Email</label>
          <input
            class="p-2 rounded-sm border-gray-900 border-2"
            v-model="RegisterDetails.email"
            type="email"
            placeholder="Digite seu email"
          />
        </div>
        <div class="flex flex-col gap-1">
          <label>Senha</label>
          <input
            class="p-2 rounded-sm border-gray-900 border-2"
            v-model="RegisterDetails.password"
            type="password"
            placeholder="Digite sua senha"
          />
        </div>
        <p>
          Já tem uma conta?
          <RouterLink class="text-white hover:underline" to="/login"
            >Entrar</RouterLink
          >
        </p>
        <button
          class="w-full h-12 bg-gray-700 duration-300 text-white rounded hover:bg-gray-900"
          type="submit"
        >
          Cadastrar
        </button>
      </form>
    </section>
    <section class="w-6/12 flex flex-col justify-center items-center max-md:hidden">
      <img class="h-[600px]" src="../assets/images/hero-signup.svg" alt="" />
    </section>
  </main>
</template>
