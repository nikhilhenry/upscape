<template>
  <div class="sort-wrapper">
    <!-- sort field -->
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
          <span>{{ formattedOption }}</span>
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

    <!-- sort order -->
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
          <span>{{ formattedOption }}</span>
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
  computed: {
    formattedOption() {
      //sort options
      let sortOptions = {
        created_at: "Created At",
        scheduled_for: "Scheduled For",
        default: "Created At",
      };
      //Return the formatted option
      return sortOptions[this.selectedOption] || sortOptions["default"];
    },
  },
  data() {
    return {
      isActive: false,
      selectedOption: "created_at",
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

.sort-wrapper {
  display: flex;
  justify-content: space-between;
  width: 100%;
  align-items: center;
  margin: 1rem 0;
}

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
