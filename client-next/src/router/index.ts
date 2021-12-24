import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

import Login from "../views/Login.vue";
import Home from "../views/Home.vue";

const routes: RouteRecordRaw[] = [
  {
    path: "/login",
    name: "Login",
    component: Login,
    meta: { layout: "empty", requiresGuest: true },
  },
  {
    path: "/",
    name: "Home",
    component: Home,
    meta: { requiresAuth: true },
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
});

export default router;
