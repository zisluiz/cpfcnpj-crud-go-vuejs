<template>  
  <MainLayout>
    <div id="app">
      <router-view></router-view>
    </div>
    </MainLayout>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import MainLayout from "./layout/MainLayout.vue";

@Options({
  components: {
    MainLayout,
  },
})
export default class App extends Vue {}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

</style>
