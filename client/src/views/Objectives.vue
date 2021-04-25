<template>
<div class="home">
  <div class="container">
    <!-- title bar -->
    <div class="title-bar">
      <h1 class="title">Objectives</h1>
      <Avatar/>
    </div>

    <!-- objectives -->
    <ul class="objective-list">
      <li v-for="(objective,index) in objectives" :key="index">
        <ObjectiveItem :objective="objective"/>
      </li>
    </ul>
  </div>
</div>
</template>

<script>
import Avatar from '@/components/Avatar'
import ObjectiveItem from '@/components/ObjectiveItem'

export default {
  name:'Objectives',
  components:{
    Avatar,
    ObjectiveItem
  },
  data(){
    return{
      objectives:[
        {
          "_id": "6065deaa354032982e90599a",
          "name": "Go to the beach",
          "completed": false,
          "scheduled_for": "2021-04-05T14:52:26Z",
          "created_at": "2021-04-01T14:54:34Z"
        },
        {
          "_id": "6065e38cd6972255cb13038d",
          "name": "should be completed",
          "completed": false,
          "scheduled_for": "2021-04-07T14:52:26Z",
          "created_at": "2021-04-01T15:15:24Z"
        }
      ]
    }
  }
}
</script>

<style lang="scss" scoped>
.objective-list{
  margin-top: 3rem;
  list-style: none;
  margin-left:0;
  padding:0;
}
</style>