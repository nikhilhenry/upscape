<template>
  <div class="login">
    <div class="container">
      <div class="container__item">
				<p class="error" v-if="error">{{error}}</p>
        <form class="form">
          <input type="password" class="form__field" placeholder="Enter Password" v-model="password" />
          <button type="button" class="btn btn--primary btn--inside uppercase" @click="login">
						<i class="fas fa-spinner fa-spin" v-if="isLoading"></i>
						<div v-else>Login</div>
					</button>
        </form>
      </div>
      
      <div class="container__item container__item--bottom">
        <p>Built by <u>Nikhil Henry.</u></p>
      </div>
    </div>    
  </div>
</template>

<script>
import loginQuery from '@/api/login'

export default {
  name:"Login",
  data(){
    return{
      password:'',
      success:null,
			error:'',
			isLoading:false
    }
  },
  methods:{
    login:async function(){
			this.isLoading=!this.isLoading
			const password = this.password
			if(password.length<8){
				this.error = 'Incorrect password.'
				return
			} 
			const response = await loginQuery(password)
			this.isLoading=!this.isLoading
      if(response.error){
				this.error=response.error
				return
			}
			console.log(response.token)

			// this.password = ''
    }
  }
}
</script>

<style lang="scss" scoped>
//** variables
$input-bg-color: #fff;
$input-text-color: #a3a3a3;
$button-bg-color:$primary;
$button-text-color: #fff;
$border-radius: .2rem;


.login {
	height: 100vh;
  width: 100vw;
}

.error{
	color:hsl(348, 100%, 61%);
}



//** helper
.container {
  height: 100%;
  width:100%;
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: center;
}

.uppercase {
	text-transform: uppercase;
}

//** button
.btn {
	display: inline-block;
	background: transparent;
	color: inherit;
	font: inherit;
	border: 0;
	outline: 0;
	padding: 0;
	transition: all 200ms ease-in;
	cursor: pointer;
	
	&--primary {
		background: $button-bg-color;
		color: $button-text-color;
		box-shadow: 0 0 10px 2px rgba(0, 0, 0, .1);
		border-radius: $border-radius;
		padding: 12px 36px;
		
		&:hover {
			background: darken($button-bg-color, 4%);
		}
		
		&:active {
			background: $button-bg-color;
			box-shadow: inset 0 0 10px 2px rgba(0, 0, 0, .2);
		}
	}
	
	&--inside {
		margin-left: -96px;
	}
}

//** form
.form {	
	&__field {
		width: 360px;
		background: #fff;
		color: $input-text-color;
		font: inherit;
		box-shadow: 0 6px 10px 0 rgba(0, 0, 0 , .1);
		border: 0;
		outline: 0;
		padding: 22px 18px;
    border-radius: $border-radius;
	}
}
</style>