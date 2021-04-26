<template>
  <div class="objective" v-bind:class="{'is-completed':completed}" @dblclick="deleteTask()">
    <div class="wrapper">
      <div class="left">
        <p class="title">{{objective.name}}</p>
        <span class="time-ago">{{timeAgo}}</span>
      </div>
      <div class="right">
        <span class="date">{{scheduledFor}}</span>
        <input id="c1" type="checkbox" class="complete" v-model="completed">
      </div>
    </div>
  </div>
</template>

<script>
import TimeAgo from 'javascript-time-ago'
import updateObjectiveById from '@/api/objectivePut.js'

export default {
  name:'ObjectiveItem',
  props:['objective'],
  data(){
    return{
      timeAgo:'',
      completed:false,
      scheduledFor:''
    }
  },
  watch:{
    completed(){
      this.completeObjective()
    }
  },
  methods:{
    formateDate:function(date){
      const formattedDate = date.toLocaleString('default',{month:'short'}).toUpperCase()+' '+new Date(date).getDate()+' '+ new Date(date).getFullYear()
      return formattedDate
    },
    completeObjective:async function(){
      const updatedObjective = await updateObjectiveById(this.objective._id,{completed:this.completed})
      this.$store.dispatch('objective/updateObjective',updatedObjective)
    }
  },
  created(){
    const timeAgo = new TimeAgo('en-US')
    this.timeAgo = timeAgo.format(new Date(this.objective.created_at))

    this.completed = this.objective.completed
    
    // format scheduled_for
    const formattedDate = this.formateDate(new Date(this.objective.scheduled_for))
    this.scheduledFor = formattedDate
  }
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