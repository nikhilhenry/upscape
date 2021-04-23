<template>
<div>
<nav class="navbar">
  <ul class="navbar-nav">
    <li class="logo">
      <a href="#" class="nav-link">
        <span class="link-text logo-text">Reflect</span>
        <img src="../assets/bmo.svg">
      </a>
    </li>
  <!-- navbar items -->
      <!-- home -->
      <li class="nav-item">
        <router-link :to="{name:'Home'}" class="nav-link" :class="{'active-class':activeRoute=='Home'}">
          <i class="las la-home icon"></i>
          <span class="link-text">Home</span>
        </router-link>
      </li>
      <!-- reflections -->
      <li class="nav-item">
        <router-link :to="{name:'Reflections'}" class="nav-link" :class="{'active-class':activeRoute=='Reflections'}">
          <i class="las la-calendar-plus icon"></i>
          <span class="link-text">Reflections</span>
        </router-link>
      </li>
      <!-- moments -->
      <li class="nav-item">
        <router-link :to="{name:'Moments'}" class="nav-link" :class="{'active-class':activeRoute=='Moments'}">
          <i class="las la-trophy icon"></i>
          <span class="link-text">Moments</span>
        </router-link>
      </li>
      <!-- mood synth -->
      <li class="nav-item">
        <router-link :to="{name:'MoodSynth'}" class="nav-link" :class="{'active-class':activeRoute=='MoodSynth'}">
          <i class="las la-play-circle icon"></i>
          <span class="link-text">Mood Synth</span>
        </router-link>
      </li>
      <!-- settings -->
      <li class="nav-item" >
        <router-link :to="{name:'Settings'}" class="nav-link" :class="{'active-class':activeRoute=='Settings'}">
          <i class="las la-cog icon"></i>
          <span class="link-text">Settings</span>
        </router-link>
      </li>
  </ul>
</nav>

<!-- router view -->
<router-view class="main"/>
</div>
</template>

<script>
export default {
  name:'NavbarLayout'
}
</script>

<style lang="scss" scoped>

body{
  color: black;
  background-color: white;
  margin: 0;
  padding: 0;
}

.main{
  margin-left: 5rem;
  padding: 1rem;
}

// active class
.active-class{
  background: $primary;
  filter:grayscale(0%)!important;
  color: $text-primary;
}

.navbar{
  position: fixed;
  background-color: $nav-background;
  transition: width 200ms ease;
  z-index: 5;
}

.navbar-nav{
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100%;

  .nav-item{
    width: 100%
  }

  .nav-item:last-child{
    margin-top: auto;
  }

  .nav-link{
    display: flex;
    align-items: center;
    height: 5rem;
    color: $text-primary;
    text-decoration: none;
    filter: grayscale(100%) opacity(0.7);
    transition: $transistion-speed;

    &:hover{
      filter:grayscale(0%);
      background: $nav-background;
      color: $text-primary;
    }

    .link-text{
      display: none;
      margin-left: 1rem;
    }

    i{
      min-width: 2rem;
      margin:0 1.5rem;
    }
  }
}

// nav-item colors
.icon {
  color: $primary;
  font-size: 32px;
  transition: $transistion-speed;
  margin: 0 auto 0;
}

// logo styling
.logo {
  font-weight: bold;
  text-transform: uppercase;
  margin: 1rem 0 1rem;
  text-align: center;
  color: $text-primary;
  font-size: 1.5rem;
  letter-spacing: 0.3ch;
  width: 100%;
  transition: $transistion-speed;
  .nav-link {filter: none!important;}

  img{
    margin: 0 auto 0;
    width:70%;
    height:70%;
  }
}

.logo img {
  transform: rotate(0deg);
  transition: $transistion-speed;
}

.logo-text
{
  display: inline;
  transition: $transistion-speed;
}


// smaller screens
@media only screen and(max-width: 600px){
  .navbar{
    bottom: 0;
    width: 100vw;
    height: 5rem;
    border-top: 1px solid $nav-background;
  }

  .logo{
    display: none;
  }

  .navbar-nav{
    flex-direction: row;
  }

  .nav-link{
    justify-content: center;
  }

  .main{
    margin:0;
    margin-bottom:3rem;
  }
}

// large screen
@media only screen and(min-width: 600px){
  .navbar{
    top: 0;
    width:5rem;
    height: 100vh;
    border-right: 1px solid $nav-background;
    &:hover{
      width:16rem;
      .link-text{
        display: inline!important;
        transition: opacity $transistion-speed;
      }
    }
  }
}
</style>