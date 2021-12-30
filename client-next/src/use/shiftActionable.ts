import {
  deleteActionable,
  updateActionable,
} from "../services/actionableService";
import { Actionable } from "../types/actionableTypes.interface";
import actionableStore from "../stores/actionable";

export default function useShiftActionable() {
  return async (actionable: Actionable) => {
    if (!actionable._id) return;
    // create new actionable
    const shiftedActionable = await updateActionable(actionable._id, {
      inbox: !actionable.inbox,
    });

    if (shiftedActionable == null) return;
    if (shiftedActionable.inbox) {
      actionableStore.deleteWeeklyItem(actionable._id);
    } else {
      actionableStore.deleteInboxItem(actionable._id);
    }
  };
}
