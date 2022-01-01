import axios from "axios";
import { useRoute, useRouter } from "vue-router";

export default function () {
  const router = useRouter();
  const route = useRoute();

  axios.interceptors.response.use(undefined, (error) => {
    const statusCode = error.response ? error.response.status : null;
    if (statusCode === 401) {
      router.push({ name: "Login", query: { redirect: route.fullPath } });
    }
  });
}
