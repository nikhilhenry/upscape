<template>
  <div class="task">
    <div class="wrapper">
      <div class="left">
        <p class="title">Eat Sushi</p>
        <span class="date">1 hour ago</span>
      </div>
      <div class="right">
        <span class="duration">60 MIN</span>
        <i class="fas fa-star"></i>
        <input id="c1" type="checkbox">
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name:'TaskItem'
}
</script>

<style lang="scss" scoped>
@import "@/assets/toggles.scss"
</style>