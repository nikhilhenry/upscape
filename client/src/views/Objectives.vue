<template>
  <div class="home">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar">
        <h1 class="title">Objectives</h1>
        <Avatar />
      </div>

      <div class="secondary-bar">
        <ObjectiveSort />
      </div>

      <!-- objectives -->
      <ul class="objective-list">
        <li v-for="(objective, index) in objectives" :key="index">
          <ObjectiveItem :objective="objective" />
        </li>
        <div ref="divAsTarget"></div>
      </ul>
    </div>
    <!-- create objective objective -->
    <div class="fab">
      <router-link :to="{ path: 'objectives/create' }" class="create-button"
        >Create Objective</router-link
      >
    </div>
    <router-view></router-view>
  </div>
</template>

<script>
import Avatar from "@/components/Avatar";
import ObjectiveItem from "@/components/ObjectiveItem";
import ObjectiveSort from "@/components/ObjectiveSort";
import { mapActions, mapGetters } from "vuex";

import getObjectives from "@/api/objectiveGet.js";

export default {
  name: "Objectives",
  components: {
    Avatar,
    ObjectiveItem,
    ObjectiveSort,
  },
  computed: {
    isLoaded() {
      return this.$store.state.objective.objectivesLoaded;
    },
    ...mapGetters({
      objectives: "objective/getObjectives",
    }),
  },
  data() {
    return {
      observer: null,
      options: { root: null, threshold: 0 },
      page: 1,
    };
  },
  methods: {
    queryObjectives: async function() {
      const objectives = await getObjectives();
      return objectives;
    },
    saveObjectivesToStore: function(objectives) {
      this.storeObjectives(objectives);
    },
    loadMore: async function() {
      this.page++;
      const newObjectives = await getObjectives(this.page);
      this.storeObjectives(newObjectives);
    },
    ...mapActions({
      storeObjectives: "objective/storeObjectives",
    }),
  },
  async mounted() {
    if (!this.isLoaded) {
      const objectives = await this.queryObjectives();
      this.saveObjectivesToStore(objectives);
    }

    // intersection observer
    this.observer = new IntersectionObserver(this.loadMore, this.options);
    this.observer.observe(this.$refs.divAsTarget);
  },
};
</script>

<style lang="scss" scoped>
.objective-list {
  margin-top: 3rem;
  list-style: none;
  margin-left: 0;
  padding: 0;
}

.fab {
  position: fixed;
  bottom: 50px;
  right: calc(50% - 7rem);
  // right: 50%;
}

.create-button {
  color: #fff;
  background-color: $primary;
  border: none;
  font-size: 1.5rem;
  font-family: "avenir";
  padding: 0.5rem 1.2rem 0.5rem;
  border-radius: 0.2rem;
  margin: 0;
  text-decoration: none;
  transition: all 200ms ease-in;
  &:hover {
    // display:none;
    background-color: darken($primary, 5%);
  }
}
</style>
