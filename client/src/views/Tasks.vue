<template>
  <div class="tasks">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar">
        <h1 class="title">Tasks</h1>
        <Avatar />
      </div>
      <div class="secondary-bar">
        <TasksDateRange
          v-bind:queryDate="queryDate"
          v-on:update:query="queryDate = $event"
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
      <ul class="task-list">
        <draggable v-model="tasks" ghost-class="ghost">
          <transition-group type="transistion" name="flip-list">
            <li v-for="task in tasks" :key="task._id">
              <TaskItem :task="task" @startTimer="openTimer" />
            </li>
          </transition-group>
        </draggable>
      </ul>
    </div>
    <TaskTimer v-if="timerActive" :taskTime="taskTime" @stop="closeTimer" />
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

<script>
import Avatar from "@/components/Avatar";
import TasksDateRange from "@/components/TasksDateRange";
import TaskItem from "@/components/TaskItem";
import tags from "@/mixins/tags.js";
import TaskTimer from "@/components/TaskTimer";

import draggable from "vuedraggable";

import getTasks from "@/api/taskGet";
import { mapGetters, mapActions } from "vuex";

export default {
  name: "Tasks",
  components: {
    Avatar,
    TasksDateRange,
    TaskItem,
    TaskTimer,
    draggable,
  },
  mixins: [tags],
  computed: {
    isLoaded() {
      return this.$store.state.task.tasksLoaded;
    },
    canCreate() {
      if (this.queryDate == "today" || this.queryDate == "tomorrow") {
        return true;
      }
      return false;
    },
    tasks: {
      get() {
        return this.$store.state.task.tasks;
      },
      set(val) {
        this.updateTaskList(val);
      },
    },
    totalDuration() {
      // if need to show remaining
      if (this.showRemaining) return this.totalRemainingDuration;
      // else show total completed
      else return this.totalCompletedDuration;
    },
    colorClass() {
      if (this.totalDuration < 2) return "light";
      if (this.totalDuration < 4.5) return "medium";
      if (this.totalDuration < 6) return "medium-heavy";
      if (this.totalDuration >= 6) return "heavy";
      return "none";
    },
    ...mapGetters({
      totalRemainingDuration: "task/getRemainingTime",
      totalCompletedDuration: "task/getTotalTime",
    }),
  },
  data() {
    return {
      queryDate: "today",
      showRemaining: true,
      timerActive: false,
    };
  },
  watch: {
    async queryDate() {
      // query for new task
      const newTaskList = await this.queryTasks();
      // set newTaskList in store
      this.setNewTaskList(newTaskList);
    },
    canCreate: function(val) {
      // if user cannot create then show completed time
      if (val) this.showRemaining = true;
      else this.showRemaining = false;
    },
  },
  methods: {
    openTimer: function(time) {
      this.taskTime = time;
      this.timerActive = true;
    },
    closeTimer: function() {
      console.log("close");
      this.timerActive = false;
    },
    queryTasks: async function() {
      const tasks = await getTasks(this.queryDate);
      // this.tasks = tasks
      return tasks;
    },
    saveTasksToStore: function(tasks) {
      this.storeTasks(tasks);
    },
    ...mapActions({
      storeTasks: "task/storeTasks",
      setNewTaskList: "task/setNewTaskList",
      updateTaskList: "task/updateTaskList",
    }),
  },
  async mounted() {
    if (!this.isLoaded) {
      //  clear if anything present in store
      this.$store.dispatch("task/setNewTaskList", []);
      const tasks = await this.queryTasks();
      this.saveTasksToStore(tasks);
    }
  },
};
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
</style>
