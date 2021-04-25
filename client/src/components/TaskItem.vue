<template>
  <div class="task" v-bind:class="{'is-completed':completed}" @dblclick="deleteTask()">
    <div class="wrapper">
      <div class="left">
        <p class="title">{{task.name}}</p>
        <span class="date">{{timeAgo}}</span>
      </div>
      <div class="right">
        <i class="fas fa-star" v-if="task.highlight"></i>
        <ul class="tags">
          <li v-for="(tag,index) in tags" :key="index">
            <span class="tag" :style="{background:tag.color}">{{tag.name}}</span>
          </li>
        </ul>
        <span class="duration">{{task.duration}} MIN</span>
        <input id="c1" type="checkbox" class="complete" v-model="completed">
      </div>
    </div>
  </div>
</template>

<script>
import TimeAgo from 'javascript-time-ago'

import deleteTaskById from '@/api/taskDelete.js'
import updateTaskById from '@/api/taskPut.js'

export default {
  name:'TaskItem',
  props: ['task'],
  data(){
    return{
      timeAgo:'',
      completed:false,
      tags:[]
    }
  },
  watch:{
    completed(){
      this.completeTask()
    }
  },
  mounted(){
    const timeAgo = new TimeAgo('en-US')
    this.timeAgo = timeAgo.format(new Date(this.task.created_at))

    this.completed = this.task.completed

    // fetching tags
    if(this.task.tag_ids){
      let tags = []
      this.task.tag_ids.forEach(tagId=>{
        const tagObject = this.$store.getters['tag/getTagById'](tagId)
        tags.push(tagObject)
      })

      // save tags to component
      this.tags = tags
    }
  },
  methods: {
    deleteTask:async function(){
      const error = await deleteTaskById(this.task._id)
      console.log(error)
      if (error) this.$store.dispatch('task/deleteTask',this.task._id)
    },
    completeTask:async function(){
      console.log('updating task...')
      const updatedTask = await updateTaskById(this.task._id,{completed:this.completed})
      this.$store.dispatch('task/updateTask',updatedTask)
    }
  }
}
</script>

<style lang="scss" scoped>
@import "@/assets/toggles.scss";

$flex-gap:2rem;

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

  .date{
    font-size: 1em;
    color:rgba($text-secondary, $alpha: .8);
    font-weight: 100;
    display:block;
    margin-bottom: .5rem;
  }

  .duration{
    font-size: 1.2rem;
    color: $text-secondary;
    font-weight: 400;
    margin-left: $flex-gap;
  }

  .fas{
    color:#ffd31d;
    font-size: 1.4rem;
    margin-left: $flex-gap;
  }

  .complete{
    margin-left: $flex-gap;
    transform: scale(1.5);
  }

  .right{
    display: flex;
    align-items: center;
    justify-content: flex-end;
    min-width: 300px;   
  }
}

.tags{
  list-style: none;
  margin:0;
  margin-left: $flex-gap;
  padding:0;
  display: flex;
  flex-direction: row;
  .tag{
    color: #fff;
    padding: .1rem .5rem .1rem;
    margin-left: .5rem;
    border-radius: .3rem;
  }
}

.is-completed{
  filter: opacity(.5)
}
</style>