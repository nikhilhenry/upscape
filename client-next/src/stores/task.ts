import { computed, reactive } from "vue";
import { Task } from "../types/TaskTypes.interface";

const state = reactive({
  tasks: [] as Array<Task>,
  tasksLoaded: false,
});

const getters = reactive({
  getTasks: computed(() => {
    return state.tasks;
  }),
  getRemainingTime: computed(() => {
    const tasks = state.tasks;

    if (!tasks) {
      return 0;
    }

    let totalTime = 0;

    // filter tasks for uncompleted
    const uncompletedTasks = tasks.filter((task) => {
      return !task.completed;
    });

    // loop through all uncompleted tasks and add time
    uncompletedTasks.forEach((task) => {
      totalTime += task.duration;
    });

    return (totalTime / 60).toFixed(2);
  }),
  getTotalTime: computed(() => {
    const tasks = state.tasks;

    if (!tasks) {
      return 0;
    }

    let totalTime = 0;

    tasks.forEach((task) => {
      totalTime += task.duration;
    });
    // round time to 2 dp in hours
    return (totalTime / 60).toFixed(2);
  }),
});

const actions = {
  storeTasks: (tasks: Task[]) => {
    state.tasks = state.tasks.concat(tasks);
    if (tasks.length) state.tasksLoaded = true;
  },
  storeTask: (task: Task) => {
    state.tasks.push(task);
  },
  deleteTask: (taskId: string) => {
    const index = state.tasks.findIndex((task) => task._id === taskId);
    // remove task from task array
    state.tasks.splice(index, 1);
  },
};

export default { state, getters, ...actions };
