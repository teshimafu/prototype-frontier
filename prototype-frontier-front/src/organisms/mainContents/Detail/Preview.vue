<template>
  <div class="preview">
    <div class="title">{{ state.portfolio.title }}</div>
    <ul>
      <li class="date">作成日:{{ showDate(state.portfolio.created_at) }}</li>
      <li class="author">作成者:{{ state.portfolio.author }}</li>
      <li>
        成果物リンク:
        <a :href="state.portfolio.link">{{ state.portfolio.link }}</a>
      </li>
      <li>
        ソースコード:
        <a :href="state.portfolio.source">{{ state.portfolio.source }}</a>
      </li>
    </ul>
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

    const showDate = (text: string) => {
      if (!text.includes("T") || text.length < 19) return "";
      const date = text.split("T")[0];
      const time = text.split("T")[1].substring(0, 8);
      return date + " " + time;
    };
    return { state, showDate };
  }
});
</script>

<style scoped>
.preview {
  width: 80%;
  margin: auto;
}
.title {
  font-size: x-large;
  margin: 1em;
}

ul li {
  list-style: none;
  text-align: left;
  margin: 0.2rem;
}
.abstruct {
  margin: 0.5em;
}

.readme {
  margin: 0.5em;
  margin: 10px;
  padding: 5px;
  text-align: left;
  border: solid 1px gray;
  border-radius: 10px;
  min-height: 2rem;
}
</style>
