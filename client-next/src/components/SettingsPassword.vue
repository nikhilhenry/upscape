<template>
  <div class="mb-6">
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

<script setup lang="ts">
import { ref } from "vue";
import { updatePassword } from "../services/userService";
import useLogout from "../use/logout";

const currentPassword = ref("");
const newPassword = ref("");

const logout = useLogout();

const updateUserPassword = async () => {
  if (!currentPassword.value || !newPassword.value) return null;
  await updatePassword({
    current_password: currentPassword.value,
    new_password: newPassword.value,
  });
  logout();
};
</script>
