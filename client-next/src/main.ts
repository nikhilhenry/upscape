import { createApp } from "vue";
import router from "./router";
import App from "./App.vue";
import axios from "axios";

// axios defaults
axios.defaults.baseURL = import.meta.env.VITE_API_URL;
axios.defaults.headers.common["auth_token"] =
  localStorage.getItem("access_token") || "";

const app = createApp(App);

app.use(router);
app.mount("#app");
