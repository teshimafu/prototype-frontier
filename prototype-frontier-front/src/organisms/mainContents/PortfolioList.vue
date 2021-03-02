<template>
  <div class="portfolio-list">
    <div id="abstruct">
      自分の作ったアプリケーションを投稿してみんなに見てもらいましょう<br />
      将来的にtwitterでランダムにアプリを投稿する機能も追加予定です
    </div>
    <button @click="createNewPortfolio">新規作品を投稿する</button>
    <PortfolioTable :portfolioList="state.portfolioList" />
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import { useRouter } from "vue-router";
import { Portfolio } from "../../models/portfolio";
import PortfolioTable from "./List/PortfolioTable.vue";
import PortfolioService from "../../services/portolioService.vue";

export default defineComponent({
  name: "PortfolioList",
  components: {
    PortfolioTable
  },
  setup() {
    const router = useRouter();
    const state = reactive<{ portfolioList: Portfolio[] }>({
      portfolioList: []
    });
    const createNewPortfolio = () => {
      router.push("/detail/new");
    };
    PortfolioService.getPortfolioList()
      .then(res => {
        state.portfolioList = res;
      })
      .catch(e => {
        if (e.response?.status == 404) {
          router.push("/error/notfound");
          return;
        }
        router.push("/error/loaderror");
      });
    return { state, createNewPortfolio };
  }
});
</script>

<style>
#abstruct {
  background-color: rgba(127, 255, 212, 0.3);
}
.portfolio-list {
  margin: auto;
}
</style>
