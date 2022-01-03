<template>
  <div class="weeklyItems">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar mb-4">
        <h1 class="title">Weekly Action Items</h1>
        <TheAvatar />
      </div>
      <!-- Weekly  -->
      <ul class="mt-16 list-none m-0 p-0" v-if="weeklyItems.length">
        <transition-group type="transition" name="list-complete">
          <li
            v-for="(weeklyItem, index) in weeklyItems"
            :key="weeklyItem._id"
            class="list-complete-item"
          >
            <ActionableItem
              :actionable="weeklyItem"
              draggable="true"
              @dragstart="dragStart($event, index)"
              @drop="dragDrop($event, index)"
              @dragenter.prevent
              @dragover.prevent
              style="-webkit-user-drag: element"
            />
          </li>
        </transition-group>
      </ul>
    </div>
    <div class="fixed bottom-12 left-1/2 -translate-x-1/2">
      <router-link
        :to="{ path: '/weekly/create', query: { type: 'weekly' } }"
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
import { getWeeklyItems } from "../services/actionableService";
import actionableStore from "../stores/actionable";
import useDragSort from "../use/dragSort";

export default defineComponent({
  components: {
    TheAvatar: TheAvatar,
    ActionableItem: ActionableItem,
  },
  setup() {
    const weeklyItems = computed({
      get: () => {
        return actionableStore.getters.getWeeklyItems;
      },
      set: (actionableItems) => {
        actionableStore.updateWeeklyItemList(actionableItems);
      },
    });

    onMounted(async () => {
      const weeklyItems = await getWeeklyItems();
      actionableStore.storeWeeklyItems(weeklyItems);
    });

    const dragFunctions = useDragSort(weeklyItems);

    return {
      weeklyItems,
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
