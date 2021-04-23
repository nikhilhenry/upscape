<template>
  <div class="avatar">
    <img :src="imgURL" alt="user_profile_pic">
  </div>
</template>

<script>
export default {
  name:'Avatar',
  data(){
    return{
    }
  },
  computed:{
    imgURL(){
     return this.$store.getters['user/getProfilePicURL']
    }
  }
}
</script>

<style lang="scss" scoped>
.avatar{
  height: 100px!important;
  width: auto;
  display: block!important;
  img{
    height: 100%;
    width:auto;
    border-radius: 35%;
    object-fit: cover;
  }
}
</style>