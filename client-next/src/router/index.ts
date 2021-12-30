import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import userStore from "../stores/user";

import Login from "../views/Login.vue";
import Home from "../views/Home.vue";
import Tasks from "../views/Tasks.vue";

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
  {
    path: "/tasks",
    name: "Tasks",
    component: Tasks,
    meta: { requiresAuth: true },
    children: [
      {
        path: "create",
        component: () => import("../views/CreateTask.vue"),
        meta: {
          requiresAuth: true,
        },
      },
    ],
  },
  {
    path: "/weekly",
    name: "Weekly",
    component: () => import("../views/Weekly.vue"),
    meta: { requiresAuth: true },
  },
  {
    path: "/settings",
    name: "Settings",
    component: () => import("../views/Settings.vue"),
    meta: { requiresAuth: true },
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
});

// route guards
router.beforeEach((to, from, next) => {
  // get access token from store

  const isAuthenticated = userStore.getters.isLoggedIn;

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
