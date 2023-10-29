<script>
import { defineComponent } from "vue";
import axios from "./api";
import store from "./store";
export default defineComponent({
  name: "App",
  setup() {
    return {
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
  mounted() {
    this.getUser();
  },
});
</script>

<template>
  <header className="bg-zinc-900">
    <nav>
      <img src="" alt="" />
      <Link from="/Market">Market</Link>
      <Link from="/Trade">Trade</Link>
      <Link from="/Contato">Contato</Link>
      <Link from="/Login">Login</Link>
      <Link from="/Register">Sing Up</Link>
    </nav>
  </header>

  <RouterView />
</template>

<style scoped></style>
