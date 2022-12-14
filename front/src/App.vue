<template>
  <div>
    <g-header />
    <router-view></router-view>
  </div>
</template>


<script>
import gHeader from './components/g-header.vue'


export default {
  components: { gHeader }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
