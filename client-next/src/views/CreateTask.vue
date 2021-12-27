<template>
  <ModalView :modalTitle="'Create Task'">
    <form @submit.prevent="submitForm">
      <div class="mb-6">
        <label class="label">Task Name</label>
        <input
          type="text"
          class="input-field"
          placeholder="Eat sushi 🍣"
          required
          name="taskName"
        />
      </div>
      <div class="mb-6">
        <label class="label">Task Duration</label>
        <input
          type="number"
          class="input-field"
          placeholder="0"
          required
          name="taskDuration"
        />
      </div>
      <div class="flex items-start mb-6">
        <div class="flex items-center h-5">
          <input
            type="checkbox"
            class="
              w-4
              h-4
              rounded
              border
              bg-gray-700
              border-gray-600
              focus:ring-offset-sky-500
            "
            name="highlight"
          />
        </div>
        <div class="ml-3 text-sm">
          <label for="remember" class="font-medium text-gray-300"
            >Highlight of the Day</label
          >
        </div>
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
  setup() {
    const submitForm = (event: any) => {
      const formData: any = new FormData(event.target);
      const { taskName, taskDuration, highlight } =
        Object.fromEntries(formData);
    };

    return {
      submitForm,
    };
  },
});
</script>

<style lang="scss">
.input-field {
  @apply bg-gray-50 border border-gray-300 text-gray-300 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500;
}

.label {
  @apply block mb-2 text-sm font-medium text-gray-300;
}
</style>
