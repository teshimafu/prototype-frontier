<template>
  <div class="portfolio-list">
    <PortfolioTable :portfolioList="state.portfolioList" />
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import { Portfolio } from "../models/portfolio";
import PortfolioTable from "@/molecules/PortfolioTable.vue";
import { PortfolioService } from "../services/portolioService";

export default defineComponent({
  name: "PortfolioList",
  components: {
    PortfolioTable
  },
  setup() {
    const state = reactive<{ portfolioList: Portfolio[] }>({
      portfolioList: []
    });
    PortfolioService.getPortfolioList().then(res => {
      state.portfolioList = res;
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
