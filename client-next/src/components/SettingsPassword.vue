<template>
  <div class="mb-6">
    <form @submit.prevent="updateUserPassword">
      <label class="label">Update Password</label>
      <label class="sub-label">Current Password</label>
      <input
        type="password"
        class="input-field"
        required
        v-model="currentPassword"
      />
      <label class="sub-label">New Password</label>
      <div class="flex">
        <input
          type="password"
          class="input-field"
          required
          v-model="newPassword"
        />
        <button
          @click="updateUserPassword"
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
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { updatePassword } from "../services/userService";
import useLogout from "../use/logout";

const currentPassword = ref("");
const newPassword = ref("");
const isLoading = ref(false);

const logout = useLogout();

const updateUserPassword = async () => {
  if (!currentPassword.value || !newPassword.value) return null;
  isLoading.value = true;
  await updatePassword({
    current_password: currentPassword.value,
    new_password: newPassword.value,
  });
  isLoading.value = false;
  logout();
};
</script>
