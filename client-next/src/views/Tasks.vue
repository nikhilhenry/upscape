<template>
  <div class="tasks">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar">
        <h1 class="title">Tasks</h1>
        <TheAvatar />
      </div>
      <div class="secondary-bar">
        <TheTaskDateRange
          v-bind:queryDate="queryDate"
          v-on:update-date-query="queryDate = $event"
        />
        <h3
          class="total-duration"
          :class="colorClass"
          @click="showRemaining = !showRemaining"
        >
          {{ totalDuration }} + 1 hrs
        </h3>
      </div>

      <!-- tasks  -->
      <ul class="task-list" v-if="tasks.length">
        <transition-group type="transition" name="list-complete">
          <li
            v-for="(task, index) in tasks"
            :key="task._id"
            class="list-complete-item"
          >
            <TaskItem
              :task="task"
              draggable="true"
              @dragstart="dragStart($event, index)"
              @drop="dragDrop($event, index)"
              @dragenter.prevent
              @dragover.prevent
            />
          </li>
        </transition-group>
      </ul>
    </div>
    <div class="fab">
      <div v-if="canCreate">
        <router-link
          :to="{ path: 'tasks/create', query: { range: queryDate } }"
          class="create-button"
          >Create Task</router-link
        >
      </div>
    </div>
    <router-view></router-view>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed, onMounted, watch, toRaw } from "vue";
import TheAvatar from "../components/TheAvatar.vue";
import TheTaskDateRange from "../components/TheTaskDateRange.vue";
import TaskItem from "../components/TaskItem.vue";
import taskStore from "../stores/task";
import { getTasks } from "../services/taskService";
import useDragSort from "../use/dragSort";
import useColorClass from "../use/colorClass";
export default defineComponent({
  components: {
    TheAvatar,
    TheTaskDateRange,
    TaskItem,
  },
  setup() {
    // query tasks tasks
    const isLoaded = computed(() => {
      return taskStore.getters.getTasks.length ? true : false;
    });
    onMounted(async () => {
      if (!isLoaded.value) {
        const tasks = await getTasks(queryDate.value);
        taskStore.storeTasks(tasks);
      }
    });

    const tasks = computed({
      get: () => {
        return taskStore.getters.getTasks;
      },
      set: (taskList) => {
        taskStore.updateTaskList(taskList);
      },
    });
    const totalDuration = computed(() => {
      return showRemaining.value
        ? taskStore.getters.getRemainingTime
        : taskStore.getters.getTotalTime;
    });
    const queryDate = ref("today");
    const showRemaining = ref(true);

    // query new tasks
    watch(queryDate, async (queryDate) => {
      const newTasks = await getTasks(queryDate);
      taskStore.setTaskList(newTasks);
    });

    // create logic
    const canCreate = computed(() => {
      return queryDate.value == "today" || queryDate.value == "tomorrow";
    });

    const dragFunctions = useDragSort(tasks);
    const colorClass = useColorClass(parseFloat(totalDuration.value));

    return {
      tasks,
      queryDate,
      showRemaining,
      totalDuration,
      canCreate,
      colorClass,
      ...dragFunctions,
    };
  },
});
</script>

<style lang="scss" scoped>
@import "../assets/time-color-classes.scss";

.secondary-bar {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;

  .total-duration {
    font-size: 24px;
    cursor: pointer;
  }
}

.task-list {
  margin-top: 3rem;
  list-style: none;
  margin-left: 0;
  padding: 0;
}

.fab {
  position: fixed;
  bottom: 50px;
  left: 50%;
  transform: translate(-50%, 0);
  // right: 50%;
}

.create-button {
  color: #fff;
  background-color: $primary;
  border: none;
  font-size: 1.5rem;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  padding: 0.5rem 1.2rem 0.5rem;
  border-radius: 0.2rem;
  margin: 0;
  text-decoration: none;
  transition: all 200ms ease-in;
  &:hover {
    // display:none;
    background-color: darken($primary, 5%);
  }
}

// draggable
.flip-list-move {
  transition: transform 0.5s;
}

.ghost {
  box-shadow: 5px 5px 2.5px -1px rgba(0, 0, 0, 0.14);
  opacity: 0.7;
}

.list-complete-item {
  transition: all 0.8s ease;
}

.list-complete-enter-from,
.list-complete-leave-to {
  opacity: 0;
  transform: translatey(30px);
}

.list-complete-leave-active {
  position: absolute;
}
</style>
