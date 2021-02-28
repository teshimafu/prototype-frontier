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
  setup(): Contect {
    const route = useRoute();
    const router = useRouter();
    const id = Number(route.params.id);
    if (isNaN(id)) {
      router.push("/");
    }
    const state = reactive<{ portfolio: Portfolio | null }>({
      portfolio: null
    });
    PortfolioService.getPortfolio(id).then(r => {
      state.portfolio = r;
    });
    return { state };
  }
});
</script>

<style scoped>
.title {
  font-size: x-large;
}

.date {
  margin: 0.5em;
}

.author {
  margin: 0.5em;
}

.abstruct {
  margin: 10px;
  padding: 5px;
  text-align: left;
  border: solid 1px gray;
  border-radius: 10px;
}
</style>
