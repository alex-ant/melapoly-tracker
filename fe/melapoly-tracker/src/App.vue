<template>
<div id="app">
  <Tracker />
</div>
</template>

<script>
import Tracker from './components/Tracker'

export default {
  name: 'app',
  components: {
    Tracker
  },
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
