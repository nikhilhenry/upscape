<template>
  <div class="tasks">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar">
        <h1 class="title">Tasks</h1>
        <Avatar/>
      </div>
      <div class="secondary-bar">
      <TasksDateRange v-bind:queryDate.sync="queryDate"/>
      <h3 class="total-duration">4.5 + 1 hrs</h3>
      </div>

       <!-- tasks  -->
      <TaskItem/>
    </div>
  </div>
</template>

<script>
import Avatar from '@/components/Avatar'
import TasksDateRange from '@/components/TasksDateRange'
import TaskItem from '@/components/TaskItem'

export default {
  name:'Tasks',
  components:{
    Avatar,
    TasksDateRange,
    TaskItem
  },
  data(){
    return{
      queryDate:'today'    
    }
  },
  methods:{
    test:function(queryDate){
      console.log(queryDate)
    }
  }
}
</script>

<style lang="scss" scoped>
.secondary-bar{
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;

  .total-duration{
    font-size: 24px;
    color:#0be881;
  }
}
</style>