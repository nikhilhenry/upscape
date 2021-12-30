<template>
  <div class="inboxItems">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar mb-4">
        <h1 class="title">Inbox Action Items</h1>
        <TheAvatar />
      </div>
      <!-- Inbox  -->
      <ul class="mt-16 list-none m-0 p-0" v-if="inboxItems.length">
        <transition-group type="transition" name="list-complete">
          <li
            v-for="(inboxItem, index) in inboxItems"
            :key="inboxItem._id"
            class="list-complete-item"
          >
            <ActionableItem
              :actionable="inboxItem"
              draggable="true"
              @dragstart="dragStart($event, index)"
              @drop="dragDrop($event, index)"
              @dragenter.prevent
              @dragover.prevent
            />
          </li>
        </transition-group>
      </ul>
    </div>
    <div class="fixed bottom-12 left-1/2 -translate-x-1/2">
      <router-link
        :to="{ path: '/inbox/create', query: { type: 'inbox' } }"
        class="create-button rounded-lg"
        >Create Action Item</router-link
      >
    </div>
    <router-view></router-view>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted } from "vue";
import ActionableItem from "../components/ActionableItem.vue";
import TheAvatar from "../components/TheAvatar.vue";
import { getInboxItems } from "../services/actionableService";
import actionableStore from "../stores/actionable";
import useDragSort from "../use/dragSort";

export default defineComponent({
  components: {
    TheAvatar: TheAvatar,
    ActionableItem: ActionableItem,
  },
  setup() {
    const inboxItems = computed({
      get: () => {
        return actionableStore.getters.getInboxItems;
      },
      set: (actionableItems) => {
        actionableStore.updateInboxItemList(actionableItems);
      },
    });

    onMounted(async () => {
      const inboxItems = await getInboxItems();
      actionableStore.storeInboxItems(inboxItems);
    });

    const dragFunctions = useDragSort(inboxItems);

    return {
      inboxItems,
      ...dragFunctions,
    };
  },
});
</script>

<style lang="scss" scoped>
.create-button {
  color: #fff;
  background-color: $primary;
  border: none;
  font-size: 1.5rem;
  padding: 0.5rem 1.2rem 0.5rem;
  margin: 0;
  text-decoration: none;
  transition: all 200ms ease-in;
  &:hover {
    // display:none;
    background-color: darken($primary, 5%);
  }
}
.flip-list-move {
  transition: transform 0.5s;
}

.ghost {
  box-shadow: 5px 5px 2.5px -1px rgba(0, 0, 0, 0.14);
  opacity: 0.7;
}

.list-complete-item {
  transition: all 0.8s ease;
}

.list-complete-enter-from,
.list-complete-leave-to {
  opacity: 0;
  transform: translatey(30px);
}

.list-complete-leave-active {
  position: absolute;
}
</style>
