<template>
  <div
    class="actionable"
    v-bind:class="{ 'is-completed': completed }"
    @dblclick="deleteActionable()"
  >
    <div
      class="
        flex
        justify-between
        items-center
        bg-skin-elevated
        px-6
        py-6
        rounded-2xl
        mb-2
        text-skin-base
      "
    >
      <div class="left">
        <p class="text-2xl">{{ actionable.name }}</p>
      </div>
      <div class="flex items-center justify-end">
        <input
          v-if="!actionable.inbox"
          id="c1"
          type="checkbox"
          class="scale-150"
          v-model="completed"
          @click="completeActionable"
        />
        <ActionableItemMenu @select-option="" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, PropType, ref, toRaw } from "vue";
import { Actionable } from "../types/actionableTypes.interface";
import {
  deleteActionable as deleteActionableById,
  updateActionable as updatedActionableById,
} from "../services/actionableService";
import actionableStore from "../stores/actionable";

export default defineComponent({
  props: {
    actionable: { type: Object as PropType<Actionable>, required: true },
  },
  setup(props) {
    const completed = ref(false);
    const itemType = ref("");

    onMounted(() => {
      completed.value = props.actionable.completed;
      itemType.value = props.actionable.inbox ? "inbox" : "weekly";
    });

    // service functions
    const completeActionable = async () => {
      const updatedActionable = await updatedActionableById(
        props.actionable._id || "",
        {
          completed: !props.actionable.completed,
        }
      );
      if (updatedActionable)
        actionableStore.updateWeeklyItem(updatedActionable);
    };

    const deleteActionable = async () => {
      const success = await deleteActionableById(props.actionable._id || "");
      if (success) {
        if (itemType.value !== "inbox")
          actionableStore.deleteInboxItem(props.actionable._id || "");
        actionableStore.deleteWeeklyItem(props.actionable._id || "");
      }
    };

    return {
      completed,
      deleteActionable,
      completeActionable,
    };
  },
});
</script>

<style lang="scss" scoped>
@import "../assets/toggles.scss";

.is-completed {
  filter: opacity(0.5);
}
</style>
