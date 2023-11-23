import Home from "./pages/Home.vue";
import Login from "./pages/Login.vue";
import Register from "./pages/Register.vue";
import Carteira from "./pages/Carteira.vue";
import Sobre from "./pages/Sobre.vue";
import Contato from "./pages/Contato.vue";
import Comprar from "./pages/Comprar.vue";
import Receber from "./pages/Receber.vue";
import Enviar from "./pages/Enviar.vue";
import Vender from "./pages/Vender.vue";
import Admin from "./pages/Admin.vue";
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
    //so mostre a carteira se o usuario estiver logado
    meta: { requiresAuth: true },
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
    path: "/vender",
    component: Vender,
  },
  {
    path: "/receber",
    component: Receber,
  },
  {
    path: "/enviar",
    component: Enviar,
  },
  {
    path: "/admin",
    component: Admin,
    // so mostre a pagina admin se o usuario for admin
    meta: { requiresAdmin: true },
  }
];

const router = VueRouter.createRouter({
  history: VueRouter.createWebHashHistory(),
  routes, // short for `routes: routes`
});

export default router;
