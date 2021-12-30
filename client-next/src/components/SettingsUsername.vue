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
          text-skin-base
          font-medium
          rounded-lg
          text-sm
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
          transition-colors
        "
      >
        <svg
          v-if="isLoading"
          class="animate-spin -ml-1 mr-3 h-5 w-5"
          fill="none"
          viewBox="0 0 24 24"
        >
          <circle
            class="opacity-25"
            cx="12"
            cy="12"
            r="10"
            stroke="currentColor"
            stroke-width="4"
          ></circle>
          <path
            class="opacity-75"
            fill="currentColor"
            d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
          ></path>
        </svg>
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
    const isLoading = ref(false);
    const updateName = async () => {
      if (!username) return null;
      isLoading.value = true;
      const response = await updateMeta({ name: username.value });
      if (!response) return null;
      userStore.storeUserMeta(response);
      isLoading.value = false;
    };

    return {
      username,
      updateName,
      isLoading,
    };
  },
});
</script>
