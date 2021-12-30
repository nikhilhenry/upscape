import { deleteTask, postTask } from "../services/taskService";
import { Task } from "../types/TaskTypes.interface";

export default function useShiftTask(task: Task) {
  return async () => {
    if (!task._id) return;
    const newTask = task;
    // set to tomorrow
    newTask.is_tomorrow = true;

    // create new task
    const shiftedTask = await postTask(newTask);
    if (shiftedTask) {
      // delete old task
      await deleteTask(task._id);
    }
  };
}
