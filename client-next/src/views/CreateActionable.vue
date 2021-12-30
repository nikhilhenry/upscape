<template>
  <ModalView :modalTitle="`Create ${itemType}`">
    <form @submit.prevent="submitForm">
      <div class="mb-6">
        <label class="label">Actionable Name</label>
        <input
          type="text"
          class="input-field"
          placeholder="Eat sushi 🍣"
          required
          name="actionableName"
        />
      </div>
      <button
        type="submit"
        class="
          inline-flex
          text-skin-base
          font-medium
          rounded-lg
          text-sm
          w-full
          sm:w-auto
          px-5
          py-2.5
          text-center
          bg-skin-button-accent
          hover:bg-skin-button-accent-hover
          focus:bg-skin-button-accent-hover
          transition-colors
        "
      >
        <svg
          v-if="isLoading"
          class="animate-spin -ml-1 mr-3 h-5 w-5"
          fill="none"
          viewBox="0 0 24 24"
        >
          <circle
            class="opacity-25"
            cx="12"
            cy="12"
            r="10"
            stroke="currentColor"
            stroke-width="4"
          ></circle>
          <path
            class="opacity-75"
            fill="currentColor"
            d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
          ></path>
        </svg>
        Submit
      </button>
    </form>
  </ModalView>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { useRoute } from "vue-router";
import ModalView from "../components/ModalView.vue";
import { postActionable } from "../services/actionableService";
import actionableStore from "../stores/actionable";
import { Actionable } from "../types/actionableTypes.interface";
export default defineComponent({
  components: { ModalView },
  setup() {
    const route = useRoute();
    const isLoading = ref(false);
    const itemType = route.query.type;

    const submitForm = async (event: any) => {
      const formData: any = new FormData(event.target);
      const { actionableName } = Object.fromEntries(formData);

      // create task
      const actionableRequest: Actionable = {
        name: actionableName,
        id:
          itemType == "inbox"
            ? actionableStore.getters.getInboxItems.length
            : actionableStore.getters.getWeeklyItems.length,
        completed: false,
        inbox: itemType == "inbox" ? true : false,
      };

      isLoading.value = !isLoading.value;
      const actionable = await postActionable(actionableRequest);
      isLoading.value = !isLoading.value;
      if (actionable) itemType == "inbox"
            ? actionableStore.storeInboxItem(actionable)
            : actionableStore.storeWeeklyItem(actionable),
    }

    return {
      submitForm,
      isLoading,
      itemType,
    };
  },
});
</script>

<style lang="scss">
.input-field {
  @apply w-full text-sm rounded-lg focus:border-skin-fill block p-2.5 bg-skin-elevated bg-opacity-80 border-skin-elevated placeholder:text-skin-muted text-skin-base;
}

.label {
  @apply block my-5 text-xl font-medium text-skin-base;
}
</style>
