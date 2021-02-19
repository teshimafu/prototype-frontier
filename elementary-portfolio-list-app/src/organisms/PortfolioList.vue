<template>
  <div class="portfolio-list">
    <PortfolioTable :portfolioList="portfolioList" />
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import { Portfolio } from "../models/Portfolio";
import PortfolioTable from "@/molecules/PortfolioTable.vue";

export default defineComponent({
  name: "PortfolioList",
  components: {
    PortfolioTable
  },
  setup() {
    const portfolioList = reactive<Portfolio[]>([
      {
        id: 1,
        title: "シロナガスクジラ",
        author: "teshima",
        createDate: "1/1",
        tags: ["1", "2", "3"],
        abstruct: "## title etc...",
        source: "https:google.com",
        url: "https://qiita.com"
      }
    ]);
    return {
      portfolioList
    };
  }
});
</script>

<style>
.portfolio-list {
  margin: auto;
}
</style>
