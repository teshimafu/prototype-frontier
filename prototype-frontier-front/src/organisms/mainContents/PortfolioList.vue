<template>
  <div class="portfolio-list">
    <div id="abstruct">
      自分の作ったアプリケーションを投稿してみんなに見てもらいましょう
      <br />将来的にtwitterでランダムにアプリを投稿する機能も追加予定です
    </div>
    <Button
      v-if="state.isEditable"
      @click="createNewPortfolio"
      :text="'新規作品を投稿する'"
    />
    <div class="search-unit">
      <SearchUnit @search="searchPortfolio" />
    </div>
    <PortfolioTable :portfolioList="state.portfolioList" />
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import { useRouter } from "vue-router";
import Button from "@/atoms/Button.vue";
import SearchUnit from "@/molecules/SearchUnit.vue";
import { Portfolio, SearchQuery } from "../../models/portfolio";
import PortfolioTable from "./List/PortfolioTable.vue";
import PortfolioService from "../../services/portolioService";
import FirebaseService from "../../services/firebaseService";

export default defineComponent({
  name: "PortfolioList",
  components: {
    Button,
    SearchUnit,
    PortfolioTable
  },
  setup() {
    const router = useRouter();
    const state = reactive<{ portfolioList: Portfolio[]; isEditable: boolean }>(
      {
        portfolioList: [],
        isEditable: false
      }
    );
    const createNewPortfolio = () => {
      router.push("/detail/new");
    };
    const searchPortfolio = (text: string) => {
      const query: SearchQuery = {
        title: text,
        author: text,
        readme: text
      };
      PortfolioService.getPortfolioList(query)
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
    };

    FirebaseService.getInstance()
      .getUser()
      .then(user => {
        if (user) {
          state.isEditable = true;
        }
      });
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
    return { state, createNewPortfolio, searchPortfolio };
  }
});
</script>

<style>
#abstruct {
  background-color: #ffffcc;
}
.search-unit {
  margin: 0.2rem;
}
.portfolio-list {
  margin: auto;
}
</style>
