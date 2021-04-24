<template>
  <div class="tasks">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar">
        <h1 class="title">Tasks</h1>
        <Avatar/>
      </div>
      <div class="secondary-bar">
      <TasksDateRange
        v-bind:queryDate="queryDate"
        v-on:update:query="queryDate = $event"
       />
      <h3 class="total-duration">4.5 + 1 hrs</h3>
      </div>

       <!-- tasks  -->
       <ul class="task-list">
        <li v-for="(task,index) in tasks" :key="index">
          <TaskItem :task="task"/>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import Avatar from '@/components/Avatar'
import TasksDateRange from '@/components/TasksDateRange'
import TaskItem from '@/components/TaskItem'

import getTasks from '@/api/taskGet'

export default {
  name:'Tasks',
  components:{
    Avatar,
    TasksDateRange,
    TaskItem
  },
  data(){
    return{
      queryDate:'today',
      tasks:[]   
    }
  },
  watch:{
    queryDate(){
      // clear existing tasks
      this.tasks = []
      // query for new task
      this.queryTasks()
    }
  },
  methods:{
    test:function(queryDate){
      console.log(queryDate)
    },
    queryTasks:async function(){
      const tasks = await getTasks(this.queryDate)
      this.tasks = tasks
    }
  },
  mounted(){
    this.queryTasks()
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

.task-list{
  margin-top: 3rem;
  list-style: none;
  margin-left:0;
  padding:0;
}
</style>