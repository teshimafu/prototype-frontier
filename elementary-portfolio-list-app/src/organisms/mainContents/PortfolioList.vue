<template>
  <div class="portfolio-list">
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
    const state = reactive<{ portfolioList: Portfolio[] }>({
      portfolioList: []
    });
    PortfolioService.getPortfolioList()
      .then(res => {
        state.portfolioList = res;
      })
      .catch(e => {
        const router = useRouter();
        if (e.response.status == 404) {
          router.push("/error/notfound");
          return;
        }
        router.push("/error/loaderror");
      });
    return { state };
  }
});
</script>

<style>
.portfolio-list {
  margin: auto;
}
</style>
