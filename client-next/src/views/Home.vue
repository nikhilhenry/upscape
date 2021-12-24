<template>
  <div class="home">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar">
        <h1 class="title">
          {{ greeting }}, <br />
          {{ firstName }} 👋
        </h1>
        <Avatar />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, ref } from "vue";
import getGreeting from "../functions/getGreeting";
import userStore from "../stores/user";

export default defineComponent({
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
