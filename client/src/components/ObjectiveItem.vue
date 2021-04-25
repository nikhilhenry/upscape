<template>
  <div class="objective">
    <div class="wrapper">
      <div class="left">
        <p class="title">Deploy App</p>
        <span class="time-ago">1 hour ago</span>
      </div>
      <div class="right">
        <span class="date">APR 26 21</span>
        <input id="c1" type="checkbox" class="complete">
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name:'ObjectiveItem'
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
  border-radius: 1rem;

  margin-bottom: .5rem;

  .title{
    font-size: 1.4rem;
    margin:1rem 0 .5rem 0rem;
  }

  .time-ago{
    font-size: 1em;
    color:rgba($text-secondary, $alpha: .8);
    font-weight: 100;
    display:block;
    margin-bottom: .5rem;
  }

  .date{
    font-size: 1.2rem;
    color: $text-secondary;
    font-weight: 400
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

.is-completed{
  filter: opacity(.5)
}
</style>