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
          inline-flex
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
        <svg
          v-if="isLoading"
          class="animate-spin -ml-1 mr-3 h-5 w-5 text-white"
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
        Submit
      </button>
    </form>
  </ModalView>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { useRoute } from "vue-router";
import ModalView from "../components/ModalView.vue";
import { postTask } from "../services/taskService";
import taskStore from "../stores/task";
import { Task } from "../types/taskTypes.interface";
export default defineComponent({
  components: { ModalView },
  setup() {
    const route = useRoute();
    const isLoading = ref(false);

    const submitForm = async (event: any) => {
      const formData: any = new FormData(event.target);
      const { taskName, taskDuration, highlight } =
        Object.fromEntries(formData);

      // create task
      const taskRequest: Task = {
        name: taskName,
        duration: parseInt(taskDuration),
        highlight: highlight ? true : false,
        id: taskStore.getters.getTasks.length,
        completed: false,
        is_tomorrow: route.query.range == "tomorrow" ? true : false,
      };

      isLoading.value = !isLoading.value;
      const task = await postTask(taskRequest);
      isLoading.value = !isLoading.value;
      if (task) taskStore.storeTask(task);
    };

    return {
      submitForm,
      isLoading,
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
