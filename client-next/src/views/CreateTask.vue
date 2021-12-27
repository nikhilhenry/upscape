<template>
  <ModalView :modalTitle="'Create Task'">
    <form>
      <div class="mb-6">
        <label class="label">Task Name</label>
        <input
          type="text"
          class="input-field"
          placeholder="Eat sushi 🍣"
          required
        />
      </div>
      <div class="mb-6">
        <label class="label">Task Duration</label>
        <input
          type="number"
          class="input-field"
          placeholder="name@flowbite.com"
          required
        />
      </div>
      <button
        type="submit"
        class="
          text-white
          bg-blue-700
          hover:bg-blue-800
          focus:ring-4 focus:ring-blue-300
          font-medium
          rounded-lg
          text-sm
          w-full
          sm:w-auto
          px-5
          py-2.5
          text-center
          dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800
        "
      >
        Submit
      </button>
    </form>
  </ModalView>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import ModalView from "../components/ModalView.vue";
export default defineComponent({
  components: { ModalView },
});
</script>

<style lang="scss">
.input-field {
  @apply bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500;
}

.label {
  @apply block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300;
}
</style>
