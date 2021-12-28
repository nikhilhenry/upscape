<template>
  <div class="home">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar">
        <h1 class="title">Settings</h1>
        <TheAvatar />
      </div>
      <section class="mt-16 grid grid-cols-2 gap-x-16">
        <div class="mb-6">
          <label class="label">Change Username</label>
          <div class="flex">
            <input
              type="text"
              class="input-field"
              placeholder="New Name"
              required
              name="taskName"
            />
            <button
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
        <div class="mb-6">
          <label class="label">Update Password</label>
          <label class="sub-label">Current Password</label>
          <input
            type="text"
            class="input-field"
            placeholder="New Name"
            required
            name="taskName"
          />
          <label class="sub-label">New Password</label>
          <div class="flex">
            <input
              type="text"
              class="input-field"
              placeholder="New Name"
              required
              name="taskName"
            />
            <button
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
      </section>
      <div class="logout fixed bottom-16 left-1/2 -translate-x-1/2">
        <button
          type="submit"
          class="
            inline-flex
            text-white
            font-medium
            rounded-lg
            text-lg
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
          Logout
        </button>
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

<style lang="scss">
.input-field {
  @apply w-full text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5 bg-gray-700 border-gray-600 placeholder-gray-400 text-white;
}

.label {
  @apply block my-5 text-xl font-medium text-gray-300;
}

.sub-label {
  @apply block my-3 text-lg font-light text-gray-300;
}
</style>
