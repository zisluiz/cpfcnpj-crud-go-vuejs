<template>
    <header>
        <h1>CPF/CNPJ CRUD APP</h1>
    </header>
    <main>
        <!--<HelloWorld msg="Welcome to Your Vue.js + TypeScript App" />-->
        <slot></slot>

        <Toast position="bottom-right" />
        <ConfirmDialog></ConfirmDialog>
    </main>
    <footer>
        Developed by zisluiz with much 
        <font-awesome-icon :icon="['fas', 'coffee']"/> and 
        <font-awesome-icon :icon="['fas', 'heart']"/>
    </footer>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";

@Options({
})
export default class App extends Vue {
}
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

footer {
    position: fixed;
    left: 0;
    bottom: 0;
    padding-top: 8px;
    height: 36px;
    width: 100%;
    background-color: var(--primary-color);
    color: white;
}

</style>
