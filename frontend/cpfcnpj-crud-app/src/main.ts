import { createApp } from "vue";
import { createRouter, createWebHistory } from "vue-router";
import App from "./App.vue";
import DocumentPage from "./DocumentPage.vue";
import HomePage from "./HomePage.vue";
import store from "./store";
import PrimeVue from "primevue/config";
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import Card from "primevue/card";
import InputText from "primevue/inputtext";
import InputMask from "primevue/inputmask";
import RadioButton from "primevue/radiobutton";
import Button from "primevue/button";
import ConfirmDialog from "primevue/confirmdialog";
import ToastService from "primevue/toastservice";
import ConfirmationService from "primevue/confirmationservice";
import Toast from "primevue/toast";
import Toolbar from "primevue/toolbar";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faCoffee } from "@fortawesome/free-solid-svg-icons/faCoffee";
import { faHeart } from "@fortawesome/free-solid-svg-icons/faHeart";

library.add(faCoffee, faHeart);

//import "primevue/resources/themes/md-light-indigo/theme.css";
import "primevue/resources/themes/saga-blue/theme.css";
import "primevue/resources/primevue.min.css";
import "primeicons/primeicons.css";
import "primeflex/primeflex.css";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", component: HomePage },
    { path: "/document/", component: DocumentPage },
    { path: "/document/:id", component: DocumentPage },
  ],
});

const app = createApp(App);
//const app = createApp({});
app.use(router);
app.use(store);
app.use(PrimeVue);
app.use(ToastService);
app.use(ConfirmationService);

app.component("DataTable", DataTable);
app.component("Column", Column);
app.component("Card", Card);
app.component("InputText", InputText);
app.component("InputMask", InputMask);
app.component("RadioButton", RadioButton);
app.component("Button", Button);
app.component("Toast", Toast);
app.component("Toolbar", Toolbar);
app.component("ConfirmDialog", ConfirmDialog);
app.component("font-awesome-icon", FontAwesomeIcon);

app.mount("#app");

app.config.globalProperties.$filters = {
  formatIdentityValue(value: string, type: string) {
    if (type == "cpf")
      return value.replace(/(\d{3})(\d{3})(\d{3})(\d{2})/, "$1.$2.$3-$4");
    else
      return value.replace(
        /(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})/,
        "$1.$2.$3/$4-$5"
      );
  },
};
