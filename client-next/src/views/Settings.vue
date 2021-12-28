<template>
  <div class="home">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar">
        <h1 class="title">Settings</h1>
        <TheAvatar />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import TheAvatar from "../components/TheAvatar.vue";

export default defineComponent({
  components: {
    TheAvatar,
  },
});
</script>
