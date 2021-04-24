<template>
  <div class="task">
    <div class="wrapper">
      <div class="left">
        <p class="title">{{task.name}}</p>
        <span class="date">1 hour ago</span>
      </div>
      <div class="right">
        <span class="duration">{{task.duration}} MIN</span>
        <i class="fas fa-star" v-if="task.highlight"></i>
        <input id="c1" type="checkbox" class="complete" v-model="task.completed">
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name:'TaskItem',
  props: ['task']
}
</script>

<style lang="scss" scoped>
@import "@/assets/toggles.scss";

.wrapper{
  display:flex;
  justify-content:space-between;
  align-items: center;
  background-color:$secondary;
  padding:.5rem 1.5rem;
  border-bottom: 3px solid $background;
  border-radius: 2rem;

  margin-bottom: .5rem;

  .title{
    font-size: 1.4rem;
    margin:1rem 0 .5rem 0rem;
  }

  .date{
    font-size: 1em;
    color:rgba($text-secondary, $alpha: .8);
    font-weight: 100;
    display:block;
    margin-bottom: .5rem;
  }

  .duration{
    font-size: 1.2rem;
    color: $text-secondary;
    font-weight: 400;
  }

  .fas{
    color:#ffd31d;
    font-size: 1.4rem;
  }

  .complete{
    transform: scale(1.5);
  }

  .right{
    display: flex;
    align-items: center;
    justify-content: space-between;
    min-width: 200px;    
  }
}
</style>