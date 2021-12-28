<template>
  <div class="h-24 w-auto block">
    <img
      :src="imgURL"
      alt="user_profile_pic"
      class="h-full w-auto rounded-3xl object-cover"
    />
  </div>
</template>
<script lang="ts">
import { computed, defineComponent } from "vue";
import userStore from "../stores/user";

export default defineComponent({
  setup() {
    const imgURL = computed(() => {
      return userStore.getters.getAvatarURL;
    });

    return {
      imgURL,
    };
  },
});
</script>
