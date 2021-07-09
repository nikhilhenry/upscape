<template>
  <div class="sort-wrapper">
    <div
      class="dropdown"
      :class="{ 'is-active': isActive }"
      @click="isActive = !isActive"
    >
      <div class="dropdown-trigger">
        <button
          class="button"
          aria-haspopup="true"
          aria-controls="dropdown-menu"
        >
          <span>{{ selectedOption }}</span>
          <span class="icon is-small">
            <i class="fas fa-angle-down" aria-hidden="true"></i>
          </span>
        </button>
      </div>
      <div class="dropdown-menu" id="dropdown-menu" role="menu">
        <div class="dropdown-content">
          <span
            class="dropdown-item"
            @click="selectOption('created_at')"
            :class="{ 'is-active': selectedOption == 'created_at' }"
          >
            Created At
          </span>
          <span
            class="dropdown-item"
            @click="selectOption('scheduled_for')"
            :class="{ 'is-active': selectedOption == 'scheduled_for' }"
          >
            Scheduled For
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ObjectiveSort",
  data() {
    return {
      isActive: false,
      selectedOption: "",
    };
  },
  methods: {
    selectOption: function(option) {
      console.log(option);
      this.selectedOption = option;
    },
  },
};
</script>

<style lang="scss" scoped>
@import "~bulma";

.button {
  background-color: $secondary;
  color: $text-secondary;
  border-color: $nav-bg-primary;
  &:hover {
    color: $text-primary;
    border-color: $nav-bg-secondary;
  }
}

.dropdown-content {
  background-color: $secondary;
  .dropdown-item {
    color: $text-primary;
    width: auto;
    padding-right: 0rem;
    &:hover {
      background-color: darken($secondary, 0.8);
      color: $text-secondary;
    }
  }
}

.dropdown-item.is-active {
  background-color: $primary;
}
</style>
