import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import PortfolioList from "../organisms/PortfolioList.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "PortolioList",
    component: PortfolioList
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;
