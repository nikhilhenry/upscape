<template>
  <div id="app">
    <component :is="this.$route.meta.layout || 'div'">
      <router-view/>
    </component>
    <Notification :notification="{message:'Hello World'}"/>
  </div>
</template>

<script>
import Notification from '@/components/Notification'

export default {
  components:{
    Notification
  }
}
</script>

<style lang="scss">
:root{
  background:$background;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color:$text-primary;
  background:$background;
  margin:0;
  padding:0;
}

body{
  margin: 0;
  padding:0;
}

</style>
