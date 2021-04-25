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
      <li>
        <ObjectiveItem/>
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