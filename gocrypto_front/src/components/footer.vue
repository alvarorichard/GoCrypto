<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "Footer",
  setup() {
    return {};
  },
});
</script>

<template>
  <footer class="w-full flex flex-col gap-2 p-2 bg-transparent backdrop-blur-2xl">
    <div class="w-full flex justify-between items-center">
      <RouterLink to="/">
        <img
          class="h-14 scale-125 duration-300 hover:scale-150"
          src="https://i.imgur.com/CAp9EhS.png"
          alt="Imagem logo"
        />
      </RouterLink>
      <h1 class="text-3xl">Footer</h1>
    </div>
    <hr class="opacity-30" />
    <div class="flex flex-col">
      <RouterLink class="max-w-[120px] flex items-center gap-2 hover:opacity-80" to="/">
        <img src="../assets/icons/instagram-icon.svg" width="28" alt="" />
        <p>Instagram</p>
      </RouterLink>
      <RouterLink class="max-w-[120px] flex items-center gap-2 hover:opacity-80" to="/">
        <img src="../assets/icons/twitter-icon.svg" width="28" alt="" />
        <p>Twitter</p>
      </RouterLink>
      <RouterLink class="max-w-[120px] flex items-center gap-2 hover:opacity-80" to="/">
        <img src="../assets/icons/facebook-icon.svg" width="28" alt="" />
        <p>Facebook</p>
      </RouterLink>
    </div>
  </footer>
</template>
