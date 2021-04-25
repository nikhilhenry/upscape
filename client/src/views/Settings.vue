<template>
<div class="home">
  <div class="container">
    <!-- title bar -->
    <div class="title-bar">
      <h1 class="title">Settings</h1>
      <Avatar/>
    </div>
      <section class="options">
        <!-- name -->
        <div class="option">
          <h3 class="title is-4">Change Name</h3>
          <div class="field has-addons">
            <div class="control">
              <input class="input" type="text" placeholder="New Name" v-model="name">
            </div>
            <div class="control">
              <a class="button is-info" @click="updateName">
                Update
              </a>
            </div>
          </div>
        </div>

        <!-- img url -->
        <div class="option">
          <h3 class="title is-4">Change Image URL</h3>
          <div class="field has-addons">
            <div class="control">
              <input class="input" type="text" placeholder="https://github.com">
            </div>
            <div class="control">
              <a class="button is-info">
                Update
              </a>
            </div>
          </div>
        </div>


        <!-- update password-->
        <div class="option">
          <h3 class="title is-4">Update Password</h3>
          <!-- old password -->
          <div class="field">
            <div class="control old-password">
              <input type="text" class="input" required placeholder="old password">
            </div>
          </div>          
          <!-- new password -->
          <div class="field has-addons">
            <div class="control">
              <input class="input" type="password" required placeholder="new password">
            </div>
            <div class="control">
              <a class="button is-info">
                Update
              </a>
            </div>
          </div>
        </div>                 

        <div class="option logout">
          <button class="button is-danger" @click="logoutUser">
            <span>Logout</span>
            <span class="icon">
              <i class="fas fa-sign-out-alt"></i>
            </span>
          </button>          
        </div>
      </section>    
  </div>
</div>
</template>

<script>
import Avatar from '@/components/Avatar'
import {mapActions} from 'vuex'

import updateUser from '@/api/userPut.js'

export default {
  name:'Settings',
  components:{
    Avatar
  },
  data(){
    return {
      name:''
    }
  },
  methods:{
    updateName:async function(){
      if(!this.name) return  null
      const response = await updateUser({name:this.name})
      this.$store.dispatch('user/storeMetaData',response)
    },
    logoutUser:async function(){
      // run store action
      await this.logout()
      
      // route user to login page
      this.$router.push({name:'Login'})
    },
    ...mapActions({
      logout:'user/logout'
    })
  }
}
</script>

<style lang="scss" scoped>
@import "~bulma";

.title{
  color:$text-primary;
}

.options{
  margin-top:3rem;

  .option{
    margin-bottom: 2rem;
    .button-main{
      margin:1.5rem 0 1rem;
      padding:.3rem 2rem .3rem;
    }
  }
}

.input{
  background:$secondary;
  border-color: darken($secondary,5%);
  color:$text-secondary;
  &::placeholder{
    color:darken($text-primary,20%);
  }
}

.old-password{
  width:21%;
}

.logout{
  position:fixed;
  bottom:calc(5rem + 20px);
}
</style>