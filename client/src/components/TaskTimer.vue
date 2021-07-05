<template>
  <div class="task-timer">
    <section class="icons">
      <div class="icon meta">
        <i class="las la-braille"></i>
      </div>
      <div class="icon pause">
        <i class="las la-pause"></i>
      </div>
      <div class="icon retry">
        <i class="las la-undo-alt"></i>
      </div>
      <div class="icon stop">
        <i class="las la-stop"></i>
      </div>
    </section>
    <section class="time">
      <div class="display">
        <p>04:56</p>
      </div>
    </section>
  </div>
</template>

<style lang="scss" scoped>
.task-timer {
  position: fixed;
  position: fixed;
  bottom: 50px;
  left: 50%;
  transform: translate(-50%, 0);
  background-color: $nav-bg-primary;
  border-radius: 0.5rem;
  display: flex;
  align-items: center;
}

.icons {
  display: flex;
  align-items: center;
  justify-content: space-between;

  .icon {
    font-size: 1.5rem;
    font-weight: 900;
    padding: 1rem 1.5rem;
    &:hover {
      background: $nav-bg-secondary;
      transition: background 200ms ease-in;
    }
  }
}

.time {
  margin: 0 1rem 0;
  p {
    font-size: 1.5rem;
    font-weight: 600;
    line-height: 0;
  }
}

// icon colors
.meta {
  color: $text-primary;
  filter: grayscale(100%) opacity(0.7);
}
.pause {
  color: $primary;
}
.retry {
  color: #4ad66d;
}
.stop {
  color: #f3722c;
}
</style>
