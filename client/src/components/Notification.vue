<template>
  <div class="notification">
    <div class="wrapper" :class="notification.success ? 'is-success':'is-danger'">
        <i class="far fa-check-circle" v-if="notification.success"></i>
        <i class="fas fa-times-circle" v-else></i>
        <h3 class="message">{{notification.message}}</h3>
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
  padding: .5rem 1.2rem .5rem;
  color: #ffffff;
  display:flex;
  align-items: center;
  // justify-content:space-around;
  min-width: 200px;
  height:50px;
  background-color:$nav-bg-primary;
  border-radius: 1rem;

  i{
    font-size: 1.5rem;
  }

  .message{
    margin-left: 10px;
  }
}

.is-success{
  background-color:#48c78e!important;
}

.is-danger{
  background-color:#f14668!important;
}
</style>

