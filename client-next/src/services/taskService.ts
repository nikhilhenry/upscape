import axios from "axios";
import task from "../stores/task";
import { Task } from "../types/TaskTypes.interface";

export const getTasks = async (): Promise<Array<Task>> => {
  try {
    const tasks: Task[] = await (await axios.get("/api/task")).data;
    return tasks;
  } catch (error) {
    console.log("failed to query tasks");
    return [];
  }
};

export const postTask = async (payload: Task): Promise<Task> => {
  try {
    const task: Task = await (await axios.post("/api/task", payload)).data;
    return task;
  } catch (error) {
    console.log("failed to create task");
    return error.data;
  }
};

export const updateTask = async (
  taskId: string,
  payload: Task
): Promise<Task> => {
  try {
    const updatedTask: Task = await axios.put(`/api/task/${taskId}`, task);
    return updatedTask;
  } catch (error) {
    console.log("task updated failed");
    return payload;
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
