<template>
  <HomeComponent />
</template>

<script>

//import { getToken } from "../services/index";
import HomeComponent from "./components/HomeComponent.vue";
export default {
  name: "App",
  components: {
    HomeComponent,
  }
};
</script>

<style>
#app {
  font-family: Roboto, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
