<template>
  <Preview v-if="state.portfolio" v-bind:portfolio="state.portfolio" />
  <Loading v-else />
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import { useRoute, useRouter } from "vue-router";
import Loading from "@/organisms/Loading.vue";
import Preview from "./Detail/Preview.vue";
import { Portfolio } from "../../models/portfolio";
import PortfolioService from "../../services/portolioService.vue";

interface Contect {
  state: {
    portfolio: Portfolio | null;
  };
}

export default defineComponent({
  components: {
    Loading,
    Preview
  },
  setup(): Contect | undefined {
    const route = useRoute();
    const router = useRouter();
    const id = Number(route.params.id);
    if (isNaN(id)) {
      router.push("/");
      return;
    }
    const state = reactive<{ portfolio: Portfolio | null }>({
      portfolio: null
    });
    PortfolioService.getPortfolio(id)
      .then(r => {
        state.portfolio = r;
      })
      .catch(e => {
        if (e.response?.status == 404) {
          router.push("/error/notfound");
          return;
        }
        router.push("/error/loaderror");
      });
    return { state };
  }
});
</script>
