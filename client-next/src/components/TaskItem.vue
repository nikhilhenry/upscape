<template>
  <div
    class="task text-skin-base"
    v-bind:class="{ 'is-completed': completed }"
    @dblclick="deleteTask()"
  >
    <div
      class="
        flex
        justify-between
        items-center
        bg-background-800
        px-6
        py-6
        rounded-2xl
        mb-2
      "
    >
      <div class="left">
        <p class="text-2xl">{{ task.name }}</p>
      </div>
      <div class="flex items-center justify-end">
        <i
          class="fas fa-star text-xl text-yellow-300"
          v-if="task.highlight"
        ></i>
        <span class="text-xl font-normal mx-8">{{ task.duration }} MIN</span>
        <input
          id="c1"
          type="checkbox"
          class="scale-150"
          v-model="completed"
          @click="completeTask"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, PropType, ref } from "vue";
import { Task } from "../types/taskTypes.interface";
import {
  deleteTask as deleteTaskById,
  updateTask as updatedTaskById,
} from "../services/taskService";
import taskStore from "../stores/task";

export default defineComponent({
  props: {
    task: { type: Object as PropType<Task>, required: true },
  },
  setup(props) {
    const completed = ref(false);

    onMounted(() => {
      completed.value = props.task.completed;
    });

    // service functions
    const completeTask = async () => {
      const updatedTask = await updatedTaskById(props.task._id || "", {
        completed: !props.task.completed,
      });
      if (updatedTask) taskStore.updateTask(updatedTask);
    };

    const deleteTask = async () => {
      const success = await deleteTaskById(props.task._id || "");
      if (success) taskStore.deleteTask(props.task._id || "");
    };

    return {
      completed,
      deleteTask,
      completeTask,
    };
  },
});
</script>

<style lang="scss" scoped>
@import "../assets/toggles.scss";

.is-completed {
  filter: opacity(0.5);
}
</style>
