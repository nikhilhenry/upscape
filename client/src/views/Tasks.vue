<template>
  <div class="tasks">
    <div class="container">
      <!-- title bar -->
      <div class="title-bar">
        <h1 class="title">Tasks</h1>
        <Avatar/>
      </div>      
    </div>
  </div>
</template>

<script>
import Avatar from '@/components/Avatar'

export default {
  name:'Tasks',
  components:{
    Avatar
  }
}
</script>