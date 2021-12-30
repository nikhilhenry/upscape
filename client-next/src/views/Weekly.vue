<template>
  <div class="tasks">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar mb-4">
        <h1 class="title">Weekly Action Items</h1>
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
    TheAvatar: TheAvatar,
  },
});
</script>

<style lang="scss" scoped>
.flip-list-move {
  transition: transform 0.5s;
}

.ghost {
  box-shadow: 5px 5px 2.5px -1px rgba(0, 0, 0, 0.14);
  opacity: 0.7;
}

.list-complete-item {
  transition: all 0.8s ease;
}

.list-complete-enter-from,
.list-complete-leave-to {
  opacity: 0;
  transform: translatey(30px);
}

.list-complete-leave-active {
  position: absolute;
}
</style>
