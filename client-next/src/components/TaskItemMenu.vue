<template>
  <div class="ml-8">
    <Menu as="div" class="relative inline-block text-left">
      <MenuButton class="hover:text-skin-inverted transition-colors"
        ><i class="fas fa-chevron-down"></i
      ></MenuButton>
      <transition
        enter-active-class="transition duration-100 ease-out"
        enter-from-class="transform scale-95 opacity-0"
        enter-to-class="transform scale-100 opacity-100"
        leave-active-class="transition duration-75 ease-in"
        leave-from-class="transform scale-100 opacity-100"
        leave-to-class="transform scale-95 opacity-0"
      >
        <MenuItems
          class="
            absolute
            right-0
            mt-2
            ml-5
            origin-top-right
            bg-skin-base
            rounded-md
            shadow-lg
            ring-1 ring-black ring-opacity-5
            focus:outline-none
            z-10
          "
        >
          <div class="px-2 py-2 flex flex-col w-max text-lg">
            <MenuItem v-slot="{ active }">
              <button
                :class="[
                  active
                    ? 'bg-skin-fill bg-opacity-50 text-skin-base'
                    : 'text-skin-muted',
                  'group flex rounded-md items-center w-full px-2 py-2',
                ]"
              >
                <i
                  class="fas fa-paper-plane w-5 h-5 mr-2 text-skin-inverted"
                ></i>
                Shift tomorrow
              </button>
            </MenuItem>
            <MenuItem v-slot="{ active }">
              <button
                :class="[
                  active
                    ? 'bg-skin-fill bg-opacity-50 text-skin-base'
                    : 'text-skin-muted',
                  'group flex rounded-md items-center w-full px-2 py-2',
                ]"
              >
                <i class="fas fa-play w-5 h-5 mr-2 text-skin-inverted"></i>
                Start Timer
              </button>
            </MenuItem>
            <MenuItem v-slot="{ active }">
              <button
                :class="[
                  active
                    ? 'bg-skin-fill bg-opacity-50 text-skin-base'
                    : 'text-skin-muted',
                  'group flex rounded-md items-center w-full px-2 py-2',
                ]"
              >
                <i class="fas fa-backspace w-5 h-5 mr-2 text-skin-inverted"></i>
                Delete Task
              </button>
            </MenuItem>
          </div>
        </MenuItems>
      </transition>
    </Menu>
  </div>
</template>

<script setup lang="ts">
import { Menu, MenuButton, MenuItems, MenuItem } from "@headlessui/vue";
</script>
