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
          <span>Dropdown button</span>
          <span class="icon is-small">
            <i class="fas fa-angle-down" aria-hidden="true"></i>
          </span>
        </button>
      </div>
      <div class="dropdown-menu" id="dropdown-menu" role="menu">
        <div class="dropdown-content">
          <a href="#" class="dropdown-item">
            Dropdown item
          </a>
          <a class="dropdown-item">
            Other dropdown item
          </a>
          <a href="#" class="dropdown-item is-active">
            Active dropdown item
          </a>
          <a href="#" class="dropdown-item">
            Other dropdown item
          </a>
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
    };
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
</style>
