<template>
  <div>
    <ul class="editor-box">
      <li>
        <InputUnit
          :title="'タイトル'"
          :input="state.portfolio.title"
          @model="setInputValue('title', $event)"
        />
      </li>
      <li>
        <InputUnit
          :title="'作成者'"
          :input="state.portfolio.author"
          @model="setInputValue('author', $event)"
        />
      </li>
      <li>
        <InputUnit
          :title="'成果物リンク'"
          :input="state.portfolio.link"
          @model="setInputValue('link', $event)"
        />
      </li>
      <li>
        <InputUnit
          :title="'ソースコード'"
          :input="state.portfolio.source"
          @model="setInputValue('source', $event)"
        />
      </li>
      <li>
        <InputUnit
          :title="'キーワード'"
          :input="state.portfolio.abstruct"
          @model="setInputValue('abstruct', $event)"
        />
      </li>
    </ul>
    <div>
      <div>詳細説明</div>
      <MarkdownEditor
        :markdownText="state.portfolio.readme"
        @returnText="setReadme"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { useRouter } from "vue-router";
import InputUnit from "@/molecules/InputUnit.vue";
import MarkdownEditor from "@/molecules/MarkdownEditor.vue";
import { InputPortfolio } from "../../../models/portfolio";
export default defineComponent({
  components: {
    InputUnit,
    MarkdownEditor
  },
  props: {
    inputPortfolio: Object as PropType<InputPortfolio>
  },
  setup(props) {
    if (!props.inputPortfolio) {
      const router = useRouter();
      router.push({ name: "NotFoundError" });
      return {};
    }
    const state: {
      portfolio: InputPortfolio;
    } = { portfolio: props.inputPortfolio };
    const setReadme = function(text: string) {
      state.portfolio.readme = text;
    };
    const setInputValue = function(
      name: keyof Omit<InputPortfolio, "id" | "created_at">,
      text: string
    ) {
      state.portfolio[name] = text;
    };
    return { state, setReadme, setInputValue };
  }
});
</script>

<style scoped>
ul li {
  list-style: none;
  margin: 0.2rem;
}
.editor-box {
  width: 80%;
  margin: auto;
}
</style>
