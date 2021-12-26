<template>
  <div class="tasks">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar">
        <h1 class="title">Tasks</h1>
        <TheAvatar />
      </div>
      <div class="secondary-bar">
        <TheTasksDateRange
          v-bind:queryDate="queryDate"
          v-on:update:query="queryDate = $event"
        />
        <h3
          class="total-duration"
          :class="colorClass"
          @click="showRemaining = !showRemaining"
        >
          {{ totalDuration }} + 1 hrs
        </h3>
      </div>

      <!-- tasks  -->
      <ul class="task-list">
            <li v-for="task in tasks" :key="task._id">
              <TaskItem :task="task" />
            </li>
          </transition-group>
      </ul>
    </div>
  </div>
</template>

<style lang="scss" scoped>
@import "../assets/time-color-classes.scss";

.secondary-bar {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;

  .total-duration {
    font-size: 24px;
    cursor: pointer;
  }
}

.task-list {
  margin-top: 3rem;
  list-style: none;
  margin-left: 0;
  padding: 0;
}

.fab {
  position: fixed;
  bottom: 50px;
  left: 50%;
  transform: translate(-50%, 0);
  // right: 50%;
}

.create-button {
  color: #fff;
  background-color: $primary;
  border: none;
  font-size: 1.5rem;
  font-family: Avenir, Helvetica, Arial, sans-serif;
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

// draggable
.flip-list-move {
  transition: transform 0.5s;
}

.ghost {
  box-shadow: 5px 5px 2.5px -1px rgba(0, 0, 0, 0.14);
  opacity: 0.7;
}
</style>