import { computed, reactive } from "vue";
import {
  postActionable,
  updateActionable,
} from "../services/actionableService";
import { Actionable } from "../types/actionableTypes.interface";

const state = reactive({
  inboxItems: [] as Array<Actionable>,
  weeklyItems: [] as Array<Actionable>,
});

const getters = reactive({
  getInboxItems: computed(() => {
    return state.inboxItems;
  }),
  getWeeklyItems: computed(() => {
    return state.weeklyItems;
  }),
});

const actions = {
  storeInboxItems: (actionables: Actionable[]) => {
    state.inboxItems = actionables;
  },
  storeInboxItem: (actionable: Actionable) => {
    state.inboxItems.push(actionable);
  },
  updateInboxItemList: (inboxItems: Actionable[]) => {
    // update order id with api
    inboxItems.forEach(async (inboxItem, index) => {
      await updateActionable(inboxItem._id!, { id: index });
    });

    state.inboxItems = [];
    state.inboxItems = inboxItems;
  },
  deleteInboxItem: (inboxItemId: string) => {
    const index = state.inboxItems.findIndex(
      (inboxItem) => inboxItem._id === inboxItemId
    );
    // remove task from task array
    state.inboxItems.splice(index, 1);
  },
  storeWeeklyItems: (actionables: Actionable[]) => {
    state.weeklyItems = actionables;
  },
  storeWeeklyItem: (actionable: Actionable) => {
    state.weeklyItems.push(actionable);
  },
  updateWeeklyItem: (updatedActionable: Actionable) => {
    const index = state.weeklyItems.findIndex(
      (weeklyItem) => weeklyItem._id === updatedActionable._id
    );
    state.weeklyItems[index] = updatedActionable;
  },
  updateWeeklyItemList: (weeklyItems: Actionable[]) => {
    // update order id with api
    weeklyItems.forEach(async (weeklyItem, index) => {
      await updateActionable(weeklyItem._id!, { id: index });
    });

    state.weeklyItems = [];
    state.weeklyItems = weeklyItems;
  },
  deleteWeeklyItem: (weeklyItemId: string) => {
    const index = state.weeklyItems.findIndex(
      (weeklyItem) => weeklyItem._id === weeklyItemId
    );
    // remove task from task array
    state.weeklyItems.splice(index, 1);
  },
};

export default { state, getters, ...actions };
