<template>
  <div v-if="state">
    <div class="title">{{ state.portfolio.title }}</div>
    <div>
      <span class="date">作成日:{{ state.portfolio.createDate }}</span>
      <span class="author">作成者:{{ state.portfolio.author }}</span>
    </div>
    <div>
      成果物リンク:
      <a :href="state.portfolio.link">{{ state.portfolio.link }}</a>
    </div>
    <div>
      ソースコード:
      <a :href="state.portfolio.source">{{ state.portfolio.source }}</a>
    </div>
    <div class="abstruct" v-html="state.portfolio.abstruct"></div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { useRouter } from "vue-router";
import marked from "marked";
import { Portfolio } from "../../../models/portfolio";

interface Contect {
  state: {
    portfolio: Portfolio;
  };
}

export default defineComponent({
  props: {
    portfolio: {
      type: Object as PropType<Portfolio>
    }
  },
  setup(props) {
    if (!props.portfolio) {
      const router = useRouter();
      router.push("/error");
      return {};
    }
    const state: {
      portfolio: Portfolio;
    } = { portfolio: props.portfolio };
    state.portfolio.abstruct = marked(state.portfolio.abstruct);
    return { state };
  }
});
</script>
