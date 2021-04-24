<template>
  <div class="date-range">
    <div class="wrapper">
      <i class="fas fa-angle-left"></i>
      <div class="content">
      <h3 class="day">Today</h3>
      <p>APR 24 2021</p>
      </div>
      <i class="fas fa-angle-right"></i>
    </div>
  </div>
</template>

<script>
export default {
  name:'TasksDateRange'
}
</script>

<style lang="scss" scoped>
.wrapper{
  display:flex;
  flex-direction: row;
  align-items: center;
  text-align: center;

  .fas{
    font-size: 2.5rem;
  }

  h3{
    margin-left: 1.2rem;
    margin-right: 1.2rem;
    margin-bottom:5px;
    margin-top:0;
    font-size:24px;
  }

  p{
    margin-top:0;
    margin-bottom: 0;
    margin-left: 1rem;
    margin-right: 1rem;
    font-size: 18px;
  }
}
</style>