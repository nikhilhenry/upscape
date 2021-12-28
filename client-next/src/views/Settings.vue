<template>
  <div class="home">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar">
        <h1 class="title">Settings</h1>
        <TheAvatar />
      </div>
      <section class="mt-16 grid grid-cols-2 gap-x-16">
        <SettingsUsername />
        <SettingsPassword />
      </section>
      <div class="logout fixed bottom-16 left-1/2 -translate-x-1/2">
        <button
          @click="logoutUser"
          type="submit"
          class="
            inline-flex
            text-white
            font-medium
            rounded-lg
            text-lg
            w-full
            sm:w-auto
            ml-4
            px-5
            py-2.5
            text-center
            bg-blue-600
            hover:bg-blue-700
            focus:ring-blue-800
          "
        >
          Logout
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import TheAvatar from "../components/TheAvatar.vue";
import SettingsUsername from "../components/SettingsUsername.vue";
import SettingsPassword from "../components/SettingsPassword.vue";
import useLogout from "../use/logout";

export default defineComponent({
  components: {
    TheAvatar,
    SettingsUsername,
    SettingsPassword,
  },
  setup() {
    const logoutUser = useLogout();

    return {
      logoutUser,
    };
  },
});
</script>

<style lang="scss">
.input-field {
  @apply w-full text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5 bg-gray-700 border-gray-600 placeholder-gray-400 text-white;
}

.label {
  @apply block my-5 text-xl font-medium text-gray-300;
}

.sub-label {
  @apply block my-3 text-lg font-light text-gray-300;
}
</style>
