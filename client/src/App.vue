<template>
  <div id="app">
    <component :is="this.$route.meta.layout || 'div'">
      <router-view/>
    </component>
    <div class="notificatinons" v-if="notification">
      <Notification :notification="notification"/>
    </div>
  </div>
</template>

<script>
import Notification from '@/components/Notification'
import {mapGetters} from 'vuex'

export default {
  components:{
    Notification
  },
  computed:{
    ...mapGetters({
      notification:'notification/getNotification'
    })
  },
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
