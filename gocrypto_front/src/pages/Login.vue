<script>
import { defineComponent } from "vue";
import axios from "../api";
import store from "../store";
import Navbar from "../components/navbar.vue";
import { RouterLink } from "vue-router";

export default defineComponent({
  name: "Login",
  components: { Navbar },
  setup() {
    return {
      msg: "Login",
      LoginDetails: {
        email: "",
        password: "",
      },
      onSu(e) {
        e.preventDefault();
        alert("Login com sucesso");
        console.log(this.LoginDetails);
        axios
          .post("/api/user/login", this.LoginDetails)
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
  components: { Navbar },
});
</script>
<template>
  <Navbar />
  <main class="w-full min-h-screen flex items-center pt-16 pb-6 px-2 bg-black">
    <section class="w-6/12 flex flex-col justify-center items-center gap-5 px-2 max-md:w-full">
      <img class="md:hidden" src="https://i.imgur.com/CAp9EhS.png" alt="Imagem logo" />
      <h1 class="text-3xl text-center text-white font-bold">Welcome</h1>
      <form class="max-w-[300px] w-full flex flex-col gap-3" @submit.prevent="onSu">
        
        <div class="flex flex-col gap-1">
          <label>Email</label>
          <input
            class="p-2 rounded-sm border-gray-900 border-2"
            v-model="LoginDetails.email" 
            type="email"
            placeholder="Digite seu email"
          />
        </div>
        <div class="flex flex-col gap-1">
          <label>Senha</label>
          <input
            class="p-2 rounded-sm border-gray-900 border-2"
            v-model="LoginDetails.password"
            type="password"
            placeholder="Digite sua senha"
          />
        </div>
        <p>
          Não tem uma conta?
          <RouterLink class="text-white hover:underline" to="/register"
            >Não tem uma conta?</RouterLink
          >
        </p>
        <button
          class="w-full h-12 bg-gray-700 duration-300 text-white rounded hover:bg-gray-900"
          type="submit"
        >
          Log in
        </button>
      </form>
    </section>
    <section class="w-6/12 flex flex-col justify-center items-center max-md:hidden">
      <img class="h-[600px]" src="../assets/images/hero-login.svg" alt="" />
    </section>
  </main>
</template>

<style scoped></style>
