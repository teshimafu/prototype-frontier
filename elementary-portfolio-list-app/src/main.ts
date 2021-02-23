import { createApp } from "vue";
import App from "@/pages/App.vue";
import router from "./router";

createApp(App)
  .use(router)
  .mount("#app");
