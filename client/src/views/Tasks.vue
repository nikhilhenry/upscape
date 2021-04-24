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
    <div class="fab">
      <button class="create-task">Create Task</button>
    </div>
  </div>
</template>

<script>
import Avatar from '@/components/Avatar'
import TasksDateRange from '@/components/TasksDateRange'
import TaskItem from '@/components/TaskItem'

import getTasks from '@/api/taskGet'
import { mapGetters,mapActions } from 'vuex'

export default {
  name:'Tasks',
  components:{
    Avatar,
    TasksDateRange,
    TaskItem
  },
  computed:{
    isLoaded(){
      return this.$store.state.task.tasksLoaded;
    },
    ...mapGetters({
      tasks:'task/getTasks'
    })
  },
  data(){
    return{
      queryDate:'today',
    }
  },
  // watch:{
  //   queryDate(){
  //     // clear existing tasks
  //     this.tasks = []
  //     // query for new task
  //     this.queryTasks()
  //   }
  // },
  methods:{
    queryTasks:async function(){
      const tasks = await getTasks(this.queryDate)
      // this.tasks = tasks
      return tasks
    },
    saveTasksToStore:function(tasks){
      this.storeTasks(tasks)
    },
    ...mapActions({
      storeTasks:'task/storeTasks'
    })
  },
  async mounted(){
    if(!this.isLoaded){
     const tasks = await this.queryTasks()
     this.saveTasksToStore(tasks)
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

.task-list{
  margin-top: 3rem;
  list-style: none;
  margin-left:0;
  padding:0;
}

.fab{
  position: fixed;
  bottom: 50px;
  right:calc(50% - 4.5rem);
  // right: 50%;
}

.create-task{
  color:#fff;
  background-color:$primary;
  border:none;
  font-size: 1.5rem;
  font-family: "avenir";
  padding: .5rem 1.2rem .5rem;
  border-radius: .2rem;
  margin:0;
  transition: all 200ms ease-in;
  &:hover{
    // display:none;
    background-color:darken($primary,5%)
  }
}

</style>