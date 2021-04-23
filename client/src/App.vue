<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<style lang="scss">
:root{
  background:$background;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color:$text-primary;
  background:$secondary;
  margin:0;
  padding:0;
}

</style>
