<template>
  <div>
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
    <div class="abstruct">{{ state.portfolio.abstruct }}</div>
    <div class="readme">
      <MarkdownPreview :text="state.portfolio.readme" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { useRouter } from "vue-router";
import { Portfolio } from "../../../models/portfolio";
import MarkdownPreview from "@/atoms/MarkdownPreview.vue";

export default defineComponent({
  components: { MarkdownPreview },
  props: {
    inputPortfolio: {
      type: Object as PropType<Portfolio>
    }
  },
  setup(props) {
    if (!props.inputPortfolio) {
      const router = useRouter();
      router.push({ name: "NotFoundError" });
      return {};
    }
    const state: {
      portfolio: Portfolio;
    } = { portfolio: props.inputPortfolio };
    if (state.portfolio.readme === undefined) {
      state.portfolio.readme = "";
    }
    return { state };
  }
});
</script>

<style scoped>
.title {
  font-size: x-large;
  margin: 1em;
}

.author {
  margin: 0.5em;
}
.date {
  margin: 0.5em;
}

.abstruct {
  margin: 0.5em;
  /* margin: 10px;
  padding: 5px;
  text-align: left;
  border: solid 1px gray;
  border-radius: 10px; */
}

.readme {
  margin: 0.5em;
  margin: 10px;
  padding: 5px;
  text-align: left;
  border: solid 1px gray;
  border-radius: 10px;
}
</style>
