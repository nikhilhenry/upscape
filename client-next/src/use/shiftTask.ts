import { deleteTask, postTask } from "../services/taskService";
import { Task } from "../types/taskTypes.interface";
import taskStore from "../stores/task";

export default function useShiftTask() {
  return async (task: Task) => {
    const { _id, created_at, completed_at, ...newTask } = task;
    if (_id == null) return;
    newTask.is_tomorrow = true;

    // create new task
    const shiftedTask = await postTask(newTask);

    if (shiftedTask == null) return;

    const success = await deleteTask(_id);
    if (success) taskStore.deleteTask(_id);
  };
}
