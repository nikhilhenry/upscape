<template>
  <div
    class="task"
    v-bind:class="{ 'is-completed': completed }"
    @dblclick="deleteTask()"
  >
    <div class="wrapper">
      <div class="left">
        <p class="title">{{ task.name }}</p>
      </div>
      <div class="right">
        <i class="fas fa-star" v-if="task.highlight"></i>
        <ul class="tags">
          <li v-for="(tag, index) in tags" :key="index">
            <span class="tag" :style="{ background: tag.color }">{{
              tag.name
            }}</span>
          </li>
        </ul>
        <span class="duration">{{ task.duration }} MIN</span>
        <input
          id="c1"
          type="checkbox"
          class="complete"
          v-model="completed"
          @click="completeTask"
        />
      </div>
    </div>
  </div>
</template>

<script>
import deleteTaskById from "@/api/taskDelete.js";
import updateTaskById from "@/api/taskPut.js";

export default {
  name: "TaskItem",
  props: ["task"],
  data() {
    return {
      completed: false,
      tags: [],
      isActive: false,
    };
  },
  mounted() {
    this.completed = this.task.completed;

    // fetching tags
    if (this.task.tag_ids) {
      let tags = [];
      this.task.tag_ids.forEach((tagId) => {
        const tagObject = this.$store.getters["tag/getTagById"](tagId);
        tags.push(tagObject);
      });

      // save tags to component
      this.tags = tags;
    }
  },
  methods: {
    deleteTask: async function() {
      const error = await deleteTaskById(this.task._id);
      console.log(error);
      if (error) this.$store.dispatch("task/deleteTask", this.task._id);
    },
    completeTask: async function() {
      const updatedTask = await updateTaskById(this.task._id, {
        completed: !this.task.completed,
      });
      this.$store.dispatch("task/updateTask", updatedTask);
    },
  },
};
</script>

<style lang="scss" scoped>
@import "@/assets/toggles.scss";

$flex-gap: 2rem;

.wrapper {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: $secondary;
  padding: 1rem 1.5rem;
  border-bottom: 3px solid $background;
  border-radius: 1rem;

  margin-bottom: 0.5rem;

  .title {
    font-size: 1.4rem;
    margin: 1rem 0 0.5rem 0rem;
  }

  .date {
    font-size: 1em;
    color: rgba($text-secondary, $alpha: 0.8);
    font-weight: 100;
    display: block;
    margin-bottom: 0.5rem;
  }

  .duration {
    font-size: 1.2rem;
    color: $text-secondary;
    font-weight: 400;
    margin-left: $flex-gap;
  }

  .la-step-forward {
    color: $primary;
    font-size: 2rem;
    margin-left: $flex-gap;
    filter: grayscale(100%) opacity(0.7);
    transition: all 200ms ease-in;
    &:hover {
      filter: none;
    }
  }

  .active {
    filter: none;
  }

  .fa-star {
    color: #ffd31d;
    font-size: 1.4rem;
    margin-left: $flex-gap;
  }

  .complete {
    margin-left: $flex-gap;
    transform: scale(1.5);
  }

  .right {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    min-width: 300px;
  }
}

.tags {
  list-style: none;
  margin: 0;
  margin-left: $flex-gap;
  padding: 0;
  display: flex;
  flex-direction: row;
  .tag {
    color: #fff;
    padding: 0.1rem 0.5rem 0.1rem;
    margin-left: 0.5rem;
    border-radius: 0.3rem;
  }
}

.is-completed {
  filter: opacity(0.5);
}
</style>
