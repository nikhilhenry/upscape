<template>
<div class="home">
  <div class="container">
    <!-- title bar -->
    <div class="title-bar">
      <h1 class="title">{{greeting}}, <br> {{firstName}} 👋</h1>
      <Avatar/>
    </div>
  </div>
</div>
</template>

<script>
import Avatar from '@/components/Avatar'

import metaQuery from '@/api/metaQuery'
import { mapActions, mapGetters } from 'vuex'
import getGreeting from '@/functions/getGreeting'

export default {
  name: 'Home',
  components:{
    Avatar
  },
  computed:{
    ...mapGetters({
      firstName:'user/getFirstName'
    })
  },
  data(){
    return{
      greeting:''
    }
  },
  methods:{
    getMeta: async function(){
      const response = await metaQuery()
      this.storeMetaData(response)
      
    },
    ...mapActions({ 
      storeMetaData:'user/storeMetaData'
    })
  },
  mounted(){
    this.greeting = getGreeting()
    if(!this.$store.state['user/metaLoaded']) this.getMeta();
  }
}
</script>
