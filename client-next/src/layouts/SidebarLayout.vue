<template>
  <div>
    <nav class="navbar">
      <ul class="navbar-nav">
        <li class="logo">
          <a href="#" class="nav-link">
            <span class="link-text logo-text">Upscape</span>
            <img src="../assets/bmo.svg" />
          </a>
        </li>
        <!-- navbar items -->
        <!-- home -->
        <li class="nav-item">
          <router-link :to="{ name: 'Home' }" class="nav-link">
            <i class="las la-home icon"></i>
            <span class="link-text">Home</span>
          </router-link>
        </li>
        <!-- tasks -->
        <li class="nav-item">
          <router-link :to="{ name: 'Tasks' }" class="nav-link">
            <i class="las la-briefcase icon"></i>
            <span class="link-text">Tasks</span>
          </router-link>
        </li>

        <!-- settings -->
        <li class="nav-item">
          <router-link :to="{ name: 'Settings' }" class="nav-link">
            <i class="las la-cog icon"></i>
            <span class="link-text">Settings</span>
          </router-link>
        </li>
      </ul>
    </nav>

    <!-- router view -->
    <router-view class="main" />
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { getMeta } from "../services/userService";
import userStore from "../stores/user";

export default defineComponent({
  setup() {
    // load meta details on mount
    onMounted(async () => {
      if (userStore.getters.isLoggedIn) {
        const userMeta = await getMeta();
        userStore.storeUserMeta(userMeta);
      }
    });
  },
});
</script>

<style lang="scss" scoped>
body {
  color: black;
  background-color: white;
  margin: 0;
  padding: 0;
}

.main {
  margin-left: 5rem;
  padding: 1rem;
}

// active class
.router-link-active {
  background: $nav-bg-secondary;
  filter: grayscale(0%) !important;
  background: $nav-bg-secondary;
  color: $nav-text-secondary;
}

.navbar {
  position: fixed;
  background-color: $nav-bg-primary;
  transition: width 200ms ease;
  z-index: 5;
}

.navbar-nav {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100%;

  .nav-item {
    width: 100%;
  }

  .nav-item:last-child {
    margin-top: auto;
  }

  .nav-link {
    display: flex;
    align-items: center;
    height: 5rem;
    color: $nav-text-primary;
    text-decoration: none;
    filter: grayscale(100%) opacity(0.7);
    transition: $transistion-speed;

    &:hover {
      filter: grayscale(0%);
      background: $nav-bg-secondary;
      color: $nav-text-secondary;
    }

    .link-text {
      display: none;
      margin-left: 1rem;
    }

    i {
      min-width: 2rem;
      margin: 0 1.5rem;
    }
  }
}

// nav-item colors
.icon {
  color: $nav-icon-primary;
  font-size: 32px;
  transition: $transistion-speed;
  margin: 0 auto 0;
}

// logo styling
.logo {
  font-weight: bold;
  text-transform: uppercase;
  margin: 1rem 0 1rem;
  text-align: center;
  color: $nav-text-secondary;
  font-size: 1.5rem;
  letter-spacing: 0.3ch;
  width: 100%;
  transition: $transistion-speed;
  .nav-link {
    filter: none !important;
  }

  img {
    margin: 0 auto 0;
    width: 70%;
    height: 70%;
    filter: brightness(0) invert(0.8);
  }
}

.logo img {
  transform: rotate(0deg);
  transition: $transistion-speed;
}

.logo-text {
  display: inline;
  transition: $transistion-speed;
}

// smaller screens
@media only screen and(max-width: 600px) {
  .navbar {
    bottom: 0;
    width: 100vw;
    height: 5rem;
    border-top: 1px solid $nav-bg-secondary;
  }

  .logo {
    display: none;
  }

  .navbar-nav {
    flex-direction: row;
  }

  .nav-link {
    justify-content: center;
  }

  .main {
    margin: 0;
    margin-bottom: 3rem;
  }
}

// large screen
@media only screen and(min-width: 600px) {
  .navbar {
    top: 0;
    width: 5rem;
    height: 100vh;
    border-right: 1px solid $nav-bg-secondary;
    &:hover {
      width: 16rem;
      .link-text {
        display: inline !important;
        transition: opacity $transistion-speed;
      }
    }
  }
}
</style>
