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
            text-skin-base
            font-medium
            rounded-lg
            text-lg
            w-full
            sm:w-auto
            ml-4
            px-5
            py-2.5
            text-center
            bg-skin-button-accent
            hover:bg-skin-button-accent-hover
            focus:bg-skin-button-accent-hover
            bg-opacity-70
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
  @apply w-full text-sm rounded-lg focus:border-skin-fill block p-2.5 bg-skin-elevated bg-opacity-80 border-skin-elevated placeholder:text-skin-muted text-skin-base;
}

.label {
  @apply block my-5 text-xl font-medium text-skin-base;
}

.sub-label {
  @apply block my-3 text-lg font-light text-skin-base;
}
</style>
