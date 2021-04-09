<template>
    <DocumentList />
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import DocumentList from "./components/DocumentList.vue";

@Options({
  components: {
    DocumentList,
  },
})
export default class HomePage extends Vue {}
</script>

<style lang="scss">
</style>
