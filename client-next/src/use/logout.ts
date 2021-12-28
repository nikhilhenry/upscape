import userStore from "../stores/user";
import { useRouter } from "vue-router";

export default function useLogout() {
  const router = useRouter();
  const logoutUser = () => {
    userStore.logout();
    router.push({ name: "Login" });
  };

  return logoutUser;
}
