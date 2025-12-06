<script lang="ts" setup>
import TrafficSimulator from './components/TrafficSimulator.vue'
</script>

<template>
  <main>
    <h1>Traffic Simulator</h1>
    <TrafficSimulator />
  </main>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  background: #1a1a1a;
  color: #fff;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
}

main {
  min-height: 100vh;
  padding: 20px;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
  color: #4CAF50;
  font-size: 2rem;
}
</style>
