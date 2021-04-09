<template>
    <DocumentForm />
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import DocumentForm from "./components/DocumentForm.vue";

@Options({
  components: {
    DocumentForm,
  },
})
export default class DocumentPage extends Vue {}
</script>

<style lang="scss">
</style>
