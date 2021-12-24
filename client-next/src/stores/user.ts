import axios from "axios";
import { computed, reactive } from "vue";
import { User } from "../types/userTypes.interface";

const state = reactive({
  token: localStorage.getItem("auth_token"),
  error: "",
  avatarURL: "",
  name: "",
});

const getters = reactive({
  isLoggedIn: computed(() => {
    return state.token ? true : false;
  }),
  getToken: () => {
    return state.token;
  },
  getAvatarURL: computed(() => {
    return state.avatarURL;
  }),
  getName: computed(() => {
    return state.name;
  }),
});

const actions = {
  storeUserMeta: (userMeta: User) => {
    state.avatarURL = userMeta.img_url;
    state.name = userMeta.name;
  },
  storeToken: async (token: string) => {
    // store token in localStorage
    localStorage.setItem("auth_token", token);

    // set axios token
    // set header globally
    axios.defaults.headers.common["auth_token"] = token;

    state.token = token;
  },
  logout: async () => {
    // remove token from localStorage
    localStorage.removeItem("auth_token");
    state.token = "";
  },
};

export default { state, getters, ...actions };
