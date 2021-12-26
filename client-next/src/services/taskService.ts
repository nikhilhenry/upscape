import axios from "axios";
import task from "../stores/task";
import { Task, TaskUpdateRequest } from "../types/TaskTypes.interface";

export const getTasks = async (range = "today"): Promise<Array<Task>> => {
  try {
    const tasks: Task[] = await (
      await axios.get(`/api/task?range=${range}`)
    ).data;
    return tasks;
  } catch (error) {
    console.log("failed to query tasks");
    return [];
  }
};

export const postTask = async (payload: Task): Promise<Task | null> => {
  try {
    const task: Task = await (await axios.post("/api/task", payload)).data;
    return task;
  } catch (error) {
    console.log("failed to create task");
    return null;
  }
};

export const updateTask = async (
  taskId: string,
  payload: TaskUpdateRequest
): Promise<Task | null> => {
  try {
    const updatedTask: Task = await axios.put(`/api/task/${taskId}`, task);
    return updatedTask;
  } catch (error) {
    console.log("task updated failed");
    return null;
  }
};

export const deleteTask = async (taskId: string): Promise<boolean> => {
  try {
    await (
      await axios.delete(`api/task/${taskId}`)
    ).data;
    return true;
  } catch (error) {
    return false;
  }
};
