import axios from "axios";
import { Router } from "vue-router";
import userStore from "../stores/user";

export const useErrorInterceptor = function (router: Router) {
  axios.interceptors.response.use(undefined, (error) => {
    const statusCode = error.response ? error.response.status : null;
    if (statusCode === 401) {
      console.log("unauthorized");
      // logout user
      userStore.logout();
      router.push({
        name: "Login",
      });
    }
  });
};
