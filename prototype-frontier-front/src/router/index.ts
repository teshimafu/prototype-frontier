import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import PortfolioList from "../organisms/mainContents/PortfolioList.vue";
import PortfolioDetail from "../organisms/mainContents/PortfolioDetail.vue";
import NotFoundError from "../organisms/mainContents/NotFoundError.vue";
import LoadError from "../organisms/mainContents/LoadError.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    redirect: "/list"
  },
  {
    path: "/list",
    name: "PortolioList",
    component: PortfolioList
  },
  {
    path: "/detail/:id",
    name: "PortolioDetail",
    component: PortfolioDetail
  },
  {
    path: "/error/notfound",
    name: "NotFoundError",
    component: NotFoundError
  },
  {
    path: "/error/loaderror",
    name: "LoadError",
    component: LoadError
  },
  {
    path: "/:pathMatch(.*)*",
    name: "InvalidPath",
    redirect: "/error/notfound"
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;
