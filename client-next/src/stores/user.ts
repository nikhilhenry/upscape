import axios from "axios";
import { computed, reactive } from "vue";

const state = reactive({
  token: localStorage.getItem("access_token"),
  error: "",
});

const getters = reactive({
  isLoggedIn: computed(() => {
    return state.token ? true : false;
  }),
  getToken: () => {
    return state.token;
  },
});

const actions = {
  storeToken: async (token: string) => {
    // store token in localStorage
    localStorage.setItem("access_token", token);

    // set axios token
    // set header globally
    axios.defaults.headers.common["access_token"] = token;

    state.token = token;
  },
  logout: async () => {
    // remove token from localStorage
    localStorage.removeItem("access_token");
    state.token = "";
  },
};

export default { state, getters, ...actions };
