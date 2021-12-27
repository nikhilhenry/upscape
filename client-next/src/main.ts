import { createApp } from "vue";
import router from "./router";
import App from "./App.vue";
import axios from "axios";
import SidebarLayout from "./layouts/SidebarLayout.vue";
import EmptyLayout from "./layouts/EmptyLayout.vue";
import "./assets/index.css";

// axios defaults
axios.defaults.baseURL = import.meta.env.VITE_API_URL;
axios.defaults.headers.common["auth_token"] =
  localStorage.getItem("auth_token") || "";

const app = createApp(App);

app.component("default-layout", SidebarLayout);
app.component("empty-layout", EmptyLayout);

app.use(router);
app.mount("#app");
