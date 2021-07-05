<template>
  <div class="task-timer">
    <section class="icons">
      <i class="las la-braille"></i>
      <i class="las la-pause"></i>
      <i class="las la-undo-alt"></i>
      <i class="las la-stop"></i>
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
}

.icons {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
</style>
