import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Login from "../views/Login.vue";
import Tasks from "../views/Tasks.vue";

import NavbarLayout from "../layouts/NavbarLayout.vue";

import store from "@/store";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
    meta: {
      requiresAuth: true,
      layout: NavbarLayout,
    },
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
    meta: {
      requiresGuest: true,
    },
  },
  {
    path: "/tasks",
    name: "Tasks",
    component: Tasks,
    meta: {
      requiresAuth: true,
      layout: NavbarLayout,
    },
    children: [
      {
        path: "create",
        component: () => import("../views/CreateTask.vue"),
        meta: {
          requiresAuth: true,
          layout: NavbarLayout,
        },
      },
    ],
  },
  {
    path: "/settings",
    name: "Settings",
    component: () => import("../views/Settings.vue"),
    meta: {
      requriesAuth: true,
      layout: NavbarLayout,
    },
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
  },
];

const router = new VueRouter({
  base: process.env.BASE_URL,
  routes,
});

// router gaurds
router.beforeEach((to, from, next) => {
  // get access token from store

  const isAuthenticated = store.getters["user/getAuthenticatedState"];

  // if auth is required
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    // check if user is authenticated
    if (isAuthenticated) next();
    // if not authenticated
    else {
      next({
        path: "/login",
        query: {
          redirect: to.fullPath,
        },
      });
    }
  }
  // if guest is required
  else if (to.matched.some((record) => record.meta.requiresGuest)) {
    // check if user is not authenticated
    if (!isAuthenticated) next();
    // if authenticated
    else {
      next({
        path: "/",
        query: {
          redirect: to.fullPath,
        },
      });
    }
  }

  // if nothing required
  else {
    next();
  }
});

export default router;
