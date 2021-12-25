<template>
  <div
    class="task"
    v-bind:class="{ 'is-completed': completed }"
    @dblclick="deleteTask()"
  >
    <div class="wrapper">
      <div class="left">
        <p class="title">{{ task.name }}</p>
      </div>
      <div class="right">
        <i class="fas fa-star" v-if="task.highlight"></i>
        <span class="duration">{{ task.duration }} MIN</span>
        <input
          id="c1"
          type="checkbox"
          class="complete"
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
import { deleteTask as deleteTaskById } from "../services/taskService";
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
    const deleteTask = async () => {
      const success = await deleteTaskById(props.task._id || "");
      if (success) taskStore.deleteTask(props.task._id || "");
    };

    return {
      completed,
      deleteTask,
    };
  },
});
</script>

<style lang="scss" scoped>
@import "@/assets/toggles.scss";

$flex-gap: 2rem;

.wrapper {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: $secondary;
  padding: 1rem 1.5rem;
  border-bottom: 3px solid $background;
  border-radius: 1rem;

  margin-bottom: 0.5rem;

  .title {
    font-size: 1.4rem;
    margin: 1rem 0 0.5rem 0rem;
  }

  .date {
    font-size: 1em;
    color: rgba($text-secondary, $alpha: 0.8);
    font-weight: 100;
    display: block;
    margin-bottom: 0.5rem;
  }

  .duration {
    font-size: 1.2rem;
    color: $text-secondary;
    font-weight: 400;
    margin-left: $flex-gap;
  }

  .la-step-forward {
    color: $primary;
    font-size: 2rem;
    margin-left: $flex-gap;
    filter: grayscale(100%) opacity(0.7);
    transition: all 200ms ease-in;
    &:hover {
      filter: none;
    }
  }

  .active {
    filter: none;
  }

  .fa-star {
    color: #ffd31d;
    font-size: 1.4rem;
    margin-left: $flex-gap;
  }

  .complete {
    margin-left: $flex-gap;
    transform: scale(1.5);
  }

  .right {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    min-width: 300px;
  }
}

.tags {
  list-style: none;
  margin: 0;
  margin-left: $flex-gap;
  padding: 0;
  display: flex;
  flex-direction: row;
  .tag {
    color: #fff;
    padding: 0.1rem 0.5rem 0.1rem;
    margin-left: 0.5rem;
    border-radius: 0.3rem;
  }
}

.is-completed {
  filter: opacity(0.5);
}
</style>
