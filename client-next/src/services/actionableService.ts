import axios from "axios";
import { Actionable } from "../types/actionableTypes.interface";

export const getInboxItems = async (): Promise<Array<Actionable>> => {
  try {
    const inboxItems: Actionable[] = await (await axios.get("/api/inbox")).data;
    return inboxItems;
  } catch (error) {
    console.log(error);
    return [];
  }
};

export const getWeeklyItems = async (): Promise<Array<Actionable>> => {
  try {
    const weeklyItems: Actionable[] = await (
      await axios.get("/api/weekly")
    ).data;
    return weeklyItems;
  } catch (error) {
    console.log(error);
    return [];
  }
};

export const postActionable = async (
  payload: Actionable
): Promise<Actionable | null> => {
  const type = payload.inbox ? "inbox" : "weekly";
  try {
    const actionable: Actionable = await (
      await axios.post(`/api/${type}`, payload)
    ).data;
    return actionable;
  } catch (error) {
    console.log(error);
    return null;
  }
};

export const updateActionable = async (
  actionableId: string,
  payload: Actionable
): Promise<Actionable | null> => {
  try {
    const updatedActionable: Actionable = await axios.put(
      `/api/actionable/${actionableId}`,
      payload
    );
    return updatedActionable;
  } catch (error) {
    console.log(error);
    return null;
  }
};

export const deleteActionable = async (
  actionableId: string
): Promise<boolean> => {
  try {
    await (
      await axios.delete(`api/actionable/${actionableId}`)
    ).data;
    return true;
  } catch (error) {
    return false;
  }
};
