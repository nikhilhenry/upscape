<template>
<div class="home">
  <div class="container">
    <h1>Good Morning, <br>
    Nikhil 👋 </h1>
  </div>
</div>
</template>

<script>
import metaQuery from '@/api/metaQuery'
import { mapActions } from 'vuex'

export default {
  name: 'Home',
  methods:{
    getMeta: async function(){
      const response = await metaQuery()
      this.storeMetaData(response)
      
    },
    ...mapActions({ 
      storeMetaData:'user/storeMetaData'
    })
  },
  created(){
    this.getMeta()
  }
}
</script>
