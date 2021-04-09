<template>
  <div>
    <Card class="filter-card">
        <template #title class="card-title">
            Filtros
        </template>
        <template #content>
        <div class="p-fluid p-formgrid p-grid">
            <div class="p-field p-col-12 p-lg-4">
                <label for="filterName" class="adjust-lineheight">Nome</label>
                <InputText id="filterName" type="text" placeholder="Nome" v-model="filterName" @keyup.enter="onFilter" />
            </div>
            <div class="p-field p-col-12 p-lg-4">
              <div class="p-formgroup-inline">
                  <div class="p-field-checkbox adjust-lineheight">
                      <RadioButton id="rb-cpf" name="cpfcnpj" value="cpf" v-model="filterIdentityType" v-on:change="changedOptionIdentity" />
                      <label for="rb-cpf">CPF</label>
                  </div>
                  <div class="p-field-checkbox">
                      <RadioButton id="rb-cnpj" name="cpfcnpj" value="cnpj" v-model="filterIdentityType" v-on:change="changedOptionIdentity" />
                      <label for="rb-cnpj">CNPJ</label>
                  </div>
              </div>  
              <InputMask id="filterCpfCnpj" type="text" placeholder="Documento" v-model="filterIdentityNumber" :mask="identityMask" @keydown.enter="onFilter" />
            </div>  
            <div class="p-field p-col-12 p-lg-4">
               <label class="adjust-lineheight">Bloqueados</label>
              <div class="p-formgroup-inline adjust-options-blocked">
                  <div class="p-field-checkbox">
                      <RadioButton id="rb-blocked-all" name="blocked" value="all" v-model="filterBlockeds" v-on:change="onFilter" />
                      <label for="city7">Todos</label>
                  </div>
                  <div class="p-field-checkbox">
                      <RadioButton id="rb-blocked-yes" name="blocked" value="yes" v-model="filterBlockeds" v-on:change="onFilter" />
                      <label for="rb-blocked-yes">Sim</label>
                  </div>
                  <div class="p-field-checkbox">
                      <RadioButton id="rb-blocked-no" name="blocked" value="no" v-model="filterBlockeds" v-on:change="onFilter" />
                      <label for="rb-blocked-no">Não</label>
                  </div>                  
              </div> 
            </div>              
            <div class="p-field p-col-12 p-lg-offset-8 p-lg-4">
                <Button type="button" label="Pesquisar" v-on:click="onFilter" />
            </div>                      
        </div> 
        </template>
    </Card>

    <Card>
        <template #title class="card-title">
            Listagem de documentos
        </template>      
      <template #content>

        <Toolbar class="p-mb-2">
            <template #left>
                <Button label="Novo" icon="pi pi-plus" class="p-button-success p-mr-2" @click="openNew" title="Novo documento" />
            </template>
            <template #right>
                <Button label="Bloquear selecionados" icon="pi pi-lock" class="p-button-warning p-mr-6" @click="blockSelecteds(true)" :disabled="!selectedDocuments || !selectedDocuments.length" />
                <Button label="Desbloquear selecionados" icon="pi pi-unlock" class="p-button-warning p-mr-2" @click="blockSelecteds(false)" :disabled="!selectedDocuments || !selectedDocuments.length" />
            </template>            
        </Toolbar>        

        <DataTable :value="documents.content" :lazy="true" :paginator="true" :rows="query.resultsPerPage"        sele
            v-model:selection="selectedDocuments" dataKey="uuid" 
            :totalRecords="documents.totalResults" :loading="loading" @page="onPage($event)" @sort="onSort($event)"           
            paginatorTemplate="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown"
            :rowsPerPageOptions="[2,10,20,50]" responsiveLayout="scroll"
            currentPageReportTemplate="Mostrando {first} até {last} de {totalRecords}">
            <Column selectionMode="multiple" headerStyle="width: 3em"></Column>
            <Column field="name" header="Nome" :sortable="true"></Column>
            <Column field="identityNumber" header="CPF/CNPJ" :sortable="true" headerStyle="text-align: center" bodyStyle="text-align: center">
              <template #body="{data}">
                  {{$filters.formatIdentityValue(data.identityNumber, data.identityType)}}
              </template>            
            </Column>
            <Column field="blocked" header="Bloqueado" :sortable="true" headerStyle="text-align: center" bodyStyle="text-align: center">
              <template #body="{data}">
                <i v-show="data.blocked" class="pi pi-lock" style="fontSize: 28px" />                
                {{ data.blocked ? "" : "Não" }}
              </template> 
            </Column> 
            <Column style="max-width:60px">
                <template #body="{data}">
                    <Button icon="pi pi-pencil" class="p-button-rounded p-button-success p-mr-2" @click="editDocument(data.uuid)" title="Editar" />
                    <Button icon="pi pi-trash" class="p-button-rounded p-button-danger" @click="deleteDocument(data)" title="Excluir" />
                </template>
            </Column>            
        </DataTable> 
      </template> 
    </Card>    
  </div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import DocumentService from '@/service/DocumentService';
import PaginatedResponse, { emptyResponse } from '@/model/PaginatedResponse';
import Query, { emptyQuery } from '@/model/Query';
import Messages, { Message } from '@/model/Messages';
import toQuery from '@/adapter/DataTableToQuery';
import { useToast } from "primevue/usetoast";

@Options({
  props: {
    
  }})
export default class DocumentList extends Vue {
  documents: PaginatedResponse | null = emptyResponse
  documentService = new DocumentService()
  loading = false
  query: Query = emptyQuery
  filterName = ""
  filterIdentityNumber = ""
  filterIdentityType = "cpf"
  filterBlockeds = "all"
  toast: any
  selectedDocuments = []
  identityMask = this.checkMask()
 
  created() {
    this.toast = useToast();
  }  

  mounted() {
    this.findDocuments(); 
  }

  findDocuments() {
    this.loading = true;
    this.selectedDocuments = []
   
    this.documentService.findDocumentsBy(this.query).then((data: PaginatedResponse | Messages) => {
      if ("content" in data)
        this.documents = data as PaginatedResponse;
      else {
        this.documents = emptyResponse;

        (data as Messages).messages.forEach((message: Message) => {
            this.toast.add({severity: message.validationType == "error" ? "error" : "warn", 
              summary:  message.validationType == "error" ? "Erro" : "Validação", detail:message.description, life: 10000});
        });
      }
      this.loading = false;
    });
  }

  onPage(event: any) {
      toQuery(event, this.query);
      this.findDocuments();
  }

  onSort(event: any) {
      toQuery(event, this.query);
      this.findDocuments();
  } 
  
  onFilter() {
    this.query.filters = []

    if (this.filterName)      
      this.query.filters.push({name: "name", value: this.filterName});

    if (this.filterIdentityNumber)      
      this.query.filters.push({name: "identityNumber", value: this.filterIdentityNumber.replace(/[^a-zA-Z0-9]/g,"")});

    if (this.filterBlockeds != "all")      
      this.query.filters.push({name: "blocked", value: this.filterBlockeds == "yes" });      

    this.findDocuments();
  }

  openNew() {
    this.$router.push('/document/')
  }

  editDocument(id: string) {    
    this.$router.push(`/document/${id}`)
  }

  deleteDocument(document: any) {
    this.$confirm.require({
        message: `Tem certeza que deseja excluir o documento "${document.name}"?`,
        header: 'Confirmação',
        icon: 'pi pi-exclamation-triangle',
        accept: () => {
          this.loading = true;

          this.documentService.delete(document.uuid).then((data: null | Messages) => {
            if (data) {
              (data as Messages).messages.forEach((message: Message) => {
                  this.toast.add({severity: message.validationType == "error" ? "error" : "warn", 
                    summary:  message.validationType == "error" ? "Erro" : "Validação", detail:message.description, life: 10000});
              });
              this.loading = false;
            } else {
              this.toast.add({severity: "success", 
                    summary:  "Exclusão realizada com sucesso", detail: `Documento "${document.name}" excluído com sucesso!`, life: 10000});              
              this.findDocuments()
            }
          });
        }
    });
  }

  blockSelecteds(block: boolean) {
          this.loading = true;

          this.documentService.blockDocuments(this.selectedDocuments, block).then((data: null | Messages) => {
            if (data) {
              (data as Messages).messages.forEach((message: Message) => {
                  this.toast.add({severity: message.validationType == "error" ? "error" : "warn", 
                    summary:  message.validationType == "error" ? "Erro" : "Validação", detail:message.description, life: 10000});
              });
              this.loading = false;
            } else {
              this.toast.add({severity: "success", 
                    summary:  "Lista atualizada com sucesso", detail: `Os documentos foram ${block ? "bloqueados" : "desbloqueados"} com sucesso!`, life: 10000});              
              this.findDocuments()
            }
          });
  }

  checkMask() {
    return this.filterIdentityType == "cpf" ? "999.999.999-99" : "99.999.999/9999-99";
  }

  changedOptionIdentity(e: any) {
    this.filterIdentityNumber = "";
  }  
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

.adjust-lineheight{
  line-height: 28px;
}

.adjust-options-blocked {
  margin-top: 10px;
}

.card-title {
    text-align: left;  
}

.filter-card {
  margin-bottom: 20px;
}

</style>
