<template>
  <div class="form">
    <Card class="form-card">
        <template #title class="card-title">
            Cadastro
        </template>
        <template #content>

          <div class="p-fluid">
              <div class="p-field">
                  <label for="name">Nome</label>
                  <InputText id="name" type="text" v-model="document.name" />
              </div>
              <div class="p-field">
                  <div class="p-field-checkbox">
                      <RadioButton id="rb-cpf" name="cpfcnpj" value="cpf" v-model="document.identityType" v-on:change="changedOptionIdentity" />
                      <label for="rb-cpf">CPF</label>
                  </div>
                  <div class="p-field-checkbox">
                      <RadioButton id="rb-cnpj" name="cpfcnpj" value="cnpj" v-model="document.identityType" v-on:change="changedOptionIdentity" />
                      <label for="rb-cnpj">CNPJ</label>
                  </div>
              </div>
              <div class="p-field">
                  <InputMask id="identityNumber" type="text" v-model="document.identityNumber" :mask="identityMask" />
              </div>       
              <div class="p-field p-grid">
                <div class="p-col-6">
                  <Button type="button" label="Salvar" class="p-button-success" :disabled="loading" v-on:click="save" />
                </div>
                <div class="p-col-6">
                  <Button type="button" label="Voltar" class="p-button-information" :disabled="loading" v-on:click="back" />
                </div>
              </div>                      
          </div> 
        </template>
    </Card>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import DocumentService from '@/service/DocumentService';
import Document from '@/model/Document';
import Messages, {Message} from '@/model/Messages';

@Options({
  props: {
    
  }})
export default class DocumentForm extends Vue {
  documentService = new DocumentService()
  loading = false
  document: Document = this.resetData()
  identityMask = this.checkMask()

  mounted() {
    if (this.$route.params.id) {
      
      this.documentService.get(this.$route.params.id as string).then((data: Document | Messages) => {        
        if ("name" in data)
          this.document = data as Document;
        else {
          this.resetData();

          (data as Messages).messages.forEach((message: Message) => {
              this.$toast.add({severity:message.validationType == "error" ? "error" : "warn", summary: message.validationType == "error" ? "Erro" : "Validação", detail:message.description, life: 10000});          
          });
        }
      });

    } else {
      this.resetData();
    }
  }

  resetData() {
    return this.document = {
      uuid: "",
      name: "",
      identityNumber: "",
      identityType: "cpf"
    }
  }

  checkMask() {
    return this.document.identityType == "cpf" ? "999.999.999-99" : "99.999.999/9999-99";
  }

  changedOptionIdentity(e: any) {
    this.document.identityNumber = "";
  }

  save() {
    const save = this.document.uuid ? this.documentService.put(this.document) : this.documentService.post(this.document)

    save.then((data: null | Messages) => {        
      if (data) {
        (data as Messages).messages.forEach((message: Message) => {
            this.$toast.add({severity:message.validationType == "error" ? "error" : "warn", summary: message.validationType == "error" ? "Erro" : "Validação", detail:message.description, life: 10000});          
        });
      } else {
        this.$toast.add({severity: "success", summary: "Documento salvo com sucesso", detail:`Documento "${this.document.name}" salvo com sucesso!`, life: 10000});                  
        this.$router.push("/");
      }
    });

  }

  back() {
    this.$router.push("/");
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

.card-title {
    text-align: left;  
}

.form {
    display: flex;
    justify-content: center;
    align-items: center;
}

.form-card {
  @media (min-width: 960px){
      min-width: 600px;
  }

  @media (max-width: 960px){
      width: 100%;
  }  
}

</style>
