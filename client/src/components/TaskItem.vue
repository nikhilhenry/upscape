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
@import "@/assets/toggles.scss";

.wrapper{
  display:flex;
  justify-content:space-between;
  background-color:#1f2225;
  padding:1rem 1rem;
  border-bottom: 3px solid $background;

  .title{
    font-size: 1.2rem;
  }

  .date{
    margin-right: 3rem;
    font-size: .8rem;
    color:rgba(#54666d, $alpha: .5);
    font-weight: 100;
  }

  .duration{
    font-size: 1rem;
    color: #54666d;
    font-weight: 400;
  }

  .fas{
    color:#ffd31d;
    font-size: 1.4rem;
  }

  .right{
    display: flex;
    align-items: center;
    justify-content: space-between;
    min-width: 200px;    
  }
}
</style>