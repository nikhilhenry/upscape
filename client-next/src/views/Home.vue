<template>
  <div class="home">
    <div class="w-full h-full my-4 mx-auto lg:max-w-7xl md:max-w-full">
      <!-- title bar -->
      <div class="title-bar">
        <h1 class="title">
          {{ greeting }}, <br />
          {{ firstName }} 👋
        </h1>
        <TheAvatar />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, ref } from "vue";
import getGreeting from "../functions/getGreeting";
import userStore from "../stores/user";
import TheAvatar from "../components/TheAvatar.vue";

export default defineComponent({
  components: {
    TheAvatar,
  },
  setup() {
    const firstName = computed(() => {
      return userStore.getters.getName;
    });

    const greeting = ref("");

    onMounted(() => {
      greeting.value = getGreeting();
    });

    return {
      greeting,
      firstName,
    };
  },
});
</script>
