<template>
  <div class="mb-6">
    <label class="label">Change Username</label>
    <div class="flex">
      <input
        type="text"
        class="input-field"
        placeholder="New Name"
        required
        name="taskName"
        v-model="username"
      />
      <button
        @click="updateName"
        type="submit"
        class="
          inline-flex
          text-white
          font-medium
          rounded-lg
          text-sm
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
        Update
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { updateMeta } from "../services/userService";
import userStore from "../stores/user";

export default defineComponent({
  setup() {
    const username = ref("");
    const updateName = async () => {
      if (!username) return null;
      const response = await updateMeta({ name: username.value });
      if (!response) return null;
      userStore.storeUserMeta(response);
    };

    return {
      username,
      updateName,
    };
  },
});
</script>
