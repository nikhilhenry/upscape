<template>
  <div class="avatar">
    <img :src="imgURL" alt="user_profile_pic" />
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

<style lang="scss" scoped>
.avatar {
  height: 100px !important;
  width: auto;
  display: block !important;
  img {
    height: 100%;
    width: auto;
    border-radius: 35%;
    object-fit: cover;
  }
}
</style>
