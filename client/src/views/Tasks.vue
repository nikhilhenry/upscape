<template>
  <div class="tasks">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar">
        <h1 class="title">Tasks</h1>
        <Avatar/>
      </div>
      <TasksDateRange v-bind:queryDate.sync="queryDate"/>      
    </div>
  </div>
</template>

<script>
import Avatar from '@/components/Avatar'
import TasksDateRange from '@/components/TasksDateRange'

export default {
  name:'Tasks',
  components:{
    Avatar,
    TasksDateRange,
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