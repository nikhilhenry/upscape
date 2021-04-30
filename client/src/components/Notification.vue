<template>
  <div class="notification">
    <div class="wrapper">
      <h2 class="message">{{notification.message}}</h2>
    </div>
  </div>
</template>

<script>
export default {
  name:'Notification',
  props:["notification"]
}
</script>
<style lang="scss" scoped>
.notification{
  position: fixed;
  right:0;
  bottom:0;
  z-index: 100;
}

.wrapper{
  margin:1rem;
}
</style>

