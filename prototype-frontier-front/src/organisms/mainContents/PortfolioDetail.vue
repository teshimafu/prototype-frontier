<template>
  <Edit v-if="state.isEdit" :inputPortfolio="state.portfolio" />
  <Preview v-else-if="state.portfolio" :inputPortfolio="state.portfolio" />
  <Loading v-else />
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import { useRoute, useRouter } from "vue-router";
import Loading from "@/organisms/Loading.vue";
import Preview from "./Detail/Preview.vue";
import Edit from "./Detail/Edit.vue";
import { InputPortfolio } from "../../models/portfolio";
import PortfolioService from "../../services/portolioService.vue";

interface Content {
  isEdit: boolean;
  portfolio: InputPortfolio | null;
}

const initPortfolio: InputPortfolio = {
  title: "",
  author: "",
  abstruct: "",
  source: "",
  link: "",
  readme: ""
};

export default defineComponent({
  components: {
    Loading,
    Edit,
    Preview
  },
  setup(): { state: Content } | undefined {
    const route = useRoute();
    const router = useRouter();
    const state = reactive<Content>({
      isEdit: false,
      portfolio: null
    });
    if ("new" === route.params.id) {
      state.portfolio = initPortfolio;
      state.isEdit = true;
      return { state };
    }
    const id = Number(route.params.id);
    if (isNaN(id)) {
      router.push("/error/notfound");
      return;
    }
    PortfolioService.getPortfolio(id)
      .then(r => {
        state.portfolio = r;
      })
      .catch(e => {
        if (e.response?.status == 404) {
          router.push("/error/notfound");
        } else {
          router.push("/error/loaderror");
        }
      });
    return { state };
  }
});
</script>
