import Home from "./pages/Home.vue";
import Login from "./pages/Login.vue";
import Register from "./pages/Register.vue";
import Carteira from "./pages/Carteira.vue";
import Sobre from "./pages/Sobre.vue";
import Contato from "./pages/Contato.vue";
import Comprar from "./pages/Comprar.vue";
import Receber from "./pages/Receber.vue";
import Enviar from "./pages/Enviar.vue";
import * as VueRouter from "vue-router";
import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    component: Home,
  },
  {
    path: "/login",
    component: Login,
  },
  {
    path: "/register",
    component: Register,
  },
  {
    path: "/carteira",
    component: Carteira,
  },
  {
    path: "/sobre-nos",
    component: Sobre,
  },
  {
    path: "/contato",
    component: Contato,
  },
  {
    path: "/comprar",
    component: Comprar,
  },
  {
    path: "/receber",
    component: Receber,
  },
  {
    path: "/enviar",
    component: Enviar,
  },
];

const router = VueRouter.createRouter({
  history: VueRouter.createWebHashHistory(),
  routes, // short for `routes: routes`
});

export default router;
