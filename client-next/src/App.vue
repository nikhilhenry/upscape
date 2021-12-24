<template>
  <div id="app">
    <component :is="layout"> <router-view /></component>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, onMounted } from "vue";
import { useRouter } from "vue-router";
import { getMeta } from "./services/userService";
import userStore from "./stores/user";

const defaultLayout = "default";

export default defineComponent({
  setup() {
    const { currentRoute } = useRouter();

    const layout = computed(
      () => `${currentRoute.value.meta.layout || defaultLayout}-layout`
    );

    // load metad details on mount
    onMounted(async () => {
      const userMeta = await getMeta();
      userStore.storeUserMeta(userMeta);
    });

    return {
      layout,
    };
  },
});
</script>

<style lang="scss">
:root {
  background: $background;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: $text-primary;
  background: $background;
  margin: 0;
  padding: 0;
}

body {
  margin: 0;
  padding: 0;
}
</style>
