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
      tasks:[
  {
    "_id": "6065441b2188168ab6e0c47a",
    "id": 1,
    "name": "Eat sushi",
    "duration": 60,
    "completed": false,
    "created_at": "2021-04-01T03:55:07Z",
    "completed_at": "0001-01-01T00:00:00Z",
    "highlight":true,
  },
  {
    "_id": "606543dbb8bd696647d73d61",
    "id": 3,
    "name": "Eat pizza",
    "duration": 10,
    "completed": false,
    "created_at": "2021-04-01T03:54:03Z",
    "completed_at": "0001-01-01T00:00:00Z"
  },
  {
    "_id": "606546f0c5d068ebc0eb3afe",
    "id": 4,
    "name": "Eat pasta",
    "duration": 5,
    "completed": false,
    "created_at": "2021-04-01T04:07:12Z",
    "completed_at": "0001-01-01T00:00:00Z"
  },
  {
    "_id": "60658b44229c35fa4eb2f160",
    "id": 11,
    "name": "the present",
    "duration": 10,
    "completed": true,
    "created_at": "2021-04-01T04:07:12Z",
    "completed_at": "2021-04-01T14:23:46Z"
  }
]    
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

.task-list{
  margin-top: 3rem;
  list-style: none;
  margin-left:0;
  padding:0;
}
</style>