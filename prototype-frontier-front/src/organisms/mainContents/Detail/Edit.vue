<template>
  <div>
    <div class="title">{{ state.portfolio.title }}</div>
    <div>
      <span class="date">作成日:{{ state.portfolio.createDate }}</span>
      <span class="author">作成者:{{ state.portfolio.author }}</span>
    </div>
    <div>
      成果物リンク:
      <a>{{ state.portfolio.link }}</a>
    </div>
    <div>
      ソースコード:
      <a>{{ state.portfolio.source }}</a>
    </div>
    <div class="abstruct">{{ state.portfolio.abstruct }}</div>
    <MarkdownEditor
      :markdownText="state.portfolio.readme"
      @returnText="setReadme"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { useRouter } from "vue-router";
import MarkdownEditor from "@/molecules/MarkdownEditor.vue";
import { InputPortfolio } from "../../../models/portfolio";
export default defineComponent({
  components: {
    MarkdownEditor
  },
  props: {
    inputPortfolio: Object as PropType<InputPortfolio>
  },
  setup(props) {
    if (!props.inputPortfolio) {
      const router = useRouter();
      router.push("/error");
      return {};
    }
    const state: {
      portfolio: InputPortfolio;
    } = { portfolio: props.inputPortfolio };
    return { state };
  },
  methods: {
    setReadme(text: string) {
      if (this.state) {
        this.state.portfolio.readme = text;
      }
    }
  }
});
</script>
