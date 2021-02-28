import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import PortfolioList from "../organisms/mainContents/PortfolioList.vue";
import PortfolioDetail from "../organisms/mainContents/PortfolioDetail.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "PortolioList",
    component: PortfolioList
  },
  {
    path: "/detail/:id",
    name: "PortolioDetail",
    component: PortfolioDetail
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;
