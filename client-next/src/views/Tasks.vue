<template>
  <div class="tasks">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar mb-4">
        <h1 class="title">Tasks</h1>
        <TheAvatar />
      </div>
      <div class="flex flex-row justify-between items-center">
        <TheTaskDateRange
          v-bind:queryDate="queryDate"
          v-on:update-date-query="queryDate = $event"
        />
        <h3
          class="text-2xl cursor-pointer"
          :class="colorClass"
          @click="showRemaining = !showRemaining"
        >
          {{ totalDuration }} + 1 hrs
        </h3>
      </div>

      <!-- tasks  -->
      <ul class="mt-12 list-none m-0 p-0" v-if="tasks.length">
        <draggable
          v-model="tasks"
          group="people"
          @start="drag = true"
          @end="drag = false"
          item-key="id"
        >
          <template #item="{ element }">
            <TaskItem :task="element" />
          </template>
        </draggable>
      </ul>
    </div>
    <div class="fixed bottom-12 left-1/2 -translate-x-1/2">
      <div v-if="canCreate">
        <router-link
          :to="{ path: '/tasks/create', query: { range: queryDate } }"
          class="create-button rounded-lg"
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
import useColorClass from "../use/colorClass";
import draggable from "vuedraggable";
export default defineComponent({
  components: {
    TheAvatar,
    TheTaskDateRange,
    TaskItem,
    draggable,
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

    const colorClass = computed(() => {
      return useColorClass(totalDuration);
    });

    const drag = ref(false);

    return {
      tasks,
      queryDate,
      showRemaining,
      totalDuration,
      canCreate,
      colorClass,
      drag,
    };
  },
});
</script>

<style lang="scss" scoped>
@import "../assets/time-color-classes.scss";

.create-button {
  color: #fff;
  background-color: $primary;
  border: none;
  font-size: 1.5rem;
  padding: 0.5rem 1.2rem 0.5rem;
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

.ghost {
  box-shadow: 5px 5px 2.5px -1px rgba(0, 0, 0, 0.14);
  opacity: 0.7;
}
</style>
